// Package topic implements topic CRUD actions.
//
// Topics are simple agenda items shown in a meeting's agenda.
// All topic actions require AgendaItem.CAN_MANAGE permission.
package topic

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("topic.create", model.MustGet("topic"))
	a.Permission = perm.AgendaItemCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"title":      map[string]any{"type": "string", "minLength": 1},
			"text":       map[string]any{"type": "string"},
			"attachment_mediafile_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"meeting_id", "title"},
	}

	// Apply mixins (replaces Python's multiple inheritance):
	// Python: TopicCreate(AttachmentMixin, SequentialNumbersMixin,
	//   CreateActionWithDependencies, CreateActionWithAgendaItemMixin,
	//   CreateActionWithListOfSpeakersMixin)
	mixin.WithListOfSpeakers(a)
	mixin.WithAgendaCreation(a)
	// TODO: mixin.WithSequentialNumbers(a)
	// TODO: mixin.WithAttachment(a)

	// Dependencies: always create agenda_item and list_of_speakers.
	// WithDependencies would be used here once those actions exist:
	// mixin.WithDependencies(a,
	//   mixin.Dependency{ActionName: "agenda_item.create", FieldMap: ...},
	//   mixin.Dependency{ActionName: "list_of_speakers.create", FieldMap: ...},
	// )

	action.Register(a)
}
