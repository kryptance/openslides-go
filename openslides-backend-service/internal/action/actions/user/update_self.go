package user

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewUpdateAction("user.update_self", model.MustGet("user"))

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"username":  map[string]any{"type": "string"},
			"pronoun":   map[string]any{"type": "string"},
			"email":     map[string]any{"type": "string"},
			"gender_id": map[string]any{"type": "integer"},
		},
	}

	// Self-only: users can only update their own profile.
	a.CheckPermissions = func(ctx context.Context, params *action.ActionParams, instance action.Instance) error {
		if params.UserID == 0 {
			return backenderr.AnonymousNotAllowed{}
		}
		return nil
	}

	// Validate email field when provided.
	mixin.WithEmailCheck(a, "email")

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// Set the id to the requesting user.
		instance["id"] = params.UserID

		// TODO: Ensure username uniqueness if changed.
		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
