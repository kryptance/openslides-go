// Package motion implements motion CRUD and state management actions.
//
// Motions are formal proposals within a meeting that follow a workflow
// of states and can be voted on. Management actions require various
// motion permissions depending on the operation.
package motion

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("motion.create", model.MustGet("motion"))
	a.Permission = perm.MotionCanCreate

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"title":      map[string]any{"type": "string", "minLength": 1},
			"text":       map[string]any{"type": "string"},
			"reason":     map[string]any{"type": "string"},
			"number":     map[string]any{"type": "string"},
			"category_id": map[string]any{"type": "integer"},
			"block_id":    map[string]any{"type": "integer"},
			"supporter_meeting_user_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"workflow_id":  map[string]any{"type": "integer"},
			"agenda_create": map[string]any{"type": "boolean"},
			"agenda_type":   map[string]any{"type": "integer"},
			"attachment_mediafile_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"meeting_id", "title"},
	}

	// Apply mixins for text hashing, dependent resource creation,
	// sequential numbering, and attachment resolution.
	mixin.WithTextHash(a, "text", "text_hash")
	mixin.WithListOfSpeakers(a)
	mixin.WithAgendaCreation(a)
	mixin.WithSequentialNumbers(a)
	mixin.WithAttachment(a)

	action.Register(a)
}
