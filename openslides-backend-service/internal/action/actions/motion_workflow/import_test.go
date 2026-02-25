package motion_workflow

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestImportSimpleCase(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/42": {
			"name":                        "test_name_fsdksjdfhdsfssdf",
			"is_active_in_organization_id": 1,
			"committee_id":                1,
		},
	})
	resp, err := tc.Request("motion_workflow.import", map[string]any{
		"name":       "test_Xcdfgee",
		"meeting_id": 42,
		"states": []any{
			map[string]any{
				"name":   "begin",
				"weight": 1,
			},
		},
	})
	tc.AssertSuccess(resp, err)
}

func TestImportEmptyData(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_workflow.import", map[string]any{})
	tc.AssertError(resp, err)
}

func TestImportMissingName(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_workflow.import", map[string]any{
		"meeting_id": 42,
	})
	tc.AssertError(resp, err)
}
