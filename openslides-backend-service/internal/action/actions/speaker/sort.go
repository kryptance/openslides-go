package speaker

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
	a := action.NewBaseAction("speaker.sort", model.MustGet("speaker"))
	a.Permission = perm.ListOfSpeakersCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"list_of_speakers_id": map[string]any{"type": "integer"},
			"speaker_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"list_of_speakers_id", "speaker_ids"},
	}

	mixin.WithSingularAction(a)

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		speakerIDsRaw, ok := instance["speaker_ids"]
		if !ok {
			return nil, fmt.Errorf("speaker_ids field is required")
		}

		speakerIDs, ok := toIntSlice(speakerIDsRaw)
		if !ok {
			return nil, fmt.Errorf("speaker_ids must be an array of integers")
		}

		var events []event.Event
		for i, sid := range speakerIDs {
			weight := i + 1
			fqid := event.FQID("speaker", sid)
			events = append(events, event.Event{
				Type:   event.TypeUpdate,
				FQID:   fqid,
				Fields: map[string]any{"weight": weight},
			})
			params.Datastore.ApplyChangedModel("speaker", sid, map[string]any{"weight": weight})
		}

		return events, nil
	}

	action.Register(a)
}
