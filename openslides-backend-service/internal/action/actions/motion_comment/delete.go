package motion_comment

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewDeleteAction("motion_comment.delete", model.MustGet("motion_comment"))
	// Custom permission: checks section write groups.
	// TODO: Implement custom CheckPermissions hook that verifies
	// the user is in the write_group_ids of the referenced section.

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
