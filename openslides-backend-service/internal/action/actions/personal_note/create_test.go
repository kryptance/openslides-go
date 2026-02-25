package personal_note

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreate(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/110": {
			"name":                        "name_meeting_110",
			"is_active_in_organization_id": 1,
			"meeting_user_ids":            []any{1},
			"group_ids":                   []any{1},
		},
		"meeting_user/1": {
			"meeting_id": 110,
			"user_id":    1,
		},
		"motion/23": {"meeting_id": 110},
		"user/1":    {"meeting_ids": []any{110}},
		"group/1":   {"meeting_id": 110},
	})

	resp, err := tc.RequestInternal("personal_note.create", map[string]any{
		"meeting_user_id":   1,
		"content_object_id": "motion/23",
		"star":              true,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("personal_note/1")
	if model["star"] != true {
		t.Errorf("expected star true, got %v", model["star"])
	}
	if model["meeting_user_id"] != 1 {
		t.Errorf("expected meeting_user_id 1, got %v", model["meeting_user_id"])
	}
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.RequestInternal("personal_note.create", map[string]any{})
	tc.AssertErrorContains(err, "meeting_user_id")
}
