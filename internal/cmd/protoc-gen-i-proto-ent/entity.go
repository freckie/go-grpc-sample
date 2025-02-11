package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
)

//go:embed entity.go.tpl
var entityTpl string

type Entity struct {
	Name    string
	Fields  []Field
	Edges   []Edge
	Package string
}

type Field struct {
	Name       string
	Type       string
	EnumValues []string
}

type Edge struct {
	Name       string
	TargetType string
}

// genEntity generates a Go Ent Entity from proto Message.
func (p *Plugin) genEntity(gen *protogen.Plugin, f *protogen.File, m *protogen.Message) *protogen.GeneratedFile {
	tpl := template.Must(template.New("entity.go.tpl").Parse(entityTpl))

	entity := Entity{
		Name:    m.GoIdent.GoName,
		Fields:  make([]Field, 0, len(m.Fields)),
		Edges:   nil,
		Package: string(f.GoPackageName),
	}

	for _, field := range m.Fields {
		// Handle Enum type
		if field.Enum != nil {
			var evs []string
			for _, ev := range field.Enum.Values {
				evs = append(evs, ev.GoIdent.GoName)
			}
			entity.Fields = append(entity.Fields, Field{
				Name:       string(field.Desc.Name()),
				Type:       "Enum",
				EnumValues: evs,
			})

		} else {
			entType, err := p.toEntType(field)
			if err != nil {
				log.Fatalf("failed to parse entity: %s", err.Error())
			}
			entity.Fields = append(entity.Fields, Field{
				Name: string(field.Desc.Name()),
				Type: entType,
			})
		}
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, entity); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}

	entityName := strings.ToLower(entity.Name)
	var filename string
	if *p.flgSkipPackage {
		filename = filepath.Join(entityName, fmt.Sprintf("%s.gen.go", entityName))
	} else {
		filename = f.GeneratedFilenamePrefix + ".gen.go"
	}
	g := gen.NewGeneratedFile(filename, f.GoImportPath)
	g.P(GENEREATED_FILE_INDICATOR)
	g.P()
	g.P(buf.String())
	return g
}
