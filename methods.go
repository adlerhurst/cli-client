package cli

// import (
// 	"encoding/csv"
// 	"errors"
// 	"fmt"
// 	"log/slog"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/spf13/pflag"
// 	"google.golang.org/protobuf/encoding/protojson"
// 	"google.golang.org/protobuf/reflect/protoreflect"
// 	"google.golang.org/protobuf/types/known/anypb"
// 	"google.golang.org/protobuf/types/known/durationpb"
// 	"google.golang.org/protobuf/types/known/structpb"
// 	"google.golang.org/protobuf/types/known/timestamppb"
// )

// type argFlag[T any] struct {
// 	*primitiveFlag[T]
// 	customFlag func(field *T, arg string) error

// 	changed bool
// }

// func newArgFlag[T any](set *pflag.FlagSet, name string, opts ...primitiveFlagOpt[T]) *argFlag[T] {
// 	parser := new(argFlag[T])
// 	parser.primitiveFlag = newPrimitiveFlag[T](set, name, opts...)

// 	return parser
// }

// func (parser *argFlag[T]) Changed() bool {
// 	return parser.changed
// }

// type primitiveFlagOpt[T any] func(*primitiveFlag[T])

// func WithDefaultValue[T any](value T) primitiveFlagOpt[T] {
// 	return func(parser *primitiveFlag[T]) {
// 		parser.defaultValue = value
// 	}
// }

// type primitiveFlag[T any] struct {
// 	Value        *T
// 	defaultValue T

// 	set  *pflag.FlagSet
// 	name string
// }

// func newPrimitiveFlag[T any](set *pflag.FlagSet, name string, opts ...primitiveFlagOpt[T]) *primitiveFlag[T] {
// 	parser := &primitiveFlag[T]{
// 		set:  set,
// 		name: name,
// 	}

// 	for _, opt := range opts {
// 		opt(parser)
// 	}

// 	return parser
// }

// func (parser *primitiveFlag[T]) applyOpts(opts []primitiveFlagOpt[T]) {
// 	for _, opt := range opts {
// 		opt(parser)
// 	}
// }

// func (parser *primitiveFlag[T]) Changed() bool {
// 	return parser.set.Changed(parser.name)
// }

// // Set implements pflag.Value.
// func (v *argFlag[T]) Set(arg string) error {
// 	v.changed = true
// 	if v.customFlag != nil {
// 		return v.customFlag(v.Value, arg)
// 	}

// 	value, ok := interface{}(v.Value).(protoreflect.ProtoMessage)
// 	if !ok {
// 		slog.Error("must implement custom parser", "type", fmt.Sprintf("%T", v.Value))
// 	}
// 	return protojson.UnmarshalOptions{
// 		// AllowPartial: true,
// 		DiscardUnknown: true,
// 	}.Unmarshal([]byte(arg), value)
// }

// // String implements pflag.Value.
// func (v *argFlag[T]) String() string {
// 	value, ok := interface{}(v.Value).(protoreflect.ProtoMessage)
// 	if !ok {
// 		return fmt.Sprint(v.Value)
// 	}
// 	return protojson.Format(value)
// }

// // Type implements pflag.Value.
// func (v *argFlag[T]) Type() string {
// 	value, ok := interface{}(v.Value).(protoreflect.ProtoMessage)
// 	if !ok {
// 		return fmt.Sprintf("%T", v.Value)
// 	}
// 	return string(value.ProtoReflect().Type().Descriptor().FullName())
// }

// type StructFlag struct {
// 	*argFlag[structpb.Struct]
// }

// func NewStructFlag(set *pflag.FlagSet, name, usage string) *StructFlag {
// 	parser := newArgFlag[structpb.Struct](set, name)
// 	parser.Value = new(structpb.Struct)
// 	set.Var(parser, name, usage)
// 	return &StructFlag{argFlag: parser}
// }

// type StructSliceFlag struct {
// 	*argFlag[[]*structpb.Struct]
// }

// func NewStructSliceFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[[]*structpb.Struct]) *StructSliceFlag {
// 	parser := newArgFlag[[]*structpb.Struct](set, name)
// 	parser.applyOpts(opts)
// 	parser.Value = new([]*structpb.Struct)
// 	set.Var(parser, name, usage)
// 	return &StructSliceFlag{argFlag: parser}
// }

// type AnyFlag struct {
// 	*argFlag[anypb.Any]
// }

// func NewAnyFlag(set *pflag.FlagSet, name, usage string) *AnyFlag {
// 	parser := newArgFlag[anypb.Any](set, name)
// 	// TODO: change to message
// 	parser.Value = new(anypb.Any)
// 	set.Var(parser, name, usage)
// 	return &AnyFlag{argFlag: parser}
// }

// type TimestampFlag struct {
// 	*argFlag[timestamppb.Timestamp]
// }

