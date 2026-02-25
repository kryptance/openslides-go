// Package motion_comment implements motion comment actions.
//
// Motion comments allow annotating motions within comment sections.
// Permission checks are custom (depends on section read/write groups).
package motion_comment

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("motion_comment.create", model.MustGet("motion_comment"))
	// Custom permission: checks section write groups.
	// TODO: Implement custom CheckPermissions hook that verifies
	// the user is in the write_group_ids of the referenced section.

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"motion_id":  map[string]any{"type": "integer"},
			"section_id": map[string]any{"type": "integer"},
			"comment":    map[string]any{"type": "string"},
		},
		"required": []string{"motion_id", "section_id"},
	}

	action.Register(a)
}
