package types

import (
	_ "embed"
	"regexp"
	"strings"

	option "github.com/adlerhurst/cli-client/protoc-gen-go-cli/option"
	"google.golang.org/protobuf/compiler/protogen"
)

type Service struct {
	*protogen.Service
	Methods []*Method
}

func NewService(service *protogen.Service) *Service {
	svc := &Service{
		Service: service,
		Methods: make([]*Method, len(service.Methods)),
	}

	for i, method := range service.Methods {
		svc.Methods[i] = MethodFromProto(svc, method)
	}

	return svc
}

var (
	//go:embed service.tmpl
	serviceDefinition string
)

func (svc *Service) Generate(plugin *protogen.Plugin, file *protogen.File) error {
	gen := plugin.NewGeneratedFile(svc.filename(file.GeneratedFilenamePrefix), file.GoImportPath)

	header(gen, file)
	if err := executeTemplate(gen, "service", serviceDefinition, svc); err != nil {
		return err
	}

	for _, method := range svc.Methods {
		err := method.Generate(plugin, gen)
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *Service) Use() string {
	re := regexp.MustCompile(`[a-z][A-Z]`)
	name := svc.name()

	matches := re.FindAllStringIndex(name, -1)
	for i := len(matches) - 1; i >= 0; i-- {
		if matches[i][0] == 0 {
			continue
		}
		name = name[:matches[i][0]+1] + "-" + name[matches[i][0]+1:]
	}

	return strings.ToLower(name)
}

func (svc *Service) Public() string {
	return title.String(svc.name())
}

func (svc *Service) Short() string {
	short, _, _ := strings.Cut(string(svc.Comments.Leading), "\n")
	return escapeComment(removeWhitespaces(short))
}

func (svc *Service) Long() string {
	return escapeComment(removeWhitespaces(string(svc.Comments.Leading)))
}

func (svc *Service) VarName() string {
	return svc.Public() + "Cmd"
}

func (svc *Service) filename(prefix string) string {
	var builder strings.Builder

	builder.WriteString(prefix)
	builder.WriteRune('_')
	builder.WriteString(string(svc.Desc.Name()))
	builder.WriteString("_cli.pb.go")

	return builder.String()
}

func (svc *Service) name() (name string) {
	command := svc.Service.Desc.Options().ProtoReflect().Get(option.E_Command.TypeDescriptor()).Message().Interface().(*option.CommandOption)
	if command != nil {
		name = command.Name
	}
	if name == "" {
		name = string(svc.Desc.Name())
	}
	return name
}
