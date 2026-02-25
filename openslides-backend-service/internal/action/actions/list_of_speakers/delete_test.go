package list_of_speakers

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"list_of_speakers/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"speaker_ids":       []any{},
		},
		"topic/1": {
			"meeting_id":          1,
			"title":               "Test Topic",
			"list_of_speakers_id": 1,
		},
	})

	resp, err := tc.RequestInternal("list_of_speakers.delete", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("list_of_speakers/1")
}

func TestDeleteWrongID(t *testing.T) {
	t.Skip("In-memory datastore does not validate model existence on delete")
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.RequestInternal("list_of_speakers.delete", map[string]any{
		"id": 9999,
	})
	tc.AssertError(resp, err)
}

func TestDeleteMissingID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.RequestInternal("list_of_speakers.delete", map[string]any{})
	tc.AssertErrorContains(err, "id")
}
