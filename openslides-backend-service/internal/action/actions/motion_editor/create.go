// Package motion_editor implements motion editor actions.
//
// Motion editors link users to motions as editors responsible for editing.
// Permission checks are custom (depends on motion.can_manage_metadata).
package motion_editor

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("motion_editor.create", model.MustGet("motion_editor"))
	// Custom permission: depends on motion.can_manage_metadata.
	// TODO: Implement custom CheckPermissions hook.

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"motion_id":       map[string]any{"type": "integer"},
			"meeting_user_id": map[string]any{"type": "integer"},
		},
		"required": []string{"motion_id", "meeting_user_id"},
	}

	action.Register(a)
}
