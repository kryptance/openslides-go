package assignment

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("assignment.update", model.MustGet("assignment"))
	a.Permission = perm.AssignmentCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":                      map[string]any{"type": "integer"},
			"title":                   map[string]any{"type": "string", "minLength": 1},
			"description":             map[string]any{"type": "string"},
			"open_posts":              map[string]any{"type": "integer", "minimum": 0},
			"phase":                   map[string]any{"type": "integer"},
			"default_poll_description": map[string]any{"type": "string"},
			"number_poll_candidates":  map[string]any{"type": "boolean"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
