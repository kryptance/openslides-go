// Package gender implements gender CRUD actions.
//
// Genders are organization-level labels for user gender identity.
// All gender actions require OML can_manage_users permission.
package gender

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("gender.create", model.MustGet("gender"))
	a.Permission = perm.OMLCanManageUsers

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"name": map[string]any{"type": "string", "minLength": 1},
		},
		"required": []string{"name"},
	}

	action.Register(a)
}
