package poll_candidate_list_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"

	// Ensure poll_candidate_list actions are registered.
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/poll_candidate"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/poll_candidate_list"
)

func TestPollCandidateListCreate(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"name":                      "meeting_1",
			"is_active_in_organization_id": 1,
		},
		"option/4": {
			"meeting_id": 1,
		},
	})
	resp, err := tc.RequestInternal("poll_candidate_list.create", map[string]any{
		"option_id":  4,
		"meeting_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll_candidate_list/1", map[string]any{
		"option_id":  4,
		"meeting_id": 1,
	})
}

func TestPollCandidateListCreateMissingOptionID(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"name":                      "meeting_1",
			"is_active_in_organization_id": 1,
		},
	})
	resp, err := tc.RequestInternal("poll_candidate_list.create", map[string]any{
		"meeting_id": 1,
	})
	tc.AssertError(resp, err)
}

func TestPollCandidateListDelete(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"name":                      "meeting_1",
			"poll_candidate_list_ids":    []any{2},
			"poll_candidate_ids":         []any{3, 4, 5},
			"is_active_in_organization_id": 1,
		},
		"user/1":  {"poll_candidate_ids": []any{3}},
		"user/2":  {"username": "test1", "poll_candidate_ids": []any{4}},
		"user/3":  {"username": "test2", "poll_candidate_ids": []any{5}},
		"poll_candidate_list/2": {
			"meeting_id":          1,
			"poll_candidate_ids":  []any{3, 4, 5},
		},
		"poll_candidate/3": {
			"meeting_id":             1,
			"poll_candidate_list_id": 2,
			"user_id":               1,
			"weight":                1,
		},
		"poll_candidate/4": {
			"meeting_id":             1,
			"poll_candidate_list_id": 2,
			"user_id":               2,
			"weight":                2,
		},
		"poll_candidate/5": {
			"meeting_id":             1,
			"poll_candidate_list_id": 2,
			"user_id":               3,
			"weight":                3,
		},
	})
	resp, err := tc.RequestInternal("poll_candidate_list.delete", map[string]any{"id": 2})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("poll_candidate_list/2")
}

func TestPollCandidateListDeleteMissingID(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.RequestInternal("poll_candidate_list.delete", map[string]any{})
	tc.AssertError(resp, err)
}
