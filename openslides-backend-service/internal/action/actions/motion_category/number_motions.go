package motion_category

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("motion_category.number_motions", model.MustGet("motion_category"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	// Number motions is a singular action: exactly one category at a time.
	mixin.WithSingularAction(a)

	// TODO: Implement custom UpdateInstance to assign sequential numbers
	// to all motions in this category based on category prefix and weight.

	action.Register(a)
}
