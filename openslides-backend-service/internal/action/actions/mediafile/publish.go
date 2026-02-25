package mediafile

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("mediafile.publish", model.MustGet("mediafile"))
	// Complex permission: requires OML can_manage_organization.
	a.Permission = perm.OMLCanManageOrganization

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"meeting_id", "ids"},
	}

	// Publish is a singular action.
	mixin.WithSingularAction(a)

	// TODO: Implement custom UpdateInstance to publish organization
	// mediafiles to a meeting by creating meeting_mediafile entries.

	action.Register(a)
}
