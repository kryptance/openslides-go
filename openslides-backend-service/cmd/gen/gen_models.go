package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"sort"
	"strings"
	"text/template"

	"github.com/OpenSlides/openslides-go/models"
)

// parsedModels holds the parsed model data for code generation.
type parsedModels struct {
	Models map[string]models.Model
}

func parseModels(r io.Reader) (*parsedModels, error) {
	m, err := models.Unmarshal(r)
	if err != nil {
		return nil, fmt.Errorf("unmarshal models.yml: %w", err)
	}
	return &parsedModels{Models: m}, nil
}

// collectionData holds data for template rendering of one collection.
type collectionData struct {
	Name   string
	Fields []fieldData
}

type fieldData struct {
	Name            string
	FieldType       string
	Required        bool
	ReadOnly        bool
	Constant        bool
	DefaultValue    string
	HasDefault      bool
	HasRelation     bool
	RelationTo      string
	RelationField   string
	OnDelete        string
	HasGenRelation  bool
	GenRelTargets   []genRelTarget
	Constraints     map[string]string
	HasConstraints  bool
}

type genRelTarget struct {
	Collection string
	Field      string
}

func writeModelDefs(w io.Writer, parsed *parsedModels) error {
	var collections []collectionData

	// Sort collection names for deterministic output.
	names := make([]string, 0, len(parsed.Models))
	for name := range parsed.Models {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		m := parsed.Models[name]
		cd := collectionData{Name: name}

		fieldNames := make([]string, 0, len(m.Fields))
		for fn := range m.Fields {
			fieldNames = append(fieldNames, fn)
		}
		sort.Strings(fieldNames)

		for _, fn := range fieldNames {
			field := m.Fields[fn]
			fd := fieldData{
				Name:     fn,
				Required: field.Required,
			}

			fd.FieldType = mapFieldType(field.Type)

			rel := field.Relation()
			if rel != nil {
				targets := rel.ToCollections()
				switch v := rel.(type) {
				case *models.AttributeRelation:
					fd.HasRelation = true
					if len(targets) > 0 {
						fd.RelationTo = targets[0].Collection
						fd.RelationField = targets[0].ToField.Name
					}
					_ = v
				case *models.AttributeGenericRelation:
					fd.HasGenRelation = true
					for _, t := range targets {
						fd.GenRelTargets = append(fd.GenRelTargets, genRelTarget{
							Collection: t.Collection,
							Field:      t.ToField.Name,
						})
					}
					_ = v
				}
			}

			cd.Fields = append(cd.Fields, fd)
		}

		collections = append(collections, cd)
	}

	return renderModelTemplate(w, collections)
}

func mapFieldType(yamlType string) string {
	switch {
	case yamlType == "number":
		return "model.FieldTypeInteger"
	case yamlType == "string":
		return "model.FieldTypeString"
	case yamlType == "text":
		return "model.FieldTypeText"
	case yamlType == "boolean":
		return "model.FieldTypeBoolean"
	case yamlType == "float":
		return "model.FieldTypeFloat"
	case strings.HasPrefix(yamlType, "decimal"):
		return "model.FieldTypeDecimal"
	case yamlType == "timestamp":
		return "model.FieldTypeTimestamp"
	case yamlType == "color":
		return "model.FieldTypeColor"
	case yamlType == "JSON":
		return "model.FieldTypeJSON"
	case yamlType == "HTMLStrict":
		return "model.FieldTypeHTMLStrict"
	case yamlType == "HTMLPermissive":
		return "model.FieldTypeHTMLPermissive"
	case yamlType == "number[]":
		return "model.FieldTypeNumberArray"
	case yamlType == "string[]":
		return "model.FieldTypeStringArray"
	case yamlType == "relation":
		return "model.FieldTypeRelation"
	case yamlType == "relation-list":
		return "model.FieldTypeRelationList"
	case yamlType == "generic-relation":
		return "model.FieldTypeGenericRelation"
	case yamlType == "generic-relation-list":
		return "model.FieldTypeGenericRelationList"
	default:
		return "model.FieldTypeJSON"
	}
}

const modelTpl = `// Code generated from models.yml. DO NOT EDIT.
package model

func init() {
{{- range .}}
	Register(&ModelDef{
		Collection: "{{.Name}}",
		Fields: map[string]*FieldDef{
		{{- range .Fields}}
			"{{.Name}}": {
				Type: {{.FieldType}},
			{{- if .Required}}
				Required: true,
			{{- end}}
			{{- if .ReadOnly}}
				ReadOnly: true,
			{{- end}}
			{{- if .Constant}}
				Constant: true,
			{{- end}}
			{{- if .HasRelation}}
				Relation: &RelationDef{
					To: RelationTarget{Collection: "{{.RelationTo}}", Field: "{{.RelationField}}"},
				{{- if ne .OnDelete ""}}
					OnDelete: {{.OnDelete}},
				{{- end}}
				},
			{{- end}}
			{{- if .HasGenRelation}}
				GenericRelation: &GenericRelationDef{
					To: []GenericRelationTarget{
					{{- range .GenRelTargets}}
						{Collection: "{{.Collection}}", Field: "{{.Field}}"},
					{{- end}}
					},
				},
			{{- end}}
			},
		{{- end}}
		},
	})
{{end}}
}
`

func renderModelTemplate(w io.Writer, collections []collectionData) error {
	tmpl, err := template.New("models").Parse(modelTpl)
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
