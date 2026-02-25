package user

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("user.participant.update", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":         map[string]any{"type": "integer"},
			"meeting_id": map[string]any{"type": "integer"},
			"structure_level_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"group_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"vote_weight": map[string]any{"type": "string"},
			"comment":     map[string]any{"type": "string"},
			"about_me":    map[string]any{"type": "string"},
			"number":      map[string]any{"type": "string"},
		},
		"required": []string{"id", "meeting_id"},
	}

	// TODO: The update should target the meeting_user record, not the user record directly.
	// Need to look up the meeting_user by user id + meeting_id, then update meeting_user fields:
	//   structure_level_ids, group_ids, vote_weight, comment, about_me, number.

	action.Register(a)
}
