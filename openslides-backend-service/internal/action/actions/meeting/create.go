// Package meeting implements meeting actions.
//
// Meetings are the central entity in OpenSlides, representing a single event
// or session within a committee. All meeting actions require
// meeting.can_manage_settings permission.
package meeting

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("meeting.create", model.MustGet("meeting"))
	a.Permission = perm.MeetingCanManageSettings

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"committee_id":  map[string]any{"type": "integer"},
			"name":          map[string]any{"type": "string", "minLength": 1},
			"description":   map[string]any{"type": "string"},
			"location":      map[string]any{"type": "string"},
			"start_time":    map[string]any{"type": "integer"},
			"end_time":      map[string]any{"type": "integer"},
			"welcome_title": map[string]any{"type": "string"},
			"welcome_text":  map[string]any{"type": "string"},
			"organization_tag_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"external_id": map[string]any{"type": "string"},
		},
		"required": []string{"committee_id", "name"},
	}

	mixin.WithListOfSpeakers(a)

	action.Register(a)
}
