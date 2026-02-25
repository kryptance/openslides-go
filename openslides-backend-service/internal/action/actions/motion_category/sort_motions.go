package motion_category

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	// sort_motions sorts motions within a category by assigning weights.
	// Uses the motion model since we are sorting motions, but the action
	// is namespaced under motion_category.
	a := action.NewUpdateAction("motion_category.sort_motions_in_category", model.MustGet("motion_category"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
			"motion_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"id", "motion_ids"},
	}

	mixin.WithSingularAction(a)

	// TODO: Implement custom UpdateInstance to assign category_weight to
	// each motion in motion_ids sequentially (linear sort on motions).

	action.Register(a)
}
