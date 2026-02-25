// Package structure_level_list_of_speakers implements structure_level_list_of_speakers actions.
//
// Structure level list of speakers manage time quotas for structure levels
// within a list of speakers.
package structure_level_list_of_speakers

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("structure_level_list_of_speakers.create", model.MustGet("structure_level_list_of_speakers"))
	a.ActionType = action.ActionTypeBackendInternal
	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id":          map[string]any{"type": "integer"},
			"list_of_speakers_id": map[string]any{"type": "integer"},
			"structure_level_id":  map[string]any{"type": "integer"},
		},
		"required": []string{"meeting_id", "list_of_speakers_id", "structure_level_id"},
	}

	action.Register(a)
}
