package personal_note

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"is_active_in_organization_id": 1,
			"meeting_user_ids":             []any{1},
		},
		"personal_note/1": {
			"star":            true,
			"note":            "blablabla",
			"meeting_user_id": 1,
			"meeting_id":      1,
		},
		"user/1": {
			"meeting_ids":      []any{1},
			"meeting_user_ids": []any{1},
		},
		"meeting_user/1": {
			"user_id":           1,
			"meeting_id":        1,
			"personal_note_ids": []any{1},
		},
	})

	resp, err := tc.RequestInternal("personal_note.update", map[string]any{
		"id":   1,
		"star": false,
		"note": "blopblop",
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("personal_note/1")
	if model["star"] != false {
		t.Errorf("expected star false, got %v", model["star"])
	}
	if model["note"] != "blopblop" {
		t.Errorf("expected note 'blopblop', got %v", model["note"])
	}
}
