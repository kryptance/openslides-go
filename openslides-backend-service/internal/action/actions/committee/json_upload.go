package committee

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
	a := action.NewBaseAction("committee.json_upload", model.MustGet("committee"))
	a.Permission = perm.OMLSuperadmin

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"data": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"name":        map[string]any{"type": "string", "minLength": 1},
						"description": map[string]any{"type": "string"},
						"manager_ids": map[string]any{
							"type":  "array",
							"items": map[string]any{"type": "integer"},
						},
						"forward_to_committee_ids": map[string]any{
							"type":  "array",
							"items": map[string]any{"type": "integer"},
						},
						"organization_tag_ids": map[string]any{
							"type":  "array",
							"items": map[string]any{"type": "integer"},
						},
						"external_id": map[string]any{"type": "string"},
					},
					"required": []string{"name"},
				},
			},
		},
		"required": []string{"data"},
	}

	// JSON upload is a singular action.
	mixin.WithSingularAction(a)

	// Apply base JSON upload mixin for row validation.
	uploadAction := mixin.NewBaseJsonUploadAction([]mixin.ImportHeader{
		{Property: "name", Type: "string", IsRequired: true},
		{Property: "description", Type: "string"},
		{Property: "external_id", Type: "string"},
	})
	mixin.WithBaseJsonUpload(a, uploadAction)

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		// TODO: Implement committee JSON upload preview logic:
		// 1. Validate each row in the data array.
		// 2. Check for duplicate committee names.
		// 3. Resolve manager_ids references.
		// 4. Store validated rows as an action_worker entry for later import.
		_ = fmt.Sprintf("committee.json_upload")
		return nil, nil
	}

	action.Register(a)
}
