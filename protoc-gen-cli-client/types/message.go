package types

import (
	_ "embed"
	"log"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Message struct {
	*protogen.Message
	Flags  []*Flag
	OneOfs []*OneOf
}

type OneOf struct {
	*oneOfFlag
	Flags []*Flag
}

func (msg *Message) NestedFlagNames() string {
	flagNames := make([]string, 0, len(msg.Flags))
	for _, flag := range msg.Flags {
		if flag.Message == nil {
			continue
		}
		flagNames = append(flagNames, `"`+flag.Name()+`"`)
	}

	return strings.Join(flagNames, ", ")
}

func (msg *Message) HasMessageFlags() bool {
	for _, flag := range msg.Flags {
		if flag.Message != nil {
			return true
		}
	}
	return false
}

func SetMessages(service *protogen.Service) {
	// for _, method := range service.Methods {
	// 	setMessage(method.Input)
	// }
}

// func SetMessagesFromFile(file *protogen.File) {
// 	for _, message := range file.Messages {
// 		setMessage(message)
// 	}
// }

var customFlags = map[protoreflect.FullName]*customFlag{
	"google.protobuf.Timestamp": {
		Type:        "TimestampParser",
		Constructor: "NewTimestampParser",
		ImportPath:  "github.com/adlerhurst/cli-client",
	},
	"google.protobuf.Duration": {
		Type:        "DurationParser",
		Constructor: "NewDurationParser",
		ImportPath:  "github.com/adlerhurst/cli-client",
	},
	"google.protobuf.Struct": {
		Type:        "StructParser",
		Constructor: "NewStructParser",
		ImportPath:  "github.com/adlerhurst/cli-client",
	},
	"google.protobuf.Any": {
		Type:        "AnyParser",
		Constructor: "NewAnyParser",
		ImportPath:  "github.com/adlerhurst/cli-client",
	},
}

func newMessage(msgs Messages, message *protogen.Message) Messages {
	if _, ok := msgs[message.GoIdent]; ok {
		return msgs
	}

	msgs[message.GoIdent] = &Message{
		Message: message,
		Flags:   make([]*Flag, 0, len(message.Fields)),
		OneOfs:  make([]*OneOf, 0, len(message.Oneofs)),
	}

	for _, oneOf := range message.Oneofs {
		if isOneOfOptional(oneOf, message.Fields) {
			continue
		}
		oo := &OneOf{
			oneOfFlag: &oneOfFlag{Oneof: oneOf},
			Flags:     make([]*Flag, 0, len(oneOf.Fields)),
		}
		msgs[message.GoIdent].OneOfs = append(msgs[message.GoIdent].OneOfs, oo)
		for _, oneOfField := range oneOf.Fields {
			flag := &Flag{
				Field: oneOfField,
				OneOf: &oneOfFlag{Oneof: oneOf},
			}
			oo.Flags = append(oo.Flags, flag)
			if oneOfField.Message != nil {
				if wellknownFlag, ok := customFlags[oneOfField.Message.Desc.FullName()]; ok {
					flag.Custom = wellknownFlag
					continue
				}
				flag.Message = &messageFlag{Message: oneOfField.Message}
				// newMessage(msgs, oneOfField.Message)
			}
		}
	}

	for _, nested := range message.Messages {
		if nested.Desc.IsMapEntry() {
			log.Println("maps are currently unsupported")
			continue
		}
		newMessage(msgs, nested)
	}

	for _, field := range message.Fields {
		if field.Desc.IsMap() {
			log.Println("maps are currently unsupported")
			continue
		}
		flag := &Flag{
			Field: field,
		}
		msgs[message.GoIdent].Flags = append(msgs[message.GoIdent].Flags, flag)

		for _, oneOf := range msgs[message.GoIdent].OneOfs {
			if field.Oneof == nil || oneOf.GoName != field.Oneof.GoName {
				continue
			}
			flag.OneOf = oneOf.oneOfFlag
			break
		}

		if field.Enum != nil {
			flag.Enum = &enumFlag{Enum: field.Enum}
			continue
		}

		if field.Message != nil {
			if wellknownFlag, ok := customFlags[field.Message.Desc.FullName()]; ok {
				flag.Custom = wellknownFlag
				continue
			}
			flag.Message = &messageFlag{Message: field.Message}
			continue
		}
	}

	return msgs
}

type Messages map[protogen.GoIdent]*Message

func MessagesFromFile(file *protogen.File) Messages {
	messages := make(Messages, len(file.Messages))

	for _, message := range file.Messages {
		messages = newMessage(messages, message)
	}

	return messages
}

var (
	//go:embed message.go.tmpl
	messageDefinition string
)

func (msgs Messages) GenerateMessages(plugin *protogen.Plugin, file *protogen.File) error {
	gen := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+"_cli_flags.go", file.GoImportPath)

	header(gen, file)

	return executeTemplate(gen, "message", messageDefinition, msgs)
}

func isOneOfOptional(oneOf *protogen.Oneof, fields []*protogen.Field) bool {
	for _, field := range fields {
		if field.Oneof == nil {
			continue
		}
		if field.Oneof.GoName == oneOf.GoName {
			return field.Desc.HasOptionalKeyword()
		}
	}
	return false
}
