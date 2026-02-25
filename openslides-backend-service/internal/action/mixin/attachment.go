package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
)

// WithAttachment wraps the UpdateInstance hook to convert attachment_mediafile_ids
// to attachment_meeting_mediafile_ids by finding or creating meeting_mediafile records.
//
// In OpenSlides, mediafiles are global objects but can be attached to meeting-scoped
// models through meeting_mediafile junction records. This mixin ensures that when
// a user provides attachment_mediafile_ids (the simpler API), the correct
// meeting_mediafile entries are resolved or created automatically.
//
// This replaces Python's AttachmentMixin.
func WithAttachment(a *action.BaseAction) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// Check if attachment_mediafile_ids is provided.
		mediafileIDs, ok := instance["attachment_mediafile_ids"]
		if !ok {
			return origUpdate(ctx, params, instance)
		}

		// Determine the meeting_id for this instance.
		meetingID, err := a.GetMeetingID(ctx, params, instance)
		if err != nil {
			return nil, fmt.Errorf("attachment mixin: get meeting_id: %w", err)
		}
		if meetingID == 0 {
			return nil, fmt.Errorf("attachment mixin: meeting_id is required when using attachment_mediafile_ids")
		}

		// Convert mediafile IDs to meeting_mediafile IDs.
		ids, ok := toIntSlice(mediafileIDs)
		if !ok {
			return nil, fmt.Errorf("attachment mixin: attachment_mediafile_ids must be an array of integers")
		}

		meetingMediafileIDs := make([]int, 0, len(ids))
		for _, mediafileID := range ids {
			// Look up existing meeting_mediafile for this mediafile+meeting combination.
			mmfID, err := findOrCreateMeetingMediafile(ctx, params, meetingID, mediafileID)
			if err != nil {
				return nil, fmt.Errorf("attachment mixin: resolve mediafile %d: %w", mediafileID, err)
			}
			meetingMediafileIDs = append(meetingMediafileIDs, mmfID)
		}

		// Replace the field with the resolved meeting_mediafile IDs.
		delete(instance, "attachment_mediafile_ids")
		instance["attachment_meeting_mediafile_ids"] = meetingMediafileIDs

		return origUpdate(ctx, params, instance)
	}
}

// findOrCreateMeetingMediafile finds an existing meeting_mediafile record for the
// given meeting and mediafile, or creates one via a sub-action.
func findOrCreateMeetingMediafile(ctx context.Context, params *action.ActionParams, meetingID, mediafileID int) (int, error) {
	// TODO: Query datastore for existing meeting_mediafile with matching
	// meeting_id and mediafile_id. If found, return its ID.
	// If not found, execute meeting_mediafile.create as a sub-action.

	createAction, err := action.Lookup("meeting_mediafile.create")
	if err != nil {
		// If the action is not yet registered, return a placeholder.
		// This allows the mixin to be used before all actions are implemented.
		return 0, fmt.Errorf("meeting_mediafile.create action not available: %w", err)
	}

	depInstance := action.Instance{
		"meeting_id":   meetingID,
		"mediafile_id": mediafileID,
	}

	subParams := &action.ActionParams{
		UserID:    params.UserID,
		Internal:  true,
		IsSubCall: true,
		Datastore: params.Datastore,
	}

	_, results, err := action.Perform(ctx, createAction, subParams, []action.Instance{depInstance})
	if err != nil {
		return 0, err
	}

	if len(results) == 0 {
		return 0, fmt.Errorf("meeting_mediafile.create returned no results")
	}

	id, ok := results[0]["id"]
	if !ok {
		return 0, fmt.Errorf("meeting_mediafile.create returned no id")
	}

	switch v := id.(type) {
	case float64:
		return int(v), nil
	case int:
		return v, nil
	default:
		return 0, fmt.Errorf("meeting_mediafile.create returned unexpected id type %T", id)
	}
}

// toIntSlice converts an any value to a slice of ints.
// Accepts []any with float64 or int elements, or []int directly.
func toIntSlice(val any) ([]int, bool) {
	switch v := val.(type) {
	case []int:
		return v, true
	case []any:
		result := make([]int, 0, len(v))
		for _, elem := range v {
			switch e := elem.(type) {
			case float64:
				result = append(result, int(e))
			case int:
				result = append(result, e)
			default:
				return nil, false
			}
		}
		return result, true
	default:
		return nil, false
	}
}
