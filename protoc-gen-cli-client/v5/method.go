package v5

import (
	"strings"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/compiler/protogen"
)

type Method struct {
	*protogen.Method
}

func (method *Method) Use() string {
	return strcase.ToKebab(method.GoName)
}

func (method *Method) Short() string {
	if len(method.Comments.Leading) == 0 {
		return ""
	}
	return strings.Split(string(method.Comments.Leading), "\n")[0]
}

func (method *Method) Long() string {
	if len(method.Comments.Leading) == 0 {
		return ""
	}
	return strings.TrimSuffix(string(method.Comments.Leading), "\n")
}

func GenerateMethod(gen *protogen.GeneratedFile, method *Method) error {
	return generate(gen, "method", methodTemplate, method)
}

func newMethod(gen *protogen.GeneratedFile, method *protogen.Method) *Method {
	return &Method{
		Method: method,
	}
}

var methodTemplate = `
	func init() {
		{{ .Parent.GoName }}Command.AddCommand({{ .Parent.GoName}}{{ .GoName }}Command)
	}
	var {{ .Parent.GoName}}{{ .GoName }}Command = &{{ ident "github.com/spf13/cobra" "Command" | qualify }}{
		Use: "{{ .Use }}",
		Short: "{{ .Short }}",
		Long: ` + "`{{ .Long }}`" + `,
	}
`
