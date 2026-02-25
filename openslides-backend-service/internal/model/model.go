// Package model defines the data model types and a global registry for all collections.
package model

import (
	"fmt"
	"sync"
)

// FieldType represents the data type of a model field.
type FieldType int

const (
	FieldTypeInteger            FieldType = iota
	FieldTypeString                       // string (maxLength=256)
	FieldTypeText                         // text (no maxLength)
	FieldTypeBoolean
	FieldTypeFloat
	FieldTypeDecimal
	FieldTypeTimestamp
	FieldTypeColor
	FieldTypeJSON
	FieldTypeHTMLStrict
	FieldTypeHTMLPermissive
	FieldTypeNumberArray
	FieldTypeStringArray
	FieldTypeRelation           // single relation (1:1 or 1:m)
	FieldTypeRelationList       // list relation (m:1 or m:n)
	FieldTypeGenericRelation    // generic single relation (FQID)
	FieldTypeGenericRelationList // generic list relation (FQID list)
)

// OnDelete defines the cascading behavior when a related model is deleted.
type OnDelete int

const (
	OnDeleteSetNull  OnDelete = iota // default: set field to null
	OnDeleteProtect                  // prevent deletion if references exist
	OnDeleteCascade                  // cascade delete to related models
)

// RelationTarget identifies the reverse side of a relation.
type RelationTarget struct {
	Collection string
	Field      string
}

// RelationDef describes a single relation field's target and on-delete behavior.
type RelationDef struct {
	To       RelationTarget
	OnDelete OnDelete
}

// GenericRelationTarget identifies one possible target of a generic relation.
type GenericRelationTarget struct {
	Collection string
	Field      string
}

// GenericRelationDef describes a generic relation field (polymorphic).
type GenericRelationDef struct {
	To       []GenericRelationTarget
	OnDelete OnDelete
}

// FieldDef defines a single field in a model.
type FieldDef struct {
	Type        FieldType
	Required    bool
	ReadOnly    bool
	Constant    bool
	Default     any
	Constraints map[string]any

	// For relation and relation-list fields.
	Relation *RelationDef

	// For generic-relation and generic-relation-list fields.
	GenericRelation *GenericRelationDef

	// EqualFields lists fields that must match between related instances.
	EqualFields []string
}

// IsRelation returns true if the field is any kind of relation.
func (f *FieldDef) IsRelation() bool {
	return f.Type == FieldTypeRelation ||
		f.Type == FieldTypeRelationList ||
		f.Type == FieldTypeGenericRelation ||
		f.Type == FieldTypeGenericRelationList
}

// IsList returns true if the field holds a list of values.
func (f *FieldDef) IsList() bool {
	return f.Type == FieldTypeRelationList ||
		f.Type == FieldTypeGenericRelationList ||
		f.Type == FieldTypeNumberArray ||
		f.Type == FieldTypeStringArray
}

// ModelDef describes a single collection's schema.
type ModelDef struct {
	Collection string
	Fields     map[string]*FieldDef
}

// GetField returns the field definition for the given name, or an error if not found.
func (m *ModelDef) GetField(name string) (*FieldDef, error) {
	f, ok := m.Fields[name]
	if !ok {
		return nil, fmt.Errorf("field %q not found in collection %q", name, m.Collection)
	}
	return f, nil
}

// RelationFields returns all relation fields of this model.
func (m *ModelDef) RelationFields() map[string]*FieldDef {
	result := make(map[string]*FieldDef)
	for name, field := range m.Fields {
		if field.IsRelation() {
			result[name] = field
		}
	}
	return result
}

// Global model registry.
var (
	registryMu sync.RWMutex
	registry   = make(map[string]*ModelDef)
)

// Register adds a model definition to the global registry.
func Register(m *ModelDef) {
	registryMu.Lock()
	defer registryMu.Unlock()
	registry[m.Collection] = m
}

// Get returns the model definition for a collection, or nil if not registered.
func Get(collection string) *ModelDef {
	registryMu.RLock()
	defer registryMu.RUnlock()
	return registry[collection]
}

// MustGet returns the model definition for a collection, panicking if not found.
func MustGet(collection string) *ModelDef {
	m := Get(collection)
	if m == nil {
		panic(fmt.Sprintf("model %q not registered", collection))
	}
	return m
}

// All returns all registered model definitions.
func All() map[string]*ModelDef {
	registryMu.RLock()
	defer registryMu.RUnlock()
	result := make(map[string]*ModelDef, len(registry))
	for k, v := range registry {
		result[k] = v
	}
	return result
}

// Reset clears the global registry. Intended for testing.
func Reset() {
	registryMu.Lock()
	defer registryMu.Unlock()
	registry = make(map[string]*ModelDef)
}
