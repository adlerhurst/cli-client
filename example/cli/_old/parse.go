package cli

// import (
// 	"encoding/base64"
// 	"errors"
// 	"fmt"
// 	"strconv"
// 	"unsafe"

// 	cli_client "github.com/adlerhurst/cli-client"
// 	"github.com/spf13/cobra"
// 	"github.com/spf13/pflag"
// )

// // type Flag[T any] struct {
// // 	set   pflag.FlagSet
// // 	name  string
// // 	value *T
// // }

// // func (f *Flag[T]) Parse(args []string) {
// // 	// f.set.AddGoFlag()
// // 	// f.set.ar
// // 	// pflag.BoolSliceVar()
// // 	// Parse args
// // }

// // func (f *Flag[T]) Value() *T {
// // 	return f.value
// // }

// func (cr *CallRequest) ParseArgs(cmd cobra.Command, args []string) error {

// }

// type cr struct {
// 	useFieldName Value[string]
// }

// func (cr *CallRequest) Flags() {
// 	cmd := cobra.Command{}
// 	set := pflag.NewFlagSet("CallRequest", pflag.ContinueOnError)

// 	set.Var(&cr.UseFieldName, "UseFieldName", "", "")
// 	set.Var(&cr.UseCustomName, "UseCustomName", "", "")

// 	// 	set.StringVar(&cr.Nested, "Nested", "", "")
// 	// 	set.StringVar(&cr.RepNest, "RepNest", "", "")

// 	// 	set.TimestampVar(&cr.CreatedAt, "CreatedAt", "", "")
// 	// 	set.StructVar(&cr.Payload, "Payload", "", "")
// 	// 	set.EnumVar[CallRequest_Wat](&cr.Wat, "Wat", "", "")

// 	// set.BoolVar(&cr.IsSomething, "IsSomething", false, "")
// 	// set.OptionalVar[int32](&cr.I32, "I32", "", "")
// 	// set.Uint32Var(&cr.Ui32, "Ui32", "", "")
// 	// set.Int64Var(&cr.I64, "I64", "", "")
// 	// set.Uint64Var(&cr.Ui64, "Ui64", "", "")
// 	// set.Float32Var(&cr.Fl, "Fl", "", "")
// 	// set.Float64Var(&cr.Dbl, "Dbl", "", "")
// 	// set.BytesVar(&cr.Beiz, "Beiz", "", "")
// 	// set.Int32Var(&cr.Si32, "Si32", "", "")
// 	// set.Int64Var(&cr.Si64, "Si64", "", "")
// 	// set.Uint32Var(&cr.F32, "F32", "", "")
// 	// set.Uint64Var(&cr.F64, "F64", "", "")
// 	// set.Int32Var(&cr.Sf32, "Sf32", "", "")
// 	// set.Int64Var(&cr.Sf64, "Sf64", "", "")
// 	// set.EnumVar[Some](&cr.Some, "Some", "", "")
// 	// set.Float64SliceVar(&cr.RepS, "RepS", "", "")
// 	// set.EnumSliceVar(&cr.RepWat, "RepWat", "", "")
// 	// set.AnyVar(&cr.Something, "Something", "", "")
// }

// type flagTree struct {
// 	nested map[string]flagTree
// }

// type fl[T any] struct {
// 	name  string
// 	value cli_client.Flag[T]
// }

// /*

// call-request:
// 	nested:
// 		Field:
// 	IsSomething:
// */

// func (cr *CallRequest_Nested) Flags() {

// }

// var _ pflag.Value = (*Value[any])(nil)

// type Value[T any] struct {
// 	addr *Value[T]

// 	v T
// }

// func newValue[T any](field *T) *Value[T] {
// }

// var ErrUnimplemented = errors.New("unimplemented")

// // Set implements pflag.Value.
// // TODO: describe allowed values
// // []byte is standard base64 encoding, as defined in RFC 4648.
// func (value *Value[T]) Set(arg string) (err error) {
// 	value.copyCheck()

