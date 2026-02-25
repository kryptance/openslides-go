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
	a := action.NewBaseAction("user.toggle_presence_by_number", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"number":     map[string]any{"type": "string"},
		},
		"required": []string{"meeting_id", "number"},
	}

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		meetingID, ok := instance["meeting_id"]
		if !ok {
			return nil, fmt.Errorf("instance has no meeting_id field")
		}

		number, ok := instance["number"].(string)
		if !ok || number == "" {
			return nil, fmt.Errorf("number field must be a non-empty string")
		}

		// TODO: Look up the meeting_user by number within the given meeting.
		// This requires a filter query on the datastore:
		//   meeting_user where meeting_id == meetingID and number == number
		// Then resolve the user_id from the meeting_user.
		// Then check current is_present_in_meeting_ids to determine toggle direction.
		// Then emit an update event with ListFields add/remove on is_present_in_meeting_ids.

		_ = meetingID
		_ = number

		return nil, fmt.Errorf("toggle_presence_by_number: not yet implemented")
	}

	action.Register(a)
}
