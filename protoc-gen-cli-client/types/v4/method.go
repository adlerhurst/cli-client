package types

import (
	"log/slog"
	"os"

	option "github.com/adlerhurst/cli-client/protoc-gen-cli-client/option"
	"google.golang.org/protobuf/compiler/protogen"
)

type Method struct {
	*protogen.Method

	Command *Command
	Input   *Message
	Parent  *Service

	option *option.MethodCommand
}

func NewMethod(svc *Service, method *protogen.Method) *Method {
	methodCmd := method.Desc.Options().
		ProtoReflect().Get(option.E_SubCommand.TypeDescriptor()).
		Message().Interface().(*option.MethodCommand)
	if methodCmd == nil {
		methodCmd = &option.MethodCommand{
			Name: string(method.Desc.Name()),
		}
	} else if methodCmd.Name == "" {
		methodCmd.Name = string(method.Desc.Name())
	}

	if methodCmd.GetIgnored() {
		slog.Error("ignoring method based on user request", "name", method.Desc.Name())
		return nil
	}

	m := &Method{
		Method: method,
		option: methodCmd,
		Parent: svc,
	}
	m.Command = svc.cmd.CommandFromMethod(m)
	if m.Command == nil {
		slog.Info("failed to load parent command for method", "name", method.Desc.Name())
		os.Exit(1)
	}
	SetMessages(method.Input)

	return m
}

func (method *Method) Generate(plugin *protogen.Plugin, gen *protogen.GeneratedFile) error {
	if method.option.GetIgnored() {
		return nil
	}

	if err := generate(gen, "set_method_pre_run", setPreRun, method); err != nil {
		return err
	}

	return method.generateRun(gen)
}

var setPreRun = `
func init() {
	{{ .Command.VarName }}.PreRun = func(cmd *cobra.Command, args []string) {
		{{ .Command.VarName }}.Flags().Parse(args)
		// TODO: _ClientFilterCmdRequest.AddFlags({{ .Command.VarName }}.Flags())
		if {{ .Command.VarName }}.Flag("help").Changed {
			{{ .Command.VarName }}.Help()
			os.Exit(0)
		}
		// TODO: _ClientFilterCmdRequest.ParseFlags(cmd.Flags(), args)
	}
}
`

func (method *Method) generateRun(gen *protogen.GeneratedFile) error {
	if method.Desc.IsStreamingClient() && method.Desc.IsStreamingServer() {
		return generate(gen, "method_run_bidirectional_stream", runBidirectionalStreamTemplate, method)
	}
	if method.Desc.IsStreamingClient() {
		return generate(gen, "method_run_client_stream", runClientStreamTemplate, method)
	}
	if method.Desc.IsStreamingServer() {
		return generate(gen, "method_run_server_stream", runServerStreamTemplate, method)
	}
	return generate(gen, "run_serial", runSerialTemplate, method)
}

var runSerialTemplate = `
var (
	{{ .Command.Name }}Options []{{ ident "google.golang.org/grpc" "CallOption" | qualify}}
	{{ .Command.Run }} = func(cmd *{{ ident "github.com/spf13/cobra" "Command" | qualify }}, args []string) {
		conn := {{ ident "github.com/adlerhurst/cli-client" "Connection" | qualify }}(cmd.Context())
		defer conn.Close()
		client := New{{ .Parent.GoName }}Client(conn)

		// TODO: set request
		res, err := client.{{ .GoName }}(cmd.Context(), nil, {{ .Command.Name }}Options...)
		if err != nil {
			{{ logger }}.Error("unable to {{ .GoName }}", "cause", err)
			{{ ident "os" "Exit" | qualify }}(1)
		}
		{{ logger }}.Info("🎉 request succeeded", "result", res)
	}
)
`

var runClientStreamTemplate = `
var (
	{{ .Command.Name }}Options []{{ ident "google.golang.org/grpc" "CallOption" | qualify}}
	{{ .Command.Run }} = func(cmd *{{ ident "github.com/spf13/cobra" "Command" | qualify }}, args []string) {
		conn := {{ ident "github.com/adlerhurst/cli-client" "Connection" | qualify }}(cmd.Context())
		defer conn.Close()
		client := New{{ .Parent.GoName }}Client(conn)

		stream, err := client.{{ .GoName }}(cmd.Context(), {{ .Command.Name }}Options...)
		if err != nil {
			{{ logger }}.Error("unable to {{ .GoName }}", "cause", err)
			{{ ident "os" "Exit" | qualify }}(1)
		}
		defer stream.CloseSend()
		// TODO: set request
		if err := stream.Send(nil); err != nil{
			{{ logger }}.Error("unable to send", "cause", err)
			os.Exit(1)
		}

		res, err := stream.CloseAndRecv()
		if err != nil{
			{{ logger }}.Error("unable to receive and close", "cause", err)
			os.Exit(1)
		}
		{{ logger }}.Info("🎉 request succeeded", "result", res.String())
	}
)
`

var runServerStreamTemplate = `
var (
	{{ .Command.Name }}Options []{{ ident "google.golang.org/grpc" "CallOption" | qualify}}
	{{ .Command.Run }} = func(cmd *{{ ident "github.com/spf13/cobra" "Command" | qualify }}, args []string) {
		conn := {{ ident "github.com/adlerhurst/cli-client" "Connection" | qualify }}(cmd.Context())
		defer conn.Close()
		client := New{{ .Parent.GoName }}Client(conn)

		// TODO: set request
		stream, err := client.{{ .GoName }}(cmd.Context(), nil, {{ .Command.Name }}Options...)
		if err != nil {
			{{ logger }}.Error("unable to {{ .GoName }}", "cause", err)
			{{ ident "os" "Exit" | qualify }}(1)
		}
		defer stream.CloseSend()
		
		res, err := stream.Recv()
		if err != nil{
			{{ logger }}.Error("unable to receive", "cause", err)
			os.Exit(1)
		}
		{{ logger }}.Info("🎉 request succeeded", "result", res.String())
	}
)
`

var runBidirectionalStreamTemplate = `
var (
	{{ .Command.Name }}Options []{{ ident "google.golang.org/grpc" "CallOption" | qualify}}
	{{ .Command.Run }} = func(cmd *{{ ident "github.com/spf13/cobra" "Command" | qualify }}, args []string) {
		conn := {{ ident "github.com/adlerhurst/cli-client" "Connection" | qualify }}(cmd.Context())
		defer conn.Close()
		client := New{{ .Parent.GoName }}Client(conn)

		stream, err := client.{{ .GoName }}(cmd.Context(), {{ .Command.Name }}Options...)
		if err != nil {
			{{ logger }}.Error("unable to {{ .GoName }}", "cause", err)
			{{ ident "os" "Exit" | qualify }}(1)
		}
		defer stream.CloseSend()

		// TODO: set request
		if err := stream.Send(nil); err != nil {
			{{ logger }}.Error("unable to send", "cause", err)
			os.Exit(1)
		}

		res, err := stream.Recv()
		if err != nil {
			{{ logger }}.Error("unable to receive", "cause", err)
			os.Exit(1)
		}
		{{ logger }}.Info("🎉 request succeeded", "result", res.String())
	}
)
`
