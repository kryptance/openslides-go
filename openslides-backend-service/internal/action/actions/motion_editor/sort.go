package motion_editor

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("motion_editor.sort", model.MustGet("motion_editor"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"motion_id": map[string]any{"type": "integer"},
			"ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"motion_id", "ids"},
	}

	mixin.WithSingularAction(a)
	mixin.WithLinearSort(a, "weight")

	action.Register(a)
}
