package structure_level

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("structure_level.update", model.MustGet("structure_level"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":           map[string]any{"type": "integer"},
			"name":         map[string]any{"type": "string", "minLength": 1},
			"color":        map[string]any{"type": "string"},
			"default_time": map[string]any{"type": "number"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
