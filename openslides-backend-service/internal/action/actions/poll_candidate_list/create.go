// Package poll_candidate_list implements poll candidate list actions.
package poll_candidate_list

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("poll_candidate_list.create", model.MustGet("poll_candidate_list"))
	a.ActionType = action.ActionTypeBackendInternal

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"option_id":  map[string]any{"type": "integer"},
		},
		"required": []string{"meeting_id", "option_id"},
	}

	action.Register(a)
}
