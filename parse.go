package cli

import (
	"time"
	"unsafe"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// var customParsers map[protoreflect.FullName]func(arg string) (any, error) = map[protoreflect.FullName]func(arg string) (any, error){
// 	// parses timestamp using RFC3339
// 	(*timestamppb.Timestamp)(nil).ProtoReflect().Descriptor().FullName(): parseTimestamp,
// 	(*durationpb.Duration)(nil).ProtoReflect().Descriptor().FullName():   parseDuration,
// 	// parses json
// 	(*structpb.Struct)(nil).ProtoReflect().Descriptor().FullName(): parseStruct,
// 	(*anypb.Any)(nil).ProtoReflect().Descriptor().FullName():       parseAny,
// }

var defaultUnmarshalOptions = protojson.UnmarshalOptions{
	// DiscardUnknown: true,
}

func ParseTimestamp(field *timestamppb.Timestamp, arg string) error {
	ts, err := time.Parse(time.RFC3339, arg)
	if err != nil {
		return err
	}
	*field = *timestamppb.New(ts)
	return nil
}

func ParseDuration(field *durationpb.Duration, arg string) error {
	duration, err := time.ParseDuration(arg)
	if err != nil {
		return err
	}
	*field = *durationpb.New(duration)
	return nil
}

func ParseStruct(field *structpb.Struct, arg string) error {
	var payload structpb.Struct
	if err := (&payload).UnmarshalJSON([]byte(arg)); err != nil {
		return err
	}

	err := defaultUnmarshalOptions.Unmarshal([]byte(arg), &payload)
	if err != nil {
		return err
	}
	*field = payload
	return nil
}

func ParseAny(field *anypb.Any, arg string) error {
	var payload anypb.Any

	err := defaultUnmarshalOptions.Unmarshal([]byte(arg), &payload)
	if err != nil {
		return err
	}
	*field = payload
	return nil
}

// func SetCustomParser(typ protoreflect.FullName, parser func(arg string) (any, error)) {
// 	customParsers[typ] = parser
// }

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

func addDefaultOpt[T protoTypes](flag *pflag.Flag) {
	var t T
	if _, ok := any(t).(bool); ok {
		flag.NoOptDefVal = "true"
	}
}

type PrimitiveValue interface {
	Add(set *pflag.FlagSet)
}

type FlagSet interface {
	PrimitiveValues() []PrimitiveValue
	NestedValues() map[string]FlagSet
}

func ParseFlags[T FlagSet](request T, args ...string) T {
	set := pflag.NewFlagSet("request", pflag.ContinueOnError)
	for _, value := range request.PrimitiveValues() {
		value.Add(set)
	}
	return request
}
