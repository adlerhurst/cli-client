package v7

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/compiler/protogen"
)

type Service struct {
	gen *protogen.GeneratedFile

	service *protogen.Service
	Methods []*Method
}

func NewService(gen *protogen.GeneratedFile, service *protogen.Service) *Service {
	svc := &Service{
		gen:     gen,
		service: service,
		Methods: make([]*Method, len(service.Methods)),
	}

	for i, method := range service.Methods {
		svc.Methods[i] = NewMethod(gen, method)
	}

	return svc
}

func (s *Service) Execute() error {
	err := generate(s.gen, "service", serviceTemplate, s, nil)
	if err != nil {
		return err
	}
	for _, method := range s.Methods {
		if err := method.Execute(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Use() string {
	return strings.ReplaceAll(string(s.service.Desc.Name()), "_", "-")
}

func (s *Service) Short() string {
	return strings.Split(string(s.service.Comments.Leading), "\n")[0]
}

func (s *Service) Long() string {
	return string(s.service.Comments.Leading)
}

func (s *Service) VariableName() string {
	return s.service.GoName + "Command"
}

func bla() {
	cmd := cobra.Command{}
	_ = cmd
	fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	cmd.Flags().AddFlagSet(fs)
	cmd.AddCommand()
}

const serviceTemplate = `
var {{ .VariableName }} = &{{ ident "github.com/spf13/cobra" "Command" | qualify }}{
	Use: {{ .Use | quote }},
	Annotations: map[string]string{
		"type": "service",
	},
	TraverseChildren: true,
	Short:            {{ .Short | quote }},
	Long:             {{ .Long | quote }},
}
`
