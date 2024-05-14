package types

import (
	"log/slog"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

var (
	timestampMsg = &Message{
		isGenerated: true,
		Type: protogen.GoIdent{
			GoName:       "Timestamp",
			GoImportPath: "google.golang.org/protobuf/types/known/timestamppb",
		},
	}
	durationMsg = &Message{
		isGenerated: true,
		Type: protogen.GoIdent{
			GoName:       "Duration",
			GoImportPath: "google.golang.org/protobuf/types/known/durationpb",
		},
	}
	emptyMsg = &Message{
		isGenerated: true,
		Type: protogen.GoIdent{
			GoName:       "Empty",
			GoImportPath: "google.golang.org/protobuf/types/known/emptypb",
		},
	}
	anyMsg = &Message{
		isGenerated: true,
		Type: protogen.GoIdent{
			GoName:       "Any",
			GoImportPath: "google.golang.org/protobuf/types/known/anypb",
		},
	}
	structMsg = &Message{
		isGenerated: true,
		Type: protogen.GoIdent{
			GoName:       "Struct",
			GoImportPath: "google.golang.org/protobuf/types/known/structpb",
		},
	}
	messages = Messages{
		timestampMsg.Type: timestampMsg,
		durationMsg.Type:  durationMsg,
		structMsg.Type:    structMsg,
		anyMsg.Type:       anyMsg,
		emptyMsg.Type:     emptyMsg,
	}
)

type Flag struct {
	Name string
	*protogen.Field
}

func (f *Flag) Type(gen *protogen.GeneratedFile) string {
	qualify := qualifier(gen)
	var builder strings.Builder

	if f.Desc.IsList() {
		builder.WriteString("[]")
	}
	if f.Message != nil {
		builder.WriteRune('*')
		builder.WriteString(qualify(f.Message.GoIdent))
	}
	if f.Enum != nil {
		return qualify(protogen.GoIdent{GoImportPath: "github.com/adlerhurst/cli-client", GoName: title.String(f.Desc.Kind().String()) + "Flag[" + qualify(f.Enum.GoIdent) + "]"})
	}
	return qualify(protogen.GoIdent{GoImportPath: "github.com/adlerhurst/cli-client", GoName: title.String(f.Desc.Kind().String()) + "Flag"})
}

type Messages map[protogen.GoIdent]*Message

type Message struct {
	*protogen.Message

	Type     protogen.GoIdent
	FlagType string
	Flags    []*Flag

	isGenerated bool
}

func SetMessages(msg *protogen.Message) {
	if _, ok := messages[msg.GoIdent]; ok {
		return
	}

	message := &Message{
		Message:  msg,
		FlagType: msg.GoIdent.GoName + "Flag",
		Type:     msg.GoIdent,
		Flags:    make([]*Flag, len(msg.Fields)),
	}
	messages[message.GoIdent] = message
	slog.Info("message added", "key", message.GoIdent.String())

	for i, field := range msg.Fields {
		message.Flags[i] = &Flag{
			Name:  field.GoName + "Flag",
			Field: field,
		}

		if field.Message != nil {
			SetMessages(field.Message)
		}
	}

}

func GenerateMessages(plugin *protogen.Plugin, file *protogen.File) error {
	gen := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+"_cli_flags.pb.go", file.GoImportPath)
	header(gen, file)

	for _, message := range messages {
		if message.isGenerated {
			continue
		}
		if err := generate(gen, "message", messageTemplate, message); err != nil {
			return err
		}
	}

	return nil
}

var messageTemplate = `
type {{ .FlagType }} struct{
	*{{ qualify .Type }}

	changed bool
	set {{ ident "github.com/spf13/pflag" "FlagSet" | qualify }}

{{ range $flag := .Flags -}}
	{{ $flag.Name }} {{ $flag.Type gen }}
{{ end -}}
}
`
