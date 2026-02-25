package assignment

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateSimple(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name":               "Test Organization",
			"active_meeting_ids": []any{110},
			"committee_ids":      []any{1},
		},
		"committee/1": {
			"name":            "Test Committee",
			"organization_id": 1,
			"meeting_ids":     []any{110},
		},
		"meeting/110": {
			"name":                                "name_zvfbAjpZ",
			"agenda_item_creation":                 "always",
			"list_of_speakers_initially_closed":    true,
			"is_active_in_organization_id":         1,
			"committee_id":                         1,
		},
		"user/1": {
			"username":        "admin",
			"is_active":       true,
			"organization_id": 1,
		},
	})
	resp, err := tc.Request("assignment.create", map[string]any{
		"title": "test_Xcdfgee", "meeting_id": 110,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("assignment/1", map[string]any{
		"title":      "test_Xcdfgee",
		"meeting_id": 110,
	})
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("assignment.create", map[string]any{})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "required field")
}

func TestCreateFullFields(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name":               "Test Organization",
			"active_meeting_ids": []any{110},
			"committee_ids":      []any{1},
		},
		"committee/1": {
			"name":            "Test Committee",
			"organization_id": 1,
			"meeting_ids":     []any{110},
		},
		"meeting/110": {
			"name":                        "name_zvfbAjpZ",
			"agenda_item_creation":         "default_yes",
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
		},
		"user/1": {
			"username":        "admin",
			"is_active":       true,
			"organization_id": 1,
		},
	})
	resp, err := tc.Request("assignment.create", map[string]any{
		"title":                    "test_Xcdfgee",
		"meeting_id":              110,
		"description":             "text_test1",
		"open_posts":              12,
		"default_poll_description": "text_test2",
		"number_poll_candidates":  true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("assignment/1", map[string]any{
		"title":                    "test_Xcdfgee",
		"meeting_id":              110,
		"description":             "text_test1",
		"open_posts":              12,
		"default_poll_description": "text_test2",
		"number_poll_candidates":  true,
	})
}
