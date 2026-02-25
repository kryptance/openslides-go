package user

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("user.reset_password_to_default", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		id, ok := instance["id"]
		if !ok {
			return nil, fmt.Errorf("instance has no id field")
		}

		var idInt int
		switch v := id.(type) {
		case float64:
			idInt = int(v)
		case int:
			idInt = v
		default:
			return nil, fmt.Errorf("id has unexpected type %T", id)
		}

		// Fetch the user's default_password from the datastore.
		userData, found := params.Datastore.GetChangedModel("user", idInt)
		if !found {
			return nil, fmt.Errorf("user/%d not found", idInt)
		}

		defaultPassword, ok := userData["default_password"].(string)
		if !ok || defaultPassword == "" {
			return nil, fmt.Errorf("user/%d has no default_password set", idInt)
		}

		// TODO: Hash the default_password using bcrypt before setting it as password.
		instance["password"] = defaultPassword // placeholder until bcrypt is integrated

		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
