// Package action_worker implements action_worker actions.
//
// Action workers track the state of asynchronous action executions.
package action_worker

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("action_worker.create", model.MustGet("action_worker"))
	a.ActionType = action.ActionTypeBackendInternal
	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"name":      map[string]any{"type": "string"},
			"state":     map[string]any{"type": "string"},
			"user_id":   map[string]any{"type": "integer"},
			"timestamp": map[string]any{"type": "integer"},
		},
		"required": []string{"name", "state", "user_id"},
	}

	action.Register(a)
}
