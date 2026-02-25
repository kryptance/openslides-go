package action

import (
	"fmt"
	"sync"
)

var (
	actionsMu sync.RWMutex
	actions   = make(map[string]*BaseAction)
)

// Register adds an action to the global actions map.
func Register(a *BaseAction) {
	actionsMu.Lock()
	defer actionsMu.Unlock()
	actions[a.Name] = a
}

// Lookup returns the action for the given name.
func Lookup(name string) (*BaseAction, error) {
	actionsMu.RLock()
	defer actionsMu.RUnlock()

	a, ok := actions[name]
	if !ok {
		return nil, fmt.Errorf("action %q is not registered", name)
	}
	return a, nil
}

// AllActions returns all registered action names.
func AllActions() []string {
	actionsMu.RLock()
	defer actionsMu.RUnlock()

	names := make([]string, 0, len(actions))
	for name := range actions {
		names = append(names, name)
	}
	return names
}

// ResetRegistry clears all registered actions. Intended for testing.
func ResetRegistry() {
	actionsMu.Lock()
	defer actionsMu.Unlock()
	actions = make(map[string]*BaseAction)
}
