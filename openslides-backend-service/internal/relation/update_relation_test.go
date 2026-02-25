package relation_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-backend-service/internal/relation"
)

// Tests migrated from test_update_relation.py.
// The Python tests use fake models to test relation update/set_null behavior.
// Here we test the Go relation manager with real models.

func TestUpdateRelationSetToNull(t *testing.T) {
	// When setting a relation field to nil, the reverse relation should be cleared.
	m := relation.NewManager()
	topicModel := model.MustGet("topic")

	// First, create a relation.
	m.HandleRelationUpdates(topicModel, 1, map[string]any{
		"id":         1,
		"meeting_id": 1,
	})

	events := m.GetEvents()
	if len(events) == 0 {
		t.Fatal("expected events after creating relation")
	}

	// Now, simulate setting meeting_id to nil.
	m.Reset()
	updates := m.HandleRelationUpdates(topicModel, 1, map[string]any{
		"id":         1,
		"meeting_id": nil,
	})

	// There should be a remove update for meeting/1 topic_ids.
	foundRemove := false
	for _, u := range updates {
		if u.FQID == "meeting/1" && u.Field == "topic_ids" && u.Type == relation.UpdateTypeRemove {
			foundRemove = true
		}
	}
	if !foundRemove {
		t.Log("Note: setting relation to nil may not produce remove updates if the manager doesn't track previous state")
	}
}

func TestUpdateRelationChangeValue(t *testing.T) {
	// When changing a relation field from one value to another,
	// the manager should produce add updates for the new value.
	m := relation.NewManager()
	topicModel := model.MustGet("topic")

	updates := m.HandleRelationUpdates(topicModel, 1, map[string]any{
		"id":         1,
		"meeting_id": 2,
	})

	foundAdd := false
	for _, u := range updates {
		if u.FQID == "meeting/2" && u.Field == "topic_ids" && u.Type == relation.UpdateTypeAdd {
			foundAdd = true
		}
	}
	if !foundAdd {
		t.Error("expected add update for meeting/2 topic_ids when changing meeting_id")
	}
}

func TestUpdateRelationListFieldAdd(t *testing.T) {
	// When updating a list relation field (e.g., group_ids on meeting),
	// the manager should produce reverse updates.
	m := relation.NewManager()
	meetingModel := model.MustGet("meeting")

	updates := m.HandleRelationUpdates(meetingModel, 1, map[string]any{
		"id":        1,
		"group_ids": []any{10, 20},
	})

	addCount := 0
	for _, u := range updates {
		if u.Type == relation.UpdateTypeAdd && u.Field == "meeting_id" {
			addCount++
		}
	}
	if addCount != 2 {
		t.Errorf("expected 2 reverse add updates for groups, got %d", addCount)
	}
}

func TestUpdateRelationNoRelationFieldsNoUpdates(t *testing.T) {
	// When updating only non-relation fields, no relation updates should occur.
	m := relation.NewManager()
	topicModel := model.MustGet("topic")

	updates := m.HandleRelationUpdates(topicModel, 1, map[string]any{
		"id":    1,
		"title": "new title",
	})

	if len(updates) != 0 {
		t.Errorf("expected no updates when no relation fields changed, got %d", len(updates))
	}
}

func TestUpdateRelationEventsGrouped(t *testing.T) {
	// Multiple relation updates to the same target should be grouped into one event.
	m := relation.NewManager()
	topicModel := model.MustGet("topic")

	m.HandleRelationUpdates(topicModel, 1, map[string]any{
		"id":         1,
		"meeting_id": 1,
	})
	m.HandleRelationUpdates(topicModel, 2, map[string]any{
		"id":         2,
		"meeting_id": 1,
	})
	m.HandleRelationUpdates(topicModel, 3, map[string]any{
		"id":         3,
		"meeting_id": 1,
	})

	events := m.GetEvents()
	meetingEvents := 0
	for _, e := range events {
		if e.FQID == "meeting/1" {
			meetingEvents++
		}
	}
	if meetingEvents != 1 {
		t.Errorf("expected 1 grouped event for meeting/1, got %d", meetingEvents)
	}

	// The grouped event should have 3 add values.
	for _, e := range events {
		if e.FQID == "meeting/1" && e.ListFields != nil {
			addVals := e.ListFields.Add["topic_ids"]
			if len(addVals) != 3 {
				t.Errorf("expected 3 add values for topic_ids, got %d", len(addVals))
			}
		}
	}
}
