// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: adlerhurst/cli/v1alpha/cli.proto

package cliv1alpha

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CLIOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CLIOption) Reset() {
	*x = CLIOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adlerhurst_cli_v1alpha_cli_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CLIOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CLIOption) ProtoMessage() {}

func (x *CLIOption) ProtoReflect() protoreflect.Message {
	mi := &file_adlerhurst_cli_v1alpha_cli_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CLIOption.ProtoReflect.Descriptor instead.
func (*CLIOption) Descriptor() ([]byte, []int) {
	return file_adlerhurst_cli_v1alpha_cli_proto_rawDescGZIP(), []int{0}
}

func (x *CLIOption) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CommandOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ignored bool   `protobuf:"varint,2,opt,name=ignored,proto3" json:"ignored,omitempty"`
}

func (x *CommandOption) Reset() {
	*x = CommandOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adlerhurst_cli_v1alpha_cli_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandOption) ProtoMessage() {}

func (x *CommandOption) ProtoReflect() protoreflect.Message {
	mi := &file_adlerhurst_cli_v1alpha_cli_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandOption.ProtoReflect.Descriptor instead.
func (*CommandOption) Descriptor() ([]byte, []int) {
	return file_adlerhurst_cli_v1alpha_cli_proto_rawDescGZIP(), []int{1}
}

func (x *CommandOption) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommandOption) GetIgnored() bool {
	if x != nil {
		return x.Ignored
	}
	return false
}

type FlagOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ignored bool   `protobuf:"varint,2,opt,name=ignored,proto3" json:"ignored,omitempty"`
}

func (x *FlagOption) Reset() {
	*x = FlagOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adlerhurst_cli_v1alpha_cli_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagOption) ProtoMessage() {}

func (x *FlagOption) ProtoReflect() protoreflect.Message {
	mi := &file_adlerhurst_cli_v1alpha_cli_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagOption.ProtoReflect.Descriptor instead.
func (*FlagOption) Descriptor() ([]byte, []int) {
	return file_adlerhurst_cli_v1alpha_cli_proto_rawDescGZIP(), []int{2}
}

func (x *FlagOption) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FlagOption) GetIgnored() bool {
	if x != nil {
		return x.Ignored
	}
	return false
}

var file_adlerhurst_cli_v1alpha_cli_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*CLIOption)(nil),
		Field:         51000,
		Name:          "adlerhurst.cli.v1alpha.cli",
		Tag:           "bytes,51000,opt,name=cli",
		Filename:      "adlerhurst/cli/v1alpha/cli.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*CommandOption)(nil),
		Field:         51000,
		Name:          "adlerhurst.cli.v1alpha.command",
		Tag:           "bytes,51000,opt,name=command",
		Filename:      "adlerhurst/cli/v1alpha/cli.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*CommandOption)(nil),
		Field:         51000,
		Name:          "adlerhurst.cli.v1alpha.sub_command",
		Tag:           "bytes,51000,opt,name=sub_command",
		Filename:      "adlerhurst/cli/v1alpha/cli.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FlagOption)(nil),
		Field:         51000,
		Name:          "adlerhurst.cli.v1alpha.flag",
		Tag:           "bytes,51000,opt,name=flag",
		Filename:      "adlerhurst/cli/v1alpha/cli.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional adlerhurst.cli.v1alpha.CLIOption cli = 51000;
	E_Cli = &file_adlerhurst_cli_v1alpha_cli_proto_extTypes[0]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional adlerhurst.cli.v1alpha.CommandOption command = 51000;
	E_Command = &file_adlerhurst_cli_v1alpha_cli_proto_extTypes[1]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional adlerhurst.cli.v1alpha.CommandOption sub_command = 51000;
	E_SubCommand = &file_adlerhurst_cli_v1alpha_cli_proto_extTypes[2]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional adlerhurst.cli.v1alpha.FlagOption flag = 51000;
	E_Flag = &file_adlerhurst_cli_v1alpha_cli_proto_extTypes[3]
)

var File_adlerhurst_cli_v1alpha_cli_proto protoreflect.FileDescriptor

