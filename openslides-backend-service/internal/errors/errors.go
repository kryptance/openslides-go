// Package errors defines the error hierarchy for the backend service.
//
// All errors have a Type() method that returns a string identifier and
// a StatusCode() method for HTTP status mapping.
package errors

import (
	"fmt"
	"strings"
)

// ActionError represents a validation or logic error in an action (HTTP 400).
type ActionError struct {
	Message string
}

func (e ActionError) Error() string { return e.Message }
func (e ActionError) Type() string  { return "action" }
func (e ActionError) StatusCode() int { return 400 }

// PermissionDenied is returned when the user lacks required permissions (HTTP 403).
type PermissionDenied struct {
	Message string
}

func (e PermissionDenied) Error() string { return e.Message }
func (e PermissionDenied) Type() string  { return "permission_denied" }
func (e PermissionDenied) StatusCode() int { return 403 }

// MissingPermission is returned when a specific permission is required but not present (HTTP 403).
type MissingPermission struct {
	Permission string
}

func (e MissingPermission) Error() string {
	return fmt.Sprintf("Missing permission: %s", e.Permission)
}

func (e MissingPermission) Type() string  { return "missing_permission" }
func (e MissingPermission) StatusCode() int { return 403 }

// AnonymousNotAllowed is returned when an anonymous user tries a non-public action (HTTP 403).
type AnonymousNotAllowed struct{}

func (e AnonymousNotAllowed) Error() string { return "Anonymous is not allowed to execute this action." }
func (e AnonymousNotAllowed) Type() string  { return "anonymous_not_allowed" }
func (e AnonymousNotAllowed) StatusCode() int { return 403 }

// ProtectedModelsError is returned when trying to delete a model that is protected by a relation (HTTP 400).
type ProtectedModelsError struct {
	FQIDs []string
}

func (e ProtectedModelsError) Error() string {
	return fmt.Sprintf("You can not delete this model, because the following models are still in use: %s", strings.Join(e.FQIDs, ", "))
}

func (e ProtectedModelsError) Type() string  { return "protected_models" }
func (e ProtectedModelsError) StatusCode() int { return 400 }

// ModelDoesNotExist is returned when a referenced model cannot be found (HTTP 400).
type ModelDoesNotExist struct {
	FQID string
}

func (e ModelDoesNotExist) Error() string {
	return fmt.Sprintf("Model %q does not exist.", e.FQID)
}

func (e ModelDoesNotExist) Type() string  { return "model_does_not_exist" }
func (e ModelDoesNotExist) StatusCode() int { return 400 }

// DatastoreLockedError signals a write conflict that can be retried (HTTP 409).
type DatastoreLockedError struct {
	Key string
}

func (e DatastoreLockedError) Error() string {
	return fmt.Sprintf("Datastore is locked: %s", e.Key)
}

func (e DatastoreLockedError) Type() string  { return "datastore_locked" }
func (e DatastoreLockedError) StatusCode() int { return 409 }

// SchemaError is returned when JSON schema validation fails (HTTP 400).
type SchemaError struct {
	Message string
}

func (e SchemaError) Error() string { return e.Message }
func (e SchemaError) Type() string  { return "schema" }
func (e SchemaError) StatusCode() int { return 400 }

// Typed is an interface for errors with a type identifier.
type Typed interface {
	error
	Type() string
}

// StatusCoder is an interface for errors with an HTTP status code.
type StatusCoder interface {
	StatusCode() int
}
