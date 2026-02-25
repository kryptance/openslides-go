package vote

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
		"option/12": {"text": "blabalbal", "meeting_id": 111},
	})

	resp, err := tc.RequestInternal("vote.create", map[string]any{
		"value":      "Y",
		"weight":     "1.000000",
		"option_id":  12,
		"user_token": "aaaabbbbccccdddd",
		"meeting_id": 111,
	})
	tc.AssertSuccess(resp, err)
	vote := tc.GetModel("vote/1")
	if vote["value"] != "Y" {
		t.Errorf("expected value Y, got %v", vote["value"])
	}
	if vote["meeting_id"] != 111 {
		t.Errorf("expected meeting_id 111, got %v", vote["meeting_id"])
	}
	if vote["weight"] != "1.000000" {
		t.Errorf("expected weight 1.000000, got %v", vote["weight"])
	}
	if vote["option_id"] != 12 {
		t.Errorf("expected option_id 12, got %v", vote["option_id"])
	}
	if vote["user_token"] != "aaaabbbbccccdddd" {
		t.Errorf("expected user_token aaaabbbbccccdddd, got %v", vote["user_token"])
	}
}
