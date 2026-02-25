package poll

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "analog",
			"pollmethod":        "YNA",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.delete", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("poll/1")
}

func TestDeleteWithOptions(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "analog",
			"pollmethod":        "YNA",
			"state":             "finished",
			"content_object_id": "motion/1",
			"option_ids":        []any{1, 2},
		},
		"option/1": {
			"meeting_id": 1,
			"poll_id":    1,
		},
		"option/2": {
			"meeting_id": 1,
			"poll_id":    1,
		},
	})

	resp, err := tc.Request("poll.delete", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("poll/1")
}

func TestDeleteWrongID(t *testing.T) {
	t.Skip("In-memory datastore does not validate model existence on delete")
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("poll.delete", map[string]any{
		"id": 9999,
	})
	tc.AssertError(resp, err)
}

func TestDeleteMissingID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("poll.delete", map[string]any{})
	tc.AssertErrorContains(err, "id")
}
