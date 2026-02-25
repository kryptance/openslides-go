package meeting

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("meeting.clone", model.MustGet("meeting"))
	a.Permission = perm.MeetingCanManageSettings

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id":      map[string]any{"type": "integer"},
			"committee_id":    map[string]any{"type": "integer"},
			"name":            map[string]any{"type": "string"},
			"set_as_template": map[string]any{"type": "boolean"},
			"start_time":      map[string]any{"type": "integer"},
			"end_time":        map[string]any{"type": "integer"},
		},
		"required": []string{"meeting_id"},
	}

	// Clone is a singular action: one meeting clone at a time.
	mixin.WithSingularAction(a)

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		meetingID, ok := instance["meeting_id"]
		if !ok {
			return nil, fmt.Errorf("instance has no meeting_id field")
		}

		// TODO: Implement deep-clone logic:
		// 1. Read the source meeting and all related models (agenda items,
		//    motions, assignments, projectors, groups, etc.).
		// 2. Reserve new IDs for the cloned meeting and all cloned sub-models.
		// 3. Remap all internal references to use the new IDs.
		// 4. Optionally override committee_id, name, set_as_template,
		//    start_time, and end_time from the instance data.
		// 5. Generate create events for the new meeting and all sub-models.
		_ = meetingID

		return nil, nil
	}

	action.Register(a)
}
