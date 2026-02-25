// Package datastore provides an extended datastore with in-memory change tracking.
package datastore

// PartialModel represents a subset of fields for a model instance.
type PartialModel map[string]any

// GetManyRequest defines a batch read request.
type GetManyRequest struct {
	Collection string
	IDs        []int
	Fields     []string
}

// FilterRequest defines a filter query.
type FilterRequest struct {
	Collection string
	Filter     Filter
	Fields     []string
}

// Filter defines conditions for querying models.
type Filter struct {
	Field    string
	Value    any
	Operator string // "=", "!=", ">", "<", ">=", "<=", "in", "not in"
}

// AndFilter combines multiple filters with AND logic.
type AndFilter struct {
	Filters []any // Filter, AndFilter, OrFilter, NotFilter
}

// OrFilter combines multiple filters with OR logic.
type OrFilter struct {
	Filters []any
}

// NotFilter negates a filter.
type NotFilter struct {
	Filter any
}
