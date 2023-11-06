package example

import (
	"encoding/csv"
	"errors"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Field interface {
	AddFlag(set *pflag.FlagSet)
	Value() any
}

type field[V any] struct {
	Name  string
	Usage string

	value        *V
	defaultValue V
}

func (f *field[V]) Value() any {
	return *f.value
}

type PathField struct {
	value *string
}

func (f *PathField) AddFlag(set *pflag.FlagSet) {
	f.value = set.String("path", "", "path to the json file containing the request")
	set.SetAnnotation("path", cobra.BashCompFilenameExt, []string{".json"})
}

func (f *PathField) Value() any {
	return *f.value
}

type StringField field[string]

func (f *StringField) AddFlag(set *pflag.FlagSet) {
	f.value = set.String(f.Name, f.defaultValue, f.Usage)
}

type StringSliceField field[[]string]

func (f *StringSliceField) AddFlag(set *pflag.FlagSet) {
	f.value = set.StringSlice(f.Name, f.defaultValue, f.Usage)
}

type BoolField field[bool]

func (f *BoolField) AddFlag(set *pflag.FlagSet) {
	f.value = set.Bool(f.Name, f.defaultValue, f.Usage)
}

type BoolSliceField field[[]bool]

func (f *BoolSliceField) AddFlag(set *pflag.FlagSet) {
	f.value = set.BoolSlice(f.Name, f.defaultValue, f.Usage)
}

type Int32Field field[int32]

func (f *Int32Field) AddFlag(set *pflag.FlagSet) {
	f.value = set.Int32(f.Name, f.defaultValue, f.Usage)
}

type Int32SliceField field[[]int32]

func (f *Int32SliceField) AddFlag(set *pflag.FlagSet) {
	f.value = set.Int32Slice(f.Name, f.defaultValue, f.Usage)
}

type Uint32Field field[uint32]

func (f *Uint32Field) AddFlag(set *pflag.FlagSet) {
	f.value = set.Uint32(f.Name, f.defaultValue, f.Usage)
}

type Uint32SliceField field[[]uint]

func (f *Uint32SliceField) AddFlag(set *pflag.FlagSet) {
	f.value = set.UintSlice(f.Name, f.defaultValue, f.Usage)
}

func (f *Uint32SliceField) Value() any {
	numbers := make([]int32, len(*f.value))

	for i, n := range *f.value {
		numbers[i] = int32(n)
	}

	return numbers
}

type Int64Field field[int64]

func (f *Int64Field) AddFlag(set *pflag.FlagSet) {
	f.value = set.Int64(f.Name, f.defaultValue, f.Usage)
}

type Int64SliceField field[[]int64]

func (f *Int64SliceField) AddFlag(set *pflag.FlagSet) {
	f.value = set.Int64Slice(f.Name, f.defaultValue, f.Usage)
}

type Uint64Field field[uint64]

func (f *Uint64Field) AddFlag(set *pflag.FlagSet) {
	f.value = set.Uint64(f.Name, f.defaultValue, f.Usage)
}

type Uint64SliceField field[[]uint]

func (f *Uint64SliceField) AddFlag(set *pflag.FlagSet) {
	f.value = set.UintSlice(f.Name, f.defaultValue, f.Usage)
}

func (f *Uint64SliceField) Value() any {
	numbers := make([]uint64, len(*f.value))

	for i, n := range *f.value {
		numbers[i] = uint64(n)
	}

	return numbers
}

type FloatField field[float32]

func (f *FloatField) AddFlag(set *pflag.FlagSet) {
	f.value = set.Float32(f.Name, f.defaultValue, f.Usage)
}

type FloatSliceField field[[]float32]

func (f *FloatSliceField) AddFlag(set *pflag.FlagSet) {
	f.value = set.Float32Slice(f.Name, f.defaultValue, f.Usage)
}

type DoubleField field[float64]

func (f *DoubleField) AddFlag(set *pflag.FlagSet) {
	f.value = set.Float64(f.Name, f.defaultValue, f.Usage)
}

type DoubleSliceField field[[]float64]

func (f *DoubleSliceField) AddFlag(set *pflag.FlagSet) {
	f.value = set.Float64Slice(f.Name, f.defaultValue, f.Usage)
}

type BytesField field[[]byte]

func (f *BytesField) AddFlag(set *pflag.FlagSet) {
	f.value = set.BytesBase64(f.Name, f.defaultValue, f.Usage+" base64 (RFC 4648) encoded")
}

type enum interface {
	~int32
	Descriptor() protoreflect.EnumDescriptor
	String() string
}

type EnumField[E enum] field[E]

func (enum *EnumField[E]) AddFlag(set *pflag.FlagSet) {
	e := new(E)
	enum.value = e
	set.Var(enum, enum.Name, enum.Usage)
}

// Set Implements pflag.Value
func (enum *EnumField[E]) Set(s string) (err error) {
	enum.value, err = parseEnum[E](s)
	return err
}

// String Implements pflag.Value
func (enum *EnumField[E]) String() string {
	return (*enum.value).String()
}

// Type Implements pflag.Value
func (enum *EnumField[E]) Type() string {
	return "enum"
}

type EnumSliceField[E enum] field[[]E]

func (enum *EnumSliceField[E]) AddFlag(set *pflag.FlagSet) {
	e := new([]E)
	enum.value = e
	set.Var(enum, enum.Name, enum.Usage)
}

// Set Implements pflag.Value
func (enum *EnumSliceField[E]) Set(s string) error {
	if s == "" {
		return nil
	}
	stringReader := strings.NewReader(s)
	csvReader := csv.NewReader(stringReader)
	records, err := csvReader.Read()
	if err != nil {
		return err
	}

	values := make([]E, len(records))

	for i, record := range records {
		e, err := parseEnum[E](record)
		if err != nil {
			return err
		}
		values[i] = *e
	}

	*enum.value = append(*enum.value, values...)

	return nil
}

// String Implements pflag.Value
func (enum *EnumSliceField[E]) String() string {
	if len(*enum.value) == 0 {
		return ""
	}
	list := make([]string, len(*enum.value))

	for i, e := range *enum.value {
		list[i] = e.String()
	}

	return "[" + strings.Join(list, ",") + "]"
}

// Type Implements pflag.Value
func (enum *EnumSliceField[E]) Type() string {
	return "enum list"
}

func parseEnum[E enum](s string) (*E, error) {
	e := new(E)
	if desc := (*e).Descriptor().Values().ByName(protoreflect.Name(s)); desc != nil {
		*e = E(desc.Number())
		return e, nil
	}

	if number, err := strconv.Atoi(s); err == nil {
		if desc := (*e).Descriptor().Values().ByNumber(protoreflect.EnumNumber(number)); desc != nil {
			*e = E(desc.Number())
			return e, nil
		}
	}

	return nil, errors.New("unknown enum variable")
}