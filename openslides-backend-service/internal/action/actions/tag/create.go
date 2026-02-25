// Package tag implements tag CRUD actions.
//
// Tags are labels that can be attached to various items within a meeting.
// All tag actions require Tag.CAN_MANAGE permission.
package tag

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("tag.create", model.MustGet("tag"))
	a.Permission = perm.TagCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"name":       map[string]any{"type": "string", "minLength": 1},
		},
		"required": []string{"meeting_id", "name"},
	}

	action.Register(a)
}
