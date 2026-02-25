package agenda_item

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
)

func TestNumberingCorrect(t *testing.T) {
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

	resp, err := tc.Request("agenda_item.numbering", map[string]any{
		"meeting_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{"item_number": "1"})
	tc.AssertModelExists("agenda_item/2", map[string]any{"item_number": "2"})
	tc.AssertModelExists("agenda_item/3", map[string]any{"item_number": "3"})
}

func TestNumberingGeneratesEvents(t *testing.T) {
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
		},
		"agenda_item/2": {
			"meeting_id":        1,
			"content_object_id": "topic/2",
			"weight":            2,
		},
	})

	resp, err := tc.Request("agenda_item.numbering", map[string]any{
		"meeting_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertEventCount(resp, 2)
	tc.AssertHasEvent(resp, event.TypeUpdate, "agenda_item/1")
	tc.AssertHasEvent(resp, event.TypeUpdate, "agenda_item/2")
}

func TestNumberingSingleItem(t *testing.T) {
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

	resp, err := tc.Request("agenda_item.numbering", map[string]any{
		"meeting_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{"item_number": "1"})
}

func TestNumberingNoItems(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"agenda_item_ids": []any{},
		},
	})

	resp, err := tc.Request("agenda_item.numbering", map[string]any{
		"meeting_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertEventCount(resp, 0)
}

func TestNumberingNilAgendaItemIds(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	// Meeting without agenda_item_ids field set.

	resp, err := tc.Request("agenda_item.numbering", map[string]any{
		"meeting_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertEventCount(resp, 0)
}

func TestNumberingMissingMeetingId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("agenda_item.numbering", map[string]any{})
	tc.AssertErrorContains(err, "meeting_id")
}

func TestNumberingWrongMeetingId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("agenda_item.numbering", map[string]any{
		"meeting_id": 9999,
	})
	tc.AssertError(resp, err)
}

func TestNumberingSequentialNumbers(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"agenda_item_ids": []any{10, 20, 30},
		},
		"agenda_item/10": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            10,
		},
		"agenda_item/20": {
			"meeting_id":        1,
			"content_object_id": "topic/2",
			"weight":            20,
		},
		"agenda_item/30": {
			"meeting_id":        1,
			"content_object_id": "topic/3",
			"weight":            30,
		},
	})

	resp, err := tc.Request("agenda_item.numbering", map[string]any{
		"meeting_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/10", map[string]any{"item_number": "1"})
	tc.AssertModelExists("agenda_item/20", map[string]any{"item_number": "2"})
	tc.AssertModelExists("agenda_item/30", map[string]any{"item_number": "3"})
}
