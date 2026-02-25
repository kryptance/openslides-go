// Package group implements group CRUD actions.
//
// Groups define sets of permissions that can be assigned to meeting users.
// All actions require UserCanManage permission.
package group

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("group.create", model.MustGet("group"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"name":       map[string]any{"type": "string", "minLength": 1},
			"permissions": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "string"},
			},
			"external_id": map[string]any{"type": "string"},
		},
		"required": []string{"meeting_id", "name"},
	}

	action.Register(a)
}
