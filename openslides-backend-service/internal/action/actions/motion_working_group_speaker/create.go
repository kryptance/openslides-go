// Package motion_working_group_speaker implements working group speaker actions.
//
// Working group speakers are users assigned to speak on behalf of a working group for a motion.
// Permission checks are custom (depends on motion.can_manage_metadata).
package motion_working_group_speaker

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("motion_working_group_speaker.create", model.MustGet("motion_working_group_speaker"))
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
