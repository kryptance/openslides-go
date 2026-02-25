package tag

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {"tag_ids": []any{111}},
		"tag/111":   {"name": "name_srtgb123", "meeting_id": 1},
	})

	resp, err := tc.Request("tag.update", map[string]any{"id": 111, "name": "name_Xcdfgee"})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("tag/111")
	if model["name"] != "name_Xcdfgee" {
		t.Errorf("expected name 'name_Xcdfgee', got %v", model["name"])
	}
}

func TestUpdateMissingID(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.Request("tag.update", map[string]any{"name": "test"})
	tc.AssertErrorContains(err, "id")
}
