package motion

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("motion.update_metadata", model.MustGet("motion"))
	a.Permission = perm.MotionCanManageMetadata

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":          map[string]any{"type": "integer"},
			"category_id": map[string]any{"type": "integer"},
			"block_id":    map[string]any{"type": "integer"},
			"supporter_meeting_user_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
