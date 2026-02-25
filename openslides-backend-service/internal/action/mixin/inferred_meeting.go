package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
)

// WithInferredMeeting wraps the UpdateInstance hook and the GetMeetingID hook
// to automatically set the meeting_id on the instance by looking it up from a
// related object.
//
// relationField is the field on the instance that references the related model
// (e.g., "motion_id"). relatedCollection is the collection of the related model
// (e.g., "motion"). The related model must have a "meeting_id" field.
//
// This replaces Python's action.get_meeting_id implementations that derive the
// meeting from a related model.
func WithInferredMeeting(a *action.BaseAction, relationField, relatedCollection string) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// If meeting_id is already set, skip inference.
		if _, ok := instance["meeting_id"]; ok {
			return origUpdate(ctx, params, instance)
		}

		meetingID, err := inferMeetingID(params, instance, relationField, relatedCollection)
		if err != nil {
			return nil, fmt.Errorf("inferred meeting: %w", err)
		}

		if meetingID > 0 {
			instance["meeting_id"] = meetingID
		}

		return origUpdate(ctx, params, instance)
	}

	// Also override GetMeetingID to use the same inference logic.
	origGetMeeting := a.GetMeetingID
	a.GetMeetingID = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (int, error) {
		// First try the original logic (direct meeting_id field).
		meetingID, err := origGetMeeting(ctx, params, instance)
		if err != nil {
			return 0, err
		}
		if meetingID > 0 {
			return meetingID, nil
		}

		// Infer from the related object.
		return inferMeetingID(params, instance, relationField, relatedCollection)
	}
}

// inferMeetingID looks up the meeting_id from a related object.
func inferMeetingID(params *action.ActionParams, instance action.Instance, relationField, relatedCollection string) (int, error) {
	relIDRaw, ok := instance[relationField]
	if !ok {
		return 0, nil
	}

	var relID int
	switch v := relIDRaw.(type) {
	case float64:
		relID = int(v)
	case int:
		relID = v
	default:
		return 0, backenderr.ActionError{
			Message: fmt.Sprintf("Field %q must be an integer", relationField),
		}
	}

	if relID == 0 {
		return 0, nil
	}

	// Fetch meeting_id from the related model.
	related, err := params.Datastore.Get(relatedCollection, relID, []string{"meeting_id"})
	if err != nil {
		return 0, fmt.Errorf("fetch %s/%d: %w", relatedCollection, relID, err)
	}

	meetingIDRaw, ok := related["meeting_id"]
	if !ok {
		return 0, nil
	}

	switch v := meetingIDRaw.(type) {
	case float64:
		return int(v), nil
	case int:
		return v, nil
	default:
		return 0, nil
	}
}
