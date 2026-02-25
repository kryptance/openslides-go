package motion

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("motion.create_forwarded", model.MustGet("motion"))
	a.Permission = perm.MotionCanForward

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id":        map[string]any{"type": "integer"},
			"title":             map[string]any{"type": "string"},
			"text":              map[string]any{"type": "string"},
			"reason":            map[string]any{"type": "string"},
			"origin_id":         map[string]any{"type": "integer"},
			"origin_meeting_id": map[string]any{"type": "integer"},
		},
		"required": []string{"meeting_id", "title"},
	}

	// Apply mixins for text hashing and list of speakers creation.
	mixin.WithTextHash(a, "text", "text_hash")
	mixin.WithListOfSpeakers(a)

	action.Register(a)
}
