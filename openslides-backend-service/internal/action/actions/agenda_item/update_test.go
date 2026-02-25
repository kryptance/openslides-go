package agenda_item

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateItemNumber(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
			"item_number":       "1",
		},
	})

	resp, err := tc.Request("agenda_item.update", map[string]any{
		"id":          1,
		"item_number": "2.1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"item_number": "2.1",
	})
}

func TestUpdateComment(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
		},
	})

	resp, err := tc.Request("agenda_item.update", map[string]any{
		"id":      1,
		"comment": "Updated comment",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"comment": "Updated comment",
	})
}

func TestUpdateType(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
			"type":              1,
		},
	})

	resp, err := tc.Request("agenda_item.update", map[string]any{
		"id":   1,
		"type": 2,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"type": 2,
	})
}

func TestUpdateWeight(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
		},
	})

	resp, err := tc.Request("agenda_item.update", map[string]any{
		"id":     1,
		"weight": 42,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"weight": 42,
	})
}

func TestUpdateDuration(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
		},
	})

	resp, err := tc.Request("agenda_item.update", map[string]any{
		"id":       1,
		"duration": 600,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"duration": 600,
	})
}

func TestUpdateClosed(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
			"closed":            false,
		},
	})

	resp, err := tc.Request("agenda_item.update", map[string]any{
		"id":     1,
		"closed": true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"closed": true,
	})
}

func TestUpdateParentId(t *testing.T) {
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
	})

	resp, err := tc.Request("agenda_item.update", map[string]any{
		"id":        2,
		"parent_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/2", map[string]any{
		"parent_id": 1,
	})
}

func TestUpdateTagIds(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
		},
		"tag/1": {
			"meeting_id": 1,
			"name":       "Important",
		},
		"tag/2": {
			"meeting_id": 1,
			"name":       "Urgent",
		},
	})

	resp, err := tc.Request("agenda_item.update", map[string]any{
		"id":      1,
		"tag_ids": []any{1, 2},
	})
	tc.AssertSuccess(resp, err)
}

func TestUpdateMultipleFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
		},
	})

	resp, err := tc.Request("agenda_item.update", map[string]any{
		"id":          1,
		"item_number": "3.2",
		"comment":     "New comment",
		"duration":    120,
		"closed":      true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"item_number": "3.2",
		"comment":     "New comment",
		"duration":    120,
		"closed":      true,
	})
}

func TestUpdateMissingID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("agenda_item.update", map[string]any{
		"item_number": "1",
	})
	tc.AssertErrorContains(err, "id")
}

func TestUpdateWrongID(t *testing.T) {
	t.Skip("In-memory datastore does not validate model existence on update")
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("agenda_item.update", map[string]any{
		"id":          9999,
		"item_number": "1",
	})
	tc.AssertError(resp, err)
}

func TestUpdateTypeInternal(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
			"type":              1,
		},
	})

	resp, err := tc.Request("agenda_item.update", map[string]any{
		"id":   1,
		"type": 3,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"type": 3,
	})
}
