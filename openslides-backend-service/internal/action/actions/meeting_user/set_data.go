package meeting_user

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("meeting_user.set_data", model.MustGet("meeting_user"))
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
		},
		"required": []string{"id"},
	}

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// TODO: Validate that the meeting_user exists and the user has permission
		// to modify the meeting-specific user data.
		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