// func NewTimestampFlag(set *pflag.FlagSet, name, usage string) *TimestampFlag {
// 	parser := newArgFlag[timestamppb.Timestamp](set, name)
// 	parser.Value = new(timestamppb.Timestamp)
// 	parser.customFlag = timestampFlag
// 	set.Var(parser, name, usage)
// 	return &TimestampFlag{argFlag: parser}
// }

// type TimestampSliceFlag struct {
// 	*argFlag[[]*timestamppb.Timestamp]
// }

// func NewTimestampSliceFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[[]*timestamppb.Timestamp]) *TimestampSliceFlag {
// 	parser := newArgFlag[[]*timestamppb.Timestamp](set, name)
// 	parser.applyOpts(opts)
// 	parser.Value = new([]*timestamppb.Timestamp)
// 	parser.customFlag = slicePtrFlag[timestamppb.Timestamp](timestampFlag)
// 	set.Var(parser, name, usage)
// 	return &TimestampSliceFlag{argFlag: parser}
// }

// func timestampFlag(field *timestamppb.Timestamp, arg string) error {
// 	timestamp, err := time.Parse(time.RFC3339, arg)
// 	if err != nil {
// 		return err
// 	}
// 	*field = *timestamppb.New(timestamp)
// 	return nil
// }

// type DurationFlag struct {
// 	*argFlag[durationpb.Duration]
// }

// func NewDurationFlag(set *pflag.FlagSet, name, usage string) *DurationFlag {
// 	parser := newArgFlag[durationpb.Duration](set, name)
// 	parser.Value = new(durationpb.Duration)
// 	parser.customFlag = durationFlag
// 	set.Var(parser, name, usage)
// 	return &DurationFlag{argFlag: parser}
// }

// type DurationSliceFlag struct {
// 	*argFlag[[]*durationpb.Duration]
// }

// func NewDurationSliceFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[[]*durationpb.Duration]) *DurationSliceFlag {
// 	parser := newArgFlag[[]*durationpb.Duration](set, name)
// 	parser.applyOpts(opts)
// 	parser.Value = new([]*durationpb.Duration)
// 	parser.customFlag = slicePtrFlag[durationpb.Duration](durationFlag)
// 	set.Var(parser, name, usage)
// 	return &DurationSliceFlag{argFlag: parser}
// }

// func durationFlag(field *durationpb.Duration, arg string) error {
// 	duration, err := time.ParseDuration(arg)
// 	if err != nil {
// 		return err
// 	}
// 	*field = *durationpb.New(duration)
// 	return nil
// }

// type enum interface {
// 	~int32
// 	Descriptor() protoreflect.EnumDescriptor
// 	String() string
// }

// type EnumFlag[E enum] struct {
// 	*argFlag[E]
// }

// func NewEnumFlag[E enum](set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[E]) *EnumFlag[E] {
// 	parser := newArgFlag[E](set, name)
// 	parser.applyOpts(opts)
// 	parser.Value = new(E)
// 	parser.customFlag = enumFlag[E]
// 	set.Var(parser, name, usage)
// 	return &EnumFlag[E]{argFlag: parser}
// }

// type EnumSliceFlag[E enum] struct {
// 	*argFlag[[]E]
// }

// func NewEnumSliceFlag[E enum](set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[[]E]) *EnumSliceFlag[E] {
// 	parser := newArgFlag[[]E](set, name)
// 	parser.applyOpts(opts)
// 	parser.Value = new([]E)
// 	parser.customFlag = sliceFlag(enumFlag[E])
// 	set.Var(parser, name, usage)
// 	return &EnumSliceFlag[E]{argFlag: parser}
// }

// func sliceFlag[T any](parser func(*T, string) error) func(*[]T, string) error {
// 	return func(field *[]T, arg string) error {
// 		stringReader := strings.NewReader(arg)
// 		csvReader := csv.NewReader(stringReader)
// 		records, err := csvReader.Read()
// 		if err != nil {
// 			return err
// 		}

// 		values := make([]T, len(records))
// 		for i, record := range records {
// 			value := new(T)
// 			err := parser(value, record)
// 			if err != nil {
// 				return err
// 			}
// 			values[i] = *value
// 		}
// 		*field = append(*field, values...)

// 		return nil
// 	}
// }

// func slicePtrFlag[T any](parser func(*T, string) error) func(*[]*T, string) error {
// 	return func(field *[]*T, arg string) error {
// 		stringReader := strings.NewReader(arg)
// 		csvReader := csv.NewReader(stringReader)
// 		records, err := csvReader.Read()
// 		if err != nil {
// 			return err
// 		}

// 		values := make([]*T, len(records))
// 		for i, record := range records {
// 			value := new(T)
// 			err := parser(value, record)
// 			if err != nil {
// 				return err
// 			}
// 			values[i] = value
// 		}
// 		*field = append(*field, values...)

