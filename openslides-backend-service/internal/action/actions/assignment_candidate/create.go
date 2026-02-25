// Package assignment_candidate implements assignment candidate actions.
//
// Assignment candidates link users to assignments (elections).
// Permission checks are custom (depends on nomination rights).
package assignment_candidate

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("assignment_candidate.create", model.MustGet("assignment_candidate"))
	// Custom permission: depends on assignment.can_nominate_self / can_nominate_other.
	// TODO: Implement custom CheckPermissions hook.

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"assignment_id":   map[string]any{"type": "integer"},
			"meeting_user_id": map[string]any{"type": "integer"},
		},
		"required": []string{"assignment_id", "meeting_user_id"},
	}

	action.Register(a)
}
