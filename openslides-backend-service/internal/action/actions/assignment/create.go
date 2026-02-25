// Package assignment implements assignment CRUD actions.
//
// Assignments (elections) allow nominating candidates and running polls.
// Management actions require AssignmentCanManage permission.
package assignment

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("assignment.create", model.MustGet("assignment"))
	a.Permission = perm.AssignmentCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id":              map[string]any{"type": "integer"},
			"title":                   map[string]any{"type": "string", "minLength": 1},
			"description":             map[string]any{"type": "string"},
			"open_posts":              map[string]any{"type": "integer", "minimum": 0},
			"phase":                   map[string]any{"type": "integer"},
			"default_poll_description": map[string]any{"type": "string"},
			"number_poll_candidates":  map[string]any{"type": "boolean"},
		},
		"required": []string{"meeting_id", "title"},
	}

	// Apply mixins for dependent resource creation.
	mixin.WithListOfSpeakers(a)
	mixin.WithAgendaCreation(a)

	action.Register(a)
}