// 		return nil
// 	}
// }

// func enumFlag[E enum](field *E, arg string) error {
// 	if desc := (*field).Descriptor().Values().ByName(protoreflect.Name(arg)); desc != nil {
// 		*field = E(desc.Number())
// 		return nil
// 	}
// 	if number, err := strconv.Atoi(arg); err == nil {
// 		if desc := (*field).Descriptor().Values().ByNumber(protoreflect.EnumNumber(number)); desc != nil {
// 			*field = E(desc.Number())
// 			return nil
// 		}
// 	}

// 	return errors.New("unknown enum variable")
// }

// type StringFlag struct {
// 	*primitiveFlag[string]
// }

// func NewStringFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[string]) *StringFlag {
// 	parser := newPrimitiveFlag[string](set, name, opts...)
// 	parser.Value = new(string)
// 	set.StringVar(parser.Value, name, parser.defaultValue, usage)
// 	return &StringFlag{primitiveFlag: parser}
// }

// type StringSliceFlag struct {
// 	*primitiveFlag[[]string]
// }

// func NewStringSliceFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[[]string]) *StringSliceFlag {
// 	parser := newPrimitiveFlag[[]string](set, name, opts...)
// 	parser.Value = new([]string)
// 	set.StringSliceVar(parser.Value, name, parser.defaultValue, usage)
// 	return &StringSliceFlag{primitiveFlag: parser}
// }

// type BoolFlag struct {
// 	*primitiveFlag[bool]
// }

// func NewBoolFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[bool]) *BoolFlag {
// 	parser := newPrimitiveFlag[bool](set, name, opts...)
// 	parser.Value = new(bool)
// 	set.BoolVar(parser.Value, name, parser.defaultValue, usage)
// 	return &BoolFlag{primitiveFlag: parser}
// }

// type BoolSliceFlag struct {
// 	*primitiveFlag[[]bool]
// }

// func NewBoolSliceFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[[]bool]) *BoolSliceFlag {
// 	parser := newPrimitiveFlag[[]bool](set, name, opts...)
// 	parser.Value = new([]bool)
// 	set.BoolSliceVar(parser.Value, name, parser.defaultValue, usage)
// 	return &BoolSliceFlag{primitiveFlag: parser}
// }

// type Int32Flag struct {
// 	*primitiveFlag[int32]
// }

// func NewInt32Flag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[int32]) *Int32Flag {
// 	parser := newPrimitiveFlag[int32](set, name, opts...)
// 	parser.Value = new(int32)
// 	set.Int32Var(parser.Value, name, parser.defaultValue, usage)
// 	return &Int32Flag{primitiveFlag: parser}
// }

// type Int32SliceFlag struct {
// 	*primitiveFlag[[]int32]
// }

// func NewInt32SliceFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[[]int32]) *Int32SliceFlag {
// 	parser := newPrimitiveFlag[[]int32](set, name, opts...)
// 	parser.Value = new([]int32)
// 	set.Int32SliceVar(parser.Value, name, parser.defaultValue, usage)
// 	return &Int32SliceFlag{primitiveFlag: parser}
// }

// type Sint32Flag struct {
// 	*primitiveFlag[int32]
// }

// func NewSint32Flag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[int32]) *Sint32Flag {
// 	return (*Sint32Flag)(NewInt32Flag(set, name, usage, opts...))
// }

// type Sfixed32Flag struct {
// 	*primitiveFlag[int32]
// }

// func NewSfixed32Flag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[int32]) *Sfixed32Flag {
// 	return (*Sfixed32Flag)(NewInt32Flag(set, name, usage, opts...))
// }

// type Uint32Flag struct {
// 	*primitiveFlag[uint32]
// }

// func NewUint32Flag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[uint32]) *Uint32Flag {
// 	parser := newPrimitiveFlag[uint32](set, name, opts...)
// 	parser.Value = new(uint32)
// 	set.Uint32Var(parser.Value, name, parser.defaultValue, usage)
// 	return &Uint32Flag{primitiveFlag: parser}
// }

// type Fixed32Flag struct {
// 	*primitiveFlag[uint32]
// }

// func NewFixed32Flag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[uint32]) *Fixed32Flag {
// 	return (*Fixed32Flag)(NewUint32Flag(set, name, usage, opts...))
// }

// type Uint32SliceFlag struct {
// 	*primitiveFlag[[]uint]
// }

// func NewUint32SliceFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[[]uint]) *Uint32SliceFlag {
// 	return &Uint32SliceFlag{primitiveFlag: newUintSliceFlag(set, name, usage, opts...)}
// }

// type Int64Flag struct {
// 	*primitiveFlag[int64]
// }