// 	var v any
// 	switch any(value.v).(type) {
// 	case string:
// 		v = arg
// 	case bool:
// 		if arg == "" {
// 			v = true
// 			break
// 		}
// 		v, err = strconv.ParseBool(arg)
// 	case int32:
// 		v, err = strconv.ParseInt(arg, 10, 32)
// 		v = int32(v.(int64))
// 	case int64:
// 		v, err = strconv.ParseInt(arg, 10, 64)
// 	case uint32:
// 		v, err = strconv.ParseUint(arg, 10, 32)
// 		v = uint32(v.(uint64))
// 	case uint64:
// 		v, err = strconv.ParseUint(arg, 10, 64)
// 	case float32:
// 		v, err = strconv.ParseFloat(arg, 32)
// 		v = float32(v.(float64))
// 	case float64:
// 		v, err = strconv.ParseFloat(arg, 64)
// 	case []byte:
// 		v, err = base64.StdEncoding.DecodeString(arg)
// 	default:
// 		// TODO: implement complex types
// 		return ErrUnimplemented
// 	}
// 	if err != nil {
// 		return err
// 	}
// 	return set(&value.v, v)
// }

// func set(ptr any, value any) error {
// 	switch v := any(ptr).(type) {
// 	case *string:
// 		*v = value.(string)
// 	case *bool:
// 		*v = value.(bool)
// 	case *int32:
// 		*v = value.(int32)
// 	case *int64:
// 		*v = value.(int64)
// 	case *uint32:
// 		*v = value.(uint32)
// 	case *uint64:
// 		*v = value.(uint64)
// 	case *float32:
// 		*v = value.(float32)
// 	case *float64:
// 		*v = value.(float64)
// 	case *[]byte:
// 		*v = value.([]byte)
// 	default:
// 		// TODO: implement complex types
// 		return ErrUnimplemented
// 	}
// 	return nil
// }

// // String implements pflag.Value.
// func (value *Value[T]) String() string {
// 	value.copyCheck()

// 	// TODO: how to represent complex types?
// 	return fmt.Sprintf("%v", value.v)
// }

// // Type implements pflag.Value.
// func (value *Value[T]) Type() string {
// 	value.copyCheck()

// 	// TODO: how to represent complex types?
// 	return fmt.Sprintf("%T", value.v)
// }

// var _ pflag.SliceValue = (*SliceValue[any])(nil)

// type SliceValue[T any] []*Value[T]

// // Append implements pflag.SliceValue.
// func (s *SliceValue[T]) Append(arg string) error {
// 	var val Value[T]
// 	if err := val.Set(arg); err != nil {
// 		return err
// 	}
// 	*s = append(*s, &val)
// 	return nil
// }

// // GetSlice implements pflag.SliceValue.
// func (s *SliceValue[T]) GetSlice() []string {
// 	res := make([]string, len(*s))

// 	for _, value := range *s {
// 		res = append(res, value.String())
// 	}

// 	return res
// }

// // Replace implements pflag.SliceValue.
// func (s *SliceValue[T]) Replace(args []string) error {
// 	res := make([]*Value[T], 0, len(args))
// 	for _, arg := range args {
// 		if err := s.Append(arg); err != nil {
// 			return err
// 		}
// 	}

// 	*s = nil
// 	*s = res
// 	return nil
// }

// // copyCheck allows uninitialized usage of stmt
// func (sf *Value[T]) copyCheck() {
// 	if sf.addr == nil {
// 		// This hack works around a failing of Go's escape analysis
// 		// that was causing b to escape and be heap allocated.
// 		// See issue 23382.
// 		// TODO: once issue 7921 is fixed, this should be reverted to
// 		// just "stmt.addr = stmt".
// 		sf.addr = (*Value[T])(noescape(unsafe.Pointer(sf)))
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
