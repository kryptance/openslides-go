package relation_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-backend-service/internal/relation"
)

func TestNewManager(t *testing.T) {
	m := relation.NewManager()
	if m == nil {
		t.Fatal("expected non-nil Manager")
	}
}

func TestHandleRelationUpdatesNoRelations(t *testing.T) {
	m := relation.NewManager()

	// Create a model def with no relation fields.
	modelDef := &model.ModelDef{
		Collection: "test",
		Fields: map[string]*model.FieldDef{
			"id":    {Type: model.FieldTypeInteger},
			"title": {Type: model.FieldTypeString},
		},
	}

	updates := m.HandleRelationUpdates(modelDef, 1, map[string]any{
		"id":    1,
		"title": "test",
	})

	if len(updates) != 0 {
		t.Errorf("expected no updates for non-relation fields, got %d", len(updates))
	}
}

func TestHandleRelationUpdatesSingleRelation(t *testing.T) {
	m := relation.NewManager()

	// topic has meeting_id -> meeting/topic_ids relation.
	topicModel := model.MustGet("topic")

	updates := m.HandleRelationUpdates(topicModel, 1, map[string]any{
		"id":         1,
		"meeting_id": 1,
	})

	// Should have a reverse update for meeting/1 adding topic/1 to topic_ids.
	found := false
	for _, u := range updates {
		if u.FQID == "meeting/1" && u.Field == "topic_ids" && u.Type == relation.UpdateTypeAdd {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected reverse relation update for meeting/1 topic_ids")
	}
}

func TestHandleRelationUpdatesListRelation(t *testing.T) {
	m := relation.NewManager()

	// meeting has group_ids relation to group.
	meetingModel := model.MustGet("meeting")

	updates := m.HandleRelationUpdates(meetingModel, 1, map[string]any{
		"id":        1,
		"group_ids": []any{1, 2, 3},
	})

	// Should have reverse updates for group/1, group/2, group/3.
	addCount := 0
	for _, u := range updates {
		if u.Type == relation.UpdateTypeAdd && u.Field == "meeting_id" {
			addCount++
		}
	}
	if addCount != 3 {
		t.Errorf("expected 3 reverse relation add updates, got %d", addCount)
	}
}

func TestGetEventsEmpty(t *testing.T) {
	m := relation.NewManager()

	events := m.GetEvents()
	if len(events) != 0 {
		t.Errorf("expected 0 events from empty manager, got %d", len(events))
	}
}

func TestGetEventsAfterRelationUpdate(t *testing.T) {
	m := relation.NewManager()

	topicModel := model.MustGet("topic")
	m.HandleRelationUpdates(topicModel, 1, map[string]any{
		"id":         1,
		"meeting_id": 1,
	})

	events := m.GetEvents()
	if len(events) == 0 {
		t.Fatal("expected relation events")
	}

	// All relation events should be update events.
	for _, e := range events {
		if e.Type != event.TypeUpdate {
			t.Errorf("expected update event, got %s", e.Type)
		}
	}
}

func TestGetEventsListFieldsStructure(t *testing.T) {
	m := relation.NewManager()

	topicModel := model.MustGet("topic")
	m.HandleRelationUpdates(topicModel, 1, map[string]any{
		"id":         1,
		"meeting_id": 1,
	})

	events := m.GetEvents()
	// Find the event for meeting/1.
	found := false
	for _, e := range events {
		if e.FQID == "meeting/1" && e.ListFields != nil {
			if addVals, ok := e.ListFields.Add["topic_ids"]; ok && len(addVals) > 0 {
				found = true
			}
		}
	}
	if !found {
		t.Error("expected list field add event for meeting/1 topic_ids")
	}
}

func TestReset(t *testing.T) {
	m := relation.NewManager()

	topicModel := model.MustGet("topic")
	m.HandleRelationUpdates(topicModel, 1, map[string]any{
		"id":         1,
		"meeting_id": 1,
	})

	m.Reset()

	events := m.GetEvents()
	if len(events) != 0 {
		t.Errorf("expected 0 events after reset, got %d", len(events))
	}
}

func TestHandleRelationUpdatesFieldNotInInstance(t *testing.T) {
	m := relation.NewManager()

	// topic has meeting_id relation, but instance doesn't include it.
	topicModel := model.MustGet("topic")
	updates := m.HandleRelationUpdates(topicModel, 1, map[string]any{
		"id":    1,
		"title": "hello",
	})

	// No relation field in instance, so no updates.
	if len(updates) != 0 {
		t.Errorf("expected no updates when no relation fields in instance, got %d", len(updates))
	}
}

func TestMultipleHandleRelationUpdates(t *testing.T) {
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
	// Both topic 1 and 2 reference meeting 1. The manager should group
	// these into one event for meeting/1.
	meetingEvents := 0
	for _, e := range events {
		if e.FQID == "meeting/1" {
			meetingEvents++
		}
	}
	if meetingEvents != 1 {
		t.Errorf("expected 1 grouped event for meeting/1, got %d", meetingEvents)
	}

	// The grouped event should have 2 values in the add list for topic_ids.
	for _, e := range events {
		if e.FQID == "meeting/1" && e.ListFields != nil {
			addVals := e.ListFields.Add["topic_ids"]
			if len(addVals) != 2 {
				t.Errorf("expected 2 add values for topic_ids, got %d", len(addVals))
			}
		}
	}
}
