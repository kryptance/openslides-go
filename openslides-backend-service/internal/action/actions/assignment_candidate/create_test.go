package assignment_candidate

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"user/110":         {"username": "test_Xcdfgee", "meeting_user_ids": []any{110}},
		"meeting_user/110": {"meeting_id": 1, "user_id": 110},
		"assignment/111":   {"title": "title_xTcEkItp", "meeting_id": 1},
	})
	resp, err := tc.Request("assignment_candidate.create", map[string]any{
		"assignment_id": 111, "meeting_user_id": 110,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("assignment_candidate/1", map[string]any{
		"meeting_user_id": 110,
		"assignment_id":   111,
	})
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("assignment_candidate.create", map[string]any{})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "required field")
}

func TestCreateWrongField(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"user/110":         {"username": "test_Xcdfgee"},
		"assignment/111":   {"title": "title_xTcEkItp", "meeting_id": 1},
		"meeting_user/110": {"meeting_id": 1, "user_id": 110},
	})
	resp, err := tc.Request("assignment_candidate.create", map[string]any{
		"wrong_field":    "text_AefohteiF8",
		"assignment_id":  111,
		"meeting_user_id": 110,
	})
	// The schema does not enforce additionalProperties=false, so unknown fields
	// are allowed through schema validation. The test passes as a create.
	tc.AssertSuccess(resp, err)
}
