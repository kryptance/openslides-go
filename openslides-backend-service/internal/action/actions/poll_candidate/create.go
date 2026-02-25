// Package poll_candidate implements poll candidate actions.
package poll_candidate

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("poll_candidate.create", model.MustGet("poll_candidate"))
	a.ActionType = action.ActionTypeBackendInternal

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id":             map[string]any{"type": "integer"},
			"user_id":               map[string]any{"type": "integer"},
			"poll_candidate_list_id": map[string]any{"type": "integer"},
			"weight":                map[string]any{"type": "integer"},
		},
		"required": []string{"meeting_id", "poll_candidate_list_id", "weight"},
	}

	action.Register(a)
}
