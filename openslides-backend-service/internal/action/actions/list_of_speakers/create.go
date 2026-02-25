// Package list_of_speakers implements list_of_speakers actions.
//
// Lists of speakers are attached to content objects (topics, motions, etc.) and
// hold the queue of speakers for that object.
package list_of_speakers

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("list_of_speakers.create", model.MustGet("list_of_speakers"))
	a.ActionType = action.ActionTypeBackendInternal

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id":        map[string]any{"type": "integer"},
			"content_object_id": map[string]any{"type": "string"},
		},
		"required": []string{"meeting_id", "content_object_id"},
	}

	action.Register(a)
}
