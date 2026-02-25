// Package chat_group implements chat_group CRUD, sort, and clear actions.
//
// Chat groups are containers for chat messages within a meeting.
// All chat_group actions require Chat.CAN_MANAGE permission.
package chat_group

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("chat_group.create", model.MustGet("chat_group"))
	a.Permission = perm.ChatCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"name":       map[string]any{"type": "string", "minLength": 1},
			"read_group_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"write_group_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"meeting_id", "name"},
	}

	action.Register(a)
}