var file_adlerhurst_cli_v1alpha_cli_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x64, 0x6c, 0x65, 0x72, 0x68, 0x75, 0x72, 0x73, 0x74, 0x2f, 0x63, 0x6c, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x61, 0x64, 0x6c, 0x65, 0x72, 0x68, 0x75, 0x72, 0x73, 0x74, 0x2e, 0x63,
	0x6c, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x09,
	0x43, 0x4c, 0x49, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a,
	0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x0a, 0x0a,
	0x46, 0x6c, 0x61, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x3a, 0x53, 0x0a, 0x03, 0x63, 0x6c, 0x69, 0x12,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb8, 0x8e,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x64, 0x6c, 0x65, 0x72, 0x68, 0x75, 0x72,
	0x73, 0x74, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43,
	0x4c, 0x49, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x63, 0x6c, 0x69, 0x3a, 0x62, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb8, 0x8e, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x61, 0x64, 0x6c, 0x65, 0x72, 0x68, 0x75, 0x72, 0x73, 0x74, 0x2e, 0x63,
	0x6c, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x3a, 0x68, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xb8, 0x8e, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x64, 0x6c, 0x65, 0x72,
	0x68, 0x75, 0x72, 0x73, 0x74, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x73, 0x75, 0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x3a, 0x57, 0x0a, 0x04, 0x66,
	0x6c, 0x61, 0x67, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xb8, 0x8e, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x64, 0x6c,
	0x65, 0x72, 0x68, 0x75, 0x72, 0x73, 0x74, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x66, 0x6c, 0x61, 0x67, 0x42, 0xf3, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6c,
	0x65, 0x72, 0x68, 0x75, 0x72, 0x73, 0x74, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x42, 0x08, 0x43, 0x6c, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6c, 0x65,
	0x72, 0x68, 0x75, 0x72, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x63, 0x6c, 0x69, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x69,
	0x2f, 0x61, 0x64, 0x6c, 0x65, 0x72, 0x68, 0x75, 0x72, 0x73, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x63, 0x6c, 0x69, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x16, 0x41, 0x64, 0x6c, 0x65, 0x72,
	0x68, 0x75, 0x72, 0x73, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0xca, 0x02, 0x16, 0x41, 0x64, 0x6c, 0x65, 0x72, 0x68, 0x75, 0x72, 0x73, 0x74, 0x5c, 0x43,
	0x6c, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x22, 0x41, 0x64, 0x6c,
	0x65, 0x72, 0x68, 0x75, 0x72, 0x73, 0x74, 0x5c, 0x43, 0x6c, 0x69, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x18, 0x41, 0x64, 0x6c, 0x65, 0x72, 0x68, 0x75, 0x72, 0x73, 0x74, 0x3a, 0x3a, 0x43, 0x6c,
	0x69, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_adlerhurst_cli_v1alpha_cli_proto_rawDescOnce sync.Once
	file_adlerhurst_cli_v1alpha_cli_proto_rawDescData = file_adlerhurst_cli_v1alpha_cli_proto_rawDesc
)

func file_adlerhurst_cli_v1alpha_cli_proto_rawDescGZIP() []byte {
	file_adlerhurst_cli_v1alpha_cli_proto_rawDescOnce.Do(func() {
		file_adlerhurst_cli_v1alpha_cli_proto_rawDescData = protoimpl.X.CompressGZIP(file_adlerhurst_cli_v1alpha_cli_proto_rawDescData)
	})
	return file_adlerhurst_cli_v1alpha_cli_proto_rawDescData
}

var file_adlerhurst_cli_v1alpha_cli_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_adlerhurst_cli_v1alpha_cli_proto_goTypes = []interface{}{
	(*CLIOption)(nil),                   // 0: adlerhurst.cli.v1alpha.CLIOption
	(*CommandOption)(nil),               // 1: adlerhurst.cli.v1alpha.CommandOption
	(*FlagOption)(nil),                  // 2: adlerhurst.cli.v1alpha.FlagOption
	(*descriptorpb.FileOptions)(nil),    // 3: google.protobuf.FileOptions
	(*descriptorpb.ServiceOptions)(nil), // 4: google.protobuf.ServiceOptions
	(*descriptorpb.MethodOptions)(nil),  // 5: google.protobuf.MethodOptions
	(*descriptorpb.FieldOptions)(nil),   // 6: google.protobuf.FieldOptions
}
var file_adlerhurst_cli_v1alpha_cli_proto_depIdxs = []int32{
	3, // 0: adlerhurst.cli.v1alpha.cli:extendee -> google.protobuf.FileOptions
	4, // 1: adlerhurst.cli.v1alpha.command:extendee -> google.protobuf.ServiceOptions
	5, // 2: adlerhurst.cli.v1alpha.sub_command:extendee -> google.protobuf.MethodOptions
	6, // 3: adlerhurst.cli.v1alpha.flag:extendee -> google.protobuf.FieldOptions
	0, // 4: adlerhurst.cli.v1alpha.cli:type_name -> adlerhurst.cli.v1alpha.CLIOption
	1, // 5: adlerhurst.cli.v1alpha.command:type_name -> adlerhurst.cli.v1alpha.CommandOption
	1, // 6: adlerhurst.cli.v1alpha.sub_command:type_name -> adlerhurst.cli.v1alpha.CommandOption
	2, // 7: adlerhurst.cli.v1alpha.flag:type_name -> adlerhurst.cli.v1alpha.FlagOption
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	4, // [4:8] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_adlerhurst_cli_v1alpha_cli_proto_init() }
func file_adlerhurst_cli_v1alpha_cli_proto_init() {
	if File_adlerhurst_cli_v1alpha_cli_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_adlerhurst_cli_v1alpha_cli_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CLIOption); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_adlerhurst_cli_v1alpha_cli_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandOption); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_adlerhurst_cli_v1alpha_cli_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagOption); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_adlerhurst_cli_v1alpha_cli_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_adlerhurst_cli_v1alpha_cli_proto_goTypes,
		DependencyIndexes: file_adlerhurst_cli_v1alpha_cli_proto_depIdxs,
		MessageInfos:      file_adlerhurst_cli_v1alpha_cli_proto_msgTypes,
		ExtensionInfos:    file_adlerhurst_cli_v1alpha_cli_proto_extTypes,
	}.Build()
	File_adlerhurst_cli_v1alpha_cli_proto = out.File
	file_adlerhurst_cli_v1alpha_cli_proto_rawDesc = nil
	file_adlerhurst_cli_v1alpha_cli_proto_goTypes = nil
	file_adlerhurst_cli_v1alpha_cli_proto_depIdxs = nil
}
