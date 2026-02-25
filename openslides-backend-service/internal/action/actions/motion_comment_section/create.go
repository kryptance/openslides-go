// Package motion_comment_section implements motion comment section actions.
//
// Comment sections define categories of comments with read/write group access control.
// All actions require MotionCanManage permission.
package motion_comment_section

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("motion_comment_section.create", model.MustGet("motion_comment_section"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"name":       map[string]any{"type": "string", "minLength": 1},
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
		"required": []string{"meeting_id", "name"},
	}

	action.Register(a)
}
