// Package action implements the action framework for the backend service.
package action

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/datastore"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-backend-service/internal/relation"
)

// ActionType defines whether an action is public or backend-internal.
type ActionType int

const (
	ActionTypePublic          ActionType = iota
	ActionTypeBackendInternal            // only callable internally
)

// Instance represents one data item in an action request.
type Instance map[string]any

// ActionParams holds the context for an action execution.
type ActionParams struct {
	UserID          int
	Internal        bool
	IsSubCall       bool
	Datastore       *datastore.Extended
	RelationManager *relation.Manager
}

// BaseAction is the core action struct with overridable hook functions.
// Mixins wrap these hooks to add behavior.
type BaseAction struct {
	Name       string
	Model      *model.ModelDef
	Schema     map[string]any
	ActionType ActionType

	// Permission required (perm.TPermission string, or OML string).
	Permission any

	// Overridable hooks.
	Prefetch         func(ctx context.Context, params *ActionParams, data []Instance) error
	CheckPermissions func(ctx context.Context, params *ActionParams, instance Instance) error
	UpdateInstance   func(ctx context.Context, params *ActionParams, instance Instance) (Instance, error)
	CreateEvents     func(ctx context.Context, params *ActionParams, instance Instance) ([]event.Event, error)
	GetMeetingID     func(ctx context.Context, params *ActionParams, instance Instance) (int, error)
	ValidateFields   func(ctx context.Context, params *ActionParams, instance Instance) error
}

// NewBaseAction creates a BaseAction with default no-op hooks.
func NewBaseAction(name string, m *model.ModelDef) *BaseAction {
	a := &BaseAction{
		Name:       name,
		Model:      m,
		ActionType: ActionTypePublic,
	}

	a.Prefetch = func(ctx context.Context, params *ActionParams, data []Instance) error {
		return nil
	}

	a.CheckPermissions = func(ctx context.Context, params *ActionParams, instance Instance) error {
		return nil
	}

	a.UpdateInstance = func(ctx context.Context, params *ActionParams, instance Instance) (Instance, error) {
		return instance, nil
	}

	a.CreateEvents = func(ctx context.Context, params *ActionParams, instance Instance) ([]event.Event, error) {
		return nil, nil
	}

	a.GetMeetingID = func(ctx context.Context, params *ActionParams, instance Instance) (int, error) {
		if mid, ok := instance["meeting_id"]; ok {
			if id, ok := mid.(float64); ok {
				return int(id), nil
			}
			if id, ok := mid.(int); ok {
				return id, nil
			}
		}
		return 0, nil
	}

	a.ValidateFields = func(ctx context.Context, params *ActionParams, instance Instance) error {
		return nil
	}

	return a
}
