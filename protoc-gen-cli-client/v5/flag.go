package v5

import (
	"fmt"
	"log/slog"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func newFlag(qualify func(protogen.GoIdent) string, field *protogen.Field) *Flag {
	return &Flag{
		Field:   field,
		qualify: qualify,
	}
}

type Flag struct {
	*protogen.Field
	qualify func(protogen.GoIdent) string
}

func (f *Flag) FlagType() string {
	if f.Desc.IsList() {
		return "SliceFlag"
	}

	return "Flag"
}

func (f *Flag) Type() string {
	var builder strings.Builder

	if f.Field.Desc.HasOptionalKeyword() {
		builder.WriteRune('*')
	}

	switch f.Desc.Kind() {
	case protoreflect.BoolKind,
		protoreflect.Int32Kind,
		protoreflect.Uint32Kind,
		protoreflect.Int64Kind,
		protoreflect.Uint64Kind,
		protoreflect.StringKind:
		builder.WriteString(f.Desc.Kind().String())
	case protoreflect.FloatKind:
		builder.WriteString(fmt.Sprintf("%T", float32(0)))
	case protoreflect.DoubleKind:
		builder.WriteString(fmt.Sprintf("%T", float64(0)))
	case protoreflect.Sint32Kind:
		builder.WriteString(fmt.Sprintf("%T", int32(0)))
	case protoreflect.Sint64Kind:
		builder.WriteString(fmt.Sprintf("%T", int64(0)))
	case protoreflect.Fixed32Kind:
		builder.WriteString(fmt.Sprintf("%T", uint32(0)))
	case protoreflect.Fixed64Kind:
		builder.WriteString(fmt.Sprintf("%T", uint64(0)))
	case protoreflect.Sfixed32Kind:
		builder.WriteString(fmt.Sprintf("%T", int32(0)))
	case protoreflect.Sfixed64Kind:
		builder.WriteString(fmt.Sprintf("%T", int64(0)))
	case protoreflect.BytesKind:
		builder.WriteString("[]byte")
	case protoreflect.GroupKind:
		panic("group is not implemented")
	case protoreflect.EnumKind:
		builder.WriteString(f.qualify(f.Enum.GoIdent))
	case protoreflect.MessageKind:
		if !f.Field.Desc.HasOptionalKeyword() {
			builder.WriteRune('*')
		}
		builder.WriteString(f.qualify(f.Message.GoIdent))
		// builder.WriteString("Message")
	default:
		slog.Info("type not supported", slog.String("type", f.Desc.Kind().String()))
		panic("undefined kind")
	}
	return builder.String()
}
