package structure_level_list_of_speakers

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("structure_level_list_of_speakers.add_time", model.MustGet("structure_level_list_of_speakers"))
	a.Permission = perm.ListOfSpeakersCanManage
	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":   map[string]any{"type": "integer"},
			"time": map[string]any{"type": "integer"},
		},
		"required": []string{"id", "time"},
	}

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// TODO: Fetch current additional_time and remaining_time from datastore.
		// Add the given time to both fields.
		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
