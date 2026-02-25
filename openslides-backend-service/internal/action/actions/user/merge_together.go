package user

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("user.merge_together", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
			"user_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"id", "user_ids"},
	}

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		id, ok := instance["id"]
		if !ok {
			return nil, fmt.Errorf("instance has no id field")
		}

		var idInt int
		switch v := id.(type) {
		case float64:
			idInt = int(v)
		case int:
			idInt = v
		default:
			return nil, fmt.Errorf("id has unexpected type %T", id)
		}

		userIDsRaw, ok := instance["user_ids"].([]any)
		if !ok {
			return nil, fmt.Errorf("user_ids must be an array")
		}

		_ = idInt
		_ = userIDsRaw

		// TODO: Merge the specified users into the target user (id).
		// For each user in user_ids:
		//   1. Transfer all meeting_user entries to the target user.
		//   2. Transfer committee memberships.
		//   3. Transfer votes and vote delegations.
		//   4. Transfer personal notes, chat messages, and other user-owned data.
		//   5. Resolve conflicts (e.g., both users in same meeting).
		//   6. Delete the source user after transferring all data.
		// TODO: Generate appropriate update/delete events.

		return nil, fmt.Errorf("user.merge_together: not yet implemented")
	}

	action.Register(a)
}
