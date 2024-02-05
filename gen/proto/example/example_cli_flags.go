// Code generated by protoc-gen-go-cli. DO NOT EDIT.

package example

import (
	pflag "github.com/spf13/pflag"
	os "os"
)

type CallRequestFlag struct {
	*CallRequest

	changed bool
	set     *pflag.FlagSet

	useFieldNameFlag  *StringParser
	useCustomNameFlag *StringParser
	nestedFlag        *CallRequest_NestedFlag
	repNestFlag       []*CallRequest_NestedFlag
	createdAtFlag     *TimestampParser
	payloadFlag       *StructParser
	watFlag           *EnumParser[CallRequest_Wat]
	isSomethingFlag   *BoolParser
	i32Flag           *Int32Parser
	ui32Flag          *Uint32Parser
	i64Flag           *Int64Parser
	ui64Flag          *Uint64Parser
	flFlag            *FloatParser
	dblFlag           *DoubleParser
	beizFlag          *BytesParser
	si32Flag          *Sint32Parser
	si64Flag          *Sint64Parser
	f32Flag           *Fixed32Parser
	f64Flag           *Fixed64Parser
	sf32Flag          *Sfixed32Parser
	sf64Flag          *Sfixed64Parser
	someFlag          *EnumParser[Some]
	repSFlag          *DoubleSliceParser
	repWatFlag        *EnumSliceParser[CallRequest_Wat]
	somethingFlag     *AnyParser
	ooTextFlag        *StringParser
	ooWatFlag         *EnumParser[CallRequest_Wat]
	ooNestedFlag      *CallRequest_NestedFlag
}

func (x *CallRequestFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("CallRequest", pflag.ContinueOnError)

	x.useFieldNameFlag = NewStringFlag(x.set, "use-field-name", "")
	x.useCustomNameFlag = NewStringFlag(x.set, "use-custom-name", "")
	x.repNestFlag = []*CallRequest_NestedFlag{}
	x.createdAtFlag = NewTimestampFlag(x.set, "created-at", "")
	x.payloadFlag = NewStructFlag(x.set, "payload", "")
	x.watFlag = NewEnumFlag[CallRequest_Wat](x.set, "wat", "")
	x.isSomethingFlag = NewBoolFlag(x.set, "is-something", "")
	x.i32Flag = NewInt32Flag(x.set, "i32", "")
	x.ui32Flag = NewUint32Flag(x.set, "ui32", "")
	x.i64Flag = NewInt64Flag(x.set, "i64", "")
	x.ui64Flag = NewUint64Flag(x.set, "ui64", "")
	x.flFlag = NewFloatFlag(x.set, "fl", "")
	x.dblFlag = NewDoubleFlag(x.set, "dbl", "")
	x.beizFlag = NewBytesFlag(x.set, "beiz", "")
	x.si32Flag = NewSint32Flag(x.set, "si32", "")
	x.si64Flag = NewSint64Flag(x.set, "si64", "")
	x.f32Flag = NewFixed32Flag(x.set, "f32", "")
	x.f64Flag = NewFixed64Flag(x.set, "f64", "")
	x.sf32Flag = NewSfixed32Flag(x.set, "sf32", "")
	x.sf64Flag = NewSfixed64Flag(x.set, "sf64", "")
	x.someFlag = NewEnumFlag[Some](x.set, "some", "")
	x.repSFlag = NewDoubleSliceFlag(x.set, "rep-s", "")
	x.repWatFlag = NewEnumSliceFlag[CallRequest_Wat](x.set, "rep-wat", "")
	x.somethingFlag = NewAnyFlag(x.set, "something", "")
	x.ooTextFlag = NewStringFlag(x.set, "oo-text", "")
	x.ooWatFlag = NewEnumFlag[CallRequest_Wat](x.set, "oo-wat", "")
	x.nestedFlag = &CallRequest_NestedFlag{CallRequest_Nested: new(CallRequest_Nested)}
	x.nestedFlag.AddFlags(x.set)
	x.ooNestedFlag = &CallRequest_NestedFlag{CallRequest_Nested: new(CallRequest_Nested)}
	x.ooNestedFlag.AddFlags(x.set)
	parent.AddFlagSet(x.set)
}

