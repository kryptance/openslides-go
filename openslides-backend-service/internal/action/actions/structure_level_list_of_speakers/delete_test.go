package structure_level_list_of_speakers

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"list_of_speakers_default_structure_level_time": 600,
			"structure_level_ids":                           []any{1},
			"list_of_speakers_ids":                          []any{2},
			"structure_level_list_of_speakers_ids":          []any{3},
		},
		"structure_level/1": {
			"meeting_id":                                  1,
			"structure_level_list_of_speakers_ids":        []any{3},
		},
		"list_of_speakers/2": {
			"meeting_id":                                  1,
			"structure_level_list_of_speakers_ids":        []any{3},
		},
		"structure_level_list_of_speakers/3": {
			"structure_level_id":  1,
			"list_of_speakers_id": 2,
			"meeting_id":          1,
		},
	})

	resp, err := tc.Request("structure_level_list_of_speakers.delete", map[string]any{"id": 3})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("structure_level_list_of_speakers/3")
}
