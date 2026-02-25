package motion_block

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("motion_block.update", model.MustGet("motion_block"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":       map[string]any{"type": "integer"},
			"title":    map[string]any{"type": "string", "minLength": 1},
			"internal": map[string]any{"type": "boolean"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
