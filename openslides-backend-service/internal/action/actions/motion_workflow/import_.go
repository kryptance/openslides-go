package motion_workflow

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("motion_workflow.import", model.MustGet("motion_workflow"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"name":       map[string]any{"type": "string", "minLength": 1},
			"states": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"name":   map[string]any{"type": "string", "minLength": 1},
						"weight": map[string]any{"type": "integer"},
					},
					"required": []string{"name"},
				},
			},
		},
		"required": []string{"meeting_id", "name"},
	}

	// Import is a singular action: one workflow at a time.
	mixin.WithSingularAction(a)

	// TODO: Implement custom UpdateInstance to create workflow with all its
	// states as sub-actions (motion_state.create for each state entry).

	action.Register(a)
}
