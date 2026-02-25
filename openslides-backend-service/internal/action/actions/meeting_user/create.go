// Package meeting_user implements meeting_user actions.
//
// Meeting users represent the relationship between a user and a meeting,
// storing meeting-specific user data like structure levels, groups, etc.
package meeting_user

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("meeting_user.create", model.MustGet("meeting_user"))
	a.Permission = perm.UserCanManage
	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id":         map[string]any{"type": "integer"},
			"user_id":            map[string]any{"type": "integer"},
			"comment":            map[string]any{"type": "string"},
			"number":             map[string]any{"type": "string"},
			"about_me":           map[string]any{"type": "string"},
			"vote_weight":        map[string]any{"type": "string"},
			"group_ids":          map[string]any{"type": "array", "items": map[string]any{"type": "integer"}},
			"structure_level_ids": map[string]any{"type": "array", "items": map[string]any{"type": "integer"}},
		},
		"required": []string{"meeting_id", "user_id"},
	}

	action.Register(a)
}
