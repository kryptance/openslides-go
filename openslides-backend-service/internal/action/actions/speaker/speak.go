package speaker

import (
	"context"
	"fmt"
	"time"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("speaker.speak", model.MustGet("speaker"))
	a.Permission = perm.ListOfSpeakersCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	// Combined begin speech: end current speaker (if any) and begin this one.
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

		now := int(time.Now().Unix())

		// Look up the speaker to find its list_of_speakers_id.
		spk, found := params.Datastore.GetChangedModel("speaker", idInt)
		if !found {
			return nil, fmt.Errorf("speaker/%d not found", idInt)
		}

		var losID int
		if lid, ok := spk["list_of_speakers_id"]; ok {
			switch v := lid.(type) {
			case float64:
				losID = int(v)
			case int:
				losID = v
			}
		}

		var events []event.Event

		// End the current speaker on this list_of_speakers if any.
		if losID > 0 {
			los, found := params.Datastore.GetChangedModel("list_of_speakers", losID)
			if found {
				if speakerIDs, ok := los["speaker_ids"]; ok {
					if sids, ok := toIntSlice(speakerIDs); ok {
						for _, sid := range sids {
							if sid == idInt {
								continue
							}
							otherSpk, found := params.Datastore.GetChangedModel("speaker", sid)
							if !found {
								continue
							}
							// Check if the other speaker is currently speaking (has begin_time but no end_time).
							if bt, ok := otherSpk["begin_time"]; ok && bt != nil {
								if et, ok := otherSpk["end_time"]; !ok || et == nil {
									fqid := event.FQID("speaker", sid)
									events = append(events, event.Event{
										Type:   event.TypeUpdate,
										FQID:   fqid,
										Fields: map[string]any{"end_time": now},
									})
									params.Datastore.ApplyChangedModel("speaker", sid, map[string]any{"end_time": now})
								}
							}
						}
					}
				}
			}
		}

		// Begin the requested speaker.
		fqid := event.FQID("speaker", idInt)
		events = append(events, event.Event{
			Type:   event.TypeUpdate,
			FQID:   fqid,
			Fields: map[string]any{"begin_time": now},
		})
		params.Datastore.ApplyChangedModel("speaker", idInt, map[string]any{"begin_time": now})

		return events, nil
	}

	action.Register(a)
}
