package meeting

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestImportSimple(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"active_meeting_ids": []any{1},
			"committee_ids":     []any{1},
		},
		"committee/1": {
			"organization_id": 1,
			"meeting_ids":     []any{1},
		},
		"meeting/1": {
			"committee_id": 1,
			"group_ids":    []any{1},
		},
		"group/1": {
			"meeting_id": 1,
			"name":       "group1_m1",
		},
	})
	resp, err := tc.Request("meeting.import", map[string]any{
		"committee_id": 1,
		"meeting": map[string]any{
			"meeting": map[string]any{
				"1": map[string]any{
					"id":       1,
					"language": "en",
					"name":     "Test",
				},
			},
		},
	})
	tc.AssertSuccess(resp, err)
}

func TestImportMissingCommitteeId(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("meeting.import", map[string]any{
		"meeting": map[string]any{},
	})
	tc.AssertError(resp, err)
}

func TestImportMissingMeeting(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("meeting.import", map[string]any{
		"committee_id": 1,
	})
	tc.AssertError(resp, err)
}
