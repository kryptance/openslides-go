package user

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("user.participant.create", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id":  map[string]any{"type": "integer"},
			"first_name":  map[string]any{"type": "string"},
			"last_name":   map[string]any{"type": "string"},
			"username":    map[string]any{"type": "string"},
			"structure_level_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"group_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"vote_weight": map[string]any{"type": "string"},
			"is_present_in_meeting_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"meeting_id"},
	}

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// TODO: Create the user account if it does not exist (or find by username).
		// TODO: Create a meeting_user entry linking user to the meeting.
		// TODO: Assign the user to the specified groups (or default group if none given).
		// TODO: Set structure_level_ids on the meeting_user.
		// TODO: Set vote_weight on the meeting_user.
		// TODO: Handle is_present_in_meeting_ids.
		// TODO: Generate username if not provided.
		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
