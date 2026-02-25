package mixin

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
)

// WithWeight wraps the UpdateInstance hook to automatically set the weight
// field to max(existing weights) + 1 when not explicitly provided.
//
// This replaces Python's WeightMixin.
func WithWeight(a *action.BaseAction, weightField string) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// If weight is not set, auto-assign next weight.
		if _, ok := instance[weightField]; !ok {
			// TODO: Query max weight from datastore and set weight + 1.
			instance[weightField] = 1
		}

		return origUpdate(ctx, params, instance)
	}
}
