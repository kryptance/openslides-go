package motion

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateForwardedCorrectOriginIdSet(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/2": {
			"name":                            "name_SNLGsvIV",
			"motions_default_workflow_id":      12,
			"committee_id":                     1,
			"is_active_in_organization_id":     1,
			"default_group_id":                 112,
			"group_ids":                        []any{112},
		},
		"group/112": {
			"name":       "YZJAwUPK",
			"meeting_id": 2,
		},
		"motion_workflow/12": {
			"name":          "name_workflow1",
			"first_state_id": 34,
			"state_ids":     []any{34},
			"meeting_id":    2,
		},
		"motion_state/34": {
			"name":       "name_state34",
			"meeting_id": 2,
		},
		"motion/12": {
			"title":      "title_FcnPUXJB",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion.create_forwarded", map[string]any{
		"title":      "test_Xcdfgee",
		"meeting_id": 2,
		"origin_id":  12,
		"text":       "test",
		"reason":     "reason_jLvcgAMx",
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateForwardedMissingMeetingId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion.create_forwarded", map[string]any{
		"title":     "test_Xcdfgee",
		"origin_id": 12,
		"text":      "test",
	})
	tc.AssertError(resp, err)
}

func TestCreateForwardedWithBasicFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/2": {
			"name":                        "Meeting 2",
			"committee_id":                1,
			"is_active_in_organization_id": 1,
		},
		"motion_workflow/12": {
			"name":          "workflow1",
			"first_state_id": 34,
			"state_ids":     []any{34},
			"meeting_id":    2,
		},
		"motion_state/34": {
			"name":       "state34",
			"meeting_id": 2,
		},
		"motion/12": {
			"title":      "title_original",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion.create_forwarded", map[string]any{
		"title":      "forwarded_motion",
		"meeting_id": 2,
		"origin_id":  12,
		"text":       "forwarded text",
	})
	tc.AssertSuccess(resp, err)
}
