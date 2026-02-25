package motion_state

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/110": {
			"name":       "name_Ycefgee",
			"state_ids":  []any{111},
			"meeting_id": 1,
		},
		"motion_state/111": {
			"name":        "name_srtgb123",
			"workflow_id": 110,
			"meeting_id":  1,
		},
	})
	resp, err := tc.Request("motion_state.update", map[string]any{
		"id":   111,
		"name": "name_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_state/111", map[string]any{"name": "name_Xcdfgee"})
}

func TestUpdateWithNextPrevious(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/110": {
			"name":       "name_Ycefgee",
			"state_ids":  []any{111, 112, 113},
			"meeting_id": 1,
		},
		"motion_state/111": {
			"name":        "name_srtgb123",
			"workflow_id": 110,
			"meeting_id":  1,
		},
		"motion_state/112": {
			"name":        "name_srtfg112",
			"workflow_id": 110,
			"meeting_id":  1,
		},
		"motion_state/113": {
			"name":        "name_srtfg113",
			"workflow_id": 110,
			"meeting_id":  1,
		},
	})
	resp, err := tc.Request("motion_state.update", map[string]any{
		"id":                 111,
		"next_state_ids":     []any{112},
		"previous_state_ids": []any{113},
	})
	tc.AssertSuccess(resp, err)
}

func TestUpdateAllowMotionForwarding(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/110": {
			"name":       "name_Ycefgee",
			"state_ids":  []any{111},
			"meeting_id": 1,
		},
		"motion_state/111": {
			"name":        "name_srtgb123",
			"workflow_id": 110,
			"meeting_id":  1,
		},
	})
	resp, err := tc.Request("motion_state.update", map[string]any{
		"id":                      111,
		"name":                    "name_Xcdfgee",
		"is_internal":             true,
		"allow_motion_forwarding": true,
	})
	tc.AssertSuccess(resp, err)
}
