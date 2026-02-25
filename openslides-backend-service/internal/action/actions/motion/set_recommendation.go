package motion

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("motion.set_recommendation", model.MustGet("motion"))
	a.Permission = perm.MotionCanManageMetadata

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":                map[string]any{"type": "integer"},
			"recommendation_id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
