package action

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

// NewUpdateAction creates a generic update action for a collection.
func NewUpdateAction(name string, m *model.ModelDef) *BaseAction {
	a := NewBaseAction(name, m)

	// Default CreateEvents for update actions: generate an update event with changed fields.
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

		// Build update fields (exclude id and meta_ prefixed fields).
		fields := make(map[string]any)
		for k, v := range instance {
			if k == "id" || len(k) > 5 && k[:5] == "meta_" {
				continue
			}
			fields[k] = v
		}

		if len(fields) == 0 {
			return nil, nil
		}

		fqid := event.FQID(m.Collection, idInt)
		e := event.Event{
			Type:   event.TypeUpdate,
			FQID:   fqid,
			Fields: fields,
		}

		return []event.Event{e}, nil
	}

	// Wrap UpdateInstance to apply changes to the extended datastore.
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *ActionParams, instance Instance) (Instance, error) {
		result, err := origUpdate(ctx, params, instance)
		if err != nil {
			return nil, err
		}

		idVal := result["id"]
		var idInt int
		switch v := idVal.(type) {
		case float64:
			idInt = int(v)
		case int:
			idInt = v
		}

		params.Datastore.ApplyChangedModel(m.Collection, idInt, map[string]any(result))

		return result, nil
	}

	return a
}
