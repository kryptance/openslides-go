// Package motion_submitter implements motion submitter actions.
//
// Motion submitters link users to motions as the submitting party.
// Permission checks are custom (depends on motion state and user role).
package motion_submitter

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("motion_submitter.create", model.MustGet("motion_submitter"))
	// Custom permission: depends on motion.can_manage or self-submission.
	// TODO: Implement custom CheckPermissions hook.

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"motion_id":       map[string]any{"type": "integer"},
			"meeting_user_id": map[string]any{"type": "integer"},
		},
		"required": []string{"motion_id"},
	}

	action.Register(a)
}
