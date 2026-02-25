// Package motion_workflow implements motion workflow actions.
//
// Motion workflows define the state machine for motion processing.
// All actions require MotionCanManage permission.
package motion_workflow

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("motion_workflow.create", model.MustGet("motion_workflow"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"name":       map[string]any{"type": "string", "minLength": 1},
		},
		"required": []string{"meeting_id", "name"},
	}

	action.Register(a)
}
