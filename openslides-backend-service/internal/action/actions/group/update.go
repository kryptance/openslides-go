package group

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("group.update", model.MustGet("group"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":   map[string]any{"type": "integer"},
			"name": map[string]any{"type": "string", "minLength": 1},
			"permissions": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "string"},
			},
			"external_id": map[string]any{"type": "string"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
