package types

import (
	_ "embed"
	"strings"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var messages = Messages{}

type Messages map[protoreflect.FullName]*Message

type Message struct {
	*protogen.Message
	Flags  []*Flag
	OneOfs []*OneOf
}

type OneOf struct {
	*oneOfFlag
	Flags []*Flag
}

func (msg *Message) NestedFlagNames() string {
	flagNames := make([]string, 0, len(msg.Flags))
	for _, flag := range msg.Flags {
		if flag.Message == nil {
			continue
		}
		flagNames = append(flagNames, `"`+flag.Name()+`"`)
	}

	return strings.Join(flagNames, ", ")
}

func (msg *Message) HasMessageFlags() bool {
	for _, flag := range msg.Flags {
		if flag.Message != nil {
			return true
		}
	}
	return false
}

func SetMessages(service *protogen.Service) {
	for _, method := range service.Methods {
		setMessage(method.Input)
	}
}

var customFlags = map[protoreflect.FullName]*customFlag{
	"google.protobuf.Timestamp": {
		Type: "timestamp",
	},
	"google.protobuf.Duration": {
		Type: "duration",
	},
	"google.protobuf.Struct": {
		Type: "struct",
	},
	"google.protobuf.Any": {
		Type: "any",
	},
}

func setMessage(message *protogen.Message) {
	if _, ok := messages[message.Desc.FullName()]; ok {
		return
	}

	msg := &Message{
		Message: message,
		Flags:   make([]*Flag, len(message.Fields)),
		OneOfs:  make([]*OneOf, 0, len(message.Oneofs)),
	}
	messages[message.Desc.FullName()] = msg

	for _, oneOf := range message.Oneofs {
		if isOneOfOptional(oneOf, message.Fields) {
			continue
		}
		oo := &OneOf{
			oneOfFlag: &oneOfFlag{Oneof: oneOf},
			Flags:     make([]*Flag, len(oneOf.Fields)),
		}
		msg.OneOfs = append(msg.OneOfs, oo)
		for i, oneOfField := range oneOf.Fields {
			flag := &Flag{
				Field: oneOfField,
				OneOf: &oneOfFlag{Oneof: oneOf},
			}
			oo.Flags[i] = flag
			if oneOfField.Message != nil {
				if wellknownFlag, ok := customFlags[oneOfField.Message.Desc.FullName()]; ok {
					flag.Custom = wellknownFlag
					continue
				}
				flag.Message = &messageFlag{Message: oneOfField.Message}
				setMessage(oneOfField.Message)
			}
		}
	}

	for i, field := range message.Fields {
		flag := &Flag{
			Field: field,
		}
		msg.Flags[i] = flag

		for _, oneOf := range msg.OneOfs {
			if field.Oneof == nil || oneOf.GoName != field.Oneof.GoName {
				continue
			}
			flag.OneOf = oneOf.oneOfFlag
			break
		}

		if field.Enum != nil {
			flag.Enum = &enumFlag{Enum: field.Enum}
			continue
		}

		if field.Message != nil {
			if wellknownFlag, ok := customFlags[field.Message.Desc.FullName()]; ok {
				flag.Custom = wellknownFlag
				continue
			}
			flag.Message = &messageFlag{Message: field.Message}
			setMessage(field.Message)
			continue
		}
	}
}

var (
	//go:embed message.go.tmpl
	messageDefinition string
	messageTemplate   = template.Must(template.New("message").Parse(messageDefinition))
)

func GenerateMessages(plugin *protogen.Plugin, file *protogen.File) (err error) {
	if len(messages) == 0 {
		return nil
	}
	gen := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+"_cli_flags.go", file.GoImportPath)

	header(gen, file)
	messages.imports(gen)

	err = executeTemplate(gen, messageTemplate, messages)
	if err != nil {
		return err
	}

	return nil
}

func (Messages) imports(gen *protogen.GeneratedFile) {
	gen.QualifiedGoIdent(protogen.GoImportPath("os").Ident("os"))
	gen.QualifiedGoIdent(protogen.GoImportPath("github.com/spf13/pflag").Ident("pflag"))
}

func isOneOfOptional(oneOf *protogen.Oneof, fields []*protogen.Field) bool {
	for _, field := range fields {
		if field.Oneof == nil {
			continue
		}
		if field.Oneof.GoName == oneOf.GoName {
			return field.Desc.HasOptionalKeyword()
		}
	}
	return false
}
