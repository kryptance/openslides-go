package projection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"

	// Ensure projection and projector actions are registered.
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/projection"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/projector"
)

// --- projection.create tests ---
// Note: projection.create is a BackendInternal action, so we use RequestInternal.

func TestCreateProjection(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1":    {"is_active_in_organization_id": 1},
		"assignment/1": {"meeting_id": 1},
		"projector/2":  {"meeting_id": 1},
	})
	resp, err := tc.RequestInternal("projection.create", map[string]any{
		"content_object_id":    "assignment/1",
		"current_projector_id": 2,
		"options":              map[string]any{},
		"stable":               true,
		"type":                 "test",
		"meeting_id":           1,
	})
	tc.AssertSuccess(resp, err)
	projection := tc.GetModel("projection/1")
	if projection["content_object_id"] != "assignment/1" {
		t.Errorf("expected content_object_id 'assignment/1', got %v", projection["content_object_id"])
	}
	if projection["meeting_id"] != 1 {
		t.Errorf("expected meeting_id 1, got %v", projection["meeting_id"])
	}
	if projection["stable"] != true {
		t.Errorf("expected stable true, got %v", projection["stable"])
	}
	if projection["type"] != "test" {
		t.Errorf("expected type 'test', got %v", projection["type"])
	}
}

func TestCreateProjectionMissingContentObjectID(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {"is_active_in_organization_id": 1},
	})
	resp, err := tc.RequestInternal("projection.create", map[string]any{
		"meeting_id": 1,
	})
	tc.AssertError(resp, err)
}

func TestCreateProjectionMissingMeetingID(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.RequestInternal("projection.create", map[string]any{
		"content_object_id": "assignment/1",
	})
	tc.AssertError(resp, err)
}

// --- projection.delete tests ---

func TestDeleteCurrentProjection(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"all_projection_ids": []any{12, 13, 14},
			"projector_ids":      []any{1},
		},
		"projector/1": {
			"current_projection_ids": []any{12},
			"meeting_id":             1,
			"preview_projection_ids": []any{13},
			"history_projection_ids": []any{14},
		},
		"projection/12": {"current_projector_id": 1, "meeting_id": 1},
		"projection/13": {"preview_projector_id": 1, "meeting_id": 1},
		"projection/14": {"history_projector_id": 1, "meeting_id": 1},
	})
	resp, err := tc.Request("projection.delete", map[string]any{"id": 12})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("projection/12")
}

func TestDeletePreviewProjection(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"all_projection_ids": []any{12, 13, 14},
			"projector_ids":      []any{1},
		},
		"projector/1": {
			"current_projection_ids": []any{12},
			"meeting_id":             1,
			"preview_projection_ids": []any{13},
			"history_projection_ids": []any{14},
		},
		"projection/12": {"current_projector_id": 1, "meeting_id": 1},
		"projection/13": {"preview_projector_id": 1, "meeting_id": 1},
		"projection/14": {"history_projector_id": 1, "meeting_id": 1},
	})
	resp, err := tc.Request("projection.delete", map[string]any{"id": 13})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("projection/13")
}

func TestDeleteProjectionMissingID(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("projection.delete", map[string]any{})
	tc.AssertError(resp, err)
}

// --- projection.update tests ---

func TestUpdateProjection(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1":    {"is_active_in_organization_id": 1},
		"projector/23": {"meeting_id": 1, "current_projection_ids": []any{33}},
		"projection/33": {
			"meeting_id":          1,
			"current_projector_id": 23,
		},
	})
	resp, err := tc.Request("projection.update", map[string]any{
		"id":      33,
		"weight":  11,
	})
	tc.AssertSuccess(resp, err)
	projection := tc.GetModel("projection/33")
	if projection["weight"] != 11 {
		t.Errorf("expected weight 11, got %v", projection["weight"])
	}
}

func TestUpdateProjectionOptions(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1":    {"is_active_in_organization_id": 1},
		"projector/23": {"meeting_id": 1, "current_projection_ids": []any{33}},
		"projection/33": {
			"meeting_id":          1,
			"current_projector_id": 23,
		},
	})
	resp, err := tc.Request("projection.update", map[string]any{
		"id":      33,
		"options": map[string]any{"bla": []any{}},
	})
	tc.AssertSuccess(resp, err)
	projection := tc.GetModel("projection/33")
	if projection["options"] == nil {
		t.Error("expected options to be set")
	}
}

func TestUpdateProjectionMissingID(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("projection.update", map[string]any{
		"weight": 11,
	})
	tc.AssertError(resp, err)
}

// --- projection.update_options tests ---
// Note: projection.update_options may not be implemented as a separate action in Go.
// The Python version has projection.update_options, but in Go projection.update handles options.
// We test updating options via projection.update.

func TestUpdateOptionsCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"committee/1": {"meeting_ids": []any{1}},
		"meeting/1": {
			"name":                         "bla",
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
		},
		"projector/23": {"meeting_id": 1, "current_projection_ids": []any{33}},
		"projection/33": {
			"meeting_id":          1,
			"current_projector_id": 23,
		},
	})
	resp, err := tc.Request("projection.update", map[string]any{
		"id":      33,
		"options": map[string]any{"bla": []any{}},
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("projection/33", map[string]any{})
}
