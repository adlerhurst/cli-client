package v7

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type Message struct {
	gen *protogen.GeneratedFile

	message *protogen.Message
	fields  []*fieldCommand
}

func NewMessage(gen *protogen.GeneratedFile, message *protogen.Message) *Message {
	msg := &Message{
		gen:     gen,
		message: message,
		fields:  make([]*fieldCommand, len(message.Fields)),
	}

	for i, field := range message.Fields {
		msg.fields[i] = NewField(gen, field)
	}

	return msg
}

func (msg *Message) Execute() error {
	return generate(msg.gen, "message", messageTemplate, msg, nil)
}

func (msg *Message) Type() string {
	return qualifier(msg.gen)(ident(msg.message.GoIdent.GoImportPath, msg.message.GoIdent.GoName))
}

const messageTemplate = `
var {{ .VariableName }} = {{ .Type }}{}`
