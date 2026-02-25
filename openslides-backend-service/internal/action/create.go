package action

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

// NewCreateAction creates a generic create action for a collection.
func NewCreateAction(name string, m *model.ModelDef) *BaseAction {
	a := NewBaseAction(name, m)

	// Default CreateEvents for create actions: generate a create event with all fields.
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

		fields := make(map[string]any, len(instance))
		for k, v := range instance {
			if k == "meta_new" {
				continue
			}
			fields[k] = v
		}

		fqid := event.FQID(m.Collection, idInt)
		e := event.Event{
			Type:   event.TypeCreate,
			FQID:   fqid,
			Fields: fields,
		}

		return []event.Event{e}, nil
	}

	// Wrap UpdateInstance to set defaults and reserve IDs.
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *ActionParams, instance Instance) (Instance, error) {
		// Reserve an ID if not already set.
		if _, ok := instance["id"]; !ok {
			ids, err := params.Datastore.ReserveIDs(m.Collection, 1)
			if err != nil {
				return nil, fmt.Errorf("reserve id for %s: %w", m.Collection, err)
			}
			instance["id"] = ids[0]
		}

		// Set defaults from model definition.
		for fieldName, fieldDef := range m.Fields {
			if _, ok := instance[fieldName]; !ok && fieldDef.Default != nil {
				instance[fieldName] = fieldDef.Default
			}
		}

		// Mark as new for datastore tracking.
		instance["meta_new"] = true

		// Call original (or mixin-wrapped) UpdateInstance.
		result, err := origUpdate(ctx, params, instance)
		if err != nil {
			return nil, err
		}

		// Apply to extended datastore.
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
