package group

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("group.update", map[string]any{
		"id":   3,
		"name": "name_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("group/3")
	if model["name"] != "name_Xcdfgee" {
		t.Errorf("expected name name_Xcdfgee, got %v", model["name"])
	}
}

func TestUpdateExternalID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("group.update", map[string]any{
		"id":          3,
		"external_id": "test",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("group/3", map[string]any{
		"external_id": "test",
	})
}
