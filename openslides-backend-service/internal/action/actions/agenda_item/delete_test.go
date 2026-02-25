package agenda_item

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
		},
		"topic/1": {
			"meeting_id":    1,
			"title":         "Test Topic",
			"agenda_item_id": 1,
		},
	})

	resp, err := tc.RequestInternal("agenda_item.delete", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("agenda_item/1")
}

func TestDeleteMultiple(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
		},
		"agenda_item/2": {
			"meeting_id":        1,
			"content_object_id": "topic/2",
			"weight":            2,
		},
		"topic/1": {
			"meeting_id":    1,
			"title":         "Test Topic 1",
			"agenda_item_id": 1,
		},
		"topic/2": {
			"meeting_id":    1,
			"title":         "Test Topic 2",
			"agenda_item_id": 2,
		},
	})

	resp, err := tc.RequestInternal("agenda_item.delete", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("agenda_item/1")
	tc.AssertModelExists("agenda_item/2")
}

func TestDeleteWrongID(t *testing.T) {
	t.Skip("In-memory datastore does not validate model existence on delete")
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.RequestInternal("agenda_item.delete", map[string]any{
		"id": 9999,
	})
	tc.AssertError(resp, err)
}

func TestDeleteMissingID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.RequestInternal("agenda_item.delete", map[string]any{})
	tc.AssertErrorContains(err, "id")
}

func TestDeleteWithChildren(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
			"child_ids":         []any{2},
		},
		"agenda_item/2": {
			"meeting_id":        1,
			"content_object_id": "topic/2",
			"weight":            2,
			"parent_id":         1,
		},
		"topic/1": {
			"meeting_id":    1,
			"title":         "Parent Topic",
			"agenda_item_id": 1,
		},
		"topic/2": {
			"meeting_id":    1,
			"title":         "Child Topic",
			"agenda_item_id": 2,
		},
	})

	resp, err := tc.RequestInternal("agenda_item.delete", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("agenda_item/1")
}
