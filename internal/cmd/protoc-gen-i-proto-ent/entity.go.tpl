package {{ .Package }}

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type {{ .Name }} struct {
	ent.Schema
}

func ({{ .Name }}) Fields() []ent.Field {
	return []ent.Field{
		{{- range .Fields }}
			{{- if eq .Type "Enum" }}
				field.Enum("{{ .Name }}").
					Values({{- range $i, $v := .EnumValues }}{{ if $i }}, {{ end }}"{{ $v }}"{{ end }}),
			{{- else }}
				field.{{ .Type }}("{{ .Name }}"),
			{{- end }}
		{{- end }}
	}
}

{{ if .Edges }}
func ({{ .Name }}) Edges() []ent.Edge {
	return nil
}
{{ end }}