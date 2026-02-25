package topic

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("topic.json_upload", model.MustGet("topic"))
	a.Permission = perm.AgendaItemCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"data": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"title":          map[string]any{"type": "string", "minLength": 1},
						"text":           map[string]any{"type": "string"},
						"agenda_comment": map[string]any{"type": "string"},
						"agenda_type":    map[string]any{"type": "integer"},
						"agenda_duration": map[string]any{"type": "integer"},
					},
					"required": []string{"title"},
				},
			},
		},
		"required": []string{"meeting_id", "data"},
	}

	// Set custom UpdateInstance before the mixin so WithBaseJsonUpload wraps around it.
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// TODO: Validate each row, check for duplicates, resolve references.
		// Build import preview with status per row.
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
		{Property: "agenda_comment", Type: "string"},
		{Property: "agenda_type", Type: "integer"},
		{Property: "agenda_duration", Type: "integer"},
	})

	mixin.WithBaseJsonUpload(a, uploadAction)

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		// json_upload does not create events, it only returns the preview.
		return nil, nil
	}

	action.Register(a)
}
