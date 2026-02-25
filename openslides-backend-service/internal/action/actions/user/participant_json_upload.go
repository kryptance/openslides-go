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
	a := action.NewBaseAction("user.participant.json_upload", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"data": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"username":            map[string]any{"type": "string"},
						"first_name":          map[string]any{"type": "string"},
						"last_name":           map[string]any{"type": "string"},
						"email":               map[string]any{"type": "string"},
						"gender":              map[string]any{"type": "string"},
						"default_password":    map[string]any{"type": "string"},
						"groups":              map[string]any{"type": "string"},
						"structure_levels":    map[string]any{"type": "string"},
						"number":              map[string]any{"type": "string"},
						"vote_weight":         map[string]any{"type": "string"},
						"comment":             map[string]any{"type": "string"},
						"is_present":          map[string]any{"type": "boolean"},
						"saml_id":             map[string]any{"type": "string"},
					},
				},
			},
		},
		"required": []string{"meeting_id", "data"},
	}

	uploadAction := mixin.NewBaseJsonUploadAction([]mixin.ImportHeader{
		{Property: "username", Type: "string", IsRequired: false},
		{Property: "first_name", Type: "string", IsRequired: false},
		{Property: "last_name", Type: "string", IsRequired: false},
		{Property: "email", Type: "string", IsRequired: false},
		{Property: "gender", Type: "string", IsRequired: false},
		{Property: "default_password", Type: "string", IsRequired: false},
		{Property: "groups", Type: "string", IsRequired: false},
		{Property: "structure_levels", Type: "string", IsRequired: false},
		{Property: "number", Type: "string", IsRequired: false},
		{Property: "vote_weight", Type: "string", IsRequired: false},
		{Property: "comment", Type: "string", IsRequired: false},
		{Property: "is_present", Type: "boolean", IsRequired: false},
		{Property: "saml_id", Type: "string", IsRequired: false},
	})

	mixin.WithBaseJsonUpload(a, uploadAction)
	mixin.WithSingularAction(a)

	// Override CreateEvents to store the validated import data as an action_worker.
	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		// TODO: Validate each row:
		//   - Look up existing users by username or saml_id.
		//   - Resolve group names to group_ids within the meeting.
		//   - Resolve structure level names to structure_level_ids within the meeting.
		//   - Validate email format.
		//   - Validate vote_weight is a valid decimal string.
		//   - Generate usernames where not provided.
		//   - Check for duplicate usernames within the import data.
		// TODO: Create an action_worker entry storing the validated rows.
		// TODO: Return the import preview as the action result.

		_ = uploadAction.Rows

		return nil, fmt.Errorf("user.participant.json_upload: not yet implemented - requires action_worker storage")
	}

	action.Register(a)
}
