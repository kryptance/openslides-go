package user

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewBaseAction("user.forget_password", model.MustGet("user"))
	// No permission check required -- this action is available without login.
	a.Permission = nil

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"email": map[string]any{"type": "string"},
		},
		"required": []string{"email"},
	}

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		email, ok := instance["email"].(string)
		if !ok || email == "" {
			return nil, fmt.Errorf("email field must be a non-empty string")
		}

		// TODO: Look up user(s) by email address in the datastore.
		// TODO: Generate a password-reset token with an expiration time.
		// TODO: Store the token in the datastore (e.g., on the user model or a separate store).
		// TODO: Send a password-reset email to the user containing the token link.
		// Note: Even if no user is found, return success to avoid email enumeration.

		// No events are generated; the side-effect is the email being sent.
		return nil, nil
	}

	action.Register(a)
}
