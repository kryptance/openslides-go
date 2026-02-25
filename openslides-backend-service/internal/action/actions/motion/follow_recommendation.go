package motion

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("motion.follow_recommendation", model.MustGet("motion"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		id, ok := instance["id"]
		if !ok {
			return nil, fmt.Errorf("instance has no id field")
		}

		var idInt int
		switch v := id.(type) {
		case float64:
			idInt = int(v)
		case int:
			idInt = v
		default:
			return nil, fmt.Errorf("id has unexpected type %T", id)
		}

		// Read the motion to get its current recommendation_id.
		motionData, found := params.Datastore.GetChangedModel("motion", idInt)
		if !found {
			return nil, fmt.Errorf("motion/%d not found", idInt)
		}

		recommendationID, ok := motionData["recommendation_id"]
		if !ok || recommendationID == nil {
			return nil, fmt.Errorf("motion/%d has no recommendation set", idInt)
		}

		// Set the recommendation as the new state.
		instance["state_id"] = recommendationID

		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
