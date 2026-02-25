package motion_supporter

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewUpdateAction("motion_supporter.update", model.MustGet("motion_supporter"))
	// Custom permission: depends on motion.can_support or motion.can_manage.
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
