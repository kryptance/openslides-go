package action_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action"

	// Import all action packages to trigger their init() registration.
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/topic"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/user"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/meeting"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/committee"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/group"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/tag"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/agenda_item"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/assignment"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/poll"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/projector"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/speaker"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/mediafile"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/organization"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/list_of_speakers"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/personal_note"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/vote"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/option"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_block"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_category"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_comment"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_comment_section"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_state"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_workflow"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_change_recommendation"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_submitter"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_supporter"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_editor"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_working_group_speaker"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/chat_group"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/chat_message"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/structure_level"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/structure_level_list_of_speakers"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/gender"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/theme"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/organization_tag"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/meeting_user"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/meeting_mediafile"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/assignment_candidate"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/projection"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/projector_message"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/projector_countdown"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/point_of_order_category"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/import_preview"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/action_worker"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/poll_candidate"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/poll_candidate_list"
)

func TestRegisterAndLookup(t *testing.T) {
	// topic.create should be registered via the topic package init().
	_, err := action.Lookup("topic.create")
	if err != nil {
		t.Fatalf("topic.create should be registered: %v", err)
	}
}

func TestLookupMultipleActions(t *testing.T) {
	actions := []string{
		"topic.create",
		"topic.update",
		"topic.delete",
		"motion.create",
		"user.create",
		"meeting.create",
		"committee.create",
		"group.create",
	}
	for _, name := range actions {
		_, err := action.Lookup(name)
		if err != nil {
			t.Errorf("action %q should be registered: %v", name, err)
		}
	}
}

func TestLookupUnknownAction(t *testing.T) {
	_, err := action.Lookup("nonexistent.action")
	if err == nil {
		t.Fatal("expected error for unknown action")
	}
}

func TestAllActions(t *testing.T) {
	all := action.AllActions()
	if len(all) < 50 {
		t.Errorf("expected at least 50 registered actions, got %d", len(all))
	}
}

func TestAllActionsContainsKnown(t *testing.T) {
	all := action.AllActions()
	allMap := make(map[string]bool, len(all))
	for _, name := range all {
		allMap[name] = true
	}

	expected := []string{"topic.create", "topic.update", "topic.delete"}
	for _, name := range expected {
		if !allMap[name] {
			t.Errorf("expected action %q to be in AllActions()", name)
		}
	}
}
