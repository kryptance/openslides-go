package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
)

// HistoryRelation defines a related object whose history should be extended
// when the main action fires.
type HistoryRelation struct {
	// OwnField is the field on the current model that references the related model.
	OwnField string

	// TargetCollection is the collection of the related model.
	TargetCollection string

	// InfoText is the history information text to record for the related object.
	InfoText string
}

// WithExtendHistory wraps the CreateEvents hook to add history information
// entries for related objects when events are generated.
//
// For example, when a motion_submitter is created, we may want to record
// a history entry on the motion itself. This mixin adds information entries
// to the write request so that the history of related objects is extended.
//
// This replaces Python's ExtendHistoryMixin.
func WithExtendHistory(a *action.BaseAction, relations ...HistoryRelation) {
	origCreateEvents := a.CreateEvents
	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		events, err := origCreateEvents(ctx, params, instance)
		if err != nil {
			return nil, err
		}

		// For each history relation, record information entries.
		for _, rel := range relations {
			relVal, ok := instance[rel.OwnField]
			if !ok {
				continue
			}

			// Resolve the related model IDs.
			relIDs := resolveRelationIDs(relVal)
			for _, relID := range relIDs {
				fqid := event.FQID(rel.TargetCollection, relID)

				// Create an information event (update with no field changes)
				// that carries the history text in the event metadata.
				infoEvent := event.Event{
					Type:   event.TypeUpdate,
					FQID:   fqid,
					Fields: map[string]any{},
				}

				// Attach history information text.
				// The WriteRequest.AddInformation method will be called by the
				// handler to persist this. For now, we embed it as a meta field.
				infoEvent.Fields["meta_info_text"] = fmt.Sprintf(
					"%s %s", a.Name, rel.InfoText,
				)
				events = append(events, infoEvent)
			}
		}

		return events, nil
	}
}

// resolveRelationIDs extracts integer IDs from a relation field value.
func resolveRelationIDs(val any) []int {
	switch v := val.(type) {
	case int:
		return []int{v}
	case float64:
		return []int{int(v)}
	case []any:
		ids := make([]int, 0, len(v))
		for _, elem := range v {
			switch e := elem.(type) {
			case float64:
				ids = append(ids, int(e))
			case int:
				ids = append(ids, e)
			}
		}
		return ids
	case []int:
		return v
	default:
		return nil
	}
}
