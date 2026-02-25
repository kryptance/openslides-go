package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
)

// WithForwardMediafiles wraps the UpdateInstance hook to duplicate mediafiles
// from a source meeting to a target meeting when forwarding motions or
// other content between meetings.
//
// When content is forwarded (e.g., a motion with attachments), the mediafiles
// must be duplicated into the target meeting as meeting_mediafile records.
// This mixin handles that duplication by creating new meeting_mediafile entries
// in the target meeting for each mediafile referenced in the source.
//
// sourceField is the field containing the list of source meeting_mediafile_ids
// to forward. targetMeetingField is the field containing the target meeting_id.
//
// This replaces Python's MediafileAttachmentMixin / forward_attachment_ids logic.
func WithForwardMediafiles(a *action.BaseAction, sourceField, targetMeetingField string) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		sourceIDs, ok := instance[sourceField]
		if !ok {
			return origUpdate(ctx, params, instance)
		}

		ids, ok := toIntSlice(sourceIDs)
		if !ok || len(ids) == 0 {
			return origUpdate(ctx, params, instance)
		}

		// Get target meeting_id.
		targetMeetingIDRaw, ok := instance[targetMeetingField]
		if !ok {
			return nil, fmt.Errorf("forward mediafiles: %s is required when %s is provided", targetMeetingField, sourceField)
		}

		var targetMeetingID int
		switch v := targetMeetingIDRaw.(type) {
		case float64:
			targetMeetingID = int(v)
		case int:
			targetMeetingID = v
		default:
			return nil, fmt.Errorf("forward mediafiles: %s must be an integer", targetMeetingField)
		}

		if targetMeetingID == 0 {
			return nil, fmt.Errorf("forward mediafiles: %s must be non-zero", targetMeetingField)
		}

		// For each source meeting_mediafile, look up the underlying mediafile_id
		// and create a new meeting_mediafile in the target meeting.
		newMeetingMediafileIDs := make([]int, 0, len(ids))
		for _, sourceMmfID := range ids {
			newID, err := duplicateMeetingMediafile(ctx, params, sourceMmfID, targetMeetingID)
			if err != nil {
				return nil, fmt.Errorf("forward mediafiles: duplicate meeting_mediafile %d: %w", sourceMmfID, err)
			}
			if newID > 0 {
				newMeetingMediafileIDs = append(newMeetingMediafileIDs, newID)
			}
		}

		// Replace the source field with the new target meeting_mediafile IDs.
		instance[sourceField] = newMeetingMediafileIDs

		return origUpdate(ctx, params, instance)
	}
}

// duplicateMeetingMediafile creates a new meeting_mediafile in the target meeting
// that references the same underlying mediafile as the source meeting_mediafile.
func duplicateMeetingMediafile(ctx context.Context, params *action.ActionParams, sourceMmfID, targetMeetingID int) (int, error) {
	// Fetch the source meeting_mediafile to get its mediafile_id.
	source, err := params.Datastore.Get("meeting_mediafile", sourceMmfID, []string{"mediafile_id", "meeting_id"})
	if err != nil {
		return 0, fmt.Errorf("fetch meeting_mediafile/%d: %w", sourceMmfID, err)
	}

	mediafileIDRaw, ok := source["mediafile_id"]
	if !ok {
		return 0, fmt.Errorf("meeting_mediafile/%d has no mediafile_id", sourceMmfID)
	}

	var mediafileID int
	switch v := mediafileIDRaw.(type) {
	case float64:
		mediafileID = int(v)
	case int:
		mediafileID = v
	}

	if mediafileID == 0 {
		return 0, nil
	}

	// Check if target meeting already has a meeting_mediafile for this mediafile.
	// TODO: Query the datastore with a filter:
	//   meeting_mediafile WHERE meeting_id = targetMeetingID AND mediafile_id = mediafileID
	// If found, return the existing ID instead of creating a duplicate.

	// Create a new meeting_mediafile in the target meeting.
	createAction, err := action.Lookup("meeting_mediafile.create")
	if err != nil {
		return 0, fmt.Errorf("meeting_mediafile.create action not available: %w", err)
	}

	depInstance := action.Instance{
		"meeting_id":   targetMeetingID,
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
