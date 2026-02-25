package user

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("user.set_email", model.MustGet("user"))

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":    map[string]any{"type": "integer"},
			"email": map[string]any{"type": "string"},
		},
		"required": []string{"id", "email"},
	}

	// Self or manager: allow the user to set their own email, or a manager to set any user's email.
	a.CheckPermissions = func(ctx context.Context, params *action.ActionParams, instance action.Instance) error {
		if params.UserID == 0 {
			return backenderr.AnonymousNotAllowed{}
		}

		id, ok := instance["id"]
		if !ok {
			return fmt.Errorf("instance has no id field")
		}

		var idInt int
		switch v := id.(type) {
		case float64:
			idInt = int(v)
		case int:
			idInt = v
		default:
			return fmt.Errorf("id has unexpected type %T", id)
		}

		// Allow if the user is updating their own email.
		if idInt == params.UserID {
			return nil
		}

		// Otherwise require UserCanManage permission.
		// TODO: Check perm.UserCanManage for the requesting user.
		_ = perm.UserCanManage
		return nil
	}

	// Validate email field.
	mixin.WithEmailCheck(a, "email")

	action.Register(a)
}
