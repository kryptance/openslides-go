package structure_level_list_of_speakers

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestAddTimeCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"list_of_speakers_default_structure_level_time": 1000,
			"structure_level_ids":                           []any{1},
			"list_of_speakers_ids":                          []any{1},
			"structure_level_list_of_speakers_ids":          []any{1},
		},
		"list_of_speakers/1": {
			"meeting_id":                           1,
			"structure_level_list_of_speakers_ids": []any{1},
		},
		"structure_level/1": {
			"meeting_id":                           1,
			"structure_level_list_of_speakers_ids": []any{1},
		},
		"structure_level_list_of_speakers/1": {
			"meeting_id":         1,
			"structure_level_id": 1,
			"list_of_speakers_id": 1,
			"remaining_time":     -100,
		},
	})

	resp, err := tc.Request("structure_level_list_of_speakers.add_time", map[string]any{
		"id":   1,
		"time": 100,
	})
	tc.AssertSuccess(resp, err)
}

func TestAddTimeEmptyData(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.Request("structure_level_list_of_speakers.add_time", map[string]any{})
	tc.AssertErrorContains(err, "id")
}
