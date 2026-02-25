package motion_submitter

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewUpdateAction("motion_submitter.update", model.MustGet("motion_submitter"))
	// Custom permission: depends on motion.can_manage.
	// TODO: Implement custom CheckPermissions hook.

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":              map[string]any{"type": "integer"},
			"meeting_user_id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
