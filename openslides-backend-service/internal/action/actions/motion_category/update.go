package motion_category

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("motion_category.update", model.MustGet("motion_category"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":     map[string]any{"type": "integer"},
			"name":   map[string]any{"type": "string", "minLength": 1},
			"prefix": map[string]any{"type": "string"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
