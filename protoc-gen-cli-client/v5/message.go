package v5

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type Message struct {
	*protogen.Message

	Fields []*Field
	Flags  map[string]*Flag
}

func newMessage(gen *protogen.GeneratedFile, message *protogen.Message) *Message {
	fields := make([]*Field, 0, len(message.Fields))
	flags := make(map[string]*Flag, len(message.Fields))

	for i, field := range message.Fields {
		if oneOf := field.Desc.ContainingOneof(); oneOf != nil {
			oneOf.Fields().Get(0).FullName()
		}

		_ = i
		// fields = append(fields, newField(gen.QualifiedGoIdent, field))
		flags[field.GoName] = newFlag(qualifier(gen), field)
	}

	return &Message{
		Message: message,
		Fields:  fields,
		Flags:   flags,
	}
}

func (msg *Message) FlagType() string {
	return msg.GoIdent.GoName + "Flag"
}

func GenerateMessages(plugin *protogen.Plugin, file *protogen.File) error {
	gen := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+"_cli_flags.pb.go", file.GoImportPath)

	header(gen, file)

	for _, message := range file.Messages {
		if err := GenerateMessage(gen, newMessage(gen, message)); err != nil {
			return err
		}
	}

	return nil
}

func GenerateMessage(gen *protogen.GeneratedFile, message *Message) error {
	// message.Fields[0].Desc.Kind().GoString()
	return generate(gen, "message", messageTemplate, message)
}

var messageTemplate = `
type {{ .FlagType }} struct {
	set {{ ident "github.com/spf13/pflag" "FlagSet" | qualify }}
	
	 {{ range $flag := .Flags -}}
	 	
	 	_{{ $flag.GoName }} {{ident "github.com/adlerhurst/cli-client" $flag.FlagType | qualify}}[{{ $flag.Type }}]
	 {{ end -}}
}

func (x *{{ .FlagType }}) Value() *{{ qualify .GoIdent }} {
	return &{{ qualify .GoIdent }}{
		{{ range $field := .Fields -}}
			{{ $field.GoName }}: x._{{ $field.GoName }}.Value(),
		{{ end -}}
	}
}
`

// var generatedRequests []string

// func GenerateRequests(plugin *protogen.Plugin, service *protogen.File) error {

// }

// func generateRequestsForService(plugin *protogen.Plugin, service *protogen.Service) error {
// 	for _, method := range service.Methods {
// 		if err := generateMessage(plugin, method.Input); err != nil {
// 			return err
// 		}
// 	}
// }

// func generateMessage(plugin *protogen.Plugin, message *protogen.Message) error {
// 	for _, subMessage := range message.Messages {
// 		if err := generateMessage(plugin, subMessage); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
