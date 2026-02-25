package meeting

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("meeting.delete_all_speakers_of_all_lists", model.MustGet("meeting"))
	a.Permission = perm.MeetingCanManageSettings

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		id, ok := instance["id"]
		if !ok {
			return nil, fmt.Errorf("instance has no id field")
		}

		var idInt int
		switch v := id.(type) {
		case float64:
			idInt = int(v)
		case int:
			idInt = v
		default:
			return nil, fmt.Errorf("id has unexpected type %T", id)
		}

		// Look up the meeting to find all list_of_speakers_ids.
		meeting, found := params.Datastore.GetChangedModel("meeting", idInt)
		if !found {
			return nil, fmt.Errorf("meeting/%d not found", idInt)
		}

		losIDs, ok := meeting["list_of_speakers_ids"]
		if !ok || losIDs == nil {
			return nil, nil
		}

		losSlice, ok := losIDs.([]any)
		if !ok {
			return nil, nil
		}

		var events []event.Event

		// For each list of speakers, find all speaker_ids and delete them.
		for _, losIDRaw := range losSlice {
			var losID int
			switch v := losIDRaw.(type) {
			case float64:
				losID = int(v)
			case int:
				losID = v
			default:
				continue
			}

			los, found := params.Datastore.GetChangedModel("list_of_speakers", losID)
			if !found {
				continue
			}

			speakerIDs, ok := los["speaker_ids"]
			if !ok || speakerIDs == nil {
				continue
			}

			speakerSlice, ok := speakerIDs.([]any)
			if !ok {
				continue
			}

			for _, spkIDRaw := range speakerSlice {
				var spkID int
				switch v := spkIDRaw.(type) {
				case float64:
					spkID = int(v)
				case int:
					spkID = v
				default:
					continue
				}

				fqid := event.FQID("speaker", spkID)
				events = append(events, event.Event{
					Type: event.TypeDelete,
					FQID: fqid,
				})
			}
		}

		return events, nil
	}

	action.Register(a)
}
