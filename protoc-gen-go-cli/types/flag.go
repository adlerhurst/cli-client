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

func (flag *Flag) IsList() bool {
	return flag.Desc.IsList()
}

func (flag *Flag) FieldName() string {
	return flag.GoName
}

func (flag *Flag) FieldNamePrivate() string {
	return flag.Desc.JSONName()
}

func (flag *Flag) Name() string {
	return strings.ReplaceAll(string(flag.Desc.Name()), "_", "-")
}

func (flag *Flag) Type() string {
	return flag.Field.GoIdent.GoName
}

func (flag *Flag) Ident() protogen.GoIdent {
	return flag.ident(false)
}

func (flag *Flag) ConstructorIdent() protogen.GoIdent {
	return flag.ident(true)
}

func (flag *Flag) ident(isConstructor bool) protogen.GoIdent {
	if flag.Custom != nil {
		name := flag.Custom.Type
		if isConstructor {
			name = flag.Custom.Constructor
		}
		return protogen.GoIdent{
			GoName:       name,
			GoImportPath: flag.Custom.ImportPath,
		}
	}
	var builder strings.Builder

	if isConstructor {
		builder.WriteString("New")
	}
	builder.WriteString(title.String(flag.Desc.Kind().String()))

	if flag.IsList() {
		builder.WriteString("Slice")
	}

	builder.WriteString("Parser")
	if flag.Enum != nil {
		builder.WriteString("[")
		builder.WriteString(flag.Enum.Type())
		builder.WriteString("]")
	}
	return protogen.GoIdent{
		GoName:       builder.String(),
		GoImportPath: "github.com/adlerhurst/cli-client",
	}
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
		fields[i] = `"` + strings.ReplaceAll(string(field.Desc.Name()), "_", "-") + `"`
	}
	return strings.Join(fields, ", ")
}

type customFlag struct {
	Type        string
	Constructor string
	ImportPath  protogen.GoImportPath
}
