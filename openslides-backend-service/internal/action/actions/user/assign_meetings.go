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
	a := action.NewBaseAction("user.assign_meetings", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
			"meeting_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"group_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"id", "meeting_ids"},
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

		meetingIDsRaw, ok := instance["meeting_ids"].([]any)
		if !ok {
			return nil, fmt.Errorf("meeting_ids must be an array")
		}

		_ = idInt
		_ = meetingIDsRaw

		// TODO: For each meeting_id in meeting_ids:
		//   1. Check if a meeting_user already exists for this user+meeting combination.
		//   2. If not, create a new meeting_user linking user to meeting.
		//   3. If group_ids are provided, assign the user to those groups in the meeting.
		//   4. If no group_ids, assign to the meeting's default group.
		// TODO: Generate create events for new meeting_user models.
		// TODO: Generate update events for group membership changes.

		return nil, fmt.Errorf("user.assign_meetings: not yet implemented")
	}

	action.Register(a)
}
