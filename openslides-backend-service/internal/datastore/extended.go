package datastore

import (
	"fmt"
	"sync"

	"github.com/OpenSlides/openslides-backend-service/internal/event"
)

// Extended provides an in-memory overlay on top of database reads.
// Changed models are tracked and merged with database results.
type Extended struct {
	mu            sync.RWMutex
	changedModels map[string]map[int]PartialModel // collection → id → fields
	toBeDeleted   map[string]bool                 // fqid → true
	nextID        map[string]int                  // collection → last reserved id
}

// NewExtended creates a new Extended datastore.
func NewExtended() *Extended {
	return &Extended{
		changedModels: make(map[string]map[int]PartialModel),
		toBeDeleted:   make(map[string]bool),
		nextID:        make(map[string]int),
	}
}

// Get returns the fields for a single model instance, merging in-memory changes.
func (e *Extended) Get(collection string, id int, fields []string) (PartialModel, error) {
	fqid := event.FQID(collection, id)

	e.mu.RLock()
	defer e.mu.RUnlock()

	if e.toBeDeleted[fqid] {
		return nil, fmt.Errorf("model %s has been deleted", fqid)
	}

	// Start with empty model - in real implementation, this fetches from DB.
	result := make(PartialModel)

	// Overlay in-memory changes.
	if models, ok := e.changedModels[collection]; ok {
		if changed, ok := models[id]; ok {
			for _, field := range fields {
				if val, ok := changed[field]; ok {
					result[field] = val
				}
			}
		}
	}

	return result, nil
}

// GetMany returns fields for multiple model instances.
func (e *Extended) GetMany(requests []GetManyRequest) (map[string]map[int]PartialModel, error) {
	result := make(map[string]map[int]PartialModel)

	for _, req := range requests {
		if _, ok := result[req.Collection]; !ok {
			result[req.Collection] = make(map[int]PartialModel)
		}
		for _, id := range req.IDs {
			pm, err := e.Get(req.Collection, id, req.Fields)
			if err != nil {
				continue // Skip deleted models.
			}
			result[req.Collection][id] = pm
		}
	}

	return result, nil
}

// ApplyChangedModel stores field changes for a model instance in memory.
func (e *Extended) ApplyChangedModel(collection string, id int, fields PartialModel) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, ok := e.changedModels[collection]; !ok {
		e.changedModels[collection] = make(map[int]PartialModel)
	}

	existing, ok := e.changedModels[collection][id]
	if !ok {
		existing = make(PartialModel)
	}

	for k, v := range fields {
		existing[k] = v
	}
	e.changedModels[collection][id] = existing
}

// ApplyToBeDeleted marks a model instance for deletion.
func (e *Extended) ApplyToBeDeleted(collection string, id int) {
	e.mu.Lock()
	defer e.mu.Unlock()

	fqid := event.FQID(collection, id)
	e.toBeDeleted[fqid] = true
}

// IsDeleted returns true if the model instance is marked for deletion.
func (e *Extended) IsDeleted(collection string, id int) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()

	fqid := event.FQID(collection, id)
	return e.toBeDeleted[fqid]
}

// GetChangedModel returns the in-memory changes for a model instance.
func (e *Extended) GetChangedModel(collection string, id int) (PartialModel, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if models, ok := e.changedModels[collection]; ok {
		if pm, ok := models[id]; ok {
			return pm, true
		}
	}
	return nil, false
}

// SetNextID sets the next ID counter for a collection if the given id is higher
// than the current counter. This is used by test utilities to ensure that
// ReserveIDs returns IDs that don't conflict with pre-set models.
func (e *Extended) SetNextID(collection string, id int) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if id > e.nextID[collection] {
		e.nextID[collection] = id
	}
}

// ReserveIDs reserves n sequential IDs for a collection.
func (e *Extended) ReserveIDs(collection string, n int) ([]int, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	start := e.nextID[collection] + 1
	ids := make([]int, n)
	for i := range n {
		ids[i] = start + i
	}
	e.nextID[collection] = start + n - 1

	return ids, nil
}

// BuildWriteRequest converts all tracked changes into a WriteRequest.
func (e *Extended) BuildWriteRequest(userID int) *event.WriteRequest {
	e.mu.RLock()
	defer e.mu.RUnlock()

	wr := event.NewWriteRequest(userID)

	// Create events for changed models.
	for collection, models := range e.changedModels {
		for id, fields := range models {
			fqid := event.FQID(collection, id)

			if e.toBeDeleted[fqid] {
				continue
			}

			// Determine if this is a create or update based on whether "id" was set.
			if _, isNew := fields["meta_new"]; isNew {
				// Remove meta_new marker.
				eventFields := make(map[string]any, len(fields))
				for k, v := range fields {
					if k != "meta_new" {
						eventFields[k] = v
					}
				}
				wr.AddEvent(event.Event{
					Type:   event.TypeCreate,
					FQID:   fqid,
					Fields: eventFields,
				})
			} else {
				wr.AddEvent(event.Event{
					Type:   event.TypeUpdate,
					FQID:   fqid,
					Fields: fields,
				})
			}
		}
	}

	// Delete events.
	for fqid := range e.toBeDeleted {
		wr.AddEvent(event.Event{
			Type: event.TypeDelete,
			FQID: fqid,
		})
	}

	return wr
}

// Reset clears all in-memory state. Used between non-atomic action executions.
func (e *Extended) Reset() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.changedModels = make(map[string]map[int]PartialModel)
	e.toBeDeleted = make(map[string]bool)
}
