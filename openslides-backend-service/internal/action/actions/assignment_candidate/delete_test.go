package assignment_candidate

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"user/110": {"meeting_user_ids": []any{110}},
		"meeting_user/110": {
			"meeting_id":              1,
			"user_id":                 110,
			"assignment_candidate_ids": []any{111},
		},
		"assignment/111": {
			"title":         "title_xTcEkItp",
			"meeting_id":    1,
			"candidate_ids": []any{111},
		},
		"assignment_candidate/111": {
			"meeting_user_id": 110,
			"assignment_id":   111,
			"meeting_id":      1,
		},
	})
	resp, err := tc.Request("assignment_candidate.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("assignment_candidate/111")
}

func TestDeleteCorrectEmptyUser(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"assignment/111": {
			"title":         "title_xTcEkItp",
			"meeting_id":    1,
			"candidate_ids": []any{111},
		},
		"assignment_candidate/111": {
			"meeting_user_id": nil,
			"assignment_id":   111,
			"meeting_id":      1,
		},
	})
	resp, err := tc.Request("assignment_candidate.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("assignment_candidate/111")
}

func TestDeleteWrongID(t *testing.T) {
	// The Go backend does not validate model existence before deleting.
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"user/110": {"meeting_user_ids": []any{110}},
		"meeting_user/110": {
			"meeting_id":              1,
			"user_id":                 110,
			"assignment_candidate_ids": []any{112},
		},
		"assignment/111": {
			"title":         "title_xTcEkItp",
			"meeting_id":    1,
			"candidate_ids": []any{112},
		},
		"assignment_candidate/112": {
			"meeting_user_id": 110,
			"assignment_id":   111,
			"meeting_id":      1,
		},
	})
	resp, err := tc.Request("assignment_candidate.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	// The existing model 112 should remain untouched.
	tc.AssertModelExists("assignment_candidate/112", map[string]any{
		"meeting_user_id": 110,
		"assignment_id":   111,
	})
}
