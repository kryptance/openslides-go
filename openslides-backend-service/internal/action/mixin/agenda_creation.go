package mixin

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
)

// WithAgendaCreation wraps the UpdateInstance hook to optionally create an
// agenda item for the new model instance.
//
// This replaces Python's AgendaItemCreationMixin.
func WithAgendaCreation(a *action.BaseAction) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		result, err := origUpdate(ctx, params, instance)
		if err != nil {
			return nil, err
		}

		// Check if agenda_create flag is set.
		agendaCreate, ok := result["agenda_create"]
		if !ok || agendaCreate != true {
			// Remove agenda fields that shouldn't be stored on the model.
			delete(result, "agenda_create")
			delete(result, "agenda_type")
			delete(result, "agenda_parent_id")
			delete(result, "agenda_comment")
			delete(result, "agenda_duration")
			delete(result, "agenda_weight")
			return result, nil
		}

		// TODO: Execute agenda_item.create as sub-action.
		// The agenda item should reference this model via content_object_id.

		// Clean up agenda fields from the instance.
		delete(result, "agenda_create")
		delete(result, "agenda_type")
		delete(result, "agenda_parent_id")
		delete(result, "agenda_comment")
		delete(result, "agenda_duration")
		delete(result, "agenda_weight")

		return result, nil
	}
}
