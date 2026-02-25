// Package committee implements committee CRUD and import actions.
//
// Committees are organizational units that contain meetings.
// All committee actions require OML superadmin permission.
package committee

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("committee.create", model.MustGet("committee"))
	a.Permission = perm.OMLSuperadmin

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"organization_id": map[string]any{"type": "integer"},
			"name":            map[string]any{"type": "string", "minLength": 1},
			"description":     map[string]any{"type": "string"},
			"manager_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"forward_to_committee_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"organization_tag_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"external_id": map[string]any{"type": "string"},
		},
		"required": []string{"organization_id", "name"},
	}

	action.Register(a)
}
