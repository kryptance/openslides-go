package list_of_speakers

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
	a := action.NewBaseAction("list_of_speakers.delete_all_speakers", model.MustGet("list_of_speakers"))
	a.Permission = perm.ListOfSpeakersCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	mixin.WithSingularAction(a)

	// Delete all speakers from this list.
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

		// Look up the list_of_speakers to find its speaker_ids.
		los, found := params.Datastore.GetChangedModel("list_of_speakers", idInt)
		if !found {
			return nil, fmt.Errorf("list_of_speakers/%d not found", idInt)
		}

		speakerIDs, ok := los["speaker_ids"]
		if !ok || speakerIDs == nil {
			return nil, nil
		}

		sids, ok := toIntSlice(speakerIDs)
		if !ok {
			return nil, nil
		}

		var events []event.Event
		for _, sid := range sids {
			fqid := event.FQID("speaker", sid)
			events = append(events, event.Event{
				Type: event.TypeDelete,
				FQID: fqid,
			})
		}

		return events, nil
	}

	action.Register(a)
}
