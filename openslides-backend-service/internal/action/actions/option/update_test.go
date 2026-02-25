package option

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"topic/1": {"meeting_id": 1},
		"poll/65": {
			"content_object_id":   "topic/1",
			"type":                "analog",
			"state":               "created",
			"pollmethod":          "YNA",
			"max_votes_amount":    1,
			"max_votes_per_option": 1,
			"min_votes_amount":    1,
			"meeting_id":          1,
			"option_ids":          []any{57},
		},
		"option/57": {
			"yes":        "0.000000",
			"no":         "0.000000",
			"abstain":    "0.000000",
			"meeting_id": 1,
			"poll_id":    65,
		},
	})

	resp, err := tc.Request("option.update", map[string]any{
		"id":      57,
		"yes":     "1.000000",
		"no":      "2.000000",
		"abstain": "3.000000",
	})
	tc.AssertSuccess(resp, err)
	option := tc.GetModel("option/57")
	if option["yes"] != "1.000000" {
		t.Errorf("expected yes 1.000000, got %v", option["yes"])
	}
	if option["no"] != "2.000000" {
		t.Errorf("expected no 2.000000, got %v", option["no"])
	}
	if option["abstain"] != "3.000000" {
		t.Errorf("expected abstain 3.000000, got %v", option["abstain"])
	}
}
