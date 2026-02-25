package motion

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("motion.sort", model.MustGet("motion"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"tree": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"id": map[string]any{"type": "integer"},
						"children": map[string]any{
							"type":  "array",
							"items": map[string]any{"type": "object"},
						},
					},
					"required": []string{"id"},
				},
			},
		},
		"required": []string{"meeting_id", "tree"},
	}

	mixin.WithSingularAction(a)
	mixin.WithTreeSort(a, "weight", "sort_parent_id")

	action.Register(a)
}
