// Package projector_countdown implements projector_countdown CRUD actions.
//
// Projector countdowns are timers displayed on a projector within a meeting.
// All projector_countdown actions require Projector.CAN_MANAGE permission.
package projector_countdown

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("projector_countdown.create", model.MustGet("projector_countdown"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id":     map[string]any{"type": "integer"},
			"title":          map[string]any{"type": "string", "minLength": 1},
			"description":    map[string]any{"type": "string"},
			"default_time":   map[string]any{"type": "number"},
			"countdown_time": map[string]any{"type": "number"},
		},
		"required": []string{"meeting_id", "title"},
	}

	action.Register(a)
}
