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
	a := action.NewBaseAction("meeting.import", model.MustGet("meeting"))
	a.Permission = perm.MeetingCanManageSettings

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"committee_id": map[string]any{"type": "integer"},
			"meeting":      map[string]any{"type": "object"},
		},
		"required": []string{"committee_id", "meeting"},
	}

	// Import is a singular action: one meeting import at a time.
	mixin.WithSingularAction(a)

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		committeeID, ok := instance["committee_id"]
		if !ok {
			return nil, fmt.Errorf("instance has no committee_id field")
		}

		meetingData, ok := instance["meeting"]
		if !ok {
			return nil, fmt.Errorf("instance has no meeting field")
		}

		// TODO: Implement meeting import from JSON logic:
		// 1. Parse the meeting JSON export object.
		// 2. Validate all collections and fields in the export data.
		// 3. Reserve new IDs for the meeting and all contained models.
		// 4. Remap all internal references to use the new IDs.
		// 5. Set the committee_id on the new meeting.
		// 6. Generate create events for the meeting and all sub-models.
		_ = committeeID
		_ = meetingData

		return nil, nil
	}

	action.Register(a)
}
