package types

import (
	option "github.com/adlerhurst/cli-client/protoc-gen-cli-client/option"
	"google.golang.org/protobuf/compiler/protogen"
)

type Service struct {
	*protogen.Service

	cmd     *Command
	methods []*Method
	option  *option.ServiceCommand
}

func NewService(svc *protogen.Service) *Service {
	service := &Service{
		Service: svc,
		methods: make([]*Method, 0, len(svc.Methods)),
		cmd:     NewCommandFromService(svc),
		option: svc.Desc.Options().
			ProtoReflect().Get(option.E_Command.TypeDescriptor()).
			Message().Interface().(*option.ServiceCommand),
	}

	for _, method := range svc.Methods {
		m := NewMethod(service, method)
		if m == nil {
			continue
		}
		service.methods = append(service.methods, m)
	}

	return service
}

func (svc *Service) Generate(plugin *protogen.Plugin, file *protogen.File) error {
	if svc.option.GetIgnored() {
		return nil
	}
	gen := plugin.NewGeneratedFile(svc.filename(file.GeneratedFilenamePrefix), file.GoImportPath)
	header(gen, file)

	if err := svc.cmd.Generate(gen); err != nil {
		return err
	}

	for _, method := range svc.methods {
		if err := method.Generate(plugin, gen); err != nil {
			return err
		}
	}

	return nil
}

func (svc *Service) filename(prefix string) string {
	return prefix + "_cli.pb.go"
}
