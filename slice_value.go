package cli

import (
	"fmt"
	"slices"
	"strings"
	"unsafe"

	"github.com/spf13/pflag"
)

var (
	_ pflag.SliceValue = (*SliceValue[string])(nil)
	_ pflag.Value      = (*SliceValue[string])(nil)
)

type SliceValue[T protoTypes] struct {
	addr *SliceValue[T]

	changed bool

	info         FlagInfo
	defaultValue []T

	// field must be a pointer to the slice
	field *[]T
}

func WithDefaultSliceValue[T protoTypes](sv *SliceValue[T], defaultValue []T) {
	sv.defaultValue = defaultValue
}

func NewSliceValue[T protoTypes](field *[]T, info FlagInfo, opts ...func(sv *SliceValue[T])) *SliceValue[T] {
	sv := &SliceValue[T]{
		field: field,
		info:  info,
	}

	for _, opt := range opts {
		opt(sv)
	}

	sv.copyCheck()
	return sv
}

// Set implements pflag.Value.
func (s *SliceValue[T]) Set(arg string) error {
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
func (s *SliceValue[T]) String() string {
	if s.field == nil {
		return ""
	}

	var builder strings.Builder
	builder.WriteString("[")
	for i, value := range *s.field {
		if i > 0 {
			builder.WriteString(",")
		}
		builder.WriteString(fmt.Sprintf("%v", value))
	}
	builder.WriteString("]")

	return builder.String()
}

// Type implements pflag.Value.
func (s *SliceValue[T]) Type() string {
	var t T
	return fmt.Sprintf("%TSlice", t)
}

// Append implements pflag.SliceValue.
func (s *SliceValue[T]) Append(arg string) error {
	s.copyCheck()

	field := new(T)
	// we only need the value for parsing so we don't add the flag info
	if err := NewValue(field, FlagInfo{}).Set(arg); err != nil {
		return err
	}
	*s.field = append(*s.field, *field)
	s.changed = true

	return nil
}

// GetSlice implements pflag.SliceValue.
func (s *SliceValue[T]) GetSlice() []string {
	s.copyCheck()

	res := make([]string, len(*s.field))

	for i, value := range *s.field {
		res[i] = fmt.Sprintf("%v", value)
	}

	return res
}

// Replace implements pflag.SliceValue.
func (s *SliceValue[T]) Replace(args []string) error {
	s.copyCheck()

	*s.field = (*s.field)[:0]
	*s.field = slices.Grow(*s.field, len(args))

	for _, arg := range args {
		if err := s.Append(arg); err != nil {
			return err
		}
	}

	return nil
}

func (s *SliceValue[T]) Add(set *pflag.FlagSet) {
	flag := set.VarPF(
		s,
		s.info.Name,
		s.info.Shorthand,
		s.info.Usage,
	)
	addDefaultOpt[T](flag)
}

// copyCheck allows uninitialized usage of stmt
func (s *SliceValue[T]) copyCheck() {
	if s.addr == nil {
		// This hack works around a failing of Go's escape analysis
		// that was causing b to escape and be heap allocated.
		// See issue 23382.
		// TODO: once issue 7921 is fixed, this should be reverted to
		// just "stmt.addr = stmt".
		s.addr = (*SliceValue[T])(noescape(unsafe.Pointer(s)))
		// TODO: condition must know if it's args are named parameters or not
		// stmt.namedArgs = make(map[placeholder]string)
	} else if s.addr != s {
		panic("statement: illegal use of non-zero Builder copied by value")
	}
}
