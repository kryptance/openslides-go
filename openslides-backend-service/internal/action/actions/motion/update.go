package motion

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("motion.update", model.MustGet("motion"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":                     map[string]any{"type": "integer"},
			"title":                  map[string]any{"type": "string", "minLength": 1},
			"text":                   map[string]any{"type": "string"},
			"reason":                 map[string]any{"type": "string"},
			"number":                 map[string]any{"type": "string"},
			"modified_final_version": map[string]any{"type": "string"},
			"category_id":            map[string]any{"type": "integer"},
			"block_id":               map[string]any{"type": "integer"},
			"supporter_meeting_user_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"workflow_id":              map[string]any{"type": "integer"},
			"recommendation_extension": map[string]any{"type": "string"},
			"attachment_mediafile_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"id"},
	}

	// Apply mixin for text hashing on update.
	mixin.WithTextHash(a, "text", "text_hash")

	action.Register(a)
}
