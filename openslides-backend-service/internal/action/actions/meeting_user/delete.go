package meeting_user

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewDeleteAction("meeting_user.delete", model.MustGet("meeting_user"))
	a.Permission = perm.UserCanManage

	action.Register(a)
}
