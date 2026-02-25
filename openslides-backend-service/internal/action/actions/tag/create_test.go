package tag

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreate(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {"tag_ids": []any{}},
	})

	resp, err := tc.Request("tag.create", map[string]any{
		"name":       "test_Xcdfgee",
		"meeting_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("tag/1", map[string]any{
		"name":       "test_Xcdfgee",
		"meeting_id": 1,
	})
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.Request("tag.create", map[string]any{})
	tc.AssertErrorContains(err, "meeting_id")
}
