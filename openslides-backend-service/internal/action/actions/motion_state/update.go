package motion_state

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("motion_state.update", model.MustGet("motion_state"))
	a.Permission = perm.MotionCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":                             map[string]any{"type": "integer"},
			"name":                           map[string]any{"type": "string", "minLength": 1},
			"weight":                         map[string]any{"type": "integer"},
			"recommendation_label":           map[string]any{"type": "string"},
			"css_class":                      map[string]any{"type": "string"},
			"restrictions":                   map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
			"allow_support":                  map[string]any{"type": "boolean"},
			"allow_create_poll":              map[string]any{"type": "boolean"},
			"allow_submitter_edit":           map[string]any{"type": "boolean"},
			"set_number":                     map[string]any{"type": "boolean"},
			"show_state_extension_field":     map[string]any{"type": "boolean"},
			"show_recommendation_extension_field": map[string]any{"type": "boolean"},
			"merge_amendment_into_final":     map[string]any{"type": "integer"},
			"next_state_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"previous_state_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"allow_motion_forwarding":     map[string]any{"type": "boolean"},
			"set_created_timestamp":       map[string]any{"type": "boolean"},
			"is_internal":                 map[string]any{"type": "boolean"},
			"allow_submitter_withdraw":    map[string]any{"type": "boolean"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
