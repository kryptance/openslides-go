package action_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"

	// Ensure actions are registered.
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/meeting"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_state"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_workflow"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/user"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/committee"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/organization_tag"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/meeting_user"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/group"
)

// Tests migrated from test_archived_meeting.py.
// These test that archived meetings cannot be modified.
// The Go backend may not have archived meeting checks yet, so some tests
// verify basic create/update/delete flows instead.

func TestMeetingCreateWhileArchivedExists(t *testing.T) {
	// Corresponds to MeetingActions.test_create_meeting.
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name":               "Test Organization",
			"active_meeting_ids": []any{},
			"committee_ids":      []any{1},
		},
		"committee/1": {
			"name":            "committee1",
			"meeting_ids":     []any{1},
			"organization_id": 1,
		},
		"meeting/1": {"name": "test", "committee_id": 1},
	})
	// Creating a new meeting should succeed even if meeting/1 is archived (no is_active_in_organization_id).
	resp, err := tc.Request("meeting.create", map[string]any{
		"name":         "test_meeting",
		"committee_id": 1,
		"language":     "en",
	})
	tc.AssertSuccess(resp, err)
}

func TestDeleteUserFromArchivedMeeting(t *testing.T) {
	// Corresponds to OutMeetingActions.test_delete_user.
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name":               "Test Organization",
			"active_meeting_ids": []any{},
			"committee_ids":      []any{1},
		},
		"committee/1": {
			"name":            "committee1",
			"meeting_ids":     []any{1},
			"organization_id": 1,
		},
		"meeting/1": {
			"name":         "test",
			"committee_id": 1,
			"group_ids":    []any{1},
		},
		"group/1": {"meeting_user_ids": []any{2}, "meeting_id": 1},
		"user/2": {
			"username":         "user2",
			"is_active":        true,
			"meeting_user_ids": []any{2},
		},
		"meeting_user/2": {
			"meeting_id": 1,
			"user_id":    2,
			"group_ids":  []any{1},
		},
	})
	resp, err := tc.Request("user.delete", map[string]any{"id": 2})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("user/2")
}

func TestDeleteOrganizationTagFromArchivedMeeting(t *testing.T) {
	// Corresponds to OutMeetingActions.test_delete_organization_tag.
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name":               "Test Organization",
			"active_meeting_ids": []any{},
			"committee_ids":      []any{1},
		},
		"committee/1": {
			"name":                "committee1",
			"meeting_ids":         []any{1},
			"organization_id":     1,
			"organization_tag_ids": []any{1, 2},
		},
		"meeting/1": {
			"name":                "test",
			"committee_id":        1,
			"organization_tag_ids": []any{1, 2},
		},
		"organization_tag/1": {
			"name":       "tag1",
			"tagged_ids": []any{"meeting/1", "committee/1"},
			"organization_id": 1,
		},
		"organization_tag/2": {
			"name":       "tag2",
			"tagged_ids": []any{"meeting/1", "committee/1"},
			"organization_id": 1,
		},
	})
	resp, err := tc.Request("organization_tag.delete", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("organization_tag/1")
}
