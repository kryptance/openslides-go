// Package user implements user-related actions.
//
// Users are the central account model in OpenSlides, representing both
// organization-level accounts and meeting participants.
package user

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("user.create", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"username":   map[string]any{"type": "string"},
			"first_name": map[string]any{"type": "string"},
			"last_name":  map[string]any{"type": "string"},
			"email":      map[string]any{"type": "string"},
			"title":      map[string]any{"type": "string"},
			"pronoun":    map[string]any{"type": "string"},
			"gender_id":  map[string]any{"type": "integer"},
			"default_password": map[string]any{"type": "string"},
			"is_active":  map[string]any{"type": "boolean"},
			"organization_management_level": map[string]any{"type": "string"},
			"committee_management_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
	}

	// Validate email field when provided.
	mixin.WithEmailCheck(a, "email")

	// TODO: Hash default_password if provided.
	// TODO: Generate username if not provided (from first_name/last_name).
	// TODO: Validate organization_management_level against allowed values.
	// TODO: Validate committee_management_ids exist.

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// TODO: If default_password is set, hash it and store as password field.
		// TODO: If username is not set, generate from first_name + last_name.
		// TODO: Ensure username uniqueness.
		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
