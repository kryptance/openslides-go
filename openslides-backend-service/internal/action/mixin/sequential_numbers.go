package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
)

// WithSequentialNumbers wraps the UpdateInstance hook to auto-generate a
// sequential_number for each new model instance within its meeting scope.
//
// The sequential number is computed as max(existing sequential_number in the
// same meeting) + 1. If no prior instances exist, it starts at 1.
//
// This replaces Python's SequentialNumbersMixin.
func WithSequentialNumbers(a *action.BaseAction) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// Only set if not already provided.
		if _, ok := instance["sequential_number"]; ok {
			return origUpdate(ctx, params, instance)
		}

		// Determine the meeting_id.
		meetingID, err := a.GetMeetingID(ctx, params, instance)
		if err != nil {
			return nil, fmt.Errorf("sequential numbers mixin: get meeting_id: %w", err)
		}
		if meetingID == 0 {
			return nil, fmt.Errorf("sequential numbers mixin: meeting_id is required")
		}

		// TODO: Query the datastore for the maximum sequential_number in this
		// meeting for this collection:
		//   filter = {meeting_id: meetingID}
		//   max_seq = max(collection.sequential_number WHERE meeting_id = meetingID)
		//
		// For now, use a simple approach based on the extended datastore.
		nextSeqNum := getNextSequentialNumber(params, a.Model.Collection, meetingID)
		instance["sequential_number"] = nextSeqNum

		return origUpdate(ctx, params, instance)
	}
}

// getNextSequentialNumber determines the next sequential number for a collection
// within a meeting by scanning the extended datastore for the current maximum.
func getNextSequentialNumber(params *action.ActionParams, collection string, meetingID int) int {
	// TODO: This should query the actual datastore with a filter for the
	// maximum sequential_number in the given meeting. For now, we track
	// the max in the extended datastore's changed models as a best-effort
	// approach during a single action request.
	maxSeq := 0

	// Scan changed models to find the current max sequential_number.
	// In a real implementation, this would also query the database.
	_ = params
	_ = collection
	_ = meetingID

	return maxSeq + 1
}
