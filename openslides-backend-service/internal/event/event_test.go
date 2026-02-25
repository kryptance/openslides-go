package event_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/event"
)

func TestFQID(t *testing.T) {
	fqid := event.FQID("topic", 42)
	expected := "topic/42"
	if fqid != expected {
		t.Errorf("expected %q, got %q", expected, fqid)
	}
}

func TestFQIDLargeID(t *testing.T) {
	fqid := event.FQID("meeting", 100000)
	expected := "meeting/100000"
	if fqid != expected {
		t.Errorf("expected %q, got %q", expected, fqid)
	}
}

func TestFQField(t *testing.T) {
	fqf := event.FQField("topic", 42, "title")
	expected := "topic/42/title"
	if fqf != expected {
		t.Errorf("expected %q, got %q", expected, fqf)
	}
}

func TestFQFieldNestedField(t *testing.T) {
	fqf := event.FQField("meeting", 1, "committee_id")
	expected := "meeting/1/committee_id"
	if fqf != expected {
		t.Errorf("expected %q, got %q", expected, fqf)
	}
}

func TestEventTypes(t *testing.T) {
	if event.TypeCreate != "create" {
		t.Errorf("expected TypeCreate 'create', got %q", event.TypeCreate)
	}
	if event.TypeUpdate != "update" {
		t.Errorf("expected TypeUpdate 'update', got %q", event.TypeUpdate)
	}
	if event.TypeDelete != "delete" {
		t.Errorf("expected TypeDelete 'delete', got %q", event.TypeDelete)
	}
}

func TestNewWriteRequest(t *testing.T) {
	wr := event.NewWriteRequest(42)
	if wr.UserID != 42 {
		t.Errorf("expected UserID 42, got %d", wr.UserID)
	}
	if wr.Information == nil {
		t.Error("expected non-nil Information map")
	}
	if wr.LockedFields == nil {
		t.Error("expected non-nil LockedFields map")
	}
	if len(wr.Events) != 0 {
		t.Errorf("expected 0 events, got %d", len(wr.Events))
	}
}

func TestWriteRequestAddEvent(t *testing.T) {
	wr := event.NewWriteRequest(1)
	e := event.Event{
		Type: event.TypeCreate,
		FQID: "topic/1",
		Fields: map[string]any{
			"title": "Test Topic",
		},
	}
	wr.AddEvent(e)

	if len(wr.Events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(wr.Events))
	}
	if wr.Events[0].Type != event.TypeCreate {
		t.Errorf("expected create event, got %s", wr.Events[0].Type)
	}
	if wr.Events[0].FQID != "topic/1" {
		t.Errorf("expected FQID 'topic/1', got %q", wr.Events[0].FQID)
	}
}

func TestWriteRequestAddInformation(t *testing.T) {
	wr := event.NewWriteRequest(1)
	wr.AddInformation("topic/1", "Created", "via API")

	info, ok := wr.Information["topic/1"]
	if !ok {
		t.Fatal("expected information for topic/1")
	}
	if len(info) != 2 {
		t.Errorf("expected 2 information entries, got %d", len(info))
	}
	if info[0] != "Created" || info[1] != "via API" {
		t.Errorf("unexpected information entries: %v", info)
	}
}

func TestWriteRequestAddInformationAppends(t *testing.T) {
	wr := event.NewWriteRequest(1)
	wr.AddInformation("topic/1", "first")
	wr.AddInformation("topic/1", "second")

	if len(wr.Information["topic/1"]) != 2 {
		t.Errorf("expected 2 information entries, got %d", len(wr.Information["topic/1"]))
	}
}

func TestWriteRequestMerge(t *testing.T) {
	wr1 := event.NewWriteRequest(1)
	wr1.AddEvent(event.Event{Type: event.TypeCreate, FQID: "topic/1"})
	wr1.AddInformation("topic/1", "Created")

	wr2 := event.NewWriteRequest(1)
	wr2.AddEvent(event.Event{Type: event.TypeCreate, FQID: "motion/1"})
	wr2.AddInformation("motion/1", "Created")
	wr2.LockedFields["motion/1"] = map[string]int{"title": 5}

	wr1.Merge(wr2)

	if len(wr1.Events) != 2 {
		t.Errorf("expected 2 events after merge, got %d", len(wr1.Events))
	}

	if _, ok := wr1.Information["motion/1"]; !ok {
		t.Error("expected information for motion/1 after merge")
	}

	if _, ok := wr1.LockedFields["motion/1"]; !ok {
		t.Error("expected locked fields for motion/1 after merge")
	}
}

func TestWriteRequestMergeLockedFieldsMerge(t *testing.T) {
	wr1 := event.NewWriteRequest(1)
	wr1.LockedFields["topic/1"] = map[string]int{"title": 3}

	wr2 := event.NewWriteRequest(1)
	wr2.LockedFields["topic/1"] = map[string]int{"text": 5}

	wr1.Merge(wr2)

	if wr1.LockedFields["topic/1"]["title"] != 3 {
		t.Error("expected original locked field to be preserved")
	}
	if wr1.LockedFields["topic/1"]["text"] != 5 {
		t.Error("expected merged locked field to be present")
	}
}

func TestEventStruct(t *testing.T) {
	e := event.Event{
		Type:   event.TypeUpdate,
		FQID:   "topic/1",
		Fields: map[string]any{"title": "Updated"},
	}

	if e.Type != event.TypeUpdate {
		t.Errorf("expected update type, got %s", e.Type)
	}
	if e.FQID != "topic/1" {
		t.Errorf("expected FQID 'topic/1', got %q", e.FQID)
	}
	if e.Fields["title"] != "Updated" {
		t.Errorf("expected title 'Updated', got %v", e.Fields["title"])
	}
}

func TestEventListFields(t *testing.T) {
	lf := &event.ListFields{
		Add:    map[string][]any{"group_ids": {1, 2}},
		Remove: map[string][]any{"group_ids": {3}},
	}
	e := event.Event{
		Type:       event.TypeUpdate,
		FQID:       "meeting_user/1",
		ListFields: lf,
	}

	if e.ListFields == nil {
		t.Fatal("expected non-nil ListFields")
	}
	if len(e.ListFields.Add["group_ids"]) != 2 {
		t.Errorf("expected 2 add values, got %d", len(e.ListFields.Add["group_ids"]))
	}
	if len(e.ListFields.Remove["group_ids"]) != 1 {
		t.Errorf("expected 1 remove value, got %d", len(e.ListFields.Remove["group_ids"]))
	}
}

func TestWriteRequestMergeEmpty(t *testing.T) {
	wr1 := event.NewWriteRequest(1)
	wr1.AddEvent(event.Event{Type: event.TypeCreate, FQID: "topic/1"})

	wr2 := event.NewWriteRequest(2)
	wr1.Merge(wr2)

	if len(wr1.Events) != 1 {
		t.Errorf("expected 1 event after merging empty, got %d", len(wr1.Events))
	}
}
