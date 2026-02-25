package structure_level_list_of_speakers

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateInitialTime(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"structure_level_ids":                  []any{1},
			"list_of_speakers_ids":                 []any{2},
			"structure_level_list_of_speakers_ids": []any{5},
		},
		"structure_level/1": {
			"meeting_id":                           1,
			"structure_level_list_of_speakers_ids": []any{5},
		},
		"list_of_speakers/2": {
			"meeting_id":                           1,
			"structure_level_list_of_speakers_ids": []any{5},
		},
		"structure_level_list_of_speakers/5": {
			"structure_level_id":  1,
			"list_of_speakers_id": 2,
			"meeting_id":          1,
			"initial_time":        600,
			"remaining_time":      600,
		},
	})

	resp, err := tc.Request("structure_level_list_of_speakers.update", map[string]any{
		"id":             5,
		"initial_time":   100,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("structure_level_list_of_speakers/5", map[string]any{
		"initial_time": 100,
	})
}

func TestUpdateRemainingTime(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"structure_level_ids":                  []any{1},
			"list_of_speakers_ids":                 []any{2},
			"structure_level_list_of_speakers_ids": []any{3},
		},
		"structure_level/1": {
			"meeting_id":                           1,
			"structure_level_list_of_speakers_ids": []any{3},
		},
		"list_of_speakers/2": {
			"meeting_id":                           1,
			"structure_level_list_of_speakers_ids": []any{3},
		},
		"structure_level_list_of_speakers/3": {
			"structure_level_id":  1,
			"list_of_speakers_id": 2,
			"meeting_id":          1,
			"initial_time":        600,
			"remaining_time":      500,
		},
	})

	resp, err := tc.Request("structure_level_list_of_speakers.update", map[string]any{
		"id":             3,
		"remaining_time": 700,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("structure_level_list_of_speakers/3", map[string]any{
		"remaining_time": 700,
	})
}
