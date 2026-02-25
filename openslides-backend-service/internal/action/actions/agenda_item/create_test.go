package agenda_item

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"topic/1": {
			"meeting_id": 1,
			"title":      "Test Topic",
		},
	})

	resp, err := tc.RequestInternal("agenda_item.create", map[string]any{
		"meeting_id":        1,
		"content_object_id": "topic/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"meeting_id":        1,
		"content_object_id": "topic/1",
	})
}

func TestCreateWithItemNumber(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"topic/1": {
			"meeting_id": 1,
			"title":      "Test Topic",
		},
	})

	resp, err := tc.RequestInternal("agenda_item.create", map[string]any{
		"meeting_id":        1,
		"content_object_id": "topic/1",
		"item_number":       "1.1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"item_number": "1.1",
	})
}

func TestCreateWithComment(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"topic/1": {
			"meeting_id": 1,
			"title":      "Test Topic",
		},
	})

	resp, err := tc.RequestInternal("agenda_item.create", map[string]any{
		"meeting_id":        1,
		"content_object_id": "topic/1",
		"comment":           "A comment",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"comment": "A comment",
	})
}

func TestCreateWithType(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"topic/1": {
			"meeting_id": 1,
			"title":      "Test Topic",
		},
	})

	resp, err := tc.RequestInternal("agenda_item.create", map[string]any{
		"meeting_id":        1,
		"content_object_id": "topic/1",
		"type":              2,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"type": 2,
	})
}

func TestCreateWithWeight(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"topic/1": {
			"meeting_id": 1,
			"title":      "Test Topic",
		},
	})

	resp, err := tc.RequestInternal("agenda_item.create", map[string]any{
		"meeting_id":        1,
		"content_object_id": "topic/1",
		"weight":            10,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"weight": 10,
	})
}

func TestCreateWithDuration(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"topic/1": {
			"meeting_id": 1,
			"title":      "Test Topic",
		},
	})

	resp, err := tc.RequestInternal("agenda_item.create", map[string]any{
		"meeting_id":        1,
		"content_object_id": "topic/1",
		"duration":          300,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"duration": 300,
	})
}

func TestCreateWithParentId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"topic/1": {
			"meeting_id": 1,
			"title":      "Parent Topic",
		},
		"topic/2": {
			"meeting_id": 1,
			"title":      "Child Topic",
		},
		"agenda_item/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"weight":            1,
		},
	})

	resp, err := tc.RequestInternal("agenda_item.create", map[string]any{
		"meeting_id":        1,
		"content_object_id": "topic/2",
		"parent_id":         1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/2", map[string]any{
		"parent_id": 1,
	})
}

func TestCreateMissingMeetingId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.RequestInternal("agenda_item.create", map[string]any{
		"content_object_id": "topic/1",
	})
	tc.AssertErrorContains(err, "meeting_id")
}

func TestCreateMissingContentObjectId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.RequestInternal("agenda_item.create", map[string]any{
		"meeting_id": 1,
	})
	tc.AssertErrorContains(err, "content_object_id")
}

func TestCreateForMotion(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "Test Motion",
		},
	})

	resp, err := tc.RequestInternal("agenda_item.create", map[string]any{
		"meeting_id":        1,
		"content_object_id": "motion/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("agenda_item/1", map[string]any{
		"content_object_id": "motion/1",
	})
}

func TestCreateWithTagIds(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"topic/1": {
			"meeting_id": 1,
			"title":      "Test Topic",
		},
		"tag/1": {
			"meeting_id": 1,
			"name":       "Important",
		},
	})

	resp, err := tc.RequestInternal("agenda_item.create", map[string]any{
		"meeting_id":        1,
		"content_object_id": "topic/1",
		"tag_ids":           []any{1},
	})
	tc.AssertSuccess(resp, err)
}
