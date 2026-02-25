// Package theme implements theme CRUD actions.
//
// Themes define the color scheme for the frontend UI.
// All theme actions require OML can_manage_organization permission.
package theme

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("theme.create", model.MustGet("theme"))
	a.Permission = perm.OMLCanManageOrganization

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"name":        map[string]any{"type": "string", "minLength": 1},
			"accent_500":  map[string]any{"type": "string"},
			"primary_500": map[string]any{"type": "string"},
			"warn_500":    map[string]any{"type": "string"},
			"headbar":     map[string]any{"type": "string"},
			"yes":         map[string]any{"type": "string"},
			"no":          map[string]any{"type": "string"},
			"abstain":     map[string]any{"type": "string"},
			"primary_100": map[string]any{"type": "string"},
			"primary_200": map[string]any{"type": "string"},
			"primary_300": map[string]any{"type": "string"},
			"primary_400": map[string]any{"type": "string"},
			"primary_600": map[string]any{"type": "string"},
			"primary_700": map[string]any{"type": "string"},
			"primary_800": map[string]any{"type": "string"},
			"primary_900": map[string]any{"type": "string"},
			"primary_a100": map[string]any{"type": "string"},
			"primary_a200": map[string]any{"type": "string"},
			"primary_a400": map[string]any{"type": "string"},
			"primary_a700": map[string]any{"type": "string"},
			"accent_100":  map[string]any{"type": "string"},
			"accent_200":  map[string]any{"type": "string"},
			"accent_300":  map[string]any{"type": "string"},
			"accent_400":  map[string]any{"type": "string"},
			"accent_600":  map[string]any{"type": "string"},
			"accent_700":  map[string]any{"type": "string"},
			"accent_800":  map[string]any{"type": "string"},
			"accent_900":  map[string]any{"type": "string"},
			"accent_a100": map[string]any{"type": "string"},
			"accent_a200": map[string]any{"type": "string"},
			"accent_a400": map[string]any{"type": "string"},
			"accent_a700": map[string]any{"type": "string"},
			"warn_100":    map[string]any{"type": "string"},
			"warn_200":    map[string]any{"type": "string"},
			"warn_300":    map[string]any{"type": "string"},
			"warn_400":    map[string]any{"type": "string"},
			"warn_600":    map[string]any{"type": "string"},
			"warn_700":    map[string]any{"type": "string"},
			"warn_800":    map[string]any{"type": "string"},
			"warn_900":    map[string]any{"type": "string"},
			"warn_a100":   map[string]any{"type": "string"},
			"warn_a200":   map[string]any{"type": "string"},
			"warn_a400":   map[string]any{"type": "string"},
			"warn_a700":   map[string]any{"type": "string"},
		},
		"required": []string{"name"},
	}

	action.Register(a)
}
