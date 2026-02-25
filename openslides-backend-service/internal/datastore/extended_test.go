package datastore_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/datastore"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
)

func TestNewExtended(t *testing.T) {
	ds := datastore.NewExtended()
	if ds == nil {
		t.Fatal("expected non-nil Extended datastore")
	}
}

func TestApplyChangedModel(t *testing.T) {
	ds := datastore.NewExtended()
	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "Test"})

	pm, found := ds.GetChangedModel("topic", 1)
	if !found {
		t.Fatal("expected topic/1 to exist")
	}
	if pm["title"] != "Test" {
		t.Errorf("expected title 'Test', got %v", pm["title"])
	}
	if pm["id"] != 1 {
		t.Errorf("expected id 1, got %v", pm["id"])
	}
}

func TestApplyChangedModelMerges(t *testing.T) {
	ds := datastore.NewExtended()
	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "First"})
	ds.ApplyChangedModel("topic", 1, map[string]any{"title": "Second", "text": "body"})

	pm, found := ds.GetChangedModel("topic", 1)
	if !found {
		t.Fatal("expected topic/1 to exist")
	}
	if pm["title"] != "Second" {
		t.Errorf("expected title 'Second', got %v", pm["title"])
	}
	if pm["text"] != "body" {
		t.Errorf("expected text 'body', got %v", pm["text"])
	}
	if pm["id"] != 1 {
		t.Errorf("expected id 1, got %v", pm["id"])
	}
}

func TestGetChangedModelNotFound(t *testing.T) {
	ds := datastore.NewExtended()
	_, found := ds.GetChangedModel("topic", 999)
	if found {
		t.Error("expected topic/999 to not be found")
	}
}

func TestGetChangedModelWrongCollection(t *testing.T) {
	ds := datastore.NewExtended()
	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "Test"})

	_, found := ds.GetChangedModel("motion", 1)
	if found {
		t.Error("expected motion/1 to not be found")
	}
}

func TestApplyToBeDeleted(t *testing.T) {
	ds := datastore.NewExtended()
	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "Test"})

	if ds.IsDeleted("topic", 1) {
		t.Error("topic/1 should not be deleted yet")
	}

	ds.ApplyToBeDeleted("topic", 1)

	if !ds.IsDeleted("topic", 1) {
		t.Error("topic/1 should be deleted")
	}
}

func TestIsDeletedOnNonexistent(t *testing.T) {
	ds := datastore.NewExtended()
	if ds.IsDeleted("topic", 999) {
		t.Error("nonexistent model should not be marked as deleted")
	}
}

func TestReserveIDs(t *testing.T) {
	ds := datastore.NewExtended()

	ids, err := ds.ReserveIDs("topic", 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ids) != 3 {
		t.Fatalf("expected 3 ids, got %d", len(ids))
	}
	if ids[0] != 1 || ids[1] != 2 || ids[2] != 3 {
		t.Errorf("expected [1, 2, 3], got %v", ids)
	}
}

func TestReserveIDsSequential(t *testing.T) {
	ds := datastore.NewExtended()

	ids1, _ := ds.ReserveIDs("topic", 2)
	ids2, _ := ds.ReserveIDs("topic", 2)

	if ids1[0] != 1 || ids1[1] != 2 {
		t.Errorf("first batch expected [1, 2], got %v", ids1)
	}
	if ids2[0] != 3 || ids2[1] != 4 {
		t.Errorf("second batch expected [3, 4], got %v", ids2)
	}
}

func TestReserveIDsPerCollection(t *testing.T) {
	ds := datastore.NewExtended()

	topicIDs, _ := ds.ReserveIDs("topic", 2)
	motionIDs, _ := ds.ReserveIDs("motion", 2)

	// Each collection starts from 1 independently.
	if topicIDs[0] != 1 {
		t.Errorf("expected topic id 1, got %v", topicIDs[0])
	}
	if motionIDs[0] != 1 {
		t.Errorf("expected motion id 1, got %v", motionIDs[0])
	}
}