// func NewInt64Flag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[int64]) *Int64Flag {
// 	parser := newPrimitiveFlag[int64](set, name, opts...)
// 	parser.Value = new(int64)
// 	set.Int64Var(parser.Value, name, parser.defaultValue, usage)
// 	return &Int64Flag{primitiveFlag: parser}
// }

// type Sint64Flag struct {
// 	*primitiveFlag[int64]
// }

// func NewSint64Flag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[int64]) *Sint64Flag {
// 	return (*Sint64Flag)(NewInt64Flag(set, name, usage, opts...))
// }

// type Sfixed64Flag struct {
// 	*primitiveFlag[int64]
// }

// func NewSfixed64Flag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[int64]) *Sfixed64Flag {
// 	return (*Sfixed64Flag)(NewInt64Flag(set, name, usage, opts...))
// }

// type Int64SliceFlag struct {
// 	*primitiveFlag[[]int64]
// }

// func NewInt64SliceFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[[]int64]) *Int64SliceFlag {
// 	parser := newPrimitiveFlag[[]int64](set, name, opts...)
// 	parser.Value = new([]int64)
// 	set.Int64SliceVar(parser.Value, name, parser.defaultValue, usage)
// 	return &Int64SliceFlag{primitiveFlag: parser}
// }

// type Uint64Flag struct {
// 	*primitiveFlag[uint64]
// }

// func NewUint64Flag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[uint64]) *Uint64Flag {
// 	parser := newPrimitiveFlag[uint64](set, name, opts...)
// 	parser.Value = new(uint64)
// 	set.Uint64Var(parser.Value, name, parser.defaultValue, usage)
// 	return &Uint64Flag{primitiveFlag: parser}
// }

// type Fixed64Flag struct {
// 	*primitiveFlag[uint64]
// }

// func NewFixed64Flag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[uint64]) *Fixed64Flag {
// 	return (*Fixed64Flag)(NewUint64Flag(set, name, usage, opts...))
// }

// type Uint64SliceFlag struct {
// 	*primitiveFlag[[]uint]
// }

// func NewUint64SliceFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[[]uint]) *Uint64SliceFlag {
// 	return &Uint64SliceFlag{primitiveFlag: newUintSliceFlag(set, name, usage, opts...)}
// }

// func newUintSliceFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[[]uint]) *primitiveFlag[[]uint] {
// 	parser := newPrimitiveFlag[[]uint](set, name, opts...)
// 	parser.Value = new([]uint)
// 	set.UintSliceVar(parser.Value, name, parser.defaultValue, usage)
// 	return parser
// }

// type FloatFlag struct {
// 	*primitiveFlag[float32]
// }

// func NewFloatFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[float32]) *FloatFlag {
// 	parser := newPrimitiveFlag[float32](set, name, opts...)
// 	parser.Value = new(float32)
// 	set.Float32Var(parser.Value, name, parser.defaultValue, usage)
// 	return &FloatFlag{primitiveFlag: parser}
// }

// type FloatSliceFlag struct {
// 	*primitiveFlag[[]float32]
// }

// func NewFloatSliceFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[[]float32]) *FloatSliceFlag {
// 	parser := newPrimitiveFlag[[]float32](set, name, opts...)
// 	parser.Value = new([]float32)
// 	set.Float32SliceVar(parser.Value, name, parser.defaultValue, usage)
// 	return &FloatSliceFlag{primitiveFlag: parser}
// }

// type DoubleFlag struct {
// 	*primitiveFlag[float64]
// }

// func NewDoubleFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[float64]) *DoubleFlag {
// 	parser := newPrimitiveFlag[float64](set, name, opts...)
// 	parser.Value = new(float64)
// 	set.Float64Var(parser.Value, name, parser.defaultValue, usage)
// 	return &DoubleFlag{primitiveFlag: parser}
// }

// type DoubleSliceFlag struct {
// 	*primitiveFlag[[]float64]
// }

// func NewDoubleSliceFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[[]float64]) *DoubleSliceFlag {
// 	parser := newPrimitiveFlag[[]float64](set, name, opts...)
// 	parser.Value = new([]float64)
// 	set.Float64SliceVar(parser.Value, name, parser.defaultValue, usage)
// 	return &DoubleSliceFlag{primitiveFlag: parser}
// }

// type BytesFlag struct {
// 	*primitiveFlag[[]byte]
// }

// func NewBytesFlag(set *pflag.FlagSet, name, usage string, opts ...primitiveFlagOpt[[]byte]) *BytesFlag {
// 	parser := newPrimitiveFlag[[]byte](set, name, opts...)
// 	parser.Value = new([]byte)
// 	set.BytesBase64Var(parser.Value, name, parser.defaultValue, usage)
// 	return &BytesFlag{primitiveFlag: parser}
// }
