package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
)

// WithLinearSort wraps the UpdateInstance hook to sort a flat list of model IDs
// by assigning sequential weight values starting from 1.
//
// The instance should contain an "ids" field with the list of model IDs in the
// desired order. Each ID is validated to exist, and its weight field is updated.
//
// This replaces Python's LinearSortMixin.
func WithLinearSort(a *action.BaseAction, weightField string) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		result, err := origUpdate(ctx, params, instance)
		if err != nil {
			return nil, err
		}

		idsRaw, ok := result["ids"]
		if !ok {
			return nil, backenderr.ActionError{
				Message: fmt.Sprintf("Action %q requires an 'ids' field", a.Name),
			}
		}

		ids, ok := toIntSlice(idsRaw)
		if !ok {
			return nil, backenderr.ActionError{
				Message: "ids field must be an array of integers",
			}
		}

		if len(ids) == 0 {
			return result, nil
		}

		// Validate all IDs exist and assign sequential weights.
		for i, id := range ids {
			// Verify the model exists in the datastore.
			_, err := params.Datastore.Get(a.Model.Collection, id, []string{"id"})
			if err != nil {
				return nil, backenderr.ActionError{
					Message: fmt.Sprintf("Model %s/%d does not exist", a.Model.Collection, id),
				}
			}

			if params.Datastore.IsDeleted(a.Model.Collection, id) {
				return nil, backenderr.ActionError{
					Message: fmt.Sprintf("Model %s/%d has been deleted", a.Model.Collection, id),
				}
			}

			// Assign weight = position + 1 (1-based).
			weight := i + 1
			params.Datastore.ApplyChangedModel(a.Model.Collection, id, map[string]any{
				weightField: weight,
			})
		}

		delete(result, "ids")
		return result, nil
	}
}
