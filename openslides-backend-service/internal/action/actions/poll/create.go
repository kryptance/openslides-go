// Package poll implements poll CRUD and state-transition actions.
//
// Polls are used for voting in meetings, attached to motions, assignments, or topics.
// All poll actions require PollCanManage permission.
package poll

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("poll.create", model.MustGet("poll"))
	a.Permission = perm.PollCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id":  map[string]any{"type": "integer"},
			"title":       map[string]any{"type": "string", "minLength": 1},
			"type":        map[string]any{"type": "string", "enum": []string{"analog", "named", "pseudoanonymous"}},
			"pollmethod":  map[string]any{"type": "string", "enum": []string{"Y", "N", "YN", "YNA"}},
			"description": map[string]any{"type": "string"},
			"content_object_id": map[string]any{"type": "string"},
			"min_votes_amount":  map[string]any{"type": "integer"},
			"max_votes_amount":  map[string]any{"type": "integer"},
			"max_votes_per_option": map[string]any{"type": "integer"},
			"global_yes":     map[string]any{"type": "boolean"},
			"global_no":      map[string]any{"type": "boolean"},
			"global_abstain": map[string]any{"type": "boolean"},
			"onehundred_percent_base": map[string]any{"type": "string"},
			"backend":                 map[string]any{"type": "string", "enum": []string{"long", "fast"}},
			"entitled_group_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"is_pseudoanonymized": map[string]any{"type": "boolean"},
		},
		"required": []string{"meeting_id", "title", "type", "pollmethod"},
	}

	action.Register(a)
}
