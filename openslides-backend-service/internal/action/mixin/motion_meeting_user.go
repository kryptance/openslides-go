package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

// MotionMeetingUserConfig defines the configuration for motion<->meeting_user
// CRUD action factories. Different motion relation types (submitter, editor,
// working_group_speaker) use the same pattern but with different field names.
type MotionMeetingUserConfig struct {
	// Collection is the junction collection name (e.g., "motion_submitter").
	Collection string

	// MotionField is the field on the junction model referencing the motion
	// (e.g., "motion_id").
	MotionField string

	// MeetingUserField is the field on the junction model referencing the
	// meeting_user (e.g., "meeting_user_id").
	MeetingUserField string

	// WeightField is the field used for ordering (e.g., "weight").
	WeightField string

	// MotionRelationField is the field on the motion that holds the list of
	// junction model IDs (e.g., "submitter_ids").
	MotionRelationField string
}

// NewMotionMeetingUserCreateAction creates a create action for a
// motion<->meeting_user junction model with standard mixins applied.
//
// The created action:
//   - Reserves an ID for the new junction record
//   - Infers the meeting_id from the related motion
//   - Auto-assigns a weight if not provided
//   - Creates the junction record
//
// This replaces Python's MotionSubmitterCreateAction, MotionEditorCreateAction,
// MotionWorkingGroupSpeakerCreateAction pattern.
func NewMotionMeetingUserCreateAction(name string, m *model.ModelDef, cfg MotionMeetingUserConfig) *action.BaseAction {
	a := action.NewCreateAction(name, m)

	// Infer meeting_id from the motion.
	WithInferredMeeting(a, cfg.MotionField, "motion")

	// Auto-assign weight.
	if cfg.WeightField != "" {
		WithWeight(a, cfg.WeightField)
	}

	return a
}

// NewMotionMeetingUserDeleteAction creates a delete action for a
// motion<->meeting_user junction model.
//
// The created action:
//   - Deletes the junction record
//   - Infers the meeting_id from the junction record's motion relation
//
// This replaces Python's MotionSubmitterDeleteAction, etc.
func NewMotionMeetingUserDeleteAction(name string, m *model.ModelDef, cfg MotionMeetingUserConfig) *action.BaseAction {
	a := action.NewDeleteAction(name, m)

	// Override GetMeetingID to infer from the junction record's motion.
	origGetMeeting := a.GetMeetingID
	a.GetMeetingID = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (int, error) {
		// First try direct meeting_id.
		meetingID, err := origGetMeeting(ctx, params, instance)
		if err != nil {
			return 0, err
		}
		if meetingID > 0 {
			return meetingID, nil
		}

		// Look up the junction record to find the motion_id, then infer meeting.
		id, ok := instance["id"]
		if !ok {
			return 0, nil
		}
		var idInt int
		switch v := id.(type) {
		case float64:
			idInt = int(v)
		case int:
			idInt = v
		}

		record, err := params.Datastore.Get(cfg.Collection, idInt, []string{cfg.MotionField})
		if err != nil {
			return 0, fmt.Errorf("get %s/%d: %w", cfg.Collection, idInt, err)
		}

		motionIDRaw, ok := record[cfg.MotionField]
		if !ok {
			return 0, nil
		}

		var motionID int
		switch v := motionIDRaw.(type) {
		case float64:
			motionID = int(v)
		case int:
			motionID = v
		}

		if motionID == 0 {
			return 0, nil
		}

		motion, err := params.Datastore.Get("motion", motionID, []string{"meeting_id"})
		if err != nil {
			return 0, fmt.Errorf("get motion/%d: %w", motionID, err)
		}

		meetingIDRaw, ok := motion["meeting_id"]
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

	return a
}

// NewMotionMeetingUserSortAction creates a sort action for a
// motion<->meeting_user junction model that reorders the junction records
// for a given motion.
//
// The instance should contain:
//   - motion_id: the ID of the motion whose junction records to sort
//   - <junction>_ids: the ordered list of junction record IDs
//
// This replaces Python's MotionSubmitterSortAction, etc.
func NewMotionMeetingUserSortAction(name string, m *model.ModelDef, cfg MotionMeetingUserConfig) *action.BaseAction {
	a := action.NewUpdateAction(name, m)

	// Sort actions accept exactly one instance.
	WithSingularAction(a)

	// Apply linear sort on the junction IDs.
	if cfg.WeightField != "" {
		WithLinearSort(a, cfg.WeightField)
	}

	return a
}
