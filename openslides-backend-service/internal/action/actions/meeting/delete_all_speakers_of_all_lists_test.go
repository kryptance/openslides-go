package meeting

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
)

func TestDeleteAllSpeakersNoLos(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/110": {
			"list_of_speakers_ids":          []any{},
			"is_active_in_organization_id":  1,
			"committee_id":                  1,
		},
	})
	resp, err := tc.Request("meeting.delete_all_speakers_of_all_lists", map[string]any{
		"id": 110,
	})
	tc.AssertSuccess(resp, err)
}

func TestDeleteAllSpeakersOneLosEmpty(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"list_of_speakers/11": {
			"meeting_id":  110,
			"speaker_ids": []any{},
		},
		"meeting/110": {
			"list_of_speakers_ids":          []any{11},
			"is_active_in_organization_id":  1,
			"committee_id":                  1,
		},
	})
	resp, err := tc.Request("meeting.delete_all_speakers_of_all_lists", map[string]any{
		"id": 110,
	})
	tc.AssertSuccess(resp, err)
}

func TestDeleteAllSpeakersOneLosOneSpeaker(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"list_of_speakers/11": {
			"meeting_id":  110,
			"speaker_ids": []any{1},
		},
		"speaker/1": {
			"list_of_speakers_id": 11,
			"meeting_id":          110,
		},
		"meeting/110": {
			"list_of_speakers_ids":          []any{11},
			"speaker_ids":                   []any{1},
			"is_active_in_organization_id":  1,
			"committee_id":                  1,
		},
	})
	resp, err := tc.Request("meeting.delete_all_speakers_of_all_lists", map[string]any{
		"id": 110,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertHasEvent(resp, event.TypeDelete, "speaker/1")
}

func TestDeleteAllSpeakersOneLosTwoSpeakers(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"list_of_speakers/11": {
			"meeting_id":  110,
			"speaker_ids": []any{1, 2},
		},
		"speaker/1": {
			"list_of_speakers_id": 11,
			"meeting_id":          110,
		},
		"speaker/2": {
			"list_of_speakers_id": 11,
			"meeting_id":          110,
		},
		"meeting/110": {
			"list_of_speakers_ids":          []any{11},
			"speaker_ids":                   []any{1, 2},
			"is_active_in_organization_id":  1,
			"committee_id":                  1,
		},
	})
	resp, err := tc.Request("meeting.delete_all_speakers_of_all_lists", map[string]any{
		"id": 110,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertHasEvent(resp, event.TypeDelete, "speaker/1")
	tc.AssertHasEvent(resp, event.TypeDelete, "speaker/2")
}

func TestDeleteAllSpeakersThreeLos(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"list_of_speakers/11": {
			"meeting_id":  110,
			"speaker_ids": []any{1, 2},
		},
		"speaker/1": {
			"list_of_speakers_id": 11,
			"meeting_id":          110,
		},
		"speaker/2": {
			"list_of_speakers_id": 11,
			"meeting_id":          110,
		},
		"list_of_speakers/12": {
			"meeting_id":  110,
			"speaker_ids": []any{},
		},
		"list_of_speakers/13": {
			"meeting_id":  110,
			"speaker_ids": []any{3},
		},
		"speaker/3": {
			"list_of_speakers_id": 13,
			"meeting_id":          110,
		},
		"meeting/110": {
			"list_of_speakers_ids":          []any{11, 12, 13},
			"speaker_ids":                   []any{1, 2, 3},
			"is_active_in_organization_id":  1,
			"committee_id":                  1,
		},
	})
	resp, err := tc.Request("meeting.delete_all_speakers_of_all_lists", map[string]any{
		"id": 110,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertHasEvent(resp, event.TypeDelete, "speaker/1")
	tc.AssertHasEvent(resp, event.TypeDelete, "speaker/2")
	tc.AssertHasEvent(resp, event.TypeDelete, "speaker/3")
}
