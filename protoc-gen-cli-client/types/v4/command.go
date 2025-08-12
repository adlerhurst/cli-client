package types

// import (
// 	"strings"

// 	option "github.com/adlerhurst/cli-client/protoc-gen-cli-client/option"
// 	"google.golang.org/protobuf/compiler/protogen"
// )

// type Command struct {
// 	VarName string
// 	Use     string
// 	Short   string
// 	Long    string
// 	Subs    []*Command
// 	Parent  *Command
// 	Run     string

// 	Name   string
// 	isUsed bool
// }

// func NewCommandFromService(svc *protogen.Service) *Command {
// 	serviceCmd := svc.Desc.Options().
// 		ProtoReflect().Get(option.E_Command.TypeDescriptor()).
// 		Message().Interface().(*option.ServiceCommand)

// 	if serviceCmd == nil {
// 		serviceCmd = &option.ServiceCommand{
// 			Name: string(svc.Desc.Name()),
// 		}
// 	} else if serviceCmd.Name == "" {
// 		serviceCmd.Name = string(svc.Desc.Name())
// 	}

// 	name := title.String(serviceCmd.GetName())
// 	cmd := &Command{
// 		VarName: name + "Cmd",
// 		Use:     serviceCmd.GetName(),
// 		Long:    string(svc.Comments.Leading),
// 		Short:   shortFromCommentSet(svc.Comments),
// 		Subs:    make([]*Command, len(serviceCmd.SubCommands)),

// 		Name: name,
// 	}

// 	for i, sub := range serviceCmd.SubCommands {
// 		cmd.Subs[i] = newCommandFromSubCommand(sub, cmd)
// 	}

// 	return cmd
// }

// func (cmd *Command) Generate(gen *protogen.GeneratedFile) error {
// 	if !cmd.isUsed {
// 		return nil
// 	}
// 	if err := generate(gen, "command", commandTemplate, cmd); err != nil {
// 		return err
// 	}

// 	if cmd.Parent != nil {
// 		if err := generate(gen, "init_sub_commands", initSubCommand, cmd); err != nil {
// 			return err
// 		}
// 	}

// 	for _, sub := range cmd.Subs {
// 		if err := sub.Generate(gen); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// var commandTemplate = `
// var {{ .VarName }} = &{{ ident "github.com/spf13/cobra" "Command" | qualify }}{
// 	Use:                "{{.Use}}",
// 	Short:              "{{.Short}}",
// 	Long:               ` + "`{{.Long}}`" + `,
// {{ if .Run -}}
// 	Run:                {{.Run}},
// {{ end -}}
// 	FParseErrWhitelist: {{ ident "github.com/spf13/cobra" "FParseErrWhitelist" | qualify }}{UnknownFlags: true},
// 	DisableFlagParsing: true,
// }
// `

// var initSubCommand = `
// func init() {
// 	{{ .Parent.VarName }}.AddCommand(
// 		{{ .VarName }},
// 	)
// }
// `

// func (cmd *Command) CommandFromMethod(method *Method) *Command {
// 	parent := cmd.methodParent(method.option.Parent)
// 	if parent == nil {
// 		return nil
// 	}

// 	name := parent.Name + strings.ReplaceAll(title.String(method.option.Name), "-", "")
// 	methodCmd := &Command{
// 		VarName: name + "Cmd",
// 		Use:     lower.String(method.option.Name),
// 		isUsed:  !method.option.GetIgnored(),
// 		Long:    string(method.Comments.Leading),
// 		Short:   shortFromCommentSet(method.Comments),
// 		Parent:  parent,
// 		Run:     name + "Run",

// 		Name: name,
// 	}
// 	parent.Subs = append(parent.Subs, methodCmd)
// 	return methodCmd
// }

// func (cmd *Command) methodParent(parent *option.MethodCommand_ParentCommand) *Command {
// 	if parent == nil || parent.GetName() == cmd.Use {
// 		cmd.isUsed = true
// 		return cmd
// 	}

// 	for _, sub := range cmd.Subs {
// 		if cmd := sub.methodParent(parent); cmd != nil {
// 			cmd.isUsed = true
// 			return cmd
// 		}
// 	}

// 	return nil
// }

// func newCommandFromSubCommand(subCommand *option.ServiceCommand_SubCommand, parent *Command) *Command {
// 	name := parent.Name + title.String(subCommand.GetName())
// 	cmd := &Command{
// 		VarName: name + "Cmd",
// 		Use:     subCommand.GetName(),
// 		Long:    subCommand.GetLongDescription(),
// 		Short:   subCommand.GetShortDescription(),
// 		Subs:    make([]*Command, len(subCommand.SubCommands)),
// 		Parent:  parent,

// 		Name: name,
// 	}

// 	for i, sub := range subCommand.SubCommands {
// 		cmd.Subs[i] = newCommandFromSubCommand(sub, cmd)
// 	}

// 	return cmd
// }

// func shortFromCommentSet(comments protogen.CommentSet) string {
// 	if len(comments.LeadingDetached) > 0 {
// 		return string(comments.LeadingDetached[0])
// 	}
// 	if lines := strings.Split(string(comments.Leading), "\n"); len(lines) > 0 {
// 		return lines[0]
// 	}

// 	return ""
// }
