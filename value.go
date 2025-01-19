package cli

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var _ pflag.Value = (*Value[string])(nil)

type protoEnum interface {
	~int32
}

type protoMessage interface {
	any
}

type protoTypes interface {
	string | bool |
		protoEnum | int32 | int64 |
		uint32 | uint64 |
		float32 | float64 |
		[]byte |
		protoMessage
}

type FlagInfo struct {
	Name      string
	Shorthand string
	Usage     string
}

type Value[T protoTypes] struct {
	addr *Value[T]

	changed bool

	info         FlagInfo
	defaultValue *T
	parse        func(field *T, arg string) error

	field *T
}

func WithDefaultValue[T protoTypes](defaultValue *T) func(v *Value[T]) {
	return func(v *Value[T]) {
		v.defaultValue = defaultValue
	}
}

func WithCustomParserValue[T protoTypes](parse func(field *T, arg string) error) func(v *Value[T]) {
	return func(v *Value[T]) {
		v.parse = parse
	}
}

// field must be a pointer to the fields value
func NewValue[T protoTypes](field *T, info FlagInfo, opts ...func(v *Value[T])) *Value[T] {
	v := &Value[T]{
		field: field,
		info:  info,
	}

	for _, opt := range opts {
		opt(v)
	}

	v.copyCheck()

	return v
}

var ErrUnimplemented = errors.New("unimplemented")

// Set implements pflag.Value.
// TODO: describe allowed values
// []byte is standard base64 encoding, as defined in RFC 4648.
func (value *Value[T]) Set(arg string) (err error) {
	value.copyCheck()

	value.changed = true

	switch v := any(value.field).(type) {
	case *string:
		*v = arg
	case *bool:
		if arg == "" {
			*v = true
			return nil
		}
		*v, err = strconv.ParseBool(arg)
		return err
	case *int32:
		parsed, err := strconv.ParseInt(arg, 10, 32)
		*v = int32(parsed)
		return err
	case protoreflect.Enum:
		var n protoreflect.EnumNumber
		if number, err := strconv.ParseInt(arg, 10, 32); err == nil {
			n = protoreflect.EnumNumber(number)
		} else if enumVal := v.Descriptor().Values().ByName(protoreflect.Name(arg)); enumVal != nil {
			n = enumVal.Number()
		} else {
			return errors.New("unable to parse enum value")
		}
		*value.field = any(v.Type().New(n)).(T)
		return nil
	case *int64:
		*v, err = strconv.ParseInt(arg, 10, 64)
		return err
	case *uint32:
		parsed, err := strconv.ParseUint(arg, 10, 32)
		*v = uint32(parsed)
		return err
	case *uint64:
		*v, err = strconv.ParseUint(arg, 10, 64)
		return err
	case *float32:
		parsed, err := strconv.ParseFloat(arg, 32)
		*v = float32(parsed)
		return err
	case *float64:
		*v, err = strconv.ParseFloat(arg, 64)
		return err
	case *[]byte:
		*v, err = base64.StdEncoding.DecodeString(arg)
		return err
	default:
		if value.parse == nil {
			return errors.New("no custom parser defined")
		}
		return value.parse(value.field, arg)
	}
	return nil
}

// String implements pflag.Value.
func (value *Value[T]) String() string {
	value.copyCheck()

	if value.field == nil {
		return ""
	}

	// TODO: how to represent complex types?
	return fmt.Sprintf("%v", *value.field)
}

// Type implements pflag.Value.
func (value *Value[T]) Type() string {
	value.copyCheck()

	var t T
	// TODO: how to represent complex types?
	return fmt.Sprintf("%T", t)
}

func (value *Value[T]) Add(set *pflag.FlagSet) {
	flag := set.VarPF(
		value,
		value.info.Name,
		value.info.Shorthand,
		value.info.Usage,
	)
	if _, ok := any(value.field).(*bool); ok {
		flag.NoOptDefVal = "true"
	}
}

// copyCheck allows uninitialized usage of stmt
func (value *Value[T]) copyCheck() {
	if value.addr == nil {
		// This hack works around a failing of Go's escape analysis
		// that was causing b to escape and be heap allocated.
		// See issue 23382.
		// TODO: once issue 7921 is fixed, this should be reverted to
		// just "stmt.addr = stmt".
		value.addr = (*Value[T])(noescape(unsafe.Pointer(value)))
		// TODO: condition must know if it's args are named parameters or not
		// stmt.namedArgs = make(map[placeholder]string)
	} else if value.addr != value {
		panic("statement: illegal use of non-zero Builder copied by value")
	}
}
