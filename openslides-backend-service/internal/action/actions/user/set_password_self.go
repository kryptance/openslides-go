package user

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewUpdateAction("user.set_password_self", model.MustGet("user"))

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"old_password": map[string]any{"type": "string"},
			"new_password": map[string]any{"type": "string"},
		},
		"required": []string{"old_password", "new_password"},
	}

	// Self-only: the action operates on the requesting user.
	a.CheckPermissions = func(ctx context.Context, params *action.ActionParams, instance action.Instance) error {
		if params.UserID == 0 {
			return backenderr.AnonymousNotAllowed{}
		}
		// The user can only change their own password.
		return nil
	}

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		oldPassword, ok := instance["old_password"].(string)
		if !ok || oldPassword == "" {
			return nil, fmt.Errorf("old_password is required")
		}

		newPassword, ok := instance["new_password"].(string)
		if !ok || newPassword == "" {
			return nil, fmt.Errorf("new_password is required")
		}

		// Set the instance id to the current user.
		instance["id"] = params.UserID

		// TODO: Fetch the user's current password hash from the datastore.
		// TODO: Verify old_password matches the stored hash using bcrypt.CompareHashAndPassword.
		// TODO: Hash new_password using bcrypt.GenerateFromPassword.
		// TODO: Set instance["password"] = hashed new password.

		// Remove action-specific fields, keep only model fields.
		delete(instance, "old_password")
		delete(instance, "new_password")
		instance["password"] = newPassword // placeholder until bcrypt is integrated

		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
