// Package speaker implements speaker actions for managing speakers on lists of speakers.
package speaker

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("speaker.create", model.MustGet("speaker"))
	a.Permission = perm.ListOfSpeakersCanBeSpeaker

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"list_of_speakers_id": map[string]any{"type": "integer"},
			"meeting_user_id":     map[string]any{"type": "integer"},
			"meeting_id":          map[string]any{"type": "integer"},
			"point_of_order":      map[string]any{"type": "boolean"},
			"note":                map[string]any{"type": "string"},
			"point_of_order_category_id": map[string]any{"type": "integer"},
			"speech_state": map[string]any{
				"type": "string",
				"enum": []string{"contribution", "pro", "contra", "interposed_question", "intervention"},
			},
		},
		"required": []string{"list_of_speakers_id", "meeting_id"},
	}

	mixin.WithWeight(a, "weight")

	action.Register(a)
}
