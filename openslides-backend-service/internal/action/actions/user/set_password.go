package user

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("user.set_password", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":             map[string]any{"type": "integer"},
			"password":       map[string]any{"type": "string", "minLength": 1},
			"set_as_default": map[string]any{"type": "boolean"},
		},
		"required": []string{"id", "password"},
	}

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		password, ok := instance["password"].(string)
		if !ok || password == "" {
			return nil, fmt.Errorf("password is required and must be non-empty")
		}

		// TODO: Hash the password using bcrypt.
		// hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		// instance["password"] = string(hashed)
		instance["password"] = password // placeholder until bcrypt is integrated

		// If set_as_default is true, also update the default_password field.
		if setDefault, ok := instance["set_as_default"].(bool); ok && setDefault {
			instance["default_password"] = password
		}
		delete(instance, "set_as_default")

		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
