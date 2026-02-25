package assignment_candidate

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("assignment_candidate.sort", model.MustGet("assignment_candidate"))
	a.Permission = perm.AssignmentCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"assignment_id": map[string]any{"type": "integer"},
			"ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"assignment_id", "ids"},
	}

	mixin.WithSingularAction(a)
	mixin.WithLinearSort(a, "weight")

	action.Register(a)
}
