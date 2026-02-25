// Package motion_change_recommendation implements change recommendation actions.
//
// Change recommendations propose modifications to specific line ranges of a motion.
// All actions require MotionCanManage permission.
package motion_change_recommendation

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("motion_change_recommendation.create", model.MustGet("motion_change_recommendation"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"motion_id":         map[string]any{"type": "integer"},
			"line_from":         map[string]any{"type": "integer", "minimum": 0},
			"line_to":           map[string]any{"type": "integer", "minimum": 0},
			"text":              map[string]any{"type": "string"},
			"rejected":          map[string]any{"type": "boolean"},
			"internal":          map[string]any{"type": "boolean"},
			"type":              map[string]any{"type": "integer"},
			"other_description": map[string]any{"type": "string"},
		},
		"required": []string{"motion_id", "line_from", "line_to", "text"},
	}

	action.Register(a)
}
