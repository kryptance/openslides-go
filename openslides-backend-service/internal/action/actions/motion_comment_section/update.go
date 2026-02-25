package motion_comment_section

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("motion_comment_section.update", model.MustGet("motion_comment_section"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":   map[string]any{"type": "integer"},
			"name": map[string]any{"type": "string", "minLength": 1},
			"read_group_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"write_group_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"submitter_can_write": map[string]any{"type": "boolean"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
