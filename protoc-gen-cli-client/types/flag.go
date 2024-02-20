package types

import (
	"strings"

	option "github.com/adlerhurst/cli-client/protoc-gen-cli-client/option"
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

func (flag *Flag) Name() (name string) {
	f := flag.Field.Desc.Options().ProtoReflect().Get(option.E_Flag.TypeDescriptor()).Message().Interface().(*option.FlagOption)
	if f != nil {
		name = f.Name
	}
	if name == "" {
		name = string(flag.Desc.Name())
	}

	return strings.ReplaceAll(name, "_", "-")
}

func (flag *Flag) Type() protogen.GoIdent {
	return flag.Field.GoIdent
}

func (flag *Flag) Ident(gen *protogen.GeneratedFile) protogen.GoIdent {
	return flag.ident(gen, false)
}

func (flag *Flag) ConstructorIdent(gen *protogen.GeneratedFile) protogen.GoIdent {
	return flag.ident(gen, true)
}

func (flag *Flag) ident(gen *protogen.GeneratedFile, isConstructor bool) protogen.GoIdent {
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
		builder.WriteString(qualifier(gen)(flag.Enum.GoIdent))
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

func (flag *messageFlag) Type() protogen.GoIdent {
	return flag.GoIdent
}

type enumFlag struct {
	*protogen.Enum
}

func (flag *enumFlag) Type() protogen.GoIdent {
	return flag.GoIdent
}

type oneOfFlag struct {
	*protogen.Oneof
}

func (flag *oneOfFlag) Type() protogen.GoIdent {
	return flag.GoIdent
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
