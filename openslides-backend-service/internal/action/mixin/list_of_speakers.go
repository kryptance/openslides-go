package mixin

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
)

// WithListOfSpeakers wraps the UpdateInstance hook to create a list of speakers
// for the new model instance.
//
// This replaces Python's ListOfSpeakersCreationMixin.
func WithListOfSpeakers(a *action.BaseAction) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		result, err := origUpdate(ctx, params, instance)
		if err != nil {
			return nil, err
		}

		// TODO: Execute list_of_speakers.create as sub-action.
		// The list of speakers should reference this model via content_object_id.

		return result, nil
	}
}
