package v7

import "google.golang.org/protobuf/compiler/protogen"

type File struct {
	gen *protogen.GeneratedFile

	file     *protogen.File
	Services []*Service
}

func NewFile(gen *protogen.GeneratedFile, file *protogen.File) *File {
	f := &File{
		gen:      gen,
		file:     file,
		Services: make([]*Service, len(file.Services)),
	}

	for i, service := range file.Services {
		f.Services[i] = NewService(gen, service)
	}

	return f
}

func (f *File) Execute() error {
	header(f.gen, f.file)
	for _, service := range f.Services {
		if err := service.Execute(); err != nil {
			return err
		}
	}
	return generate(f.gen, "file", fileTemplate, f, nil)
}

const fileTemplate = `
func init() {
{{- range $service := .Services }}
{{- range $method := $service.Methods }}
	{{ $service.VariableName }}.AddCommand({{ $method.VariableName }})
{{- end }}
{{- end }}
}`
