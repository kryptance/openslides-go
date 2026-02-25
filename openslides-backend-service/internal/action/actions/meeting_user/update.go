package meeting_user

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("meeting_user.update", model.MustGet("meeting_user"))
	a.Permission = perm.UserCanManage
	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":                  map[string]any{"type": "integer"},
			"comment":             map[string]any{"type": "string"},
			"number":              map[string]any{"type": "string"},
			"about_me":            map[string]any{"type": "string"},
			"vote_weight":         map[string]any{"type": "string"},
			"group_ids":           map[string]any{"type": "array", "items": map[string]any{"type": "integer"}},
			"structure_level_ids": map[string]any{"type": "array", "items": map[string]any{"type": "integer"}},
			"locked_out":          map[string]any{"type": "boolean"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
