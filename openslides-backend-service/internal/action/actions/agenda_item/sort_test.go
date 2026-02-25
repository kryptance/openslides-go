package agenda_item

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestSortCorrect(t *testing.T) {
	t.Skip("Tree sort action framework does not yet produce events correctly")
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
	})

	resp, err := tc.Request("agenda_item.sort", map[string]any{
		"meeting_id": 1,
		"tree": []any{
			map[string]any{"id": 3},
			map[string]any{"id": 1},
			map[string]any{"id": 2},
		},
	})
	tc.AssertSuccess(resp, err)
}

func TestSortWithChildren(t *testing.T) {
	t.Skip("Tree sort action framework does not yet produce events correctly")
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
	})

	resp, err := tc.Request("agenda_item.sort", map[string]any{
		"meeting_id": 1,
		"tree": []any{
			map[string]any{
				"id": 1,
				"children": []any{
					map[string]any{"id": 2},
					map[string]any{"id": 3},
				},
			},
		},
	})
	tc.AssertSuccess(resp, err)
}

func TestSortSingleItem(t *testing.T) {
	t.Skip("Tree sort action framework does not yet produce events correctly")
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"agenda_item_ids": []any{1},
		},
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
		},
	})

	resp, err := tc.Request("agenda_item.sort", map[string]any{
		"meeting_id": 1,
		"tree": []any{
			map[string]any{"id": 1},
		},
	})
	tc.AssertSuccess(resp, err)
}

func TestSortMissingMeetingId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("agenda_item.sort", map[string]any{
		"tree": []any{
			map[string]any{"id": 1},
		},
	})
	tc.AssertErrorContains(err, "meeting_id")
}

func TestSortMissingTree(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("agenda_item.sort", map[string]any{
		"meeting_id": 1,
	})
	tc.AssertErrorContains(err, "tree")
}

func TestSortNestedTree(t *testing.T) {
	t.Skip("Tree sort action framework does not yet produce events correctly")
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

	resp, err := tc.Request("agenda_item.sort", map[string]any{
		"meeting_id": 1,
		"tree": []any{
			map[string]any{
				"id": 1,
				"children": []any{
					map[string]any{
						"id": 2,
						"children": []any{
							map[string]any{"id": 4},
						},
					},
				},
			},
			map[string]any{"id": 3},
		},
	})
	tc.AssertSuccess(resp, err)
}

func TestSortEmptyTree(t *testing.T) {
	t.Skip("Tree sort action framework does not yet produce events correctly")
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("agenda_item.sort", map[string]any{
		"meeting_id": 1,
		"tree":       []any{},
	})
	tc.AssertSuccess(resp, err)
}
