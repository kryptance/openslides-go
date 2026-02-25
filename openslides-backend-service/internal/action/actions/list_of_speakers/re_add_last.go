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
	a := action.NewBaseAction("list_of_speakers.re_add_last", model.MustGet("list_of_speakers"))
	a.Permission = perm.ListOfSpeakersCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	mixin.WithSingularAction(a)

	// Re-add the last finished speaker by clearing their end_time.
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

		// Look up the list_of_speakers to find its speakers.
		los, found := params.Datastore.GetChangedModel("list_of_speakers", idInt)
		if !found {
			return nil, fmt.Errorf("list_of_speakers/%d not found", idInt)
		}

		speakerIDs, ok := los["speaker_ids"]
		if !ok || speakerIDs == nil {
			return nil, fmt.Errorf("no speakers on list_of_speakers/%d", idInt)
		}

		sids, ok := toIntSlice(speakerIDs)
		if !ok || len(sids) == 0 {
			return nil, fmt.Errorf("no speakers on list_of_speakers/%d", idInt)
		}

		// Find the last finished speaker (has end_time, highest end_time).
		var lastSpeakerID int
		var lastEndTime int
		for _, sid := range sids {
			spk, found := params.Datastore.GetChangedModel("speaker", sid)
			if !found {
				continue
			}
			if et, ok := spk["end_time"]; ok && et != nil {
				var endTime int
				switch v := et.(type) {
				case float64:
					endTime = int(v)
				case int:
					endTime = v
				}
				if endTime > lastEndTime {
					lastEndTime = endTime
					lastSpeakerID = sid
				}
			}
		}

		if lastSpeakerID == 0 {
			return nil, fmt.Errorf("no finished speaker found on list_of_speakers/%d", idInt)
		}

		// Clear end_time and begin_time to re-add the speaker to the waiting list.
		fqid := event.FQID("speaker", lastSpeakerID)
		fields := map[string]any{
			"begin_time": nil,
			"end_time":   nil,
		}
		params.Datastore.ApplyChangedModel("speaker", lastSpeakerID, fields)

		return []event.Event{{
			Type:   event.TypeUpdate,
			FQID:   fqid,
			Fields: fields,
		}}, nil
	}

	action.Register(a)
}
