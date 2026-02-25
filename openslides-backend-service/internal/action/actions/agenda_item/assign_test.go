package agenda_item

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestAssignCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"agenda_item_ids": []any{1, 2, 3},
		},
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
			"item_number":       "1",
		},
		"agenda_item/2": {
			"meeting_id":        1,
			"content_object_id": "topic/2",
			"weight":            2,
			"item_number":       "2",
		},
		"agenda_item/3": {
			"meeting_id":        1,
			"content_object_id": "topic/3",
			"weight":            3,
			"item_number":       "3",
		},
	})

	resp, err := tc.Request("agenda_item.assign", map[string]any{
		"meeting_id": 1,
		"ids":        []any{2, 3},
		"parent_id":  1,
	})
	tc.AssertSuccess(resp, err)
}

func TestAssignWithoutParent(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"agenda_item_ids": []any{1, 2},
		},
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
			"parent_id":         nil,
		},
		"agenda_item/2": {
			"meeting_id":        1,
			"content_object_id": "topic/2",
			"weight":            2,
			"parent_id":         1,
		},
	})

	resp, err := tc.Request("agenda_item.assign", map[string]any{
		"meeting_id": 1,
		"ids":        []any{2},
	})
	tc.AssertSuccess(resp, err)
}

func TestAssignMissingMeetingId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("agenda_item.assign", map[string]any{
		"ids": []any{1, 2},
	})
	tc.AssertErrorContains(err, "meeting_id")
}

func TestAssignMissingIds(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("agenda_item.assign", map[string]any{
		"meeting_id": 1,
	})
	tc.AssertErrorContains(err, "ids")
}

func TestAssignMultipleItems(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"agenda_item_ids": []any{1, 2, 3, 4},
		},
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
		"agenda_item/3": {
			"meeting_id":        1,
			"content_object_id": "topic/3",
			"weight":            3,
		},
		"agenda_item/4": {
			"meeting_id":        1,
			"content_object_id": "topic/4",
			"weight":            4,
		},
	})

	resp, err := tc.Request("agenda_item.assign", map[string]any{
		"meeting_id": 1,
		"ids":        []any{2, 3, 4},
		"parent_id":  1,
	})
	tc.AssertSuccess(resp, err)
}

func TestAssignEmptyIds(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("agenda_item.assign", map[string]any{
		"meeting_id": 1,
		"ids":        []any{},
	})
	tc.AssertSuccess(resp, err)
}
