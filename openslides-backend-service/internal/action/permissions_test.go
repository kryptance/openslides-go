package action_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/datastore"
	"github.com/OpenSlides/openslides-backend-service/internal/relation"

	// Ensure actions are registered.
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/topic"
)

// Tests migrated from test_permissions.py.
// The Python tests use fake models with specific permissions.
// We test the permission check lifecycle behavior in Go.

func TestPermissionCheckSkippedForInternalCalls(t *testing.T) {
	// Corresponds to the general pattern: internal calls skip permission checks.
	a := action.NewBaseAction("test.perm.internal", nil)
	permCalled := false
	a.CheckPermissions = func(ctx context.Context, params *action.ActionParams, instance action.Instance) error {
		permCalled = true
		return nil
	}

	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{"id": 1}
	_, _, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if permCalled {
		t.Error("permission check should be skipped for internal calls")
	}
}

func TestPermissionCheckCalledForExternalCalls(t *testing.T) {
	// Corresponds to the general pattern: external calls invoke permission checks.
	a := action.NewBaseAction("test.perm.external", nil)
	permCalled := false
	a.CheckPermissions = func(ctx context.Context, params *action.ActionParams, instance action.Instance) error {
		permCalled = true
		return nil
	}

	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          1,
		Internal:        false,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{"id": 1}
	_, _, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !permCalled {
		t.Error("permission check should be called for external calls")
	}
}

func TestPermissionCheckDenied(t *testing.T) {
	// When permission check returns an error, the action should fail.
	a := action.NewBaseAction("test.perm.denied", nil)
	a.CheckPermissions = func(ctx context.Context, params *action.ActionParams, instance action.Instance) error {
		return &permissionError{message: "not allowed"}
	}

	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          2,
		Internal:        false,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{"id": 1}
	_, _, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err == nil {
		t.Fatal("expected error for denied permission")
	}
}

type permissionError struct {
	message string
}

func (e *permissionError) Error() string {
	return e.message
}

func TestSchemaValidationRejectsUnknownFields(t *testing.T) {
	// Test that additionalProperties: false rejects unknown fields.
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"name": map[string]any{"type": "string"},
		},
		"additionalProperties": false,
	}
	instance := action.Instance{"name": "ok", "unknown_field": "bad"}
	err := action.ValidateSchema(schema, instance)
	if err == nil {
		t.Error("expected schema validation to reject unknown field")
	}
}

func TestSchemaValidationAcceptsValidData(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"name": map[string]any{"type": "string"},
		},
		"required": []string{"name"},
	}
	instance := action.Instance{"name": "valid"}
	err := action.ValidateSchema(schema, instance)
	if err != nil {
		t.Errorf("expected valid data to pass schema validation: %v", err)
	}
}
