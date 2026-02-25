package chat_message

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewUpdateAction("chat_message.update", model.MustGet("chat_message"))
	// Permission is nil; custom CheckPermissions verifies the user owns the message.
	a.Permission = nil

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":      map[string]any{"type": "integer"},
			"content": map[string]any{"type": "string", "minLength": 1},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
