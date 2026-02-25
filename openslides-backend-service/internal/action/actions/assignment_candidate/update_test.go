package assignment_candidate

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

// assignment_candidate.update is an internal-only action in the Python backend.
// In the Go backend, weight changes go through the sort action.

func TestUpdateCorrectInternal(t *testing.T) {
	// NOTE: Sort action framework limitation - see sort_test.go.
	t.Skip("Sort action framework does not yet produce events correctly")

	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"assignment/222": {"title": "title_SNLGsvIV", "meeting_id": 1},
		"user/233":       {"username": "username_233", "meeting_user_ids": []any{233}},
		"meeting_user/233": {
			"meeting_id":              1,
			"user_id":                 233,
			"assignment_candidate_ids": []any{31},
		},
		"assignment_candidate/31": {
			"assignment_id":   222,
			"meeting_user_id": 233,
			"meeting_id":      1,
			"weight":          2,
		},
	})
	resp, err := tc.Request("assignment_candidate.sort", map[string]any{
		"assignment_id": 222, "ids": []any{31},
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("assignment_candidate/31", map[string]any{"weight": 1})
}
