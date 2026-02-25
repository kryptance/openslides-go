package meeting

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func setupArchiveMeeting(tc *testutil.ActionTestCase) {
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"active_meeting_ids": []any{1},
		},
		"committee/1": {
			"name":            "test_committee",
			"organization_id": 1,
		},
		"meeting/1": {
			"committee_id":                 1,
			"is_active_in_organization_id": 1,
		},
	})
}

func TestArchiveSimple(t *testing.T) {
	tc := testutil.New(t)
	setupArchiveMeeting(tc)
	resp, err := tc.Request("meeting.archive", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
}

func TestArchiveTwoMeetings(t *testing.T) {
	tc := testutil.New(t)
	setupArchiveMeeting(tc)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"active_meeting_ids": []any{1, 2},
		},
	})
	resp, err := tc.Request("meeting.archive", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
}

func TestArchiveMissingId(t *testing.T) {
	tc := testutil.New(t)
	setupArchiveMeeting(tc)
	resp, err := tc.Request("meeting.archive", map[string]any{})
	tc.AssertError(resp, err)
}

func TestArchiveLockedMeeting(t *testing.T) {
	tc := testutil.New(t)
	setupArchiveMeeting(tc)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"locked_from_inside": true,
		},
	})
	resp, err := tc.Request("meeting.archive", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
}
