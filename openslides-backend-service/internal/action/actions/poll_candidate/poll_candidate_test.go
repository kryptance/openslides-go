package poll_candidate_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"

	// Ensure poll_candidate actions are registered.
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/poll_candidate"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/poll_candidate_list"
)

func TestPollCandidateCreate(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"name":                      "meeting_1",
			"is_active_in_organization_id": 1,
			"poll_candidate_list_ids":    []any{2},
		},
		"poll_candidate_list/2": {"meeting_id": 1},
		"user/1": {
			"username":  "admin",
			"is_active": true,
		},
	})
	resp, err := tc.RequestInternal("poll_candidate.create", map[string]any{
		"user_id":               1,
		"poll_candidate_list_id": 2,
		"weight":                12,
		"meeting_id":            1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll_candidate/1", map[string]any{
		"meeting_id":             1,
		"user_id":               1,
		"poll_candidate_list_id": 2,
		"weight":                12,
	})
}

func TestPollCandidateCreateMissingWeight(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"name":                      "meeting_1",
			"is_active_in_organization_id": 1,
		},
		"poll_candidate_list/2": {"meeting_id": 1},
	})
	resp, err := tc.RequestInternal("poll_candidate.create", map[string]any{
		"user_id":               1,
		"poll_candidate_list_id": 2,
		"meeting_id":            1,
	})
	tc.AssertError(resp, err)
}

func TestPollCandidateDelete(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"name":                      "meeting_1",
			"poll_candidate_list_ids":    []any{2},
			"poll_candidate_ids":         []any{3},
			"is_active_in_organization_id": 1,
		},
		"user/1":                {"poll_candidate_ids": []any{3}},
		"poll_candidate_list/2": {"meeting_id": 1, "poll_candidate_ids": []any{3}},
		"poll_candidate/3": {
			"meeting_id":             1,
			"poll_candidate_list_id": 2,
			"user_id":               1,
			"weight":                1,
		},
	})
	resp, err := tc.RequestInternal("poll_candidate.delete", map[string]any{"id": 3})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("poll_candidate/3")
}

func TestPollCandidateDeleteMissingID(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.RequestInternal("poll_candidate.delete", map[string]any{})
	tc.AssertError(resp, err)
}
