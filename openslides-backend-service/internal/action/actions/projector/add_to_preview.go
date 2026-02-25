package projector

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
	a := action.NewBaseAction("projector.add_to_preview", model.MustGet("projector"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":                map[string]any{"type": "integer"},
			"content_object_id": map[string]any{"type": "string"},
			"meeting_id":        map[string]any{"type": "integer"},
			"stable":            map[string]any{"type": "boolean"},
			"type":              map[string]any{"type": "string"},
			"options":           map[string]any{"type": "object"},
		},
		"required": []string{"id", "content_object_id", "meeting_id"},
	}

	mixin.WithSingularAction(a)

	// Add a projection to the preview queue of a projector.
	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		projectorID := toInt(instance["id"])
		if projectorID == 0 {
			return nil, fmt.Errorf("instance has no id field")
		}

		contentObjectID, _ := instance["content_object_id"].(string)
		meetingID := toInt(instance["meeting_id"])
		stable, _ := instance["stable"].(bool)
		projType, _ := instance["type"].(string)
		options, _ := instance["options"].(map[string]any)

		// Create a new projection in the preview queue.
		projectionIDs, err := params.Datastore.ReserveIDs("projection", 1)
		if err != nil {
			return nil, fmt.Errorf("reserve projection id: %w", err)
		}
		projectionID := projectionIDs[0]

		projFields := map[string]any{
			"id":                    projectionID,
			"meeting_id":            meetingID,
			"content_object_id":     contentObjectID,
			"preview_projector_id":  projectorID,
			"stable":                stable,
		}
		if projType != "" {
			projFields["type"] = projType
		}
		if options != nil {
			projFields["options"] = options
		}

		projFQID := event.FQID("projection", projectionID)
		params.Datastore.ApplyChangedModel("projection", projectionID, projFields)

		return []event.Event{{
			Type:   event.TypeCreate,
			FQID:   projFQID,
			Fields: projFields,
		}}, nil
	}

	action.Register(a)
}
