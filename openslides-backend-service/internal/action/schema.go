package action

import (
	"fmt"

	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
)

// ValidateSchema validates an instance against an action's JSON schema definition.
// Performs structural checks: required fields, type checking, and constraint validation.
func ValidateSchema(schema map[string]any, instance Instance) error {
	if schema == nil {
		return nil
	}

	// Check required fields.
	if required, ok := schema["required"]; ok {
		if requiredList, ok := required.([]string); ok {
			for _, field := range requiredList {
				if _, exists := instance[field]; !exists {
					return backenderr.SchemaError{
						Message: fmt.Sprintf("required field %q is missing", field),
					}
				}
			}
		}
	}

	// Validate properties.
	properties, ok := schema["properties"].(map[string]any)
	if !ok {
		return nil
	}

	for key, val := range instance {
		if key == "id" || len(key) > 5 && key[:5] == "meta_" {
			continue
		}

		propDef, ok := properties[key]
		if !ok {
			// Check additionalProperties setting.
			if addProps, ok := schema["additionalProperties"]; ok {
				if addPropsBool, ok := addProps.(bool); ok && !addPropsBool {
					return backenderr.SchemaError{
						Message: fmt.Sprintf("unexpected field %q", key),
					}
				}
			}
			continue
		}

		propMap, ok := propDef.(map[string]any)
		if !ok {
			continue
		}

		if err := validateProperty(key, val, propMap); err != nil {
			return err
		}
	}

	return nil
}

func validateProperty(field string, value any, prop map[string]any) error {
	if value == nil {
		return nil
	}

	propType, ok := prop["type"].(string)
	if !ok {
		return nil
	}

	switch propType {
	case "integer":
		switch v := value.(type) {
		case float64:
			// Check constraints.
			if minVal, ok := prop["minimum"]; ok {
				if minFloat, ok := minVal.(float64); ok && v < minFloat {
					return backenderr.SchemaError{
						Message: fmt.Sprintf("field %q value %v is less than minimum %v", field, v, minFloat),
					}
				}
			}
		case int:
			// ok
		default:
			return backenderr.SchemaError{
				Message: fmt.Sprintf("field %q must be an integer, got %T", field, value),
			}
		}

	case "string":
		str, ok := value.(string)
		if !ok {
			return backenderr.SchemaError{
				Message: fmt.Sprintf("field %q must be a string, got %T", field, value),
			}
		}
		if minLen, ok := prop["minLength"]; ok {
			if minLenFloat, ok := minLen.(float64); ok && float64(len(str)) < minLenFloat {
				return backenderr.SchemaError{
					Message: fmt.Sprintf("field %q must have at least %v characters", field, minLen),
				}
			}
			if minLenInt, ok := minLen.(int); ok && len(str) < minLenInt {
				return backenderr.SchemaError{
					Message: fmt.Sprintf("field %q must have at least %d characters", field, minLenInt),
				}
			}
		}

	case "boolean":
		if _, ok := value.(bool); !ok {
			return backenderr.SchemaError{
				Message: fmt.Sprintf("field %q must be a boolean, got %T", field, value),
			}
		}

	case "array":
		arr, ok := value.([]any)
		if !ok {
			return backenderr.SchemaError{
				Message: fmt.Sprintf("field %q must be an array, got %T", field, value),
			}
		}
		// Validate array items.
		if items, ok := prop["items"].(map[string]any); ok {
			for i, item := range arr {
				if err := validateProperty(fmt.Sprintf("%s[%d]", field, i), item, items); err != nil {
					return err
				}
			}
		}

	case "object":
		if _, ok := value.(map[string]any); !ok {
			return backenderr.SchemaError{
				Message: fmt.Sprintf("field %q must be an object, got %T", field, value),
			}
		}

	case "number":
		switch value.(type) {
		case float64, int:
			// ok
		default:
			return backenderr.SchemaError{
				Message: fmt.Sprintf("field %q must be a number, got %T", field, value),
			}
		}
	}

	// Validate enum.
	if enum, ok := prop["enum"].([]string); ok {
		str, ok := value.(string)
		if ok {
			found := false
			for _, e := range enum {
				if e == str {
					found = true
					break
				}
			}
			if !found {
				return backenderr.SchemaError{
					Message: fmt.Sprintf("field %q value %q is not one of the allowed values %v", field, str, enum),
				}
			}
		}
	}

	return nil
}
