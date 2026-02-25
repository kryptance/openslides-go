// Package organization_tag implements organization_tag CRUD actions.
//
// Organization tags are labels that can be applied across the entire organization.
// All organization_tag actions require OML can_manage_organization permission.
package organization_tag

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("organization_tag.create", model.MustGet("organization_tag"))
	a.Permission = perm.OMLCanManageOrganization

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"name":  map[string]any{"type": "string", "minLength": 1},
			"color": map[string]any{"type": "string", "minLength": 1},
		},
		"required": []string{"name", "color"},
	}

	action.Register(a)
}
