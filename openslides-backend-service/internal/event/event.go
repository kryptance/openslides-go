// Package event defines event types for the write request to the datastore.
package event

import "fmt"

// Type represents the kind of database operation.
type Type string

const (
	TypeCreate Type = "create"
	TypeUpdate Type = "update"
	TypeDelete Type = "delete"
)

// Event represents a single change to a model instance.
type Event struct {
	Type       Type
	FQID       string
	Fields     map[string]any
	ListFields *ListFields
}

// ListFields represents add/remove operations on list fields.
type ListFields struct {
	Add    map[string][]any
	Remove map[string][]any
}

// FQID builds a fully qualified ID from collection and id.
func FQID(collection string, id int) string {
	return fmt.Sprintf("%s/%d", collection, id)
}

// FQField builds a fully qualified field from collection, id and field.
func FQField(collection string, id int, field string) string {
	return fmt.Sprintf("%s/%d/%s", collection, id, field)
}

// WriteRequest contains all events for one atomic write operation.
type WriteRequest struct {
	Events          []Event
	Information     map[string][]string
	UserID          int
	LockedFields    map[string]map[string]int
}

// NewWriteRequest creates a new empty WriteRequest.
func NewWriteRequest(userID int) *WriteRequest {
	return &WriteRequest{
		UserID:       userID,
		Information:  make(map[string][]string),
		LockedFields: make(map[string]map[string]int),
	}
}

// AddEvent appends an event to the write request.
func (wr *WriteRequest) AddEvent(e Event) {
	wr.Events = append(wr.Events, e)
}

// AddInformation adds history information for an FQID.
func (wr *WriteRequest) AddInformation(fqid string, info ...string) {
	wr.Information[fqid] = append(wr.Information[fqid], info...)
}

// Merge combines another WriteRequest's events into this one.
func (wr *WriteRequest) Merge(other *WriteRequest) {
	wr.Events = append(wr.Events, other.Events...)
	for fqid, info := range other.Information {
		wr.Information[fqid] = append(wr.Information[fqid], info...)
	}
	for key, fields := range other.LockedFields {
		if _, ok := wr.LockedFields[key]; !ok {
			wr.LockedFields[key] = make(map[string]int)
		}
		for field, pos := range fields {
			wr.LockedFields[key][field] = pos
		}
	}
}
