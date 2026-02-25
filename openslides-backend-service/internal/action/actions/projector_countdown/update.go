package projector_countdown

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("projector_countdown.update", model.MustGet("projector_countdown"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":             map[string]any{"type": "integer"},
			"title":          map[string]any{"type": "string", "minLength": 1},
			"description":    map[string]any{"type": "string"},
			"default_time":   map[string]any{"type": "number"},
			"countdown_time": map[string]any{"type": "number"},
			"running":        map[string]any{"type": "boolean"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
