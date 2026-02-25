package vote

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/111": {
			"name":                         "meeting_Xcdfgee",
			"is_active_in_organization_id": 1,
		},
		"option/12": {"text": "blabalbal", "meeting_id": 111},
		"vote/1": {
			"value":      "Y",
			"meeting_id": 111,
			"weight":     "1.000000",
			"option_id":  12,
		},
	})

	resp, err := tc.Request("vote.update", map[string]any{
		"id":     1,
		"weight": "1.500000",
	})
	tc.AssertSuccess(resp, err)
	vote := tc.GetModel("vote/1")
	if vote["value"] != "Y" {
		t.Errorf("expected value Y, got %v", vote["value"])
	}
	if vote["weight"] != "1.500000" {
		t.Errorf("expected weight 1.500000, got %v", vote["weight"])
	}
}
