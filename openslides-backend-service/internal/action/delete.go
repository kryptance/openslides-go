package action

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/event"
	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

// NewDeleteAction creates a generic delete action for a collection.
// It handles cascade and protect on-delete behaviors.
func NewDeleteAction(name string, m *model.ModelDef) *BaseAction {
	a := NewBaseAction(name, m)

	// Default CreateEvents for delete actions: generate a delete event.
	a.CreateEvents = func(ctx context.Context, params *ActionParams, instance Instance) ([]event.Event, error) {
		id, ok := instance["id"]
		if !ok {
			return nil, fmt.Errorf("instance has no id field")
		}

		var idInt int
		switch v := id.(type) {
		case float64:
			idInt = int(v)
		case int:
			idInt = v
		default:
			return nil, fmt.Errorf("id has unexpected type %T", id)
		}

		fqid := event.FQID(m.Collection, idInt)
		e := event.Event{
			Type: event.TypeDelete,
			FQID: fqid,
		}

		return []event.Event{e}, nil
	}

	// Wrap UpdateInstance to handle on-delete cascading and protection.
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *ActionParams, instance Instance) (Instance, error) {
		id, ok := instance["id"]
		if !ok {
			return nil, fmt.Errorf("instance has no id field")
		}

		var idInt int
		switch v := id.(type) {
		case float64:
			idInt = int(v)
		case int:
			idInt = v
		default:
			return nil, fmt.Errorf("id has unexpected type %T", id)
		}

		// Mark as deleted in the extended datastore.
		params.Datastore.ApplyToBeDeleted(m.Collection, idInt)

		// Check on_delete policies for all relation fields.
		var protectedFQIDs []string
		for fieldName, fieldDef := range m.RelationFields() {
			if fieldDef.Relation == nil {
				continue
			}

			switch fieldDef.Relation.OnDelete {
			case model.OnDeleteProtect:
				// Check if any related models exist.
				relatedModel, found := params.Datastore.GetChangedModel(m.Collection, idInt)
				if found {
					if val, ok := relatedModel[fieldName]; ok && val != nil {
						targetCollection := fieldDef.Relation.To.Collection
						switch v := val.(type) {
						case []any:
							for _, relID := range v {
								protectedFQIDs = append(protectedFQIDs, fmt.Sprintf("%s/%v", targetCollection, relID))
							}
						case float64:
							if v != 0 {
								protectedFQIDs = append(protectedFQIDs, fmt.Sprintf("%s/%v", targetCollection, int(v)))
							}
						case int:
							if v != 0 {
								protectedFQIDs = append(protectedFQIDs, fmt.Sprintf("%s/%d", targetCollection, v))
							}
						}
					}
				}

			case model.OnDeleteCascade:
				// TODO: Execute cascading deletes via sub-actions.

			case model.OnDeleteSetNull:
				// Handled by the relation manager.
			}
		}

		if len(protectedFQIDs) > 0 {
			return nil, backenderr.ProtectedModelsError{FQIDs: protectedFQIDs}
		}

		result, err := origUpdate(ctx, params, instance)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	return a
}