func (x *CallRequestFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := fieldIndexes(args, "nested", "rep-nest", "oo-nested")

	if err := x.set.Parse(flagIndexes.primitives().args); err != nil {
		DefaultConfig.Logger.Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if flagIdx := flagIndexes.lastByName("nested"); flagIdx != nil {
		x.nestedFlag.ParseFlags(x.set, flagIdx.args)
	}

	if flagIdx := flagIndexes.lastByName("oo-nested"); flagIdx != nil {
		x.ooNestedFlag.ParseFlags(x.set, flagIdx.args)
	}

	for _, flagIdx := range flagIndexes.byName("rep-nest") {
		x.repNestFlag = append(x.repNestFlag, &CallRequest_NestedFlag{CallRequest_Nested: new(CallRequest_Nested)})
		x.repNestFlag[len(x.repNestFlag)-1].AddFlags(x.set)
		x.repNestFlag[len(x.repNestFlag)-1].ParseFlags(x.set, flagIdx.args)
	}

	if x.useFieldNameFlag.Changed() {
		x.changed = true
		x.UseFieldName = *x.useFieldNameFlag.Value
	}
	if x.useCustomNameFlag.Changed() {
		x.changed = true
		x.UseCustomName = *x.useCustomNameFlag.Value
	}

	if x.nestedFlag.Changed() {
		x.changed = true
		x.Nested = x.nestedFlag.CallRequest_Nested
	}

	if len(x.repNestFlag) > 0 {
		x.changed = true
		x.RepNest = make([]*CallRequest_Nested, len(x.repNestFlag))
		for i, value := range x.repNestFlag {
			x.RepNest[i] = value.CallRequest_Nested
		}
	}

	if x.createdAtFlag.Changed() {
		x.changed = true
		x.CreatedAt = x.createdAtFlag.Value
	}
	if x.payloadFlag.Changed() {
		x.changed = true
		x.Payload = x.payloadFlag.Value
	}
	if x.watFlag.Changed() {
		x.changed = true
		x.Wat = *x.watFlag.Value
	}
	if x.isSomethingFlag.Changed() {
		x.changed = true
		x.IsSomething = *x.isSomethingFlag.Value
	}
	if x.i32Flag.Changed() {
		x.changed = true
		x.I32 = x.i32Flag.Value
	}
	if x.ui32Flag.Changed() {
		x.changed = true
		x.Ui32 = *x.ui32Flag.Value
	}
	if x.i64Flag.Changed() {
		x.changed = true
		x.I64 = *x.i64Flag.Value
	}
	if x.ui64Flag.Changed() {
		x.changed = true
		x.Ui64 = *x.ui64Flag.Value
	}
	if x.flFlag.Changed() {
		x.changed = true
		x.Fl = *x.flFlag.Value
	}
	if x.dblFlag.Changed() {
		x.changed = true
		x.Dbl = *x.dblFlag.Value
	}
	if x.beizFlag.Changed() {
		x.changed = true
		x.Beiz = *x.beizFlag.Value
	}
	if x.si32Flag.Changed() {
		x.changed = true
		x.Si32 = *x.si32Flag.Value
	}
	if x.si64Flag.Changed() {
		x.changed = true
		x.Si64 = *x.si64Flag.Value
	}
	if x.f32Flag.Changed() {
		x.changed = true
		x.F32 = *x.f32Flag.Value
	}
	if x.f64Flag.Changed() {
		x.changed = true
		x.F64 = *x.f64Flag.Value
	}
	if x.sf32Flag.Changed() {
		x.changed = true
		x.Sf32 = *x.sf32Flag.Value
	}
	if x.sf64Flag.Changed() {
		x.changed = true
		x.Sf64 = *x.sf64Flag.Value
	}
	if x.someFlag.Changed() {
		x.changed = true
		x.Some = *x.someFlag.Value
	}
	if x.repSFlag.Changed() {
		x.changed = true
		x.RepS = *x.repSFlag.Value
	}
	if x.repWatFlag.Changed() {
		x.changed = true
		x.RepWat = *x.repWatFlag.Value
	}
	if x.somethingFlag.Changed() {
		x.changed = true
		x.Something = x.somethingFlag.Value
	}

	switch fieldIndexes(args, "oo-text", "oo-wat", "oo-nested").last().flag {
	case "oo-text":
		if x.ooTextFlag.Changed() {
			x.changed = true
			x.Oo = &CallRequest_OoText{OoText: *x.ooTextFlag.Value}
		}
	case "oo-wat":
		if x.ooWatFlag.Changed() {
			x.changed = true
			x.Oo = &CallRequest_OoWat{OoWat: *x.ooWatFlag.Value}
		}
	case "oo-nested":
		if x.ooNestedFlag.Changed() {
			x.changed = true
			x.Oo = &CallRequest_OoNested{OoNested: x.ooNestedFlag.CallRequest_Nested}
		}
	}
}

func (x *CallRequestFlag) Changed() bool {
	return x.changed
}

type CallRequest_NestedFlag struct {
	*CallRequest_Nested

	changed bool
	set     *pflag.FlagSet

	fieldFlag *StringParser
}

func (x *CallRequest_NestedFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("nested", pflag.ContinueOnError)

	x.fieldFlag = NewStringFlag(x.set, "field", "")
	parent.AddFlagSet(x.set)
}

func (x *CallRequest_NestedFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := fieldIndexes(args)

	if err := x.set.Parse(flagIndexes.primitives().args); err != nil {
		DefaultConfig.Logger.Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if x.fieldFlag.Changed() {
		x.changed = true
		x.Field = *x.fieldFlag.Value
	}
}

func (x *CallRequest_NestedFlag) Changed() bool {
	return x.changed
}

type NestedRequestFlag struct {
	*NestedRequest

	changed bool
	set     *pflag.FlagSet

	nestedFlag *NestedRequest_NestedFlag
	idFlag     *StringParser
}

func (x *NestedRequestFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("NestedRequest", pflag.ContinueOnError)

	x.idFlag = NewStringFlag(x.set, "id", "")
	x.nestedFlag = &NestedRequest_NestedFlag{NestedRequest_Nested: new(NestedRequest_Nested)}
	x.nestedFlag.AddFlags(x.set)
	parent.AddFlagSet(x.set)
}

func (x *NestedRequestFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := fieldIndexes(args, "nested")

	if err := x.set.Parse(flagIndexes.primitives().args); err != nil {
		DefaultConfig.Logger.Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if flagIdx := flagIndexes.lastByName("nested"); flagIdx != nil {
		x.nestedFlag.ParseFlags(x.set, flagIdx.args)
	}

	if x.nestedFlag.Changed() {
		x.changed = true
		x.Nested = x.nestedFlag.NestedRequest_Nested
	}

	if x.idFlag.Changed() {
		x.changed = true
		x.Id = *x.idFlag.Value
	}
}

func (x *NestedRequestFlag) Changed() bool {
	return x.changed
}

type NestedRequest_NestedFlag struct {
	*NestedRequest_Nested

	changed bool
	set     *pflag.FlagSet

	idFlag    *StringParser
	depthFlag *Int32Parser
}

func (x *NestedRequest_NestedFlag) AddFlags(parent *pflag.FlagSet) {
	x.set = pflag.NewFlagSet("NestedRequest_Nested", pflag.ContinueOnError)

	x.idFlag = NewStringFlag(x.set, "id", "")
	x.depthFlag = NewInt32Flag(x.set, "depth", "")
	parent.AddFlagSet(x.set)
}

func (x *NestedRequest_NestedFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	flagIndexes := fieldIndexes(args)

	if err := x.set.Parse(flagIndexes.primitives().args); err != nil {
		DefaultConfig.Logger.Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if x.idFlag.Changed() {
		x.changed = true
		x.Id = *x.idFlag.Value
	}
	if x.depthFlag.Changed() {
		x.changed = true
		x.Depth = *x.depthFlag.Value
	}
}

func (x *NestedRequest_NestedFlag) Changed() bool {
	return x.changed
}
