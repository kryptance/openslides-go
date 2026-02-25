package presenter

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

// Presenter is the interface that all presenter implementations must satisfy.
type Presenter interface {
	Handle(ctx context.Context, userID int, data json.RawMessage) (any, error)
}

var (
	presentersMu sync.RWMutex
	presenters   = make(map[string]Presenter)
)

// RegisterPresenter adds a presenter to the global registry.
func RegisterPresenter(name string, p Presenter) {
	presentersMu.Lock()
	defer presentersMu.Unlock()
	presenters[name] = p
}

// LookupPresenter returns the presenter for the given name.
func LookupPresenter(name string) (Presenter, error) {
	presentersMu.RLock()
	defer presentersMu.RUnlock()

	p, ok := presenters[name]
	if !ok {
		return nil, fmt.Errorf("presenter %q is not registered", name)
	}
	return p, nil
}
