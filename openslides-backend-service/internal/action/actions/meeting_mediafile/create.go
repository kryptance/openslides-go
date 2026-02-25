// Package meeting_mediafile implements meeting mediafile actions.
//
// Meeting mediafiles link organization-level mediafiles to specific meetings.
// Permission checks are custom (depends on mediafile ownership and meeting access).
package meeting_mediafile

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("meeting_mediafile.create", model.MustGet("meeting_mediafile"))
	// Custom permission: depends on mediafile ownership and meeting access.
	// TODO: Implement custom CheckPermissions hook.

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"mediafile_id": map[string]any{"type": "integer"},
			"meeting_id":   map[string]any{"type": "integer"},
			"is_public":    map[string]any{"type": "boolean"},
		},
		"required": []string{"mediafile_id", "meeting_id", "is_public"},
	}

	action.Register(a)
}
