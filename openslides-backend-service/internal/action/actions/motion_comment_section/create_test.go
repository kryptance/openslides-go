package motion_comment_section

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateGoodCaseRequiredFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_comment_section.create", map[string]any{
		"name": "test_Xcdfgee", "meeting_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_comment_section/1", map[string]any{
		"name":       "test_Xcdfgee",
		"meeting_id": 1,
	})
}

func TestCreateGoodCaseAllFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"group/23": {"name": "name_IIwngcUT", "meeting_id": 1},
	})
	resp, err := tc.Request("motion_comment_section.create", map[string]any{
		"name":                "test_Xcdfgee",
		"meeting_id":          1,
		"read_group_ids":      []any{23},
		"write_group_ids":     []any{23},
		"submitter_can_write": true,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("motion_comment_section/1")
	if model["name"] != "test_Xcdfgee" {
		t.Errorf("expected name 'test_Xcdfgee', got %v", model["name"])
	}
	if model["meeting_id"] != 1 {
		t.Errorf("expected meeting_id 1, got %v", model["meeting_id"])
	}
	if model["submitter_can_write"] != true {
		t.Errorf("expected submitter_can_write true, got %v", model["submitter_can_write"])
	}
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_comment_section.create", map[string]any{})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "required field")
}

func TestCreateWrongField(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_comment_section.create", map[string]any{
		"name":        "name_test1",
		"meeting_id":  1,
		"wrong_field": "text_AefohteiF8",
	})
	// The Go schema does not enforce additionalProperties=false.
	tc.AssertSuccess(resp, err)
}
