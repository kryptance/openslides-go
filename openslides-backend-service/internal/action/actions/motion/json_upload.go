package motion

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("motion.json_upload", model.MustGet("motion"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"data": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"title":  map[string]any{"type": "string", "minLength": 1},
						"text":   map[string]any{"type": "string"},
						"number": map[string]any{"type": "string"},
						"reason": map[string]any{"type": "string"},
						"submitters": map[string]any{
							"type":  "array",
							"items": map[string]any{"type": "string"},
						},
						"category_name": map[string]any{"type": "string"},
						"block_title":   map[string]any{"type": "string"},
						"supporters": map[string]any{
							"type":  "array",
							"items": map[string]any{"type": "string"},
						},
					},
					"required": []string{"title"},
				},
			},
		},
		"required": []string{"meeting_id", "data"},
	}

	// Set custom UpdateInstance before the mixin so WithBaseJsonUpload wraps around it.
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// TODO: Validate each row, check for duplicate numbers, resolve
		// submitter/supporter usernames to user IDs, resolve category names
		// to category IDs, resolve block titles to block IDs.
		rows := instance["data"]
		_ = rows
		instance["result"] = map[string]any{
			"state": "done",
			"rows":  []any{},
		}
		return instance, nil
	}

	uploadAction := mixin.NewBaseJsonUploadAction([]mixin.ImportHeader{
		{Property: "title", Type: "string", IsRequired: true},
		{Property: "text", Type: "string"},
		{Property: "number", Type: "string"},
		{Property: "reason", Type: "string"},
		{Property: "submitters", Type: "string"},
		{Property: "category_name", Type: "string"},
		{Property: "block_title", Type: "string"},
		{Property: "supporters", Type: "string"},
	})

	mixin.WithBaseJsonUpload(a, uploadAction)

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		// json_upload does not create events, it only returns the preview.
		return nil, nil
	}

	action.Register(a)
}
