package user

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewBaseAction("user.save_saml_account", model.MustGet("user"))
	a.ActionType = action.ActionTypeBackendInternal

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"username":   map[string]any{"type": "string"},
			"saml_id":    map[string]any{"type": "string"},
			"first_name": map[string]any{"type": "string"},
			"last_name":  map[string]any{"type": "string"},
			"email":      map[string]any{"type": "string"},
			"title":      map[string]any{"type": "string"},
			"pronoun":    map[string]any{"type": "string"},
			"gender_id":  map[string]any{"type": "integer"},
			"is_active":  map[string]any{"type": "boolean"},
		},
		"required": []string{"username", "saml_id"},
	}

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		samlID, ok := instance["saml_id"].(string)
		if !ok || samlID == "" {
			return nil, fmt.Errorf("saml_id is required and must be non-empty")
		}

		// TODO: Look up existing user by saml_id in the datastore.
		// If found: update the existing user with the provided fields.
		// If not found: create a new user with the provided fields and saml_id.

		// TODO: Filter the instance to only include known user model fields.
		// TODO: Set is_active to true by default for new SAML accounts.
		// TODO: Handle username generation/uniqueness for new accounts.

		// Placeholder: determine if this is a create or update.
		// For now, emit a placeholder that must be replaced with actual logic.
		_ = samlID

		return nil, fmt.Errorf("user.save_saml_account: not yet implemented - requires datastore lookup by saml_id")
	}

	action.Register(a)
}
