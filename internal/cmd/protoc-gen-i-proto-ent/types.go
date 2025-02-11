package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// toEntType converts Protobuf types into Go Ent field types
func (p *Plugin) toEntType(field *protogen.Field) (string, error) {
	kind := field.Desc.Kind()

	// non-basic types
	if kind == protoreflect.MessageKind && field.Message != nil {
		switch field.Message.Desc.FullName() {
		case "google.protobuf.Timestamp":
			return "Time", nil
		}
	}

	// basic types
	switch kind {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
		return "Int", nil
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
		return "Int64", nil
	case protoreflect.StringKind:
		return "String", nil
	case protoreflect.BoolKind:
		return "Bool", nil
	case protoreflect.DoubleKind, protoreflect.FloatKind:
		return "Float", nil
	default:
		return "", fmt.Errorf("unknown type of field [%s]", field.Desc.FullName())
	}
}
