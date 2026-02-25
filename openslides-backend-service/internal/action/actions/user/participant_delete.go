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
	a := action.NewBaseAction("user.participant.delete", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":         map[string]any{"type": "integer"},
			"meeting_id": map[string]any{"type": "integer"},
		},
		"required": []string{"id", "meeting_id"},
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

		meetingID, ok := instance["meeting_id"]
		if !ok {
			return nil, fmt.Errorf("instance has no meeting_id field")
		}

		_ = idInt
		_ = meetingID

		// TODO: Look up the meeting_user for user id + meeting_id.
		// TODO: Delete the meeting_user entry (not the user itself).
		// TODO: Remove the user from is_present_in_meeting_ids for this meeting.
		// TODO: Remove meeting-specific data (structure_level assignments, group memberships, etc.).
		// TODO: Handle cascading cleanup of personal_notes, chat_messages, etc. in this meeting.

		return nil, fmt.Errorf("user.participant.delete: not yet implemented")
	}

	action.Register(a)
}
