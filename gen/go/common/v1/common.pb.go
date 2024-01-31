// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: common/v1/common.proto

package commonv1

import (
	annotations "google.golang.org/genproto/googleapis/api/annotations"
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

type HumansWhoCanRpc int32

const (
	HumansWhoCanRpc_HUMANS_WHO_CAN_RPC_UNSPECIFIED                HumansWhoCanRpc = 0
	HumansWhoCanRpc_HUMANS_WHO_CAN_RPC_CONTRIBUTORS_WITH_APPROVAL HumansWhoCanRpc = 1
	HumansWhoCanRpc_HUMANS_WHO_CAN_RPC_CONTRIBUTORS               HumansWhoCanRpc = 2
	HumansWhoCanRpc_HUMANS_WHO_CAN_RPC_ALL_HUMANS                 HumansWhoCanRpc = 3
)

// Enum value maps for HumansWhoCanRpc.
var (
	HumansWhoCanRpc_name = map[int32]string{
		0: "HUMANS_WHO_CAN_RPC_UNSPECIFIED",
		1: "HUMANS_WHO_CAN_RPC_CONTRIBUTORS_WITH_APPROVAL",
		2: "HUMANS_WHO_CAN_RPC_CONTRIBUTORS",
		3: "HUMANS_WHO_CAN_RPC_ALL_HUMANS",
	}
	HumansWhoCanRpc_value = map[string]int32{
		"HUMANS_WHO_CAN_RPC_UNSPECIFIED":                0,
		"HUMANS_WHO_CAN_RPC_CONTRIBUTORS_WITH_APPROVAL": 1,
		"HUMANS_WHO_CAN_RPC_CONTRIBUTORS":               2,
		"HUMANS_WHO_CAN_RPC_ALL_HUMANS":                 3,
	}
)

func (x HumansWhoCanRpc) Enum() *HumansWhoCanRpc {
	p := new(HumansWhoCanRpc)
	*p = x
	return p
}

func (x HumansWhoCanRpc) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HumansWhoCanRpc) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v1_common_proto_enumTypes[0].Descriptor()
}

func (HumansWhoCanRpc) Type() protoreflect.EnumType {
	return &file_common_v1_common_proto_enumTypes[0]
}

func (x HumansWhoCanRpc) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HumansWhoCanRpc.Descriptor instead.
func (HumansWhoCanRpc) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{0}
}

var file_common_v1_common_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*HumansWhoCanRpc)(nil),
		Field:         50000,
		Name:          "common.v1.humans_who_can_rpc",
		Tag:           "varint,50000,opt,name=humans_who_can_rpc,enum=common.v1.HumansWhoCanRpc",
		Filename:      "common/v1/common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*annotations.HttpRule)(nil),
		Field:         72295728,
		Name:          "common.v1.http",
		Tag:           "bytes,72295728,opt,name=http",
		Filename:      "common/v1/common.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional common.v1.HumansWhoCanRpc humans_who_can_rpc = 50000;
	E_HumansWhoCanRpc = &file_common_v1_common_proto_extTypes[0]
	// optional google.api.HttpRule http = 72295728;
	E_Http = &file_common_v1_common_proto_extTypes[1]
)

var File_common_v1_common_proto protoreflect.FileDescriptor

var file_common_v1_common_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb0, 0x01, 0x0a,
	0x0f, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x73, 0x57, 0x68, 0x6f, 0x43, 0x61, 0x6e, 0x52, 0x70, 0x63,
	0x12, 0x22, 0x0a, 0x1e, 0x48, 0x55, 0x4d, 0x41, 0x4e, 0x53, 0x5f, 0x57, 0x48, 0x4f, 0x5f, 0x43,
	0x41, 0x4e, 0x5f, 0x52, 0x50, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x31, 0x0a, 0x2d, 0x48, 0x55, 0x4d, 0x41, 0x4e, 0x53, 0x5f, 0x57,
	0x48, 0x4f, 0x5f, 0x43, 0x41, 0x4e, 0x5f, 0x52, 0x50, 0x43, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52,
	0x49, 0x42, 0x55, 0x54, 0x4f, 0x52, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x41, 0x50, 0x50,
	0x52, 0x4f, 0x56, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x48, 0x55, 0x4d, 0x41, 0x4e,
	0x53, 0x5f, 0x57, 0x48, 0x4f, 0x5f, 0x43, 0x41, 0x4e, 0x5f, 0x52, 0x50, 0x43, 0x5f, 0x43, 0x4f,
	0x4e, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x4f, 0x52, 0x53, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d,
	0x48, 0x55, 0x4d, 0x41, 0x4e, 0x53, 0x5f, 0x57, 0x48, 0x4f, 0x5f, 0x43, 0x41, 0x4e, 0x5f, 0x52,
	0x50, 0x43, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x48, 0x55, 0x4d, 0x41, 0x4e, 0x53, 0x10, 0x03, 0x3a,
	0x69, 0x0a, 0x12, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x73, 0x5f, 0x77, 0x68, 0x6f, 0x5f, 0x63, 0x61,
	0x6e, 0x5f, 0x72, 0x70, 0x63, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x73,
	0x57, 0x68, 0x6f, 0x43, 0x61, 0x6e, 0x52, 0x70, 0x63, 0x52, 0x0f, 0x68, 0x75, 0x6d, 0x61, 0x6e,
	0x73, 0x57, 0x68, 0x6f, 0x43, 0x61, 0x6e, 0x52, 0x70, 0x63, 0x3a, 0x4b, 0x0a, 0x04, 0x68, 0x74,
	0x74, 0x70, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xb0, 0xca, 0xbc, 0x22, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x42, 0xa2, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x73, 0x74, 0x75, 0x72, 0x6c, 0x61, 0x2f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa,
	0x02, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_v1_common_proto_rawDescOnce sync.Once
	file_common_v1_common_proto_rawDescData = file_common_v1_common_proto_rawDesc
)

func file_common_v1_common_proto_rawDescGZIP() []byte {
	file_common_v1_common_proto_rawDescOnce.Do(func() {
		file_common_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_v1_common_proto_rawDescData)
	})
	return file_common_v1_common_proto_rawDescData
}

var file_common_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_v1_common_proto_goTypes = []interface{}{
	(HumansWhoCanRpc)(0),               // 0: common.v1.HumansWhoCanRpc
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
	(*annotations.HttpRule)(nil),       // 2: google.api.HttpRule
}
var file_common_v1_common_proto_depIdxs = []int32{
	1, // 0: common.v1.humans_who_can_rpc:extendee -> google.protobuf.MethodOptions
	1, // 1: common.v1.http:extendee -> google.protobuf.MethodOptions
	0, // 2: common.v1.humans_who_can_rpc:type_name -> common.v1.HumansWhoCanRpc
	2, // 3: common.v1.http:type_name -> google.api.HttpRule
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_v1_common_proto_init() }
func file_common_v1_common_proto_init() {
	if File_common_v1_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_common_proto_goTypes,
		DependencyIndexes: file_common_v1_common_proto_depIdxs,
		EnumInfos:         file_common_v1_common_proto_enumTypes,
		ExtensionInfos:    file_common_v1_common_proto_extTypes,
	}.Build()
	File_common_v1_common_proto = out.File
	file_common_v1_common_proto_rawDesc = nil
	file_common_v1_common_proto_goTypes = nil
	file_common_v1_common_proto_depIdxs = nil
}