func TestReset(t *testing.T) {
	ds := datastore.NewExtended()
	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "Test"})
	ds.ApplyToBeDeleted("motion", 1)

	ds.Reset()

	_, found := ds.GetChangedModel("topic", 1)
	if found {
		t.Error("expected topic/1 to be cleared after reset")
	}

	if ds.IsDeleted("motion", 1) {
		t.Error("expected motion/1 deletion flag to be cleared after reset")
	}
}

func TestGet(t *testing.T) {
	ds := datastore.NewExtended()
	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "Test", "text": "body"})

	pm, err := ds.Get("topic", 1, []string{"title", "text"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pm["title"] != "Test" {
		t.Errorf("expected title 'Test', got %v", pm["title"])
	}
	if pm["text"] != "body" {
		t.Errorf("expected text 'body', got %v", pm["text"])
	}
}

func TestGetDeletedModel(t *testing.T) {
	ds := datastore.NewExtended()
	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "Test"})
	ds.ApplyToBeDeleted("topic", 1)

	_, err := ds.Get("topic", 1, []string{"title"})
	if err == nil {
		t.Fatal("expected error when getting deleted model")
	}
}

func TestGetFieldSubset(t *testing.T) {
	ds := datastore.NewExtended()
	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "Test", "text": "body"})

	pm, err := ds.Get("topic", 1, []string{"title"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pm["title"] != "Test" {
		t.Errorf("expected title 'Test', got %v", pm["title"])
	}
	// text was not requested.
	if _, ok := pm["text"]; ok {
		t.Error("expected text field to not be returned when not requested")
	}
}

func TestGetMany(t *testing.T) {
	ds := datastore.NewExtended()
	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "Topic 1"})
	ds.ApplyChangedModel("topic", 2, map[string]any{"id": 2, "title": "Topic 2"})

	result, err := ds.GetMany([]datastore.GetManyRequest{
		{Collection: "topic", IDs: []int{1, 2}, Fields: []string{"title"}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result["topic"]) != 2 {
		t.Errorf("expected 2 topic results, got %d", len(result["topic"]))
	}
}

func TestBuildWriteRequest(t *testing.T) {
	ds := datastore.NewExtended()
	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "Test", "meta_new": true})

	wr := ds.BuildWriteRequest(42)
	if wr.UserID != 42 {
		t.Errorf("expected UserID 42, got %d", wr.UserID)
	}
	if len(wr.Events) == 0 {
		t.Fatal("expected at least one event")
	}

	// The model with meta_new should produce a create event.
	found := false
	for _, e := range wr.Events {
		if e.Type == event.TypeCreate && e.FQID == "topic/1" {
			found = true
			// meta_new should not appear in event fields.
			if _, ok := e.Fields["meta_new"]; ok {
				t.Error("meta_new should not be in event fields")
			}
		}
	}
	if !found {
		t.Error("expected create event for topic/1")
	}
}

func TestBuildWriteRequestUpdate(t *testing.T) {
	ds := datastore.NewExtended()
	// No meta_new means it's an update.
	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "Updated"})

	wr := ds.BuildWriteRequest(1)
	found := false
	for _, e := range wr.Events {
		if e.Type == event.TypeUpdate && e.FQID == "topic/1" {
			found = true
		}
	}
	if !found {
		t.Error("expected update event for topic/1")
	}
}

func TestBuildWriteRequestDelete(t *testing.T) {
	ds := datastore.NewExtended()
	ds.ApplyToBeDeleted("topic", 1)

	wr := ds.BuildWriteRequest(1)
	found := false
	for _, e := range wr.Events {
		if e.Type == event.TypeDelete && e.FQID == "topic/1" {
			found = true
		}
	}
	if !found {
		t.Error("expected delete event for topic/1")
	}
}

func TestBuildWriteRequestDeletedModelNotCreated(t *testing.T) {
	ds := datastore.NewExtended()
	// Model that has changes AND is deleted: should produce delete, not create/update.
	ds.ApplyChangedModel("topic", 1, map[string]any{"id": 1, "title": "Test"})
	ds.ApplyToBeDeleted("topic", 1)

	wr := ds.BuildWriteRequest(1)
	for _, e := range wr.Events {
		if e.FQID == "topic/1" && (e.Type == event.TypeCreate || e.Type == event.TypeUpdate) {
			t.Errorf("deleted model should not produce %s event", e.Type)
		}
	}
}
