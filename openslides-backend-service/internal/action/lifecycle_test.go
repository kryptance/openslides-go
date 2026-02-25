package action_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/datastore"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-backend-service/internal/relation"

	// Ensure action packages are registered.
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/topic"
)

func TestPerformCreateAction(t *testing.T) {
	a := action.NewCreateAction("test.create", model.MustGet("topic"))
	ds := datastore.NewExtended()
	rm := relation.NewManager()

	ds.ApplyChangedModel("meeting", 1, map[string]any{"id": 1, "name": "test"})

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{"meeting_id": 1, "title": "test topic"}
	events, results, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(events) == 0 {
		t.Fatal("expected events")
	}
	if len(results) == 0 {
		t.Fatal("expected results")
	}
	// Verify the event is a create event.
	if events[0].Type != event.TypeCreate {
		t.Errorf("expected create event, got %s", events[0].Type)
	}
}

func TestPerformCreateActionAssignsID(t *testing.T) {
	a := action.NewCreateAction("test.create", model.MustGet("topic"))
	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{"meeting_id": 1, "title": "auto id topic"}
	_, results, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) == 0 {
		t.Fatal("expected results")
	}
	id, ok := results[0]["id"]
	if !ok {
		t.Fatal("expected result to contain id")
	}
	if id == nil || id == 0 {
		t.Fatal("expected non-zero id in result")
	}
}

func TestPerformCreateActionAppliesDefaults(t *testing.T) {
	// Use organization model which has Default values set (e.g. default_language = "en").
	a := action.NewCreateAction("test.create.org", model.MustGet("organization"))
	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{"name": "Test Org"}
	_, _, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check that the created model in the datastore has defaults applied.
	pm, found := ds.GetChangedModel("organization", 1)
	if !found {
		t.Fatal("expected organization model to exist in datastore")
	}
	if pm["default_language"] != "en" {
		t.Errorf("expected default_language to be 'en', got %v", pm["default_language"])
	}
}

func TestPerformUpdateAction(t *testing.T) {
	a := action.NewUpdateAction("test.update", model.MustGet("topic"))
	ds := datastore.NewExtended()
	rm := relation.NewManager()

	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "old", "meeting_id": 1})

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{"id": 1, "title": "new title"}
	events, _, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(events) == 0 {
		t.Fatal("expected events")
	}
	if events[0].Type != event.TypeUpdate {
		t.Errorf("expected update event, got %s", events[0].Type)
	}
}

func TestPerformUpdateActionUpdatesDatastore(t *testing.T) {
	a := action.NewUpdateAction("test.update", model.MustGet("topic"))
	ds := datastore.NewExtended()
	rm := relation.NewManager()

	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "old", "meeting_id": 1})

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{"id": 1, "title": "new title"}
	_, _, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	pm, found := ds.GetChangedModel("topic", 1)
	if !found {
		t.Fatal("expected topic model to exist in datastore")
	}
	if pm["title"] != "new title" {
		t.Errorf("expected title 'new title', got %v", pm["title"])
	}
}

func TestPerformUpdateActionNoChanges(t *testing.T) {
	a := action.NewUpdateAction("test.update", model.MustGet("topic"))
	ds := datastore.NewExtended()
	rm := relation.NewManager()

	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "old", "meeting_id": 1})

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	// Only id, no field changes. CreateEvents returns nil, nil for empty fields.
	instance := action.Instance{"id": 1}
	events, _, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// The update event should have no fields, so it should produce no event.
	if len(events) != 0 {
		t.Errorf("expected no events for update with no changed fields, got %d", len(events))
	}
}

func TestPerformDeleteAction(t *testing.T) {
	a := action.NewDeleteAction("test.delete", model.MustGet("topic"))
	ds := datastore.NewExtended()
	rm := relation.NewManager()

	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "test", "meeting_id": 1})

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{"id": 1}
	events, _, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(events) == 0 {
		t.Fatal("expected events")
	}
	if events[0].Type != event.TypeDelete {
		t.Errorf("expected delete event, got %s", events[0].Type)
	}
}

func TestPerformDeleteActionMarksDeleted(t *testing.T) {
	a := action.NewDeleteAction("test.delete", model.MustGet("topic"))
	ds := datastore.NewExtended()
	rm := relation.NewManager()

	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "test", "meeting_id": 1})

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{"id": 1}
	_, _, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !ds.IsDeleted("topic", 1) {
		t.Fatal("expected topic/1 to be marked as deleted")
	}
}

func TestSchemaValidationRequired(t *testing.T) {
	a := action.NewCreateAction("test.create.schema", model.MustGet("topic"))
	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": map[string]any{"type": "string"},
		},
		"required": []string{"title"},
	}
	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	// Missing required field "title".
	instance := action.Instance{"meeting_id": 1}
	_, _, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err == nil {
		t.Fatal("expected schema validation error for missing required field")
	}
}

func TestSchemaValidationValid(t *testing.T) {
	a := action.NewCreateAction("test.create.schema.valid", model.MustGet("topic"))
	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title":      map[string]any{"type": "string"},
			"meeting_id": map[string]any{"type": "integer"},
		},
		"required": []string{"title"},
	}
	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{"title": "hello", "meeting_id": 1}
	_, _, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPermissionCheckSkippedForInternal(t *testing.T) {
	a := action.NewBaseAction("test.perm", model.MustGet("topic"))
	permCalled := false
	a.CheckPermissions = func(ctx context.Context, params *action.ActionParams, instance action.Instance) error {
		permCalled = true
		return nil
	}

	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{"id": 1}
	_, _, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if permCalled {
		t.Error("permission check should be skipped for internal calls")
	}
}

func TestPermissionCheckCalledForNonInternal(t *testing.T) {
	a := action.NewBaseAction("test.perm.ext", model.MustGet("topic"))
	permCalled := false
	a.CheckPermissions = func(ctx context.Context, params *action.ActionParams, instance action.Instance) error {
		permCalled = true
		return nil
	}

	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          1,
		Internal:        false,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{"id": 1}
	_, _, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !permCalled {
		t.Error("permission check should be called for non-internal calls")
	}
}

func TestMultipleInstances(t *testing.T) {
	a := action.NewCreateAction("test.multi", model.MustGet("topic"))
	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	instances := []action.Instance{
		{"meeting_id": 1, "title": "topic 1"},
		{"meeting_id": 1, "title": "topic 2"},
		{"meeting_id": 1, "title": "topic 3"},
	}
	events, results, err := action.Perform(context.Background(), a, params, instances)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Should get at least 3 create events (one per instance) plus relation events.
	createEvents := 0
	for _, e := range events {
		if e.Type == event.TypeCreate {
			createEvents++
		}
	}
	if createEvents != 3 {
		t.Errorf("expected 3 create events, got %d", createEvents)
	}
	if len(results) != 3 {
		t.Errorf("expected 3 results, got %d", len(results))
	}
}
