package meeting

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func setupUpdateMeeting(tc *testutil.ActionTestCase) {
	tc.SetModels(map[string]map[string]any{
		"committee/1": {
			"name": "test_committee",
		},
		"group/1": {},
		"meeting/1": {
			"name":                         "test_name",
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
			"default_group_id":             1,
			"projector_ids":                []any{1},
			"reference_projector_id":       1,
			"language":                     "en",
		},
		"projector/1": {
			"name":                                    "Projector 1",
			"meeting_id":                              1,
			"used_as_reference_projector_meeting_id":   1,
		},
	})
}

func TestUpdateExportFields(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateMeeting(tc)
	resp, err := tc.Request("meeting.update", map[string]any{
		"id":                            1,
		"export_csv_encoding":           "utf-8",
		"export_csv_separator":          ",",
		"export_pdf_pagenumber_alignment": "center",
		"export_pdf_fontsize":           11,
		"export_pdf_pagesize":           "A4",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting/1", map[string]any{
		"export_csv_encoding":             "utf-8",
		"export_csv_separator":            ",",
		"export_pdf_pagenumber_alignment": "center",
		"export_pdf_fontsize":             11,
		"export_pdf_pagesize":             "A4",
	})
}

func TestUpdateEmailFields(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateMeeting(tc)
	resp, err := tc.Request("meeting.update", map[string]any{
		"id":                  1,
		"users_email_sender":  "test@example.com",
		"users_email_replyto": "test2@example.com",
		"users_email_subject": "blablabla",
		"users_email_body":    "testtesttest",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting/1", map[string]any{
		"users_email_sender":  "test@example.com",
		"users_email_replyto": "test2@example.com",
		"users_email_subject": "blablabla",
		"users_email_body":    "testtesttest",
	})
}

func TestUpdateName(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateMeeting(tc)
	resp, err := tc.Request("meeting.update", map[string]any{
		"id":   1,
		"name": "new_name",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting/1", map[string]any{
		"name": "new_name",
	})
}

func TestUpdateDescription(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateMeeting(tc)
	resp, err := tc.Request("meeting.update", map[string]any{
		"id":          1,
		"description": "A new description",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting/1", map[string]any{
		"description": "A new description",
	})
}

func TestUpdateLocation(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateMeeting(tc)
	resp, err := tc.Request("meeting.update", map[string]any{
		"id":       1,
		"location": "Room 101",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting/1", map[string]any{
		"location": "Room 101",
	})
}

func TestUpdateTimes(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateMeeting(tc)
	resp, err := tc.Request("meeting.update", map[string]any{
		"id":         1,
		"start_time": 1608120653,
		"end_time":   1608121653,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting/1", map[string]any{
		"start_time": 1608120653,
		"end_time":   1608121653,
	})
}

func TestUpdateWelcomeFields(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateMeeting(tc)
	resp, err := tc.Request("meeting.update", map[string]any{
		"id":            1,
		"welcome_title": "Welcome!",
		"welcome_text":  "Hello everyone",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting/1", map[string]any{
		"welcome_title": "Welcome!",
		"welcome_text":  "Hello everyone",
	})
}

func TestUpdateConferenceSettings(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateMeeting(tc)
	resp, err := tc.Request("meeting.update", map[string]any{
		"id":                        1,
		"conference_show":           true,
		"conference_auto_connect":   true,
		"conference_open_microphone": true,
		"conference_open_video":     true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting/1", map[string]any{
		"conference_show":            true,
		"conference_auto_connect":    true,
		"conference_open_microphone": true,
		"conference_open_video":      true,
	})
}

func TestUpdateApplauseSettings(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateMeeting(tc)
	resp, err := tc.Request("meeting.update", map[string]any{
		"id":                1,
		"applause_enable":   true,
		"applause_type":     "applause-type-particles",
		"applause_timeout":  6,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting/1", map[string]any{
		"applause_enable":  true,
		"applause_type":    "applause-type-particles",
		"applause_timeout": 6,
	})
}

func TestUpdateAgendaSettings(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateMeeting(tc)
	resp, err := tc.Request("meeting.update", map[string]any{
		"id":                      1,
		"agenda_show_subtitles":   true,
		"agenda_enable_numbering": false,
		"agenda_numeral_system":   "arabic",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting/1", map[string]any{
		"agenda_show_subtitles":   true,
		"agenda_enable_numbering": false,
		"agenda_numeral_system":   "arabic",
	})
}

func TestUpdateListOfSpeakersSettings(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateMeeting(tc)
	resp, err := tc.Request("meeting.update", map[string]any{
		"id": 1,
		"list_of_speakers_amount_last_on_projector": 3,
		"list_of_speakers_couple_countdown":         true,
		"list_of_speakers_initially_closed":         true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting/1", map[string]any{
		"list_of_speakers_amount_last_on_projector": 3,
		"list_of_speakers_couple_countdown":         true,
		"list_of_speakers_initially_closed":         true,
	})
}

func TestUpdateMotionsSettings(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateMeeting(tc)
	resp, err := tc.Request("meeting.update", map[string]any{
		"id":                               1,
		"motions_preamble":                 "The assembly may decide:",
		"motions_default_line_numbering":   "none",
		"motions_line_length":              90,
		"motions_reason_required":          false,
		"motions_show_sequential_number":   true,
		"motions_recommendations_by":       "ABK",
		"motions_export_title":             "Motions",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting/1", map[string]any{
		"motions_preamble":               "The assembly may decide:",
		"motions_default_line_numbering": "none",
		"motions_line_length":            90,
		"motions_export_title":           "Motions",
	})
}

func TestUpdateUsersSettings(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateMeeting(tc)
	resp, err := tc.Request("meeting.update", map[string]any{
		"id":                           1,
		"users_enable_presence_view":   true,
		"users_enable_vote_weight":     true,
		"users_allow_self_set_present": true,
		"users_pdf_welcometitle":       "Welcome!",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting/1", map[string]any{
		"users_enable_presence_view":   true,
		"users_enable_vote_weight":     true,
		"users_allow_self_set_present": true,
		"users_pdf_welcometitle":       "Welcome!",
	})
}

func TestUpdateMissingId(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateMeeting(tc)
	resp, err := tc.Request("meeting.update", map[string]any{
		"name": "new_name",
	})
	tc.AssertError(resp, err)
}
