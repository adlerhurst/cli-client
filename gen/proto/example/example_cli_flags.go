// Code generated by protoc-gen-go-cli. DO NOT EDIT.

package example

import (
	pflag "github.com/spf13/pflag"
	os "os"
)

type CallRequestFlag struct {
	*CallRequest

	changed bool
}

func (x *CallRequestFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	set := pflag.NewFlagSet("CallRequest", pflag.ContinueOnError)
	parent.AddFlagSet(set)
	flagIndexes := fieldIndexes(args, "nested", "repNest", "ooNested")

	useFieldNameFlag := NewStringFlag(set, "useFieldName", "")
	useCustomNameFlag := NewStringFlag(set, "useCustomName", "")
	nestedFlag := &CallRequest_NestedFlag{CallRequest_Nested: new(CallRequest_Nested)}
	repNestFlag := make([]*CallRequest_NestedFlag, len(flagIndexes.byName("repNest")))
	createdAtFlag := NewTimestampFlag(set, "createdAt", "")
	payloadFlag := NewStructFlag(set, "payload", "")
	watFlag := NewEnumFlag[CallRequest_Wat](set, "wat", "")
	isSomethingFlag := NewBoolFlag(set, "isSomething", "")
	i32Flag := NewInt32Flag(set, "i32", "")
	ui32Flag := NewUint32Flag(set, "ui32", "")
	i64Flag := NewInt64Flag(set, "i64", "")
	ui64Flag := NewUint64Flag(set, "ui64", "")
	flFlag := NewFloatFlag(set, "fl", "")
	dblFlag := NewDoubleFlag(set, "dbl", "")
	beizFlag := NewBytesFlag(set, "beiz", "")
	si32Flag := NewSint32Flag(set, "si32", "")
	si64Flag := NewSint64Flag(set, "si64", "")
	f32Flag := NewFixed32Flag(set, "f32", "")
	f64Flag := NewFixed64Flag(set, "f64", "")
	sf32Flag := NewSfixed32Flag(set, "sf32", "")
	sf64Flag := NewSfixed64Flag(set, "sf64", "")
	someFlag := NewEnumFlag[Some](set, "some", "")
	repSFlag := NewDoubleSliceFlag(set, "repS", "")
	repWatFlag := NewEnumSliceFlag[CallRequest_Wat](set, "repWat", "")
	somethingFlag := NewAnyFlag(set, "something", "")
	ooTextFlag := NewStringFlag(set, "ooText", "")
	ooWatFlag := NewEnumFlag[CallRequest_Wat](set, "ooWat", "")
	ooNestedFlag := &CallRequest_NestedFlag{CallRequest_Nested: new(CallRequest_Nested)}

	if err := set.Parse(flagIndexes.primitives().args); err != nil {
		DefaultConfig.Logger.Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if flagIdx := flagIndexes.lastByName("nested"); flagIdx != nil {
		nestedFlag.ParseFlags(set, flagIdx.args)
	}

	if flagIdx := flagIndexes.lastByName("ooNested"); flagIdx != nil {
		ooNestedFlag.ParseFlags(set, flagIdx.args)
	}

	for i, flagIdx := range flagIndexes.byName("repNest") {
		repNestFlag[i] = new(CallRequest_NestedFlag)
		repNestFlag[i].ParseFlags(set, flagIdx.args)
	}

	if useFieldNameFlag.Changed() {
		x.changed = true
		x.UseFieldName = *useFieldNameFlag.Value
	}
	if useCustomNameFlag.Changed() {
		x.changed = true
		x.UseCustomName = *useCustomNameFlag.Value
	}

	if nestedFlag.Changed() {
		x.changed = true
		x.Nested = nestedFlag.CallRequest_Nested
	}

	if len(repNestFlag) > 0 {
		x.changed = true
		x.RepNest = make([]*CallRequest_Nested, len(repNestFlag))
		for i, value := range repNestFlag {
			x.RepNest[i] = value.CallRequest_Nested
		}
	}

	if createdAtFlag.Changed() {
		x.changed = true
		x.CreatedAt = createdAtFlag.Value
	}
	if payloadFlag.Changed() {
		x.changed = true
		x.Payload = payloadFlag.Value
	}
	if watFlag.Changed() {
		x.changed = true
		x.Wat = *watFlag.Value
	}
	if isSomethingFlag.Changed() {
		x.changed = true
		x.IsSomething = *isSomethingFlag.Value
	}
	if i32Flag.Changed() {
		x.changed = true
		x.I32 = i32Flag.Value
	}
	if ui32Flag.Changed() {
		x.changed = true
		x.Ui32 = *ui32Flag.Value
	}
	if i64Flag.Changed() {
		x.changed = true
		x.I64 = *i64Flag.Value
	}
	if ui64Flag.Changed() {
		x.changed = true
		x.Ui64 = *ui64Flag.Value
	}
	if flFlag.Changed() {
		x.changed = true
		x.Fl = *flFlag.Value
	}
	if dblFlag.Changed() {
		x.changed = true
		x.Dbl = *dblFlag.Value
	}
	if beizFlag.Changed() {
		x.changed = true
		x.Beiz = *beizFlag.Value
	}
	if si32Flag.Changed() {
		x.changed = true
		x.Si32 = *si32Flag.Value
	}
	if si64Flag.Changed() {
		x.changed = true
		x.Si64 = *si64Flag.Value
	}
	if f32Flag.Changed() {
		x.changed = true
		x.F32 = *f32Flag.Value
	}
	if f64Flag.Changed() {
		x.changed = true
		x.F64 = *f64Flag.Value
	}
	if sf32Flag.Changed() {
		x.changed = true
		x.Sf32 = *sf32Flag.Value
	}
	if sf64Flag.Changed() {
		x.changed = true
		x.Sf64 = *sf64Flag.Value
	}
	if someFlag.Changed() {
		x.changed = true
		x.Some = *someFlag.Value
	}
	if repSFlag.Changed() {
		x.changed = true
		x.RepS = *repSFlag.Value
	}
	if repWatFlag.Changed() {
		x.changed = true
		x.RepWat = *repWatFlag.Value
	}
	if somethingFlag.Changed() {
		x.changed = true
		x.Something = somethingFlag.Value
	}

	switch fieldIndexes(args, "ooText", "ooWat", "ooNested").last().flag {
	case "ooText":
		if ooTextFlag.Changed() {
			x.changed = true
			x.Oo = &CallRequest_OoText{OoText: *ooTextFlag.Value}
		}
	case "ooWat":
		if ooWatFlag.Changed() {
			x.changed = true
			x.Oo = &CallRequest_OoWat{OoWat: *ooWatFlag.Value}
		}
	case "ooNested":
		if ooNestedFlag.Changed() {
			x.changed = true
			x.Oo = &CallRequest_OoNested{OoNested: ooNestedFlag.CallRequest_Nested}
		}
	}
}

func (x *CallRequestFlag) Changed() bool {
	return x.changed
}

type CallRequest_NestedFlag struct {
	*CallRequest_Nested

	changed bool
}

func (x *CallRequest_NestedFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	set := pflag.NewFlagSet("CallRequest_Nested", pflag.ContinueOnError)
	parent.AddFlagSet(set)
	flagIndexes := fieldIndexes(args)

	fieldFlag := NewStringFlag(set, "field", "")

	if err := set.Parse(flagIndexes.primitives().args); err != nil {
		DefaultConfig.Logger.Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if fieldFlag.Changed() {
		x.changed = true
		x.Field = *fieldFlag.Value
	}
}

func (x *CallRequest_NestedFlag) Changed() bool {
	return x.changed
}

type NestedRequestFlag struct {
	*NestedRequest

	changed bool
}

func (x *NestedRequestFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	set := pflag.NewFlagSet("NestedRequest", pflag.ContinueOnError)
	parent.AddFlagSet(set)
	flagIndexes := fieldIndexes(args, "nested")

	nestedFlag := &NestedRequest_NestedFlag{NestedRequest_Nested: new(NestedRequest_Nested)}
	idFlag := NewStringFlag(set, "id", "")

	if err := set.Parse(flagIndexes.primitives().args); err != nil {
		DefaultConfig.Logger.Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if flagIdx := flagIndexes.lastByName("nested"); flagIdx != nil {
		nestedFlag.ParseFlags(set, flagIdx.args)
	}

	if nestedFlag.Changed() {
		x.changed = true
		x.Nested = nestedFlag.NestedRequest_Nested
	}

	if idFlag.Changed() {
		x.changed = true
		x.Id = *idFlag.Value
	}
}

func (x *NestedRequestFlag) Changed() bool {
	return x.changed
}

type NestedRequest_NestedFlag struct {
	*NestedRequest_Nested

	changed bool
}

func (x *NestedRequest_NestedFlag) ParseFlags(parent *pflag.FlagSet, args []string) {
	set := pflag.NewFlagSet("NestedRequest_Nested", pflag.ContinueOnError)
	parent.AddFlagSet(set)
	flagIndexes := fieldIndexes(args)

	idFlag := NewStringFlag(set, "id", "")
	depthFlag := NewInt32Flag(set, "depth", "")

	if err := set.Parse(flagIndexes.primitives().args); err != nil {
		DefaultConfig.Logger.Error("failed to parse flags", "cause", err)
		os.Exit(1)
	}

	if idFlag.Changed() {
		x.changed = true
		x.Id = *idFlag.Value
	}
	if depthFlag.Changed() {
		x.changed = true
		x.Depth = *depthFlag.Value
	}
}

func (x *NestedRequest_NestedFlag) Changed() bool {
	return x.changed
}
