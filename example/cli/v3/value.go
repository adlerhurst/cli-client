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

type Value[T protoTypes] struct {
	addr *Value[T]

	changed bool

	field *T
}

// field must be a pointer to the fields value
func newValue[T protoTypes](field *T) *Value[T] {
	v := &Value[T]{
		field: field,
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
		var t T
		parse, ok := customParsers[any(t).(protoreflect.ProtoMessage).ProtoReflect().Descriptor().FullName()]
		if !ok {
			return ErrUnimplemented
		}

		field, err := parse(arg)
		if err != nil {
			return err
		}
		*value.field, ok = field.(T)
		if !ok {
			return errors.New("unable to parse value")
		}
		return nil
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

func (value *Value[T]) Add(set *pflag.FlagSet, name, shorthand, usage string) {
	flag := set.VarPF(value, name, shorthand, usage)
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

// func parse[T protoTypes](arg string) (value T, err error) {
// 	if arg == "ADMIN" {
// 		arg = "ADMIN"
// 	}

// 	switch v := any(value).(type) {
// 	case *string:
// 		*v = arg
// 		return *v, nil
// 	case *bool:
// 		if arg == "" {
// 			*v = true
// 			return *v, nil
// 		}
// 		*v, err = strconv.ParseBool(arg)
// 		return err
// 	case *int32:
// 		parsed, err := strconv.ParseInt(arg, 10, 32)
// 		return parsed.(T), err
// 	case protoreflect.Enum:
// 		var n protoreflect.EnumNumber
// 		if number, err := strconv.ParseInt(arg, 10, 32); err == nil {
// 			n = protoreflect.EnumNumber(number)
// 		} else if enumVal := (*v).Descriptor().Values().ByName(protoreflect.Name(arg)); enumVal != nil {
// 			n = enumVal.Number()
// 		} else {
// 			return value, errors.New("unable to parse enum value")
// 		}
// 		return v.Type().New(n).(T), nil
// 	case *int64:
// 		*v, err = strconv.ParseInt(arg, 10, 64)
// 		return err
// 	case *uint32:
// 		parsed, err := strconv.ParseUint(arg, 10, 32)
// 		*v = uint32(parsed)
// 		return err
// 	case *uint64:
// 		*v, err = strconv.ParseUint(arg, 10, 64)
// 		return err
// 	case *float32:
// 		parsed, err := strconv.ParseFloat(arg, 32)
// 		*v = float32(parsed)
// 		return err
// 	case *float64:
// 		*v, err = strconv.ParseFloat(arg, 64)
// 		return err
// 	case *[]byte:
// 		*v, err = base64.StdEncoding.DecodeString(arg)
// 		return err
// 	default:
// 		// TODO: implement complex types
// 		return ErrUnimplemented
// 	}

// 	return value, nil
// }

// func parseEnum[T protoEnum](arg string) (T, error) {
// 	var value T
// 	if arg == "ADMIN" {
// 		arg = "ADMIN"
// 	}

// 	v, ok := any(value).(protoreflect.Enum)
// 	if ok {
// 		var n protoreflect.EnumNumber
// 		if number, err := strconv.ParseInt(arg, 10, 32); err == nil {
// 			n = protoreflect.EnumNumber(number)
// 		} else if enumVal := (*v).Descriptor().Values().ByName(protoreflect.Name(arg)); enumVal != nil {
// 			n = enumVal.Number()
// 		} else {
// 			return value, errors.New("unable to parse enum value")
// 		}
// 		return v.Type().New(n).(T), nil
// 	}

// 	return value, nil
// }

// func set(value any, arg string) (err error) {
// 	if arg == "ADMIN" {
// 		arg = "admin"
// 	}
// 	v, ok := value.(*protoreflect.Enum)
// 	if ok {
// 		var n protoreflect.EnumNumber
// 		if number, err := strconv.ParseInt(arg, 10, 32); err == nil {
// 			n = protoreflect.EnumNumber(number)
// 		} else if enumVal := (*v).Descriptor().Values().ByName(protoreflect.Name(arg)); enumVal != nil {
// 			n = enumVal.Number()
// 		} else {
// 			return errors.New("unable to parse enum value")
// 		}
// 		(*v) = (*v).Type().New(n)
// 		return nil
// 	}
// 	switch v := value.(type) {
// 	case *string:
// 		*v = arg
// 		return nil
// 	case *bool:
// 		if arg == "" {
// 			*v = true
// 			return nil
// 		}
// 		*v, err = strconv.ParseBool(arg)
// 		return err
// 	case *int32:
// 		parsed, err := strconv.ParseInt(arg, 10, 32)
// 		*v = int32(parsed)
// 		return err
// 	case *protoreflect.Enum:
// 		var n protoreflect.EnumNumber
// 		if number, err := strconv.ParseInt(arg, 10, 32); err == nil {
// 			n = protoreflect.EnumNumber(number)
// 		} else if enumVal := (*v).Descriptor().Values().ByName(protoreflect.Name(arg)); enumVal != nil {
// 			n = enumVal.Number()
// 		} else {
// 			return errors.New("unable to parse enum value")
// 		}
// 		(*v) = (*v).Type().New(n)
// 		return nil
// 	case *int64:
// 		*v, err = strconv.ParseInt(arg, 10, 64)
// 		return err
// 	case *uint32:
// 		parsed, err := strconv.ParseUint(arg, 10, 32)
// 		*v = uint32(parsed)
// 		return err
// 	case *uint64:
// 		*v, err = strconv.ParseUint(arg, 10, 64)
// 		return err
// 	case *float32:
// 		parsed, err := strconv.ParseFloat(arg, 32)
// 		*v = float32(parsed)
// 		return err
// 	case *float64:
// 		*v, err = strconv.ParseFloat(arg, 64)
// 		return err
// 	case *[]byte:
// 		*v, err = base64.StdEncoding.DecodeString(arg)
// 		return err
// 	default:
// 		// TODO: implement complex types
// 		return ErrUnimplemented
// 	}
// }
