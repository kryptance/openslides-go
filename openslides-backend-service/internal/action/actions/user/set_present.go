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
	a := action.NewUpdateAction("user.set_present", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":         map[string]any{"type": "integer"},
			"meeting_id": map[string]any{"type": "integer"},
			"present":    map[string]any{"type": "boolean"},
		},
		"required": []string{"id", "meeting_id", "present"},
	}

	// Override CreateEvents to add/remove the meeting_id from is_present_in_meeting_ids.
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

		present, ok := instance["present"].(bool)
		if !ok {
			return nil, fmt.Errorf("present field must be a boolean")
		}

		fqid := event.FQID("user", idInt)
		e := event.Event{
			Type: event.TypeUpdate,
			FQID: fqid,
			ListFields: &event.ListFields{
				Add:    make(map[string][]any),
				Remove: make(map[string][]any),
			},
		}

		if present {
			e.ListFields.Add["is_present_in_meeting_ids"] = []any{meetingID}
		} else {
			e.ListFields.Remove["is_present_in_meeting_ids"] = []any{meetingID}
		}

		return []event.Event{e}, nil
	}

	action.Register(a)
}
