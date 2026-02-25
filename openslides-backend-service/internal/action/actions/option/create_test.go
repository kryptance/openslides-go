package option

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/111": {
			"name":                         "meeting_Xcdfgee",
			"is_active_in_organization_id": 1,
		},
		"poll/1": {
			"meeting_id": 111,
		},
	})

	resp, err := tc.RequestInternal("option.create", map[string]any{
		"text":       "testtesttest",
		"meeting_id": 111,
		"weight":     10,
		"poll_id":    1,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("option/1")
	if model["text"] != "testtesttest" {
		t.Errorf("expected text testtesttest, got %v", model["text"])
	}
	if model["meeting_id"] != 111 {
		t.Errorf("expected meeting_id 111, got %v", model["meeting_id"])
	}
}
