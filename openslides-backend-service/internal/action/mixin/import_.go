package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
)

// ImportState represents the state of an import row or entry.
type ImportState int

const (
	// ImportStateNew indicates a row that will create a new model.
	ImportStateNew ImportState = iota

	// ImportStateDone indicates a row that has been successfully processed.
	ImportStateDone

	// ImportStateError indicates a row that has errors and cannot be imported.
	ImportStateError

	// ImportStateWarning indicates a row with non-fatal warnings.
	ImportStateWarning

	// ImportStateGenerated indicates a row with auto-generated values.
	ImportStateGenerated
)

// String returns the string representation of the ImportState.
func (s ImportState) String() string {
	switch s {
	case ImportStateNew:
		return "new"
	case ImportStateDone:
		return "done"
	case ImportStateError:
		return "error"
	case ImportStateWarning:
		return "warning"
	case ImportStateGenerated:
		return "generated"
	default:
		return "unknown"
	}
}

// ParseImportState parses a string into an ImportState.
func ParseImportState(s string) (ImportState, error) {
	switch s {
	case "new":
		return ImportStateNew, nil
	case "done":
		return ImportStateDone, nil
	case "error":
		return ImportStateError, nil
	case "warning":
		return ImportStateWarning, nil
	case "generated":
		return ImportStateGenerated, nil
	default:
		return 0, fmt.Errorf("unknown import state: %q", s)
	}
}

// LookupEntry represents a single entry in a lookup table, tracking an imported
// value's resolution state.
type LookupEntry struct {
	// Value is the imported value used for lookup.
	Value any

	// ID is the resolved model ID (0 if not found).
	ID int

	// State indicates the current state of the lookup.
	State ImportState

	// Info contains additional information or error messages.
	Info string
}

// Lookup provides a mapping from imported values to existing model IDs.
// It is used during import to resolve references (e.g., looking up a user
// by username to find their ID).
type Lookup struct {
	// Collection is the target collection being looked up.
	Collection string

	// Field is the field used for matching (e.g., "username", "name").
	Field string

	// entries maps the lookup value (string) to its resolved entry.
	entries map[string]*LookupEntry
}

// NewLookup creates a new Lookup for the given collection and field.
func NewLookup(collection, field string) *Lookup {
	return &Lookup{
		Collection: collection,
		Field:      field,
		entries:    make(map[string]*LookupEntry),
	}
}

// Add adds a value to the lookup for later resolution.
func (l *Lookup) Add(value string) {
	if _, exists := l.entries[value]; !exists {
		l.entries[value] = &LookupEntry{
			Value: value,
			State: ImportStateNew,
		}
	}
}

// Resolve attempts to find matching models in the datastore for all added values.
// TODO: This should query the datastore using filter requests.
func (l *Lookup) Resolve(params *action.ActionParams) error {
	_ = params
	// TODO: For each entry, query the datastore:
	//   filter: {l.Field: entry.Value}
	//   If exactly one match: set entry.ID and state to Done.
	//   If multiple matches: set state to Error with info message.
	//   If no match: leave state as New.
	return nil
}

// Get returns the lookup entry for a value, or nil if not found.
func (l *Lookup) Get(value string) *LookupEntry {
	return l.entries[value]
}

// GetID returns the resolved ID for a value, or 0 if not resolved.
func (l *Lookup) GetID(value string) int {
	if entry, ok := l.entries[value]; ok {
		return entry.ID
	}
	return 0
}

// HasErrors returns true if any entry is in error state.
func (l *Lookup) HasErrors() bool {
	for _, entry := range l.entries {
		if entry.State == ImportStateError {
			return true
		}
	}
	return false
}

// Entries returns all lookup entries.
func (l *Lookup) Entries() map[string]*LookupEntry {
	return l.entries
}

// BaseImportAction provides shared functionality for import execution actions.
// Import actions process previously uploaded and validated import data.
//
// This replaces Python's BaseImportAction.
type BaseImportAction struct {
	// ImportName is the name of the import preview action that generated the data
	// (e.g., "topic.json_upload").
	ImportName string

	// Lookups holds the lookup tables used to resolve references.
	Lookups map[string]*Lookup
}

