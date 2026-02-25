package motion

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("motion.reset_state", model.MustGet("motion"))
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

		// Read the motion to find its current state_id, then look up the
		// workflow to determine the first_state_id.
		motionData, found := params.Datastore.GetChangedModel("motion", idInt)
		if !found {
			return nil, fmt.Errorf("motion/%d not found", idInt)
		}

		stateID, ok := motionData["state_id"]
		if !ok || stateID == nil {
			return nil, fmt.Errorf("motion/%d has no state_id", idInt)
		}

		var stateIDInt int
		switch v := stateID.(type) {
		case float64:
			stateIDInt = int(v)
		case int:
			stateIDInt = v
		default:
			return nil, fmt.Errorf("state_id has unexpected type %T", stateID)
		}

		// Look up the motion_state to find the workflow_id.
		stateData, found := params.Datastore.GetChangedModel("motion_state", stateIDInt)
		if !found {
			return nil, fmt.Errorf("motion_state/%d not found", stateIDInt)
		}

		workflowID, ok := stateData["workflow_id"]
		if !ok || workflowID == nil {
			return nil, fmt.Errorf("motion_state/%d has no workflow_id", stateIDInt)
		}

		var workflowIDInt int
		switch v := workflowID.(type) {
		case float64:
			workflowIDInt = int(v)
		case int:
			workflowIDInt = v
		default:
			return nil, fmt.Errorf("workflow_id has unexpected type %T", workflowID)
		}

		// Look up the workflow to find the first_state_id.
		workflowData, found := params.Datastore.GetChangedModel("motion_workflow", workflowIDInt)
		if !found {
			return nil, fmt.Errorf("motion_workflow/%d not found", workflowIDInt)
		}

		firstStateID, ok := workflowData["first_state_id"]
		if !ok || firstStateID == nil {
			return nil, fmt.Errorf("motion_workflow/%d has no first_state_id", workflowIDInt)
		}

		// Reset to the workflow's first state.
		instance["state_id"] = firstStateID

		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
