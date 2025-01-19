package cli

import (
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/spf13/pflag"
)

type indexTree struct {
	flags    *pflag.FlagSet
	startIdx int
	endIdx   int

	sub []indexTree
}

func (it *indexTree) parse(args []string) error {
	if err := it.flags.Parse(args[it.startIdx : it.endIdx+1]); err != nil {
		return err
	}
	for _, sub := range it.sub {
		if err := sub.parse(args); err != nil {
			return err
		}
	}
	return nil
}

var _ pflag.Value = (*Value)(nil)

type Value struct {
	addr *Value

	changed bool

	field reflect.Value
}

// field must be a pointer to the fields value
func newValue(field any) *Value {
	return newValueReflect(reflect.ValueOf(field))
}

func newValueReflect(value reflect.Value) *Value {
	return &Value{
		field: value,
	}
}

var ErrUnimplemented = errors.New("unimplemented")

// Set implements pflag.Value.
// TODO: describe allowed values
// []byte is standard base64 encoding, as defined in RFC 4648.
func (value *Value) Set(arg string) (err error) {
	value.copyCheck()

	value.changed = true

	return set(value.field.Elem(), arg)
}

// String implements pflag.Value.
func (value *Value) String() string {
	value.copyCheck()

	// TODO: how to represent complex types?
	return fmt.Sprintf("%v", value.field.Elem().Interface())
}

// Type implements pflag.Value.
func (value *Value) Type() string {
	value.copyCheck()

	// TODO: how to represent complex types?
	return value.field.Elem().Type().String()
}

func (value *Value) Add(set *pflag.FlagSet, name, shorthand, usage string) {
	flag := set.VarPF(value, name, shorthand, usage)

	// used to allow flags to be set without a value
	if value.field.Elem().Kind() == reflect.Bool {
		flag.NoOptDefVal = "true"
	}
}

// copyCheck allows uninitialized usage of stmt
func (sf *Value) copyCheck() {
	if sf.addr == nil {
		// This hack works around a failing of Go's escape analysis
		// that was causing b to escape and be heap allocated.
		// See issue 23382.
		// TODO: once issue 7921 is fixed, this should be reverted to
		// just "stmt.addr = stmt".
		sf.addr = (*Value)(noescape(unsafe.Pointer(sf)))
		// TODO: condition must know if it's args are named parameters or not
		// stmt.namedArgs = make(map[placeholder]string)
	} else if sf.addr != sf {
		panic("statement: illegal use of non-zero Builder copied by value")
	}
}

var (
	_ pflag.SliceValue = (*SliceValue)(nil)
	_ pflag.Value      = (*SliceValue)(nil)
)

type SliceValue struct {
	addr *SliceValue

	changed bool

	// field must be a pointer to the slice
	field reflect.Value

	values      []*Value
	elementType reflect.Type
}

func newSliceValue(field any) *SliceValue {
	return &SliceValue{
		field:       reflect.ValueOf(field),
		elementType: reflect.TypeOf(field).Elem(),
	}
}

// Set implements pflag.Value.
func (s *SliceValue) Set(arg string) error {
	s.copyCheck()

	argSplit := strings.Split(arg, ",")
	for _, arg := range argSplit {
		if err := s.Append(arg); err != nil {
			return err
		}
	}

	s.changed = true
	return nil
}

// String implements pflag.Value.
func (s *SliceValue) String() string {
	return fmt.Sprintf("%v", s.field.Elem().Interface())
}

// Type implements pflag.Value.
func (s *SliceValue) Type() string {
	return s.field.Elem().Type().String()
}

// Append implements pflag.SliceValue.
func (s *SliceValue) Append(arg string) error {
	s.copyCheck()

	element := reflect.New(s.elementType.Elem())
	value := newValueReflect(element)

	if err := value.Set(arg); err != nil {
		return err
	}

	s.values = append(s.values, value)
	s.field.Elem().Set(reflect.Append(s.field.Elem(), element.Elem()))

	s.changed = true

	return nil
}

// GetSlice implements pflag.SliceValue.
func (s *SliceValue) GetSlice() []string {
	s.copyCheck()

	res := make([]string, len(s.values))

	for i, value := range s.values {
		res[i] = value.String()
	}

	return res
}

// Replace implements pflag.SliceValue.
func (s *SliceValue) Replace(args []string) error {
	s.copyCheck()

	s.reset(len(args))

	for _, arg := range args {
		if err := s.Append(arg); err != nil {
			return err
		}
	}

	return nil
}

func (s *SliceValue) Add(set *pflag.FlagSet, name, shorthand, usage string) {
	flag := set.VarPF(s, name, shorthand, usage)
	if s.field.Kind() == reflect.Bool {
		flag.NoOptDefVal = "true"
	}
}

func (s *SliceValue) reset(newLen int) {
	s.field.Elem().Set(reflect.MakeSlice(s.elementType, 0, newLen))
	s.values = nil
	s.values = make([]*Value, 0, newLen)
}

// copyCheck allows uninitialized usage of stmt
func (sf *SliceValue) copyCheck() {
	if sf.addr == nil {
		// This hack works around a failing of Go's escape analysis
		// that was causing b to escape and be heap allocated.
		// See issue 23382.
		// TODO: once issue 7921 is fixed, this should be reverted to
		// just "stmt.addr = stmt".
		sf.addr = (*SliceValue)(noescape(unsafe.Pointer(sf)))
		// TODO: condition must know if it's args are named parameters or not
		// stmt.namedArgs = make(map[placeholder]string)
	} else if sf.addr != sf {
		panic("statement: illegal use of non-zero Builder copied by value")
	}
}

// noescape hides a pointer from escape analysis. It is the identity function
// but escape analysis doesn't think the output depends on the input.
// noescape is inlined and currently compiles down to zero instructions.
// USE CAREFULLY!
// This was copied from the runtime; see issues 23382 and 7921.
//
//go:nosplit
//go:nocheckptr
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	//nolint: staticcheck
	return unsafe.Pointer(x ^ 0)
}

func set(ptr reflect.Value, arg string) error {
	asdf := ptr.Interface()
	switch asdf.(type) {
	// switch (ptr).(type) {
	case string:
		ptr.SetString(arg)
		return nil
	case bool:
		if arg == "" {
			ptr.SetBool(true)
			return nil
		}
		value, err := strconv.ParseBool(arg)
		ptr.SetBool(value)
		return err
	case int32:
		value, err := strconv.ParseInt(arg, 10, 32)
		ptr.SetInt(value)
		return err
	case int64:
		value, err := strconv.ParseInt(arg, 10, 64)
		ptr.SetInt(value)
		return err
	case uint32:
		value, err := strconv.ParseUint(arg, 10, 32)
		ptr.SetUint(value)
		return err
	case uint64:
		value, err := strconv.ParseUint(arg, 10, 64)
		ptr.SetUint(value)
		return err
	case float32:
		value, err := strconv.ParseFloat(arg, 32)
		ptr.SetFloat(value)
		return err
	case float64:
		value, err := strconv.ParseFloat(arg, 64)
		ptr.SetFloat(value)
		return err
	case []byte:
		value, err := base64.StdEncoding.DecodeString(arg)
		ptr.SetBytes(value)
		return err
	default:
		// TODO: implement complex types
		return ErrUnimplemented
	}
}
