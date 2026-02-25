package main

// Blank imports to trigger init() registration of all actions and presenters.
//
// Each action package registers its actions in init() via action.Register().
// Each presenter file registers via presenter.RegisterPresenter() in init().
import (
	// Actions - sorted by collection.
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/action_worker"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/agenda_item"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/assignment"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/assignment_candidate"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/chat_group"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/chat_message"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/committee"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/gender"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/group"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/import_preview"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/list_of_speakers"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/mediafile"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/meeting"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/meeting_mediafile"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/meeting_user"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_block"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_category"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_change_recommendation"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_comment"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_comment_section"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_editor"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_state"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_submitter"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_supporter"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_workflow"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion_working_group_speaker"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/option"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/organization"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/organization_tag"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/personal_note"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/point_of_order_category"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/poll"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/projection"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/projector"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/projector_countdown"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/projector_message"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/speaker"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/structure_level"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/structure_level_list_of_speakers"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/tag"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/theme"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/topic"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/user"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/vote"

	// Presenters - registered via internal/presenter/ init() functions.
	_ "github.com/OpenSlides/openslides-backend-service/internal/presenter"
)
