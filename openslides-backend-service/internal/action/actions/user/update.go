package user

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("user.update", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":         map[string]any{"type": "integer"},
			"username":   map[string]any{"type": "string"},
			"first_name": map[string]any{"type": "string"},
			"last_name":  map[string]any{"type": "string"},
			"email":      map[string]any{"type": "string"},
			"title":      map[string]any{"type": "string"},
			"pronoun":    map[string]any{"type": "string"},
			"gender_id":  map[string]any{"type": "integer"},
			"default_password": map[string]any{"type": "string"},
			"is_active":  map[string]any{"type": "boolean"},
			"organization_management_level": map[string]any{"type": "string"},
		},
		"required": []string{"id"},
	}

	// Validate email field when provided.
	mixin.WithEmailCheck(a, "email")

	// TODO: Ensure username uniqueness if changed.
	// TODO: Validate organization_management_level against allowed values.
	// TODO: Hash default_password if provided.

	action.Register(a)
}
