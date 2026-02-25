// Package vote implements vote CRUD actions.
//
// Votes represent individual voting records within a poll option.
// The create and clear actions are backend-internal only.
package vote

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("vote.create", model.MustGet("vote"))
	a.ActionType = action.ActionTypeBackendInternal

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"option_id":  map[string]any{"type": "integer"},
			"user_id":    map[string]any{"type": "integer"},
			"value":      map[string]any{"type": "string"},
			"weight":     map[string]any{"type": "string"},
			"meeting_id": map[string]any{"type": "integer"},
			"user_token": map[string]any{"type": "string"},
		},
		"required": []string{"option_id", "value", "meeting_id"},
	}

	action.Register(a)
}
