package user

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("user.account.json_upload", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"data": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"username":         map[string]any{"type": "string"},
						"first_name":       map[string]any{"type": "string"},
						"last_name":        map[string]any{"type": "string"},
						"email":            map[string]any{"type": "string"},
						"title":            map[string]any{"type": "string"},
						"pronoun":          map[string]any{"type": "string"},
						"gender":           map[string]any{"type": "string"},
						"default_password": map[string]any{"type": "string"},
						"is_active":        map[string]any{"type": "boolean"},
						"saml_id":          map[string]any{"type": "string"},
					},
				},
			},
		},
		"required": []string{"data"},
	}

	uploadAction := mixin.NewBaseJsonUploadAction([]mixin.ImportHeader{
		{Property: "username", Type: "string", IsRequired: false},
		{Property: "first_name", Type: "string", IsRequired: false},
		{Property: "last_name", Type: "string", IsRequired: false},
		{Property: "email", Type: "string", IsRequired: false},
		{Property: "title", Type: "string", IsRequired: false},
		{Property: "pronoun", Type: "string", IsRequired: false},
		{Property: "gender", Type: "string", IsRequired: false},
		{Property: "default_password", Type: "string", IsRequired: false},
		{Property: "is_active", Type: "boolean", IsRequired: false},
		{Property: "saml_id", Type: "string", IsRequired: false},
	})

	mixin.WithBaseJsonUpload(a, uploadAction)
	mixin.WithSingularAction(a)

	// Override CreateEvents to store the validated import data as an action_worker.
	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		// TODO: Validate each row:
		//   - Check username uniqueness (against existing users and within the import data).
		//   - Validate email format.
		//   - Look up existing users by username or saml_id for update vs. create determination.
		//   - Generate usernames where not provided (from first_name/last_name).
		// TODO: Create an action_worker entry storing the validated rows.
		// TODO: Return the import preview (rows with state, headers) as the action result.

		_ = uploadAction.Rows

		return nil, fmt.Errorf("user.account.json_upload: not yet implemented - requires action_worker storage")
	}

	action.Register(a)
}
