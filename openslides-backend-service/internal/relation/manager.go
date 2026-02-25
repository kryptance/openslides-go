// Package relation implements bidirectional relation management.
//
// When a relation field changes, the reverse field on the related model must also
// be updated. The Manager computes these reverse updates automatically.
package relation

import (
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

// FieldUpdate describes an update to a relation field on a related model.
type FieldUpdate struct {
	FQID  string
	Field string
	Type  UpdateType // add or remove
	Value any
}

// UpdateType indicates whether to add or remove from a relation.
type UpdateType string

const (
	UpdateTypeAdd    UpdateType = "add"
	UpdateTypeRemove UpdateType = "remove"
)

// Manager computes reverse relation updates.
type Manager struct {
	updates []FieldUpdate
}

// NewManager creates a new relation Manager.
func NewManager() *Manager {
	return &Manager{}
}

// HandleRelationUpdates processes an instance's relation fields and computes
// the necessary reverse-side updates.
//
// For each relation field that has changed, it calculates the diff (add/remove)
// and generates updates for the reverse field on the target collection.
func (m *Manager) HandleRelationUpdates(modelDef *model.ModelDef, id int, instance map[string]any) []FieldUpdate {
	var updates []FieldUpdate

	for fieldName, fieldDef := range modelDef.RelationFields() {
		val, ok := instance[fieldName]
		if !ok {
			continue
		}

		if fieldDef.Relation != nil {
			relUpdates := m.handleSingleRelation(fieldDef, modelDef.Collection, id, fieldName, val)
			updates = append(updates, relUpdates...)
		}

		if fieldDef.GenericRelation != nil {
			relUpdates := m.handleGenericRelation(fieldDef, modelDef.Collection, id, fieldName, val)
			updates = append(updates, relUpdates...)
		}
	}

	m.updates = append(m.updates, updates...)
	return updates
}

func (m *Manager) handleSingleRelation(fieldDef *model.FieldDef, collection string, id int, fieldName string, newValue any) []FieldUpdate {
	target := fieldDef.Relation.To
	reverseField := target.Field
	targetCollection := target.Collection

	var updates []FieldUpdate

	// Determine the new IDs being set.
	newIDs := toIntSlice(newValue)

	// For each new related ID, add the reverse reference.
	for _, relatedID := range newIDs {
		fqid := event.FQID(targetCollection, relatedID)
		updates = append(updates, FieldUpdate{
			FQID:  fqid,
			Field: reverseField,
			Type:  UpdateTypeAdd,
			Value: id,
		})
	}

	// TODO: Fetch current values from DB to compute removals.
	// For now, only additions are tracked.

	return updates
}

func (m *Manager) handleGenericRelation(fieldDef *model.FieldDef, collection string, id int, fieldName string, newValue any) []FieldUpdate {
	// TODO: Implement generic relation handling.
	return nil
}

// GetEvents converts accumulated field updates into events.
func (m *Manager) GetEvents() []event.Event {
	// Group updates by FQID.
	byFQID := make(map[string]map[string]*listFieldOp)
	for _, u := range m.updates {
		if _, ok := byFQID[u.FQID]; !ok {
			byFQID[u.FQID] = make(map[string]*listFieldOp)
		}
		op, ok := byFQID[u.FQID][u.Field]
		if !ok {
			op = &listFieldOp{}
			byFQID[u.FQID][u.Field] = op
		}
		switch u.Type {
		case UpdateTypeAdd:
			op.add = append(op.add, u.Value)
		case UpdateTypeRemove:
			op.remove = append(op.remove, u.Value)
		}
	}

	var events []event.Event
	for fqid, fields := range byFQID {
		lf := &event.ListFields{
			Add:    make(map[string][]any),
			Remove: make(map[string][]any),
		}
		for field, op := range fields {
			if len(op.add) > 0 {
				lf.Add[field] = op.add
			}
			if len(op.remove) > 0 {
				lf.Remove[field] = op.remove
			}
		}
		events = append(events, event.Event{
			Type:       event.TypeUpdate,
			FQID:       fqid,
			ListFields: lf,
		})
	}

	return events
}

// Reset clears all accumulated updates.
func (m *Manager) Reset() {
	m.updates = nil
}

type listFieldOp struct {
	add    []any
	remove []any
}

// toIntSlice converts a value to a slice of ints.
func toIntSlice(val any) []int {
	if val == nil {
		return nil
	}

	switch v := val.(type) {
	case float64:
		if v == 0 {
			return nil
		}
		return []int{int(v)}
	case int:
		if v == 0 {
			return nil
		}
		return []int{v}
	case []any:
		result := make([]int, 0, len(v))
		for _, item := range v {
			switch id := item.(type) {
			case float64:
				result = append(result, int(id))
			case int:
				result = append(result, id)
			}
		}
		return result
	case []int:
		return v
	}
	return nil
}
