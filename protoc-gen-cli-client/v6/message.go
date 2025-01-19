package v6

import (
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GenerateMessages(plugin *protogen.Plugin, file *protogen.File) error {
	gen := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+"_cli_flags.v6.pb.go", file.GoImportPath)

	header(gen, file)

	for _, service := range file.Services {
		for _, method := range service.Methods {
			if err := generateMessage(gen, method.Input); err != nil {
				return err
			}
		}
	}

	return nil
}

var generatedMessages = map[protoreflect.FullName]bool{
	(*timestamppb.Timestamp)(nil).ProtoReflect().Descriptor().FullName(): true,
	(*durationpb.Duration)(nil).ProtoReflect().Descriptor().FullName():   true,
	(*structpb.Struct)(nil).ProtoReflect().Descriptor().FullName():       true,
	(*anypb.Any)(nil).ProtoReflect().Descriptor().FullName():             true,
}

func alreadyGenerated(name protoreflect.FullName) bool {
	return generatedMessages[name]
}

func generateMessage(gen *protogen.GeneratedFile, msg *protogen.Message) error {
	if alreadyGenerated(msg.Desc.FullName()) {
		return nil
	}

	for _, subMsg := range msg.Messages {
		if subMsg.Desc.IsMapEntry() {
			continue
		}
		if err := generateMessage(gen, subMsg); err != nil {
			return err
		}
	}

	return generate(
		gen,
		"message",
		messageTemplate,
		&message{Message: msg},
		template.FuncMap{
			"valueConstructor": func(field *protogen.Field) string {
				if field.Desc.IsList() {
					return qualifier(gen)(ident("github.com/adlerhurst/cli-client", "NewSliceValue"))
				}
				return qualifier(gen)(ident("github.com/adlerhurst/cli-client", "NewValue"))
			},
			"messageGenerated": func(msg *message) string {
				generatedMessages[msg.Desc.FullName()] = true
				return ""
			},
		},
	)
}

type message struct {
	*protogen.Message
}

func (msg *message) PrimitiveFields() []*protogen.Field {
	fields := make([]*protogen.Field, 0, len(msg.Fields))
	for _, field := range msg.Fields {
		if field.Oneof != nil {
			continue
		}
		switch field.Desc.Kind() {
		case protoreflect.MessageKind, protoreflect.GroupKind:
			continue
		default:
			fields = append(fields, field)
		}
	}

	return fields
}

func (msg *message) MessageFields() []*protogen.Field {
	fields := make([]*protogen.Field, 0, len(msg.Fields))
	for _, field := range msg.Fields {
		if (field.Oneof != nil && !field.Desc.HasOptionalKeyword()) || field.Message == nil || field.Desc.IsList() || field.Desc.IsMap() {
			continue
		}
		fields = append(fields, field)
	}
	return fields
}

// func (msg *{{ qualify .GoIdent}}) ParseArgs(cmd *{{ident "github.com/spf13/cobra" "Command" | qualify}}, args []string) error {

//		return nil
//	}
var messageTemplate = `

func (msg *{{ qualify .GoIdent}}) SetFlags(set *{{ident "github.com/spf13/pflag" "FlagSet" | qualify}}) {
	{{ $ = . -}}
	{{ range $field := .PrimitiveFields -}}
	{{ valueConstructor $field }}(&msg.{{ $field.GoName }}).Add(set, "{{ $field.GoName }}", "", "")
	{{ end -}}

	{{ messageGenerated . -}}

	{{ range $field := .MessageFields -}}
	// {{ $field.GoName }}
	_{{ $field.GoName }}Flags := {{ident "github.com/spf13/pflag" "NewFlagSet" | qualify}}("{{ $field.GoName }}", {{ident "github.com/spf13/pflag" "ContinueOnError" | qualify}})
	msg.{{ $field.GoName }}.SetFlags(_{{ $field.GoName }}Flags)
	{{ end -}}
}
`
