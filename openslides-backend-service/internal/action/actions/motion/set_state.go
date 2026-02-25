package motion

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("motion.set_state", model.MustGet("motion"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":       map[string]any{"type": "integer"},
			"state_id": map[string]any{"type": "integer"},
		},
		"required": []string{"id", "state_id"},
	}

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// Set the new state_id on the instance.
		instance["state_id"] = instance["state_id"]

		// Clear recommendation_extension_reference_ids when changing state,
		// as the references may no longer be valid in the new state.
		instance["recommendation_extension_reference_ids"] = nil

		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
