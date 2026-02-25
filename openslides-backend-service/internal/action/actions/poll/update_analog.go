package poll

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("poll.update_analog", model.MustGet("poll"))
	a.Permission = perm.PollCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
			"options": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"id": map[string]any{"type": "integer"},
						"Y":  map[string]any{"type": "string"},
						"N":  map[string]any{"type": "string"},
						"A":  map[string]any{"type": "string"},
					},
					"required": []string{"id"},
				},
			},
		},
		"required": []string{"id", "options"},
	}

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// TODO: Validate that the poll type is "analog".
		// TODO: Process each option entry: update the corresponding option's
		//       vote values (Y, N, A) via option.update sub-actions.

		// The options field is processed separately and should not be written
		// to the poll model itself.
		delete(instance, "options")

		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
