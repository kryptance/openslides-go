package meeting_mediafile

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewUpdateAction("meeting_mediafile.update", model.MustGet("meeting_mediafile"))
	// Custom permission: depends on mediafile ownership and meeting access.
	// TODO: Implement custom CheckPermissions hook.

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":        map[string]any{"type": "integer"},
			"is_public": map[string]any{"type": "boolean"},
			"access_group_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"inherited_access_group_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
