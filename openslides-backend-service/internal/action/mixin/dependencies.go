// Package mixin provides composable hook wrappers for actions.
//
// Instead of Python's multiple inheritance with MRO, Go uses hook-wrapping
// functions. Each mixin wraps an existing hook (e.g., UpdateInstance) to add
// behavior before or after the original hook.
package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
)

// Dependency defines a dependent action to execute after the main action.
type Dependency struct {
	ActionName string
	FieldMap   map[string]string // source field → target field
}

// WithDependencies wraps the UpdateInstance hook to execute dependent actions
// after the original hook completes.
//
// This replaces Python's CreateActionWithDependencies/CreateActionWithInternalDependencies.
func WithDependencies(a *action.BaseAction, deps ...Dependency) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		result, err := origUpdate(ctx, params, instance)
		if err != nil {
			return nil, err
		}

		for _, dep := range deps {
			depAction, err := action.Lookup(dep.ActionName)
			if err != nil {
				return nil, fmt.Errorf("dependency %s: %w", dep.ActionName, err)
			}

			// Build dependent instance from field map.
			depInstance := make(action.Instance)
			for srcField, dstField := range dep.FieldMap {
				if val, ok := result[srcField]; ok {
					depInstance[dstField] = val
				}
			}

			// Execute as sub-call, sharing datastore and relation manager.
			subParams := &action.ActionParams{
				UserID:          params.UserID,
				Internal:        true,
				IsSubCall:       true,
				Datastore:       params.Datastore,
				RelationManager: params.RelationManager,
			}

			_, _, err = action.Perform(ctx, depAction, subParams, []action.Instance{depInstance})
			if err != nil {
				return nil, fmt.Errorf("dependency %s: %w", dep.ActionName, err)
			}
		}

		return result, nil
	}
}
