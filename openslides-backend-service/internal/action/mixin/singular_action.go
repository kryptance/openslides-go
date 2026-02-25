package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
)

// WithSingularAction wraps the Prefetch hook to enforce that the action is
// called with exactly one data instance. If the data slice has any other length,
// an ActionError is returned.
//
// This is used for actions that are inherently singular, such as sort actions
// or bulk operations that accept a single payload describing the entire operation.
//
// This replaces Python's SingularActionMixin.
func WithSingularAction(a *action.BaseAction) {
	origPrefetch := a.Prefetch
	a.Prefetch = func(ctx context.Context, params *action.ActionParams, data []action.Instance) error {
		if len(data) != 1 {
			return backenderr.ActionError{
				Message: fmt.Sprintf(
					"Action %q must be called with exactly 1 data instance, got %d",
					a.Name, len(data),
				),
			}
		}
		return origPrefetch(ctx, params, data)
	}
}
