package user

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewBaseAction("user.forget_password_confirm", model.MustGet("user"))
	// No permission check required -- token-based authentication.
	a.Permission = nil

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"token":        map[string]any{"type": "string"},
			"user_id":      map[string]any{"type": "integer"},
			"new_password": map[string]any{"type": "string"},
		},
		"required": []string{"token", "user_id", "new_password"},
	}

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		token, ok := instance["token"].(string)
		if !ok || token == "" {
			return nil, fmt.Errorf("token field must be a non-empty string")
		}

		userID, ok := instance["user_id"]
		if !ok {
			return nil, fmt.Errorf("user_id is required")
		}

		var userIDInt int
		switch v := userID.(type) {
		case float64:
			userIDInt = int(v)
		case int:
			userIDInt = v
		default:
			return nil, fmt.Errorf("user_id has unexpected type %T", userID)
		}

		newPassword, ok := instance["new_password"].(string)
		if !ok || newPassword == "" {
			return nil, fmt.Errorf("new_password must be a non-empty string")
		}

		// TODO: Verify the token is valid and not expired for the given user_id.
		// TODO: Hash newPassword using bcrypt.
		// TODO: Invalidate the token after use.

		fqid := event.FQID("user", userIDInt)
		e := event.Event{
			Type: event.TypeUpdate,
			FQID: fqid,
			Fields: map[string]any{
				"password": newPassword, // TODO: Use hashed password.
			},
		}

		return []event.Event{e}, nil
	}

	action.Register(a)
}