// NewBaseImportAction creates a new BaseImportAction.
func NewBaseImportAction(importName string) *BaseImportAction {
	return &BaseImportAction{
		ImportName: importName,
		Lookups:    make(map[string]*Lookup),
	}
}

// AddLookup registers a lookup table for use during import.
func (b *BaseImportAction) AddLookup(name string, lookup *Lookup) {
	b.Lookups[name] = lookup
}

// WithBaseImport wraps the Prefetch hook of an action to set up the import context,
// loading the import preview data and resolving lookups.
func WithBaseImport(a *action.BaseAction, importAction *BaseImportAction) {
	origPrefetch := a.Prefetch
	a.Prefetch = func(ctx context.Context, params *action.ActionParams, data []action.Instance) error {
		if err := origPrefetch(ctx, params, data); err != nil {
			return err
		}

		// TODO: Load the import preview data from the action_worker model.
		// Validate that the import is in the correct state (not already imported).

		// Resolve all lookups.
		for name, lookup := range importAction.Lookups {
			if err := lookup.Resolve(params); err != nil {
				return fmt.Errorf("resolve lookup %q: %w", name, err)
			}
		}

		return nil
	}
}

// BaseJsonUploadAction provides shared functionality for JSON upload (preview)
// actions that validate import data before the actual import.
//
// This replaces Python's BaseJsonUploadAction.
type BaseJsonUploadAction struct {
	// Rows holds the validated import rows with their state information.
	Rows []ImportRow

	// Headers defines the expected column headers for the import.
	Headers []ImportHeader
}

// ImportRow represents a single row of import data with validation state.
type ImportRow struct {
	// State indicates the validation state of this row.
	State ImportState

	// Messages contains validation messages (errors, warnings).
	Messages []string

	// Data contains the row data as field name to value mapping.
	Data map[string]any
}

// ImportHeader defines a column header for the import format.
type ImportHeader struct {
	// Property is the field name this column maps to.
	Property string

	// Type is the expected data type ("string", "integer", "boolean", etc.).
	Type string

	// IsRequired indicates whether this column must have a value.
	IsRequired bool
}

// NewBaseJsonUploadAction creates a new BaseJsonUploadAction.
func NewBaseJsonUploadAction(headers []ImportHeader) *BaseJsonUploadAction {
	return &BaseJsonUploadAction{
		Headers: headers,
	}
}

// WithBaseJsonUpload wraps the UpdateInstance hook of an action to perform
// JSON upload validation. Each row in the import data is validated against
// the headers and the results are stored for the import execution step.
func WithBaseJsonUpload(a *action.BaseAction, uploadAction *BaseJsonUploadAction) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		result, err := origUpdate(ctx, params, instance)
		if err != nil {
			return nil, err
		}

		// Extract import data rows.
		dataRaw, ok := result["data"]
		if !ok {
			return nil, backenderr.ActionError{
				Message: "Import data must contain a 'data' field with the rows to import.",
			}
		}

		rows, ok := dataRaw.([]any)
		if !ok {
			return nil, backenderr.ActionError{
				Message: "Import 'data' field must be an array.",
			}
		}

		// Validate each row.
		validatedRows := make([]ImportRow, 0, len(rows))
		for i, rowRaw := range rows {
			rowMap, ok := rowRaw.(map[string]any)
			if !ok {
				validatedRows = append(validatedRows, ImportRow{
					State:    ImportStateError,
					Messages: []string{fmt.Sprintf("Row %d: invalid format", i+1)},
				})
				continue
			}

			row := ImportRow{
				State: ImportStateNew,
				Data:  rowMap,
			}

			// Check required headers.
			for _, header := range uploadAction.Headers {
				if header.IsRequired {
					val, exists := rowMap[header.Property]
					if !exists || val == nil || val == "" {
						row.State = ImportStateError
						row.Messages = append(row.Messages, fmt.Sprintf(
							"Required field %q is missing or empty", header.Property,
						))
					}
				}
			}

			validatedRows = append(validatedRows, row)
		}

		uploadAction.Rows = validatedRows

		// Store validated result for the import action to pick up.
		// The result is stored as an action_worker entry.
		result["rows"] = validatedRows
		result["headers"] = uploadAction.Headers

		return result, nil
	}
}
