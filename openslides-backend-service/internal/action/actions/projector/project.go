package projector

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("projector.project", model.MustGet("projector"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"content_object_id": map[string]any{"type": "string"},
			"meeting_id":        map[string]any{"type": "integer"},
			"stable":            map[string]any{"type": "boolean"},
			"type":              map[string]any{"type": "string"},
			"options":           map[string]any{"type": "object"},
		},
		"required": []string{"ids", "content_object_id", "meeting_id"},
	}

	mixin.WithSingularAction(a)

	// Project content onto one or more projectors.
	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		idsRaw, ok := instance["ids"]
		if !ok {
			return nil, fmt.Errorf("ids field is required")
		}

		projectorIDs, ok := toIntSlice(idsRaw)
		if !ok {
			return nil, fmt.Errorf("ids must be an array of integers")
		}

		contentObjectID, _ := instance["content_object_id"].(string)
		meetingID := toInt(instance["meeting_id"])
		stable, _ := instance["stable"].(bool)
		projType, _ := instance["type"].(string)
		options, _ := instance["options"].(map[string]any)

		var events []event.Event

		for _, projID := range projectorIDs {
			// Create a new projection for each projector.
			projectionIDs, err := params.Datastore.ReserveIDs("projection", 1)
			if err != nil {
				return nil, fmt.Errorf("reserve projection id: %w", err)
			}
			projectionID := projectionIDs[0]

			projFields := map[string]any{
				"id":                   projectionID,
				"meeting_id":           meetingID,
				"content_object_id":    contentObjectID,
				"current_projector_id": projID,
				"stable":               stable,
			}
			if projType != "" {
				projFields["type"] = projType
			}
			if options != nil {
				projFields["options"] = options
			}

			projFQID := event.FQID("projection", projectionID)
			events = append(events, event.Event{
				Type:   event.TypeCreate,
				FQID:   projFQID,
				Fields: projFields,
			})
			params.Datastore.ApplyChangedModel("projection", projectionID, projFields)
		}

		return events, nil
	}

	action.Register(a)
}
