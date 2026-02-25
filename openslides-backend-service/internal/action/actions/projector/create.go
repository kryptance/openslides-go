// Package projector implements projector actions for managing projectors in meetings.
//
// Projectors display projected content in a meeting. All projector actions
// require Projector.CAN_MANAGE permission.
package projector

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("projector.create", model.MustGet("projector"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id":               map[string]any{"type": "integer"},
			"name":                     map[string]any{"type": "string", "minLength": 1},
			"width":                    map[string]any{"type": "integer"},
			"aspect_ratio_numerator":   map[string]any{"type": "integer"},
			"aspect_ratio_denominator": map[string]any{"type": "integer"},
			"color":                    map[string]any{"type": "string"},
			"background_color":         map[string]any{"type": "string"},
			"header_background_color":  map[string]any{"type": "string"},
			"header_font_color":        map[string]any{"type": "string"},
			"header_h1_color":          map[string]any{"type": "string"},
			"chyron_background_color":  map[string]any{"type": "string"},
			"chyron_font_color":        map[string]any{"type": "string"},
			"show_header_footer":       map[string]any{"type": "boolean"},
			"show_title":               map[string]any{"type": "boolean"},
			"show_logo":                map[string]any{"type": "boolean"},
			"show_clock":               map[string]any{"type": "boolean"},
			"is_internal":              map[string]any{"type": "boolean"},
		},
		"required": []string{"meeting_id", "name"},
	}

	action.Register(a)
}
