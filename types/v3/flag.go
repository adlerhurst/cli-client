package types

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

type Flag struct {
	*protogen.Field

	Custom  *customFlag
	Message *messageFlag
	Enum    *enumFlag
	OneOf   *oneOfFlag
}

func (flag *Flag) FlagConstructor() string {
	var builder strings.Builder

	builder.WriteString(flag.VarName())
	builder.WriteString(" := ")

	builder.WriteString("New")
	builder.WriteString(title.String(flag.typ()))

	if flag.IsList() {
		builder.WriteString("Slice")
	}

	builder.WriteString("Flag")
	if flag.Enum != nil {
		builder.WriteString("[")
		builder.WriteString(flag.Enum.Type())
		builder.WriteString("]")
	}
	builder.WriteString(`(set, "`)
	builder.WriteString(flag.FieldNamePrivate())
	builder.WriteString(`", "")`)

	return builder.String()
}

func (flag *Flag) IsList() bool {
	return flag.Desc.IsList()
}

func (flag *Flag) IsPrimitive() bool {
	return flag.Message == nil
}

func (flag *Flag) FieldName() string {
	return flag.GoName
}

func (flag *Flag) Name() string {
	return flag.GoName
}

func (flag *Flag) FieldNamePrivate() string {
	return flag.Desc.JSONName()
}

func (flag *Flag) Type() string {
	return flag.Field.GoIdent.GoName
}

func (flag *Flag) typ() string {
	if flag.Custom != nil {
		return flag.Custom.Type
	}
	return flag.Desc.Kind().String()
}

func (flag *Flag) IsPtr() bool {
	return flag.Desc.HasOptionalKeyword() || flag.Message != nil || flag.Custom != nil
}

func (flag *Flag) VarName() string {
	return flag.FieldNamePrivate() + "Flag"
}

type messageFlag struct {
	*protogen.Message
}

func (flag *messageFlag) Type() string {
	return flag.GoIdent.GoName
}

type enumFlag struct {
	*protogen.Enum
}

func (flag *enumFlag) Type() string {
	return flag.GoIdent.GoName
}

type oneOfFlag struct {
	*protogen.Oneof
}

func (flag *oneOfFlag) Type() string {
	return flag.GoIdent.GoName
}

func (flag *oneOfFlag) FieldName() string {
	return flag.GoName
}

func (flag *oneOfFlag) FieldNames() string {
	fields := make([]string, len(flag.Fields))
	for i, field := range flag.Fields {
		fields[i] = `"` + field.Desc.JSONName() + `"`
	}
	return strings.Join(fields, ", ")
}

type customFlag struct {
	Type string
}
