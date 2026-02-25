// Package motion_supporter implements motion supporter actions.
//
// Motion supporters allow users to express support for a motion.
// Permission checks are custom (depends on motion.can_support and state).
package motion_supporter

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("motion_supporter.create", model.MustGet("motion_supporter"))
	// Custom permission: depends on motion.can_support and current state.
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
