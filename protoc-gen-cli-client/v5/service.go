package v5

import (
	"strings"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/compiler/protogen"
)

type Service struct {
	*protogen.Service
}

func (svc *Service) Use() string {
	return strcase.ToKebab(svc.GoName)
}

func (svc *Service) Short() string {
	if len(svc.Comments.Leading) == 0 {
		return ""
	}
	return strings.Split(string(svc.Comments.Leading), "\n")[0]
}

func (svc *Service) Long() string {
	if len(svc.Comments.Leading) == 0 {
		return ""
	}
	return strings.TrimSuffix(string(svc.Comments.Leading), "\n")
}

func GenerateServices(plugin *protogen.Plugin, file *protogen.File) error {
	gen := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+"_cli_command.pb.go", file.GoImportPath)

	header(gen, file)

	for _, service := range file.Services {
		if len(service.Methods) == 0 {
			continue
		}

		if err := GenerateService(gen, newService(gen, service)); err != nil {
			return err
		}

		for _, method := range service.Methods {
			if err := GenerateMethod(gen, newMethod(gen, method)); err != nil {
				return err
			}
		}
	}

	return nil
}

func GenerateService(gen *protogen.GeneratedFile, service *Service) error {
	return generate(gen, "service", serviceTemplate, service)
}

func newService(gen *protogen.GeneratedFile, service *protogen.Service) *Service {
	return &Service{
		Service: service,
	}
}

var serviceTemplate = `
	var {{ .GoName}}Command = &{{ ident "github.com/spf13/cobra" "Command" | qualify }}{
		Use: "{{ .Use }}",
		Short: "{{ .Short }}",
		Long: ` + "`{{ .Long }}`" + `,
	}
`
