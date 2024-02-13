package types

import (
	_ "embed"
	"log"
	"regexp"
	"strings"

	option "github.com/adlerhurst/cli-client/protoc-gen-cli-client/option"

	"google.golang.org/protobuf/compiler/protogen"
)

type Method struct {
	Parent *Service
	*protogen.Method
	Request protogen.GoIdent
}

func MethodFromProto(parent *Service, method *protogen.Method) *Method {
	log.Println("ueli", method.Desc.IsStreamingClient(), method.Desc.IsStreamingServer())
	m := &Method{
		Parent:  parent,
		Method:  method,
		Request: method.Input.GoIdent,
	}

	return m
}

var (
	//go:embed method.tmpl
	methodDefinition string
)

func (method *Method) Generate(plugin *protogen.Plugin, gen *protogen.GeneratedFile) error {
	return executeTemplate(gen, "method", methodDefinition, method)
}

func (method *Method) IsServerSideStream() bool {
	return method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient()
}

func (method *Method) IsClientSideStream() bool {
	return method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer()
}

func (method *Method) IsBidirectionalStream() bool {
	return method.Desc.IsStreamingClient() && method.Desc.IsStreamingServer()
}

func (method *Method) IsNoStream() bool {
	return !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer()
}

func (method *Method) Use() string {
	re := regexp.MustCompile(`[a-z][A-Z]`)
	name := method.name()

	matches := re.FindAllStringIndex(name, -1)
	for i := len(matches) - 1; i >= 0; i-- {
		if matches[i][0] == 0 {
			continue
		}
		name = name[:matches[i][0]+1] + "-" + name[matches[i][0]+1:]
	}

	return strings.ToLower(name)
}

func (method *Method) Short() string {
	short, _, _ := strings.Cut(string(method.Comments.Leading), "\n")
	return escapeComment(removeWhitespaces(short))
}

func (method *Method) Long() string {
	return escapeComment(removeWhitespaces(string(method.Comments.Leading) + string(method.Comments.Trailing)))
}

func (method *Method) VarName() string {
	return method.public() + "Cmd"
}

func (method *Method) public() string {
	return method.Parent.Public() + strings.ReplaceAll(title.String(method.name()), " ", "")
}

func (method *Method) name() (name string) {
	subCommand := method.Method.Desc.Options().ProtoReflect().Get(option.E_SubCommand.TypeDescriptor()).Message().Interface().(*option.CommandOption)
	if subCommand != nil {
		name = subCommand.Name
	}
	if name == "" {
		name = string(method.Desc.Name())
	}
	return name
}
