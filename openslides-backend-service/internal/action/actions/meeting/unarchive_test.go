package meeting

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func setupUnarchiveMeeting(tc *testutil.ActionTestCase) {
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"active_meeting_ids": []any{},
		},
		"committee/1": {
			"name":            "test_committee",
			"organization_id": 1,
		},
		"meeting/1": {
			"committee_id": 1,
		},
	})
}

func TestUnarchiveSimple(t *testing.T) {
	tc := testutil.New(t)
	setupUnarchiveMeeting(tc)
	resp, err := tc.Request("meeting.unarchive", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
}

func TestUnarchiveTwoMeetings(t *testing.T) {
	tc := testutil.New(t)
	setupUnarchiveMeeting(tc)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"active_meeting_ids": []any{2},
		},
	})
	resp, err := tc.Request("meeting.unarchive", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
}

func TestUnarchiveLockedMeeting(t *testing.T) {
	tc := testutil.New(t)
	setupUnarchiveMeeting(tc)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"locked_from_inside": true,
		},
	})
	resp, err := tc.Request("meeting.unarchive", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
}
