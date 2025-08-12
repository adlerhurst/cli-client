package v7

import (
	"log"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

type Method struct {
	gen *protogen.GeneratedFile

	method *protogen.Method

	request *Message
}

func NewMethod(gen *protogen.GeneratedFile, method *protogen.Method) *Method {
	return &Method{
		gen:     gen,
		method:  method,
		request: NewMessage(gen, method.Input),
	}
}

func (m *Method) Execute() error {
	return generate(m.gen, "method", methodTemplate, m, nil)
	// TODO: add command to service command
	// TODO: {--path to payload | --json (json object as params) | --request (flags)}
}

func (m *Method) Use() string {
	return strings.ReplaceAll(string(m.method.Desc.Name()), "_", "-")
}

func (m *Method) Short() string {
	return strings.Split(string(m.method.Comments.Leading), "\n")[0]
}

func (m *Method) Long() string {
	return string(m.method.Comments.Leading)
}

func (m *Method) VariableName() string {
	return m.method.GoName + "Command"
}

func (m *Method) RequestName() string {
	return m.request.message.GoIdent.GoName
}

func (m *Method) Request() string {
	log.Printf("request (%q): %q\n", m.request.message.GoIdent.String(), qualifier(m.gen)(ident(m.request.message.GoIdent.GoImportPath, m.request.message.GoIdent.GoName)))

	return qualifier(m.gen)(ident(m.request.message.GoIdent.GoImportPath, m.request.message.GoIdent.GoName))
}

const methodTemplate = `
var (
	_{{ .RequestName }} {{ .Request }}

	{{ .VariableName }} = &{{ ident "github.com/spf13/cobra" "Command" | qualify }}{
		Use: {{ .Use | quote }},
		TraverseChildren:   true,
		DisableFlagParsing: true,
		FParseErrWhitelist: {{ ident "github.com/spf13/cobra" "FParseErrWhitelist" | qualify }}{
			UnknownFlags: true,
		},
		Annotations: map[string]string{
			"type": "method",
		},
		Short:  {{ .Short | quote }},
		Long:   {{ .Long | quote }},
		Args:	cobra.MinimumNArgs(1),
		PreRun: {{ ident "github.com/adlerhurst/cli-client/protoc-gen-cli-client/v7/parse" "PreRun" | qualify }}(&_{{ .RequestName }}),
		Run: func(cmd *{{ ident "github.com/spf13/cobra" "Command" | qualify }}, args []string) {
			{{ ident "fmt" "Println" | qualify }}("{{ .VariableName }} called with args: ", args)
		},
	}
)`
