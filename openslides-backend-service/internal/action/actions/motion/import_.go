package motion

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
	a := action.NewBaseAction("motion.import", model.MustGet("motion"))
	a.Permission = perm.MotionCanManage

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

	importAction := mixin.NewBaseImportAction("motion.json_upload")
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

			// Reserve an ID for the new motion.
			ids, err := params.Datastore.ReserveIDs("motion", 1)
			if err != nil {
				return nil, fmt.Errorf("reserve id for motion: %w", err)
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
			if v, ok := data["text"]; ok {
				fields["text"] = v
			}
			if v, ok := data["number"]; ok {
				fields["number"] = v
			}
			if v, ok := data["reason"]; ok {
				fields["reason"] = v
			}

			// TODO: Resolve submitters, supporters, category_name, and
			// block_title from lookup tables and create the corresponding
			// relation events (motion_submitter.create, etc.).

			events = append(events, event.Event{
				Type:   event.TypeCreate,
				FQID:   event.FQID("motion", id),
				Fields: fields,
			})
		}

		return events, nil
	}

	action.Register(a)
}
