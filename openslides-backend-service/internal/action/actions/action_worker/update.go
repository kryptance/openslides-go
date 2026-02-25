package action_worker

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewUpdateAction("action_worker.update", model.MustGet("action_worker"))
	a.ActionType = action.ActionTypeBackendInternal
	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":        map[string]any{"type": "integer"},
			"state":     map[string]any{"type": "string"},
			"result":    map[string]any{"type": "object"},
			"timestamp": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
