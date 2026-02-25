package action_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
)

func TestValidateSchemaRequiredMissing(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": map[string]any{"type": "string"},
		},
		"required": []string{"title"},
	}
	instance := action.Instance{"other_field": "value"}
	err := action.ValidateSchema(schema, instance)
	if err == nil {
		t.Fatal("expected error for missing required field")
	}
}

func TestValidateSchemaRequiredPresent(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": map[string]any{"type": "string"},
		},
		"required": []string{"title"},
	}
	instance := action.Instance{"title": "hello"}
	err := action.ValidateSchema(schema, instance)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateSchemaStringType(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"name": map[string]any{"type": "string"},
		},
	}

	// Valid string.
	err := action.ValidateSchema(schema, action.Instance{"name": "valid"})
	if err != nil {
		t.Fatalf("unexpected error for valid string: %v", err)
	}

	// Invalid type.
	err = action.ValidateSchema(schema, action.Instance{"name": 123})
	if err == nil {
		t.Fatal("expected error for non-string value")
	}
}

func TestValidateSchemaIntegerType(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"count": map[string]any{"type": "integer"},
		},
	}

	// Valid integer (as float64, which is how JSON deserializes).
	err := action.ValidateSchema(schema, action.Instance{"count": float64(5)})
	if err != nil {
		t.Fatalf("unexpected error for valid float64 integer: %v", err)
	}

	// Valid integer (as int).
	err = action.ValidateSchema(schema, action.Instance{"count": 5})
	if err != nil {
		t.Fatalf("unexpected error for valid int: %v", err)
	}

	// Invalid type.
	err = action.ValidateSchema(schema, action.Instance{"count": "not a number"})
	if err == nil {
		t.Fatal("expected error for non-integer value")
	}
}

func TestValidateSchemaBooleanType(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"active": map[string]any{"type": "boolean"},
		},
	}

	// Valid boolean.
	err := action.ValidateSchema(schema, action.Instance{"active": true})
	if err != nil {
		t.Fatalf("unexpected error for valid boolean: %v", err)
	}

	// Invalid type.
	err = action.ValidateSchema(schema, action.Instance{"active": "true"})
	if err == nil {
		t.Fatal("expected error for non-boolean value")
	}
}

func TestValidateSchemaMinLength(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": map[string]any{"type": "string", "minLength": float64(3)},
		},
	}

	// Valid: meets minLength.
	err := action.ValidateSchema(schema, action.Instance{"title": "abc"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Invalid: too short.
	err = action.ValidateSchema(schema, action.Instance{"title": "ab"})
	if err == nil {
		t.Fatal("expected error for string shorter than minLength")
	}
}

func TestValidateSchemaMinLengthInt(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": map[string]any{"type": "string", "minLength": 1},
		},
	}

	// Valid: meets minLength.
	err := action.ValidateSchema(schema, action.Instance{"title": "a"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Invalid: empty string.
	err = action.ValidateSchema(schema, action.Instance{"title": ""})
	if err == nil {
		t.Fatal("expected error for empty string with minLength=1")
	}
}

func TestValidateSchemaArrayType(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
	}

	// Valid array.
	err := action.ValidateSchema(schema, action.Instance{"ids": []any{float64(1), float64(2)}})
	if err != nil {
		t.Fatalf("unexpected error for valid array: %v", err)
	}

	// Invalid: not an array.
	err = action.ValidateSchema(schema, action.Instance{"ids": "not array"})
	if err == nil {
		t.Fatal("expected error for non-array value")
	}
}

func TestValidateSchemaArrayItemType(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
	}

	// Invalid: array item has wrong type.
	err := action.ValidateSchema(schema, action.Instance{"ids": []any{"not", "integers"}})
	if err == nil {
		t.Fatal("expected error for array items with wrong type")
	}
}

