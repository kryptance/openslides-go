// Package motion_block implements motion block CRUD actions.
//
// Motion blocks group related motions together for collective voting.
// All actions require MotionCanManage permission.
package motion_block

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("motion_block.create", model.MustGet("motion_block"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"title":      map[string]any{"type": "string", "minLength": 1},
			"internal":   map[string]any{"type": "boolean"},
		},
		"required": []string{"meeting_id", "title"},
	}

	// Apply mixins for dependent resource creation.
	mixin.WithListOfSpeakers(a)
	mixin.WithAgendaCreation(a)

	action.Register(a)
}
