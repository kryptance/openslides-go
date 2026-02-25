package poll

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("poll.update", model.MustGet("poll"))
	a.Permission = perm.PollCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":          map[string]any{"type": "integer"},
			"title":       map[string]any{"type": "string", "minLength": 1},
			"description": map[string]any{"type": "string"},
			"min_votes_amount":     map[string]any{"type": "integer"},
			"max_votes_amount":     map[string]any{"type": "integer"},
			"max_votes_per_option": map[string]any{"type": "integer"},
			"global_yes":     map[string]any{"type": "boolean"},
			"global_no":      map[string]any{"type": "boolean"},
			"global_abstain": map[string]any{"type": "boolean"},
			"onehundred_percent_base": map[string]any{"type": "string"},
			"entitled_group_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"backend": map[string]any{"type": "string", "enum": []string{"long", "fast"}},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
