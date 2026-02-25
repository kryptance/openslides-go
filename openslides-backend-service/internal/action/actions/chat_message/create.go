// Package chat_message implements chat_message CRUD actions.
//
// Chat messages are individual messages within a chat group.
// Permission checks are custom: users need write access to the chat group.
package chat_message

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("chat_message.create", model.MustGet("chat_message"))
	// Permission is nil; custom CheckPermissions verifies write access to the chat group.
	a.Permission = nil

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"chat_group_id": map[string]any{"type": "integer"},
			"content":       map[string]any{"type": "string", "minLength": 1},
		},
		"required": []string{"chat_group_id", "content"},
	}

	action.Register(a)
}
