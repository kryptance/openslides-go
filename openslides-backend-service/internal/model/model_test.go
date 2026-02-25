package model_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func TestModelRegistry(t *testing.T) {
	all := model.All()
	if len(all) < 49 {
		t.Errorf("expected at least 49 models, got %d", len(all))
	}
}

func TestMustGetPanicsOnUnknown(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for unknown model")
		}
	}()
	model.MustGet("nonexistent_model_that_does_not_exist")
}

func TestMustGetReturnsModel(t *testing.T) {
	m := model.MustGet("topic")
	if m == nil {
		t.Fatal("expected non-nil model for 'topic'")
	}
	if m.Collection != "topic" {
		t.Errorf("expected collection 'topic', got %q", m.Collection)
	}
}

func TestGetReturnsNilForUnknown(t *testing.T) {
	m := model.Get("nonexistent_model_xyz")
	if m != nil {
		t.Fatal("expected nil for unknown model")
	}
}

func TestTopicModel(t *testing.T) {
	m := model.MustGet("topic")
	if m.Collection != "topic" {
		t.Errorf("expected collection 'topic', got %q", m.Collection)
	}
	// Check fields exist.
	fields := []string{"id", "title", "text", "meeting_id"}
	for _, f := range fields {
		if _, err := m.GetField(f); err != nil {
			t.Errorf("topic model should have field %q: %v", f, err)
		}
	}
}

func TestTopicModelFieldNotFound(t *testing.T) {
	m := model.MustGet("topic")
	_, err := m.GetField("nonexistent_field")
	if err == nil {
		t.Error("expected error for nonexistent field")
	}
}

func TestTopicModelIDField(t *testing.T) {
	m := model.MustGet("topic")
	f, err := m.GetField("id")
	if err != nil {
		t.Fatalf("expected id field: %v", err)
	}
	if f.Type != model.FieldTypeInteger {
		t.Errorf("expected id type Integer, got %v", f.Type)
	}
	if !f.Required {
		t.Error("expected id to be required")
	}
	if !f.Constant {
		t.Error("expected id to be constant")
	}
}

func TestTopicModelTitleField(t *testing.T) {
	m := model.MustGet("topic")
	f, err := m.GetField("title")
	if err != nil {
		t.Fatalf("expected title field: %v", err)
	}
	if f.Type != model.FieldTypeString {
		t.Errorf("expected title type String, got %v", f.Type)
	}
	if !f.Required {
		t.Error("expected title to be required")
	}
}

func TestTopicModelMeetingIDRelation(t *testing.T) {
	m := model.MustGet("topic")
	f, err := m.GetField("meeting_id")
	if err != nil {
		t.Fatalf("expected meeting_id field: %v", err)
	}
	if f.Type != model.FieldTypeRelation {
		t.Errorf("expected meeting_id type Relation, got %v", f.Type)
	}
	if f.Relation == nil {
		t.Fatal("expected meeting_id to have a relation definition")
	}
	if f.Relation.To.Collection != "meeting" {
		t.Errorf("expected relation target collection 'meeting', got %q", f.Relation.To.Collection)
	}
	if f.Relation.To.Field != "topic_ids" {
		t.Errorf("expected relation target field 'topic_ids', got %q", f.Relation.To.Field)
	}
}

func TestUserModel(t *testing.T) {
	m := model.MustGet("user")
	if m.Collection != "user" {
		t.Errorf("expected collection 'user', got %q", m.Collection)
	}

	fields := []string{"id", "username", "first_name", "last_name", "is_active", "email"}
	for _, f := range fields {
		if _, err := m.GetField(f); err != nil {
			t.Errorf("user model should have field %q: %v", f, err)
		}
	}
}

func TestMeetingModel(t *testing.T) {
	m := model.MustGet("meeting")
	if m.Collection != "meeting" {
		t.Errorf("expected collection 'meeting', got %q", m.Collection)
	}

	fields := []string{"id", "name", "committee_id"}
	for _, f := range fields {
		if _, err := m.GetField(f); err != nil {
			t.Errorf("meeting model should have field %q: %v", f, err)
		}
	}
}

func TestRelationFields(t *testing.T) {
	m := model.MustGet("topic")
	relFields := m.RelationFields()
	if len(relFields) == 0 {
		t.Fatal("expected topic to have relation fields")
	}
	// meeting_id should be a relation field.
	if _, ok := relFields["meeting_id"]; !ok {
		t.Error("expected 'meeting_id' to be in relation fields")
	}
}

func TestFieldDefIsRelation(t *testing.T) {
	m := model.MustGet("topic")

	// meeting_id is a relation.
	meetingField, _ := m.GetField("meeting_id")
	if !meetingField.IsRelation() {
		t.Error("expected meeting_id to be a relation")
	}

	// title is not a relation.
	titleField, _ := m.GetField("title")
	if titleField.IsRelation() {
		t.Error("expected title not to be a relation")
	}
}

func TestFieldDefIsList(t *testing.T) {
	m := model.MustGet("meeting")

	// group_ids is a RelationList.
	groupField, err := m.GetField("group_ids")
	if err != nil {
		t.Fatalf("expected group_ids field: %v", err)
	}
	if !groupField.IsList() {
		t.Error("expected group_ids to be a list field")
	}

	// committee_id is a single Relation.
	committeeField, err := m.GetField("committee_id")
	if err != nil {
		t.Fatalf("expected committee_id field: %v", err)
	}
	if committeeField.IsList() {
		t.Error("expected committee_id not to be a list field")
	}
}

func TestAllModelsHaveIDField(t *testing.T) {
	all := model.All()
	for name, m := range all {
		_, err := m.GetField("id")
		if err != nil {
			t.Errorf("model %q should have an 'id' field: %v", name, err)
		}
	}
}

func TestKnownCollections(t *testing.T) {
	collections := []string{
		"organization", "user", "meeting_user", "gender",
		"organization_tag", "theme", "committee", "meeting",
		"group", "personal_note", "tag", "agenda_item",
		"list_of_speakers", "speaker", "topic", "motion",
		"motion_submitter", "motion_comment", "motion_comment_section",
		"motion_category", "motion_block", "motion_change_recommendation",
		"motion_state", "motion_workflow", "poll", "option", "vote",
		"assignment", "assignment_candidate", "mediafile",
		"projector", "projection", "projector_message",
		"projector_countdown", "chat_group", "chat_message",
		"structure_level", "structure_level_list_of_speakers",
		"point_of_order_category", "meeting_mediafile",
	}
	for _, c := range collections {
		m := model.Get(c)
		if m == nil {
			t.Errorf("expected model %q to be registered", c)
		}
	}
}
