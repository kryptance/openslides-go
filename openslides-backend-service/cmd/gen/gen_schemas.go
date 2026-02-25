package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"sort"
	"strings"
	"text/template"
)

func writeSchemaDefs(w io.Writer, parsed *parsedModels) error {
	var collections []schemaCollectionData

	names := make([]string, 0, len(parsed.Models))
	for name := range parsed.Models {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		m := parsed.Models[name]
		cd := schemaCollectionData{Name: name}

		fieldNames := make([]string, 0, len(m.Fields))
		for fn := range m.Fields {
			fieldNames = append(fieldNames, fn)
		}
		sort.Strings(fieldNames)

		for _, fn := range fieldNames {
			field := m.Fields[fn]
			sf := schemaFieldData{
				Name:     fn,
				JSONType: mapJSONType(field.Type),
				Required: field.Required,
			}
			cd.Fields = append(cd.Fields, sf)
		}

		collections = append(collections, cd)
	}

	return renderSchemaTemplate(w, collections)
}

type schemaCollectionData struct {
	Name   string
	Fields []schemaFieldData
}

type schemaFieldData struct {
	Name     string
	JSONType string
	Required bool
}

func mapJSONType(yamlType string) string {
	switch {
	case yamlType == "number", yamlType == "timestamp":
		return "integer"
	case yamlType == "string", yamlType == "text", yamlType == "color",
		yamlType == "HTMLStrict", yamlType == "HTMLPermissive":
		return "string"
	case strings.HasPrefix(yamlType, "decimal"):
		return "string"
	case yamlType == "boolean":
		return "boolean"
	case yamlType == "float":
		return "number"
	case yamlType == "JSON":
		return "object"
	case yamlType == "number[]":
		return "array"
	case yamlType == "string[]":
		return "array"
	case yamlType == "relation":
		return "integer"
	case yamlType == "relation-list":
		return "array"
	case yamlType == "generic-relation":
		return "string"
	case yamlType == "generic-relation-list":
		return "array"
	default:
		return "object"
	}
}

const schemaTpl = `// Code generated from models.yml. DO NOT EDIT.
package model

// SchemaProperties maps collection names to their JSON Schema property definitions.
// Used for action payload validation.
var SchemaProperties = map[string]map[string]map[string]any{
{{- range .}}
	"{{.Name}}": {
	{{- range .Fields}}
		"{{.Name}}": {"type": "{{.JSONType}}"},
	{{- end}}
	},
{{- end}}
}
`

func renderSchemaTemplate(w io.Writer, collections []schemaCollectionData) error {
	tmpl, err := template.New("schemas").Parse(schemaTpl)
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, collections); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("format source: %w\n\nGenerated:\n%s", err, buf.String())
	}

	_, err = w.Write(formatted)
	return err
}
