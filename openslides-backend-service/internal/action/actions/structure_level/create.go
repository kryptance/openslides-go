// Package structure_level implements structure_level CRUD actions.
//
// Structure levels represent organizational hierarchy levels within a meeting.
// All structure_level actions require User.CAN_MANAGE permission.
package structure_level

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("structure_level.create", model.MustGet("structure_level"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id":   map[string]any{"type": "integer"},
			"name":         map[string]any{"type": "string", "minLength": 1},
			"color":        map[string]any{"type": "string"},
			"default_time": map[string]any{"type": "number"},
		},
		"required": []string{"meeting_id", "name"},
	}

	action.Register(a)
}
