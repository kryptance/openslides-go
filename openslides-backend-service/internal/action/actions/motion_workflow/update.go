package motion_workflow

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("motion_workflow.update", model.MustGet("motion_workflow"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":   map[string]any{"type": "integer"},
			"name": map[string]any{"type": "string", "minLength": 1},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
