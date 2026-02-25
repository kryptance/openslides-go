package relation_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-backend-service/internal/relation"
)

// Tests migrated from test_create_relation.py.
// The Python tests use fake models to test relation creation behavior.
// Here we test the Go relation manager with real models.

func TestCreateRelationSingleField(t *testing.T) {
	// When creating a topic with meeting_id, the relation manager should
	// produce a reverse update for meeting/topic_ids.
	m := relation.NewManager()
	topicModel := model.MustGet("topic")

	updates := m.HandleRelationUpdates(topicModel, 1, map[string]any{
		"id":         1,
		"meeting_id": 1,
		"title":      "test",
	})

	foundMeetingUpdate := false
	for _, u := range updates {
		if u.FQID == "meeting/1" && u.Field == "topic_ids" && u.Type == relation.UpdateTypeAdd {
			foundMeetingUpdate = true
		}
	}
	if !foundMeetingUpdate {
		t.Error("expected reverse relation update for meeting/1 topic_ids when creating topic with meeting_id")
	}
}

func TestCreateRelationListField(t *testing.T) {
	// When setting group_ids on a meeting, the relation manager should
	// produce reverse updates for each group's meeting_id.
	m := relation.NewManager()
	meetingModel := model.MustGet("meeting")

	updates := m.HandleRelationUpdates(meetingModel, 1, map[string]any{
		"id":        1,
		"group_ids": []any{1, 2},
	})

	addCount := 0
	for _, u := range updates {
		if u.Type == relation.UpdateTypeAdd && u.Field == "meeting_id" {
			addCount++
		}
	}
	if addCount != 2 {
		t.Errorf("expected 2 reverse relation add updates for groups, got %d", addCount)
	}
}

func TestCreateRelationNoRelationFields(t *testing.T) {
	// When the instance has no relation fields, no updates should be produced.
	m := relation.NewManager()
	topicModel := model.MustGet("topic")

	updates := m.HandleRelationUpdates(topicModel, 1, map[string]any{
		"id":    1,
		"title": "no relations",
	})

	if len(updates) != 0 {
		t.Errorf("expected no relation updates for non-relation fields, got %d", len(updates))
	}
}

func TestCreateRelationProducesEvents(t *testing.T) {
	// After handling relation updates, the manager should produce update events.
	m := relation.NewManager()
	topicModel := model.MustGet("topic")

	m.HandleRelationUpdates(topicModel, 1, map[string]any{
		"id":         1,
		"meeting_id": 1,
	})

	events := m.GetEvents()
	if len(events) == 0 {
		t.Fatal("expected events after handling relation updates")
	}
}

func TestCreateRelationMultipleCreatesGrouped(t *testing.T) {
	// Creating multiple topics for the same meeting should group updates.
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
}
