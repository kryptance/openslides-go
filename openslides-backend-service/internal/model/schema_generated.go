// Code generated from models.yml. DO NOT EDIT.
//
// To regenerate:
//   go run ./cmd/gen/ -schemas -path lib/openslides-go/meta/models.yml > internal/model/schema_generated.go
package model

// FieldTypeToJSONType maps a FieldType to its JSON Schema type string.
func FieldTypeToJSONType(ft FieldType) string {
	switch ft {
	case FieldTypeInteger, FieldTypeTimestamp:
		return "integer"
	case FieldTypeString, FieldTypeText, FieldTypeColor, FieldTypeHTMLStrict, FieldTypeHTMLPermissive:
		return "string"
	case FieldTypeBoolean:
		return "boolean"
	case FieldTypeFloat, FieldTypeDecimal:
		return "number"
	case FieldTypeJSON:
		return "object"
	case FieldTypeNumberArray:
		return "array"
	case FieldTypeStringArray:
		return "array"
	case FieldTypeRelation:
		return "integer"
	case FieldTypeRelationList:
		return "array"
	case FieldTypeGenericRelation:
		return "string"
	case FieldTypeGenericRelationList:
		return "array"
	default:
		return "string"
	}
}

// FieldTypeToJSONItems returns the items definition for array-type fields.
func FieldTypeToJSONItems(ft FieldType) map[string]any {
	switch ft {
	case FieldTypeRelationList, FieldTypeNumberArray:
		return map[string]any{"type": "integer"}
	case FieldTypeStringArray, FieldTypeGenericRelationList:
		return map[string]any{"type": "string"}
	default:
		return nil
	}
}

// BuildJSONSchemaProperties generates JSON Schema property definitions from a ModelDef.
// This is used for action payload validation.
func BuildJSONSchemaProperties(m *ModelDef) map[string]map[string]any {
	props := make(map[string]map[string]any, len(m.Fields))
	for name, field := range m.Fields {
		prop := map[string]any{
			"type": FieldTypeToJSONType(field.Type),
		}

		// Add items for array types.
		if items := FieldTypeToJSONItems(field.Type); items != nil {
			prop["items"] = items
		}

		// Add constraints.
		if field.Constraints != nil {
			for k, v := range field.Constraints {
				prop[k] = v
			}
		}

		props[name] = prop
	}
	return props
}

// BuildCreateSchema generates a JSON Schema for a create action.
func BuildCreateSchema(m *ModelDef) map[string]any {
	props := BuildJSONSchemaProperties(m)

	// Remove read-only and constant fields from create schema.
	delete(props, "id")
	for name, field := range m.Fields {
		if field.ReadOnly {
			delete(props, name)
		}
	}

	// Determine required fields.
	var required []string
	for name, field := range m.Fields {
		if field.Required && name != "id" && !field.ReadOnly {
			required = append(required, name)
		}
	}

	schema := map[string]any{
		"type":       "object",
		"properties": props,
	}
	if len(required) > 0 {
		schema["required"] = required
	}

	return schema
}

// BuildUpdateSchema generates a JSON Schema for an update action.
func BuildUpdateSchema(m *ModelDef) map[string]any {
	props := BuildJSONSchemaProperties(m)

	// Keep id as required, remove constant and read-only fields (except id).
	for name, field := range m.Fields {
		if name == "id" {
			continue
		}
		if field.Constant || field.ReadOnly {
			delete(props, name)
		}
	}

	schema := map[string]any{
		"type":       "object",
		"properties": props,
		"required":   []string{"id"},
	}

	return schema
}
