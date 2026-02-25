// Package motion_category implements motion category actions.
//
// Motion categories organize motions into a hierarchical structure.
// All actions require MotionCanManage permission.
package motion_category

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("motion_category.create", model.MustGet("motion_category"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"name":       map[string]any{"type": "string", "minLength": 1},
			"prefix":     map[string]any{"type": "string"},
		},
		"required": []string{"meeting_id", "name"},
	}

	action.Register(a)
}
