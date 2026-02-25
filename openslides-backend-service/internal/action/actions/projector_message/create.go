// Package projector_message implements projector_message CRUD actions.
//
// Projector messages are text messages displayed on a projector within a meeting.
// All projector_message actions require Projector.CAN_MANAGE permission.
package projector_message

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("projector_message.create", model.MustGet("projector_message"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"message":    map[string]any{"type": "string"},
		},
		"required": []string{"meeting_id"},
	}

	action.Register(a)
}
