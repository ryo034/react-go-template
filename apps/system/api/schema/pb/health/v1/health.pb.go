// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: health/v1/health.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServingStatus int32

const (
	ServingStatus_SERVING_STATUS_UNSPECIFIED ServingStatus = 0
	ServingStatus_SERVING_STATUS_SERVING     ServingStatus = 1
	ServingStatus_SERVING_STATUS_NOT_SERVING ServingStatus = 2
)

// Enum value maps for ServingStatus.
var (
	ServingStatus_name = map[int32]string{
		0: "SERVING_STATUS_UNSPECIFIED",
		1: "SERVING_STATUS_SERVING",
		2: "SERVING_STATUS_NOT_SERVING",
	}
	ServingStatus_value = map[string]int32{
		"SERVING_STATUS_UNSPECIFIED": 0,
		"SERVING_STATUS_SERVING":     1,
		"SERVING_STATUS_NOT_SERVING": 2,
	}
)

func (x ServingStatus) Enum() *ServingStatus {
	p := new(ServingStatus)
	*p = x
	return p
}

func (x ServingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_health_v1_health_proto_enumTypes[0].Descriptor()
}

func (ServingStatus) Type() protoreflect.EnumType {
	return &file_health_v1_health_proto_enumTypes[0]
}

func (x ServingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServingStatus.Descriptor instead.
func (ServingStatus) EnumDescriptor() ([]byte, []int) {
	return file_health_v1_health_proto_rawDescGZIP(), []int{0}
}

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_v1_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_health_v1_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_health_v1_health_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=health.v1.ServingStatus" json:"status,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_v1_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_health_v1_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_health_v1_health_proto_rawDescGZIP(), []int{1}
}

func (x *CheckResponse) GetStatus() ServingStatus {
	if x != nil {
		return x.Status
	}
	return ServingStatus_SERVING_STATUS_UNSPECIFIED
}

var File_health_v1_health_proto protoreflect.FileDescriptor

var file_health_v1_health_proto_rawDesc = []byte{
	0x0a, 0x16, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x22, 0x28, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x41, 0x0a,
	0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2a, 0x6b, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1e, 0x0a,
	0x1a, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x32, 0x4b, 0x0a,
	0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a,
	0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x79, 0x6f, 0x30, 0x33, 0x34, 0x2f,
	0x72, 0x65, 0x61, 0x63, 0x74, 0x2d, 0x67, 0x6f, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x62, 0x2f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_health_v1_health_proto_rawDescOnce sync.Once
	file_health_v1_health_proto_rawDescData = file_health_v1_health_proto_rawDesc
)

func file_health_v1_health_proto_rawDescGZIP() []byte {
	file_health_v1_health_proto_rawDescOnce.Do(func() {
		file_health_v1_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_v1_health_proto_rawDescData)
	})
	return file_health_v1_health_proto_rawDescData
}

var file_health_v1_health_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_health_v1_health_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_health_v1_health_proto_goTypes = []interface{}{
	(ServingStatus)(0),    // 0: health.v1.ServingStatus
	(*CheckRequest)(nil),  // 1: health.v1.CheckRequest
	(*CheckResponse)(nil), // 2: health.v1.CheckResponse
}
var file_health_v1_health_proto_depIdxs = []int32{
	0, // 0: health.v1.CheckResponse.status:type_name -> health.v1.ServingStatus
	1, // 1: health.v1.HealthService.Check:input_type -> health.v1.CheckRequest
	2, // 2: health.v1.HealthService.Check:output_type -> health.v1.CheckResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_health_v1_health_proto_init() }
func file_health_v1_health_proto_init() {
	if File_health_v1_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_health_v1_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_health_v1_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
			RawDescriptor: file_health_v1_health_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_v1_health_proto_goTypes,
		DependencyIndexes: file_health_v1_health_proto_depIdxs,
		EnumInfos:         file_health_v1_health_proto_enumTypes,
		MessageInfos:      file_health_v1_health_proto_msgTypes,
	}.Build()
	File_health_v1_health_proto = out.File
	file_health_v1_health_proto_rawDesc = nil
	file_health_v1_health_proto_goTypes = nil
	file_health_v1_health_proto_depIdxs = nil
}