func TestValidateSchemaObjectType(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"config": map[string]any{"type": "object"},
		},
	}

	// Valid object.
	err := action.ValidateSchema(schema, action.Instance{"config": map[string]any{"key": "value"}})
	if err != nil {
		t.Fatalf("unexpected error for valid object: %v", err)
	}

	// Invalid type.
	err = action.ValidateSchema(schema, action.Instance{"config": "not object"})
	if err == nil {
		t.Fatal("expected error for non-object value")
	}
}

func TestValidateSchemaNumberType(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"weight": map[string]any{"type": "number"},
		},
	}

	// Valid float64.
	err := action.ValidateSchema(schema, action.Instance{"weight": float64(3.14)})
	if err != nil {
		t.Fatalf("unexpected error for valid float64: %v", err)
	}

	// Valid int.
	err = action.ValidateSchema(schema, action.Instance{"weight": 42})
	if err != nil {
		t.Fatalf("unexpected error for valid int: %v", err)
	}

	// Invalid string.
	err = action.ValidateSchema(schema, action.Instance{"weight": "heavy"})
	if err == nil {
		t.Fatal("expected error for non-number value")
	}
}

func TestValidateSchemaMinimumConstraint(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"weight": map[string]any{"type": "integer", "minimum": float64(0)},
		},
	}

	// Valid: at minimum.
	err := action.ValidateSchema(schema, action.Instance{"weight": float64(0)})
	if err != nil {
		t.Fatalf("unexpected error for value at minimum: %v", err)
	}

	// Invalid: below minimum.
	err = action.ValidateSchema(schema, action.Instance{"weight": float64(-1)})
	if err == nil {
		t.Fatal("expected error for value below minimum")
	}
}

func TestValidateSchemaNilSchema(t *testing.T) {
	err := action.ValidateSchema(nil, action.Instance{"anything": "goes"})
	if err != nil {
		t.Fatalf("nil schema should accept anything, got: %v", err)
	}
}

func TestValidateSchemaNilValue(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": map[string]any{"type": "string"},
		},
	}

	// nil value should be accepted (nullable).
	err := action.ValidateSchema(schema, action.Instance{"title": nil})
	if err != nil {
		t.Fatalf("nil value should be accepted: %v", err)
	}
}

func TestValidateSchemaAdditionalPropertiesFalse(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": map[string]any{"type": "string"},
		},
		"additionalProperties": false,
	}

	// Known field is fine.
	err := action.ValidateSchema(schema, action.Instance{"title": "ok"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Unknown field should fail.
	err = action.ValidateSchema(schema, action.Instance{"unknown_field": "nope"})
	if err == nil {
		t.Fatal("expected error for unknown field with additionalProperties=false")
	}
}

func TestValidateSchemaEnum(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"state": map[string]any{"type": "string", "enum": []string{"open", "closed"}},
		},
	}

	// Valid enum value.
	err := action.ValidateSchema(schema, action.Instance{"state": "open"})
	if err != nil {
		t.Fatalf("unexpected error for valid enum: %v", err)
	}

	// Invalid enum value.
	err = action.ValidateSchema(schema, action.Instance{"state": "invalid"})
	if err == nil {
		t.Fatal("expected error for invalid enum value")
	}
}

func TestValidateSchemaIdFieldSkipped(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": map[string]any{"type": "string"},
		},
		"additionalProperties": false,
	}

	// The "id" field should always be skipped in validation.
	err := action.ValidateSchema(schema, action.Instance{"id": 1, "title": "ok"})
	if err != nil {
		t.Fatalf("id field should be skipped in validation: %v", err)
	}
}

func TestValidateSchemaMetaFieldSkipped(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": map[string]any{"type": "string"},
		},
		"additionalProperties": false,
	}

	// Fields starting with "meta_" should be skipped.
	err := action.ValidateSchema(schema, action.Instance{"meta_new": true, "title": "ok"})
	if err != nil {
		t.Fatalf("meta_ fields should be skipped in validation: %v", err)
	}
}
