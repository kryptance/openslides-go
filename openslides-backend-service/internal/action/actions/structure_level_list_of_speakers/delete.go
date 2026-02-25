package structure_level_list_of_speakers

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewDeleteAction("structure_level_list_of_speakers.delete", model.MustGet("structure_level_list_of_speakers"))
	a.Permission = perm.ListOfSpeakersCanManage

	action.Register(a)
}
