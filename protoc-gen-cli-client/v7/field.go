package v7

import (
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/compiler/protogen"
)

type fieldCommand struct {
	field *protogen.Field
}

func NewField(gen *protogen.GeneratedFile, field *protogen.Field) *fieldCommand {
	return &fieldCommand{
		field: field,
	}
}

func (f *fieldCommand) Execute() error {
	return nil
}

type String struct {
	field *protogen.Field
}

func NewString(gen *protogen.GeneratedFile, field *protogen.Field) *String {
	return &String{
		field: field,
	}
}
func (s *String) Execute() error {
	var placeholder string
	pflag.StringVar(&placeholder, string(s.field.Desc.Name()), s.field.Desc.Default().String(), string(s.field.Comments.Leading))
	return nil
}
