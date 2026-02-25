package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
)

// WithForbidAnonymousGroup wraps the UpdateInstance hook to prevent the
// anonymous group from being included in group lists.
//
// In OpenSlides, each meeting has a special "anonymous" group
// (meeting.anonymous_group_id). Certain actions should reject payloads that
// include this group in group_ids or similar relation fields.
//
// groupField specifies the field name that contains the group IDs to validate
// (e.g., "group_ids").
//
// This replaces Python's ForbidAnonymousGroupMixin.
func WithForbidAnonymousGroup(a *action.BaseAction, groupField string) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		groupIDs, ok := instance[groupField]
		if !ok {
			return origUpdate(ctx, params, instance)
		}

		ids, ok := toIntSlice(groupIDs)
		if !ok {
			return origUpdate(ctx, params, instance)
		}

		if len(ids) == 0 {
			return origUpdate(ctx, params, instance)
		}

		// Determine the meeting_id to look up the anonymous group.
		meetingID, err := a.GetMeetingID(ctx, params, instance)
		if err != nil {
			return nil, fmt.Errorf("forbid anonymous group: get meeting_id: %w", err)
		}
		if meetingID == 0 {
			return origUpdate(ctx, params, instance)
		}

		// Fetch the meeting's anonymous_group_id.
		meeting, err := params.Datastore.Get("meeting", meetingID, []string{"anonymous_group_id"})
		if err != nil {
			return nil, fmt.Errorf("forbid anonymous group: fetch meeting %d: %w", meetingID, err)
		}

		anonGroupID, ok := meeting["anonymous_group_id"]
		if !ok || anonGroupID == nil {
			// No anonymous group configured, nothing to check.
			return origUpdate(ctx, params, instance)
		}

		var anonID int
		switch v := anonGroupID.(type) {
		case float64:
			anonID = int(v)
		case int:
			anonID = v
		default:
			return origUpdate(ctx, params, instance)
		}

		// Check if the anonymous group is in the provided group list.
		for _, id := range ids {
			if id == anonID {
				return nil, backenderr.ActionError{
					Message: fmt.Sprintf(
						"The anonymous group (id %d) is not allowed in the %s field.",
						anonID, groupField,
					),
				}
			}
		}

		return origUpdate(ctx, params, instance)
	}
}
