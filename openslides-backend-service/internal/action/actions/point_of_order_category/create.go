// Package point_of_order_category implements point_of_order_category CRUD actions.
//
// Point of order categories define types of points of order in a meeting.
// All point_of_order_category actions require Meeting.CAN_MANAGE_SETTINGS permission.
package point_of_order_category

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("point_of_order_category.create", model.MustGet("point_of_order_category"))
	a.Permission = perm.MeetingCanManageSettings

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"text":       map[string]any{"type": "string", "minLength": 1},
			"rank":       map[string]any{"type": "integer"},
		},
		"required": []string{"meeting_id", "text", "rank"},
	}

	action.Register(a)
}
