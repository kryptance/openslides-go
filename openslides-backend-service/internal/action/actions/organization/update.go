// Package organization implements organization-level actions.
//
// The organization is the top-level entity in OpenSlides. There is typically
// only one organization instance. All organization actions require OML
// superadmin permission.
package organization

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("organization.update", model.MustGet("organization"))
	a.Permission = perm.OMLSuperadmin

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":                           map[string]any{"type": "integer"},
			"name":                         map[string]any{"type": "string"},
			"description":                  map[string]any{"type": "string"},
			"legal_notice":                 map[string]any{"type": "string"},
			"privacy_policy":               map[string]any{"type": "string"},
			"login_text":                   map[string]any{"type": "string"},
			"theme_id":                     map[string]any{"type": "integer"},
			"enable_electronic_voting":     map[string]any{"type": "boolean"},
			"reset_password_verbose_errors": map[string]any{"type": "boolean"},
			"limit_of_meetings":            map[string]any{"type": "integer"},
			"limit_of_users":               map[string]any{"type": "integer"},
			"url":                          map[string]any{"type": "string"},
			"saml_enabled":                 map[string]any{"type": "boolean"},
			"saml_login_button_text":       map[string]any{"type": "string"},
			"saml_attr_mapping":            map[string]any{"type": "object"},
			"genders": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "string"},
			},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
