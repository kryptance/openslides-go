package personal_note

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/111": {
			"personal_note_ids":            []any{1},
			"is_active_in_organization_id": 1,
			"meeting_user_ids":             []any{1},
		},
		"user/1": {
			"meeting_user_ids": []any{1},
			"meeting_ids":      []any{111},
		},
		"personal_note/1": {
			"star":            true,
			"note":            "blablabla",
			"meeting_user_id": 1,
			"meeting_id":      111,
		},
		"meeting_user/1": {
			"user_id":           1,
			"meeting_id":        111,
			"personal_note_ids": []any{1},
		},
	})

	resp, err := tc.RequestInternal("personal_note.delete", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("personal_note/1")
}
