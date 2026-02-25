// Package option implements option CRUD actions.
//
// Options represent individual choices within a poll.
// The create and delete actions are backend-internal only.
package option

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("option.create", model.MustGet("option"))
	a.ActionType = action.ActionTypeBackendInternal

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"poll_id":           map[string]any{"type": "integer"},
			"content_object_id": map[string]any{"type": "string"},
			"text":              map[string]any{"type": "string"},
			"weight":            map[string]any{"type": "integer"},
			"meeting_id":        map[string]any{"type": "integer"},
		},
		"required": []string{"poll_id", "meeting_id"},
	}

	action.Register(a)
}
