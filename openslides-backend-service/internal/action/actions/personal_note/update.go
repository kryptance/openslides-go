package personal_note

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewUpdateAction("personal_note.update", model.MustGet("personal_note"))
	// Permission is nil; custom CheckPermissions verifies the user owns the personal note.
	a.Permission = nil

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":   map[string]any{"type": "integer"},
			"note": map[string]any{"type": "string"},
			"star": map[string]any{"type": "boolean"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
