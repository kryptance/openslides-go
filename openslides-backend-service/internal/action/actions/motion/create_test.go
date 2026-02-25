package motion

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateGoodCaseRequiredFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/12": {
			"meeting_id":    1,
			"first_state_id": 34,
			"state_ids":     []any{34},
		},
		"motion_state/34": {
			"workflow_id": 12,
			"meeting_id":  1,
		},
	})
	resp, err := tc.Request("motion.create", map[string]any{
		"title":      "test_Xcdfgee",
		"meeting_id": 1,
		"workflow_id": 12,
		"text":       "test",
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion.create", map[string]any{})
	tc.AssertError(resp, err)
}

func TestCreateWorkflowId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/12": {
			"meeting_id":    1,
			"first_state_id": 34,
			"state_ids":     []any{34},
		},
		"motion_state/34": {
			"workflow_id": 12,
			"meeting_id":  1,
		},
	})
	resp, err := tc.Request("motion.create", map[string]any{
		"title":       "title_test1",
		"meeting_id":  1,
		"workflow_id": 12,
		"text":        "test",
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateWithTextAndTitle(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/12": {
			"meeting_id":    1,
			"first_state_id": 34,
			"state_ids":     []any{34},
		},
		"motion_state/34": {
			"workflow_id": 12,
			"meeting_id":  1,
		},
	})
	resp, err := tc.Request("motion.create", map[string]any{
		"title":       "title_test1",
		"meeting_id":  1,
		"workflow_id": 12,
		"text":        "test body text",
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateWithNumber(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/12": {
			"meeting_id":    1,
			"first_state_id": 34,
			"state_ids":     []any{34},
		},
		"motion_state/34": {
			"workflow_id": 12,
			"meeting_id":  1,
		},
	})
	resp, err := tc.Request("motion.create", map[string]any{
		"title":       "title_test1",
		"meeting_id":  1,
		"workflow_id": 12,
		"text":        "test",
		"number":      "001",
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateWithReason(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/12": {
			"meeting_id":    1,
			"first_state_id": 34,
			"state_ids":     []any{34},
		},
		"motion_state/34": {
			"workflow_id": 12,
			"meeting_id":  1,
		},
	})
	resp, err := tc.Request("motion.create", map[string]any{
		"title":       "title_test1",
		"meeting_id":  1,
		"workflow_id": 12,
		"text":        "test",
		"reason":      "test_reason",
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateWithCategoryId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/12": {
			"meeting_id":    1,
			"first_state_id": 34,
			"state_ids":     []any{34},
		},
		"motion_state/34": {
			"workflow_id": 12,
			"meeting_id":  1,
		},
		"motion_category/124": {
			"name":       "name_wbtlHQro",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion.create", map[string]any{
		"title":       "title_test1",
		"meeting_id":  1,
		"workflow_id": 12,
		"text":        "test",
		"category_id": 124,
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateWithBlockId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/12": {
			"meeting_id":    1,
			"first_state_id": 34,
			"state_ids":     []any{34},
		},
		"motion_state/34": {
			"workflow_id": 12,
			"meeting_id":  1,
		},
		"motion_block/78": {
			"title":      "title_kXTvKvjc",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion.create", map[string]any{
		"title":       "title_test1",
		"meeting_id":  1,
		"workflow_id": 12,
		"text":        "test",
		"block_id":    78,
	})
	tc.AssertSuccess(resp, err)
}
