package projector

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("projector.update", model.MustGet("projector"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":                       map[string]any{"type": "integer"},
			"name":                     map[string]any{"type": "string"},
			"width":                    map[string]any{"type": "integer"},
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
			"scroll":                   map[string]any{"type": "integer"},
			"current_projection_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
