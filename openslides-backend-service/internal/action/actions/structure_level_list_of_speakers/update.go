package structure_level_list_of_speakers

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("structure_level_list_of_speakers.update", model.MustGet("structure_level_list_of_speakers"))
	a.Permission = perm.ListOfSpeakersCanManage
	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":                 map[string]any{"type": "integer"},
			"initial_time":       map[string]any{"type": "integer"},
			"remaining_time":     map[string]any{"type": "integer"},
			"additional_time":    map[string]any{"type": "integer"},
			"current_start_time": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
