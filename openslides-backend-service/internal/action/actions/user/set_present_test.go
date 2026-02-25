package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
)

func TestSetPresentAddCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"name":                         "Test Meeting",
			"committee_id":                 1,
			"is_active_in_organization_id": 1,
		},
		"user/111": {"username": "username_srtgb123"},
		"committee/1": {},
	})
	resp, err := tc.Request("user.set_present", map[string]any{
		"id":         111,
		"meeting_id": 1,
		"present":    true,
	})
	tc.AssertSuccess(resp, err)
	// Verify the event adds meeting_id to is_present_in_meeting_ids
	tc.AssertHasEvent(resp, event.TypeUpdate, "user/111")
}

func TestSetPresentDelCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"present_user_ids":             []any{111},
			"committee_id":                 1,
			"is_active_in_organization_id": 1,
		},
		"user/111": {
			"username":                  "username_srtgb123",
			"is_present_in_meeting_ids": []any{1},
		},
		"committee/1": {},
	})
	resp, err := tc.Request("user.set_present", map[string]any{
		"id":         111,
		"meeting_id": 1,
		"present":    false,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertHasEvent(resp, event.TypeUpdate, "user/111")
}

func TestSetPresentGeneratesListFieldsAdd(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/111": {"username": "testuser"},
		"meeting/1": {
			"committee_id":                 1,
			"is_active_in_organization_id": 1,
		},
		"committee/1": {},
	})
	resp, err := tc.Request("user.set_present", map[string]any{
		"id":         111,
		"meeting_id": 1,
		"present":    true,
	})
	tc.AssertSuccess(resp, err)
	if resp == nil || len(resp.Events) == 0 {
		t.Fatal("expected events")
	}
	e := resp.Events[0]
	if e.ListFields == nil {
		t.Fatal("expected list fields")
	}
	if len(e.ListFields.Add["is_present_in_meeting_ids"]) != 1 {
		t.Fatalf("expected 1 add entry, got %d", len(e.ListFields.Add["is_present_in_meeting_ids"]))
	}
}

func TestSetPresentGeneratesListFieldsRemove(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/111": {
			"username":                  "testuser",
			"is_present_in_meeting_ids": []any{1},
		},
		"meeting/1": {
			"committee_id":                 1,
			"is_active_in_organization_id": 1,
			"present_user_ids":             []any{111},
		},
		"committee/1": {},
	})
	resp, err := tc.Request("user.set_present", map[string]any{
		"id":         111,
		"meeting_id": 1,
		"present":    false,
	})
	tc.AssertSuccess(resp, err)
	if resp == nil || len(resp.Events) == 0 {
		t.Fatal("expected events")
	}
	e := resp.Events[0]
	if e.ListFields == nil {
		t.Fatal("expected list fields")
	}
	if len(e.ListFields.Remove["is_present_in_meeting_ids"]) != 1 {
		t.Fatalf("expected 1 remove entry, got %d", len(e.ListFields.Remove["is_present_in_meeting_ids"]))
	}
}

func TestSetPresentEventCount(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/111": {"username": "testuser"},
		"meeting/1": {
			"committee_id":                 1,
			"is_active_in_organization_id": 1,
		},
		"committee/1": {},
	})
	resp, err := tc.Request("user.set_present", map[string]any{
		"id":         111,
		"meeting_id": 1,
		"present":    true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertEventCount(resp, 1)
}
