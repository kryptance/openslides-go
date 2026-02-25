package meeting

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func setupDeleteMeeting(tc *testutil.ActionTestCase) {
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"committee_ids":      []any{1},
			"active_meeting_ids": []any{1},
		},
		"committee/1": {
			"organization_id": 1,
			"name":            "test_committee",
			"meeting_ids":     []any{1},
		},
		"group/11": {
			"meeting_id": 1,
		},
		"user/1": {
			"username": "user1",
		},
		"meeting/1": {
			"name":                         "test",
			"committee_id":                 1,
			"is_active_in_organization_id": 1,
			"group_ids":                    []any{11},
		},
	})
}

func TestDeleteSimple(t *testing.T) {
	tc := testutil.New(t)
	setupDeleteMeeting(tc)
	resp, err := tc.Request("meeting.delete", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("meeting/1")
}

func TestDeleteWithTagAndMotion(t *testing.T) {
	tc := testutil.New(t)
	setupDeleteMeeting(tc)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"tag_ids":    []any{42},
			"motion_ids": []any{1},
		},
		"tag/42": {
			"meeting_id": 1,
			"tagged_ids": []any{"motion/1"},
		},
		"motion/1": {
			"meeting_id": 1,
			"tag_ids":    []any{42},
		},
	})
	resp, err := tc.Request("meeting.delete", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("meeting/1")
}

func TestDeleteWithHistoryProjection(t *testing.T) {
	tc := testutil.New(t)
	setupDeleteMeeting(tc)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"all_projection_ids":           []any{42},
			"projector_ids":                []any{1},
			"is_active_in_organization_id": 1,
		},
		"projector/1": {
			"meeting_id":              1,
			"history_projection_ids":  []any{42},
			"current_projection_ids":  []any{42},
		},
		"projection/42": {
			"meeting_id":          1,
			"content_object_id":   "meeting/1",
			"history_projector_id": 1,
			"current_projector_id": 1,
			"stable":              false,
		},
	})
	resp, err := tc.Request("meeting.delete", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("meeting/1")
}

func TestDeleteArchivedMeeting(t *testing.T) {
	tc := testutil.New(t)
	setupDeleteMeeting(tc)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"active_meeting_ids": []any{},
		},
		"meeting/1": {
			"is_active_in_organization_id": nil,
		},
	})
	resp, err := tc.Request("meeting.delete", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("meeting/1")
}

func TestDeleteMissingId(t *testing.T) {
	tc := testutil.New(t)
	setupDeleteMeeting(tc)
	resp, err := tc.Request("meeting.delete", map[string]any{})
	tc.AssertError(resp, err)
}
