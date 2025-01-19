package cli

import (
	"time"
	"unsafe"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var customParsers map[protoreflect.FullName]func(arg string) (any, error) = map[protoreflect.FullName]func(arg string) (any, error){
	// parses timestamp using RFC3339
	(*timestamppb.Timestamp)(nil).ProtoReflect().Descriptor().FullName(): parseTimestamp,
	(*durationpb.Duration)(nil).ProtoReflect().Descriptor().FullName():   parseDuration,
	// parses json
	(*structpb.Struct)(nil).ProtoReflect().Descriptor().FullName(): parseStruct,
	(*anypb.Any)(nil).ProtoReflect().Descriptor().FullName():       parseAny,
}

func parseTimestamp(arg string) (any, error) {
	ts, err := time.Parse(time.RFC3339, arg)
	if err != nil {
		return nil, err
	}
	return timestamppb.New(ts), nil
}

func parseDuration(arg string) (any, error) {
	duration, err := time.ParseDuration(arg)
	if err != nil {
		return nil, err
	}
	return durationpb.New(duration), nil
}

func parseStruct(arg string) (any, error) {
	payload := new(structpb.Struct)
	if err := payload.UnmarshalJSON([]byte(arg)); err != nil {
		return nil, err
	}
	return payload, nil
}

func parseAny(arg string) (any, error) {
	payload := new(anypb.Any)

	err := protojson.UnmarshalOptions{
		// DiscardUnknown: true,

	}.Unmarshal([]byte(arg), payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func SetCustomParser(typ protoreflect.FullName, parser func(arg string) (any, error)) {
	customParsers[typ] = parser
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

func addDefaultOpt[T protoTypes](flag *pflag.Flag) {
	var t T
	if _, ok := any(t).(bool); ok {
		flag.NoOptDefVal = "true"
	}
}
