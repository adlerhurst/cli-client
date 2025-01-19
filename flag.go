package cli

// import (
// 	"errors"
// 	"fmt"
// 	"log/slog"
// 	"reflect"
// 	"strconv"
// 	"unsafe"

// 	"github.com/spf13/pflag"
// 	"google.golang.org/protobuf/encoding/protojson"
// 	"google.golang.org/protobuf/reflect/protoreflect"
// )

// var _ pflag.SliceValue = (*SliceFlag[any])(nil)

// type SliceFlag[T any] struct {
// 	flag

// 	addr *SliceFlag[T]

// 	values []T
// }

// func (sf SliceFlag[T]) Value() []T {
// 	sf.copyCheck()

// 	return sf.values
// }

// // Append implements pflag.SliceValue.
// func (sf *SliceFlag[T]) Append(arg string) (err error) {
// 	sf.copyCheck()

// 	var value T
// 	if value, err = parseValue[T](arg); err != nil {
// 		return err
// 	}
// 	sf.values = append(sf.values, value)
// 	sf.changed = true

// 	return nil
// }

// // GetSlice implements pflag.SliceValue.
// func (sf *SliceFlag[T]) GetSlice() []string {
// 	sf.copyCheck()

// 	if len(sf.values) == 0 {
// 		return nil
// 	}

// 	values := make([]string, len(sf.values))
// 	for i, value := range sf.values {
// 		values[i] = valueToString(value)
// 	}

// 	return values
// }

// // Replace implements pflag.SliceValue.
// func (sf *SliceFlag[T]) Replace(args []string) (err error) {
// 	sf.copyCheck()

// 	values := make([]T, len(args))

// 	defer func() {
// 		if err != nil {
// 			return
// 		}
// 		sf.changed = true
// 	}()

// 	for i, arg := range args {
// 		if values[i], err = parseValue[T](arg); err != nil {
// 			return err
// 		}
// 	}

// 	sf.values = nil
// 	sf.values = values
// 	return nil
// }

// var _ pflag.Value = (*Flag[any])(nil)

// type Flag[T any] struct {
// 	flag

// 	addr *Flag[T]

// 	value        T
// 	defaultValue T
// }

// func (f Flag[T]) Value() T {
// 	f.copyCheck()

// 	return f.value
// }

// var ErrParse = errors.New("unable to parse value")

// // Set implements [pflag.Value].
// func (f *Flag[T]) Set(arg string) (err error) {
// 	f.copyCheck()

// 	defer func() {
// 		if err != nil {
// 			return
// 		}
// 		f.changed = true
// 	}()

// 	f.value, err = parseValue[T](arg)
// 	return err
// }

// // String implements [pflag.Value].
// func (f *Flag[T]) String() string {
// 	f.copyCheck()

// 	return valueToString(f.value)
// }

// // Type implements [pflag.Value].
// func (f *Flag[T]) Type() string {
// 	f.copyCheck()

// 	value, ok := any(f.value).(protoreflect.ProtoMessage)
// 	if !ok {
// 		return fmt.Sprintf("%T", f.value)
// 	}
// 	return string(value.ProtoReflect().Type().Descriptor().FullName())
// }

// type flag struct {
// 	name    string
// 	changed bool
// }

// // copyCheck allows uninitialized usage of stmt
// func (f *Flag[T]) copyCheck() {
// 	if f.addr == nil {
// 		// This hack works around a failing of Go's escape analysis
// 		// that was causing b to escape and be heap allocated.
// 		// See issue 23382.
// 		// TODO: once issue 7921 is fixed, this should be reverted to
// 		// just "stmt.addr = stmt".
// 		f.addr = (*Flag[T])(noescape(unsafe.Pointer(f)))
// 		// TODO: condition must know if it's args are named parameters or not
// 		// stmt.namedArgs = make(map[placeholder]string)
// 	} else if f.addr != f {
// 		panic("statement: illegal use of non-zero Builder copied by value")
// 	}
// }

// // copyCheck allows uninitialized usage of stmt
// func (sf *SliceFlag[T]) copyCheck() {
// 	if sf.addr == nil {
// 		// This hack works around a failing of Go's escape analysis
// 		// that was causing b to escape and be heap allocated.
// 		// See issue 23382.
// 		// TODO: once issue 7921 is fixed, this should be reverted to
// 		// just "stmt.addr = stmt".
// 		sf.addr = (*SliceFlag[T])(noescape(unsafe.Pointer(sf)))
// 		// TODO: condition must know if it's args are named parameters or not
// 		// stmt.namedArgs = make(map[placeholder]string)
// 	} else if sf.addr != sf {
// 		panic("statement: illegal use of non-zero Builder copied by value")
// 	}
// }

// // noescape hides a pointer from escape analysis. It is the identity function
// // but escape analysis doesn't think the output depends on the input.
// // noescape is inlined and currently compiles down to zero instructions.
// // USE CAREFULLY!
// // This was copied from the runtime; see issues 23382 and 7921.
// //
// //go:nosplit
// //go:nocheckptr
// func noescape(p unsafe.Pointer) unsafe.Pointer {
// 	x := uintptr(p)
// 	//nolint: staticcheck
// 	return unsafe.Pointer(x ^ 0)
// }

// func valueToString(value any) string {
// 	if reflect.ValueOf(value).IsZero() {
// 		return ""
// 	}

// 	v, ok := value.(protoreflect.ProtoMessage)
// 	if !ok {
// 		return fmt.Sprintf("%T", value)
// 	}

// 	return string(v.ProtoReflect().Type().Descriptor().FullName())
// }

// func parseValue[T any](arg string) (T, error) {
// 	var value T
// 	if v, ok := any(value).(protoreflect.EnumType); ok {
// 		enumValue, err := parseEnumValue(arg, v.Descriptor())
// 		if err != nil {
// 			return value, err
// 		}
// 		return enumValue.(T), nil
// 	}
// 	v, ok := any(&value).(protoreflect.ProtoMessage)
// 	if !ok {
// 		slog.Error("must implement custom parser", "type", fmt.Sprintf("%T", value))
// 		return value, ErrParse
// 	}
// 	err := protojson.UnmarshalOptions{
// 		// AllowPartial: true,
// 		DiscardUnknown: true,
// 	}.Unmarshal([]byte(arg), v)

// 	return value, err
// }

// func parseEnumValue(arg string, value protoreflect.EnumDescriptor) (protoreflect.EnumValueDescriptor, error) {
// 	v := value.Values().ByName(protoreflect.Name(arg))
// 	if v != nil {
// 		return v, nil
// 	}
// 	if i, err := strconv.Atoi(arg); err == nil {
// 		v := value.Values().ByNumber(protoreflect.EnumNumber(i))
// 		if v != nil {
// 			return v, nil
// 		}
// 	}
// 	return nil, ErrParse
// }
