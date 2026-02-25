package meeting

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func setupCloneMeeting(tc *testutil.ActionTestCase) {
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"active_meeting_ids":   []any{1},
			"organization_tag_ids": []any{1},
			"user_ids":            []any{1},
			"template_meeting_ids": []any{1},
		},
		"organization_tag/1": {
			"name":            "TEST",
			"color":           "#eeeeee",
			"organization_id": 1,
		},
		"committee/1": {
			"organization_id": 1,
			"meeting_ids":     []any{1},
		},
		"committee/2": {
			"organization_id": 1,
		},
		"meeting/1": {
			"template_for_organization_id":        1,
			"committee_id":                        1,
			"language":                            "en",
			"name":                                "Test",
			"default_group_id":                    1,
			"admin_group_id":                      2,
			"motions_default_amendment_workflow_id": 1,
			"motions_default_workflow_id":          1,
			"reference_projector_id":              1,
			"projector_countdown_default_time":    60,
			"projector_countdown_warning_time":    5,
			"projector_ids":                       []any{1},
			"group_ids":                           []any{1, 2},
			"motion_state_ids":                    []any{1},
			"motion_workflow_ids":                 []any{1},
			"is_active_in_organization_id":        1,
		},
		"group/1": {
			"meeting_id":                   1,
			"name":                         "default group",
			"weight":                       1,
			"default_group_for_meeting_id": 1,
		},
		"group/2": {
			"meeting_id":                 1,
			"name":                       "admin group",
			"weight":                     1,
			"admin_group_for_meeting_id": 1,
		},
		"motion_workflow/1": {
			"meeting_id":                         1,
			"name":                               "blup",
			"first_state_id":                     1,
			"default_amendment_workflow_meeting_id": 1,
			"default_workflow_meeting_id":         1,
			"state_ids":                          []any{1},
			"sequential_number":                  1,
		},
		"motion_state/1": {
			"css_class":                1,
			"meeting_id":              1,
			"workflow_id":             1,
			"name":                    "test",
			"weight":                  1,
			"first_state_of_workflow_id": 1,
		},
		"projector/1": {
			"sequential_number":                   1,
			"meeting_id":                          1,
			"used_as_reference_projector_meeting_id": 1,
			"name":                                "Default projector",
		},
	})
}

func TestCloneWithoutUsers(t *testing.T) {
	tc := testutil.New(t)
	setupCloneMeeting(tc)
	resp, err := tc.Request("meeting.clone", map[string]any{
		"meeting_id":      1,
		"set_as_template": true,
	})
	tc.AssertSuccess(resp, err)
}

func TestCloneSimple(t *testing.T) {
	tc := testutil.New(t)
	setupCloneMeeting(tc)
	resp, err := tc.Request("meeting.clone", map[string]any{
		"meeting_id": 1,
	})
	tc.AssertSuccess(resp, err)
}

func TestCloneToOtherCommittee(t *testing.T) {
	tc := testutil.New(t)
	setupCloneMeeting(tc)
	resp, err := tc.Request("meeting.clone", map[string]any{
		"meeting_id":   1,
		"committee_id": 2,
	})
	tc.AssertSuccess(resp, err)
}

func TestCloneWithName(t *testing.T) {
	tc := testutil.New(t)
	setupCloneMeeting(tc)
	resp, err := tc.Request("meeting.clone", map[string]any{
		"meeting_id": 1,
		"name":       "Cloned Meeting",
	})
	tc.AssertSuccess(resp, err)
}

func TestCloneWithTimes(t *testing.T) {
	tc := testutil.New(t)
	setupCloneMeeting(tc)
	resp, err := tc.Request("meeting.clone", map[string]any{
		"meeting_id": 1,
		"start_time": 1700000000,
		"end_time":   1700100000,
	})
	tc.AssertSuccess(resp, err)
}

func TestCloneMissingMeetingId(t *testing.T) {
	tc := testutil.New(t)
	setupCloneMeeting(tc)
	resp, err := tc.Request("meeting.clone", map[string]any{})
	tc.AssertError(resp, err)
}
