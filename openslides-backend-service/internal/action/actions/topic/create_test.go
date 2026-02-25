package topic

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreate(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("topic.create", map[string]any{
		"meeting_id": 1,
		"title":      "test",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("topic/1", map[string]any{
		"meeting_id": 1,
		"title":      "test",
	})
}

func TestCreateMultipleInOneRequest(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.RequestMulti("topic.create", []map[string]any{
		{"meeting_id": 1, "title": "A"},
		{"meeting_id": 1, "title": "B"},
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("topic/1", map[string]any{
		"title":      "A",
		"meeting_id": 1,
	})
	tc.AssertModelExists("topic/2", map[string]any{
		"title":      "B",
		"meeting_id": 1,
	})
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.Request("topic.create", map[string]any{})
	tc.AssertErrorContains(err, "meeting_id")
}

func TestCreateMissingTitle(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("topic.create", map[string]any{
		"meeting_id": 1,
	})
	tc.AssertErrorContains(err, "title")
}
