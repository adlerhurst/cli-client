package cli

import (
	"fmt"
	"log/slog"
	"unsafe"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var _ pflag.Value = (*Flag[any])(nil)

type Flag[T any] struct {
	addr *Flag[T]

	Value        T
	defaultValue T

	name string

	changed bool
}

// Set implements [pflag.Value].
func (f *Flag[T]) Set(arg string) error {
	f.copyCheck()

	f.changed = true

	value, ok := interface{}(f.Value).(protoreflect.ProtoMessage)
	if !ok {
		slog.Error("must implement custom parser", "type", fmt.Sprintf("%T", f.Value))
	}
	return protojson.UnmarshalOptions{
		// AllowPartial: true,
		DiscardUnknown: true,
	}.Unmarshal([]byte(arg), value)
}

// String implements [pflag.Value].
func (f *Flag[T]) String() string {
	f.copyCheck()

	value, ok := any(f.Value).(protoreflect.ProtoMessage)
	if !ok {
		return fmt.Sprint(f.Value)
	}
	return protojson.Format(value)
}

// Type implements [pflag.Value].
func (f *Flag[T]) Type() string {
	f.copyCheck()

	value, ok := any(f.Value).(protoreflect.ProtoMessage)
	if !ok {
		return fmt.Sprintf("%T", f.Value)
	}
	return string(value.ProtoReflect().Type().Descriptor().FullName())
}

// copyCheck allows uninitialized usage of stmt
func (stmt *Flag[T]) copyCheck() {
	if stmt.addr == nil {
		// This hack works around a failing of Go's escape analysis
		// that was causing b to escape and be heap allocated.
		// See issue 23382.
		// TODO: once issue 7921 is fixed, this should be reverted to
		// just "stmt.addr = stmt".
		stmt.addr = (*Flag[T])(noescape(unsafe.Pointer(stmt)))
		// TODO: condition must know if it's args are named parameters or not
		// stmt.namedArgs = make(map[placeholder]string)
	} else if stmt.addr != stmt {
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
