package assignment

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
	a := action.NewBaseAction("assignment.import", model.MustGet("assignment"))
	a.Permission = perm.AssignmentCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
			"import_preview": map[string]any{
				"type": "array",
			},
		},
		"required": []string{"id", "import_preview"},
	}

	importAction := mixin.NewBaseImportAction("assignment.json_upload")
	mixin.WithBaseImport(a, importAction)

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		previewRaw, ok := instance["import_preview"]
		if !ok {
			return nil, fmt.Errorf("missing import_preview in instance")
		}

		rows, ok := previewRaw.([]any)
		if !ok {
			return nil, fmt.Errorf("import_preview must be an array")
		}

		var events []event.Event
		for _, rowRaw := range rows {
			row, ok := rowRaw.(map[string]any)
			if !ok {
				continue
			}

			data, ok := row["data"].(map[string]any)
			if !ok {
				continue
			}

			// Skip rows that are in error state.
			if state, ok := row["state"].(string); ok && state == "error" {
				continue
			}

			// Reserve an ID for the new assignment.
			ids, err := params.Datastore.ReserveIDs("assignment", 1)
			if err != nil {
				return nil, fmt.Errorf("reserve id for assignment: %w", err)
			}
			id := ids[0]

			fields := make(map[string]any)
			fields["id"] = id
			if v, ok := instance["meeting_id"]; ok {
				fields["meeting_id"] = v
			}
			if v, ok := data["title"]; ok {
				fields["title"] = v
			}
			if v, ok := data["description"]; ok {
				fields["description"] = v
			}
			if v, ok := data["open_posts"]; ok {
				fields["open_posts"] = v
			}
			if v, ok := data["phase"]; ok {
				fields["phase"] = v
			}
			if v, ok := data["default_poll_description"]; ok {
				fields["default_poll_description"] = v
			}
			if v, ok := data["number_poll_candidates"]; ok {
				fields["number_poll_candidates"] = v
			}

			events = append(events, event.Event{
				Type:   event.TypeCreate,
				FQID:   event.FQID("assignment", id),
				Fields: fields,
			})
		}

		return events, nil
	}

	action.Register(a)
}
