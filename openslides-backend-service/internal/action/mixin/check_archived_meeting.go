package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
)

// WithCheckArchivedMeeting wraps the UpdateInstance hook to prevent modifications
// to entities belonging to archived meetings.
//
// Before calling the original hook, it resolves the meeting_id from the instance
// and checks whether the meeting's is_archived flag is set. If the meeting is
// archived, an ActionError is returned.
//
// This replaces Python's check_for_archived_meeting logic.
func WithCheckArchivedMeeting(a *action.BaseAction) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		meetingID, err := a.GetMeetingID(ctx, params, instance)
		if err != nil {
			return nil, fmt.Errorf("check archived meeting: get meeting_id: %w", err)
		}

		if meetingID == 0 {
			// No meeting context, skip check.
			return origUpdate(ctx, params, instance)
		}

		// Look up the meeting's is_archived flag.
		meeting, err := params.Datastore.Get("meeting", meetingID, []string{"is_archived"})
		if err != nil {
			return nil, fmt.Errorf("check archived meeting: fetch meeting %d: %w", meetingID, err)
		}

		if isArchived, ok := meeting["is_archived"]; ok {
			if archived, ok := isArchived.(bool); ok && archived {
				return nil, backenderr.ActionError{
					Message: fmt.Sprintf(
						"Meeting %d is archived. You cannot modify data in an archived meeting.",
						meetingID,
					),
				}
			}
		}

		return origUpdate(ctx, params, instance)
	}
}
