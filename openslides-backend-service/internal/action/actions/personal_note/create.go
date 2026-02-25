// Package personal_note implements personal_note CRUD actions.
//
// Personal notes are private annotations attached to content objects by individual users.
// Permission checks are custom: only the owning user may create/update/delete their notes.
package personal_note

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("personal_note.create", model.MustGet("personal_note"))
	// Permission is nil; custom CheckPermissions verifies the user owns the meeting_user.
	a.Permission = nil

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_user_id":   map[string]any{"type": "integer"},
			"content_object_id": map[string]any{"type": "string"},
			"note":              map[string]any{"type": "string"},
			"star":              map[string]any{"type": "boolean"},
		},
		"required": []string{"meeting_user_id", "content_object_id"},
	}

	action.Register(a)
}
