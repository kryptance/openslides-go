package assignment

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("assignment.json_upload", model.MustGet("assignment"))
	a.Permission = perm.AssignmentCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"data": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"title":                  map[string]any{"type": "string", "minLength": 1},
						"description":            map[string]any{"type": "string"},
						"open_posts":             map[string]any{"type": "integer", "minimum": 0},
						"phase":                  map[string]any{"type": "string", "enum": []string{"search", "voting", "finished"}},
						"default_poll_description": map[string]any{"type": "string"},
						"number_poll_candidates":  map[string]any{"type": "boolean"},
					},
					"required": []string{"title"},
				},
			},
		},
		"required": []string{"meeting_id", "data"},
	}

	// Set custom UpdateInstance before the mixin so WithBaseJsonUpload wraps around it.
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// TODO: Validate each row, check for duplicates, validate phase values.
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
		{Property: "description", Type: "string"},
		{Property: "open_posts", Type: "integer"},
		{Property: "phase", Type: "string"},
		{Property: "default_poll_description", Type: "string"},
		{Property: "number_poll_candidates", Type: "boolean"},
	})

	mixin.WithBaseJsonUpload(a, uploadAction)

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		// json_upload does not create events, it only returns the preview.
		return nil, nil
	}

	action.Register(a)
}
