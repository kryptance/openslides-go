package motion_working_group_speaker

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewDeleteAction("motion_working_group_speaker.delete", model.MustGet("motion_working_group_speaker"))
	// Custom permission: depends on motion.can_manage_metadata.
	// TODO: Implement custom CheckPermissions hook.

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
