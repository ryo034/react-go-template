// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: shared/v1/shared.proto

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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZipCode    string `protobuf:"bytes,1,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	Country    string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Prefecture string `protobuf:"bytes,3,opt,name=prefecture,proto3" json:"prefecture,omitempty"`
	City       string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Street     string `protobuf:"bytes,5,opt,name=street,proto3" json:"street,omitempty"`
	Building   string `protobuf:"bytes,6,opt,name=building,proto3" json:"building,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_v1_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_shared_v1_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_shared_v1_shared_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetPrefecture() string {
	if x != nil {
		return x.Prefecture
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetBuilding() string {
	if x != nil {
		return x.Building
	}
	return ""
}

var File_shared_v1_shared_proto protoreflect.FileDescriptor

var file_shared_v1_shared_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x76, 0x31, 0x22, 0xa6, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x66, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x66, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x49, 0x5a, 0x47,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x79, 0x6f, 0x30, 0x33,
	0x34, 0x2f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x2d, 0x67, 0x6f, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shared_v1_shared_proto_rawDescOnce sync.Once
	file_shared_v1_shared_proto_rawDescData = file_shared_v1_shared_proto_rawDesc
)

func file_shared_v1_shared_proto_rawDescGZIP() []byte {
	file_shared_v1_shared_proto_rawDescOnce.Do(func() {
		file_shared_v1_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_v1_shared_proto_rawDescData)
	})
	return file_shared_v1_shared_proto_rawDescData
}

var file_shared_v1_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shared_v1_shared_proto_goTypes = []interface{}{
	(*Address)(nil), // 0: shared.v1.Address
}
var file_shared_v1_shared_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shared_v1_shared_proto_init() }
func file_shared_v1_shared_proto_init() {
	if File_shared_v1_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shared_v1_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_shared_v1_shared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_v1_shared_proto_goTypes,
		DependencyIndexes: file_shared_v1_shared_proto_depIdxs,
		MessageInfos:      file_shared_v1_shared_proto_msgTypes,
	}.Build()
	File_shared_v1_shared_proto = out.File
	file_shared_v1_shared_proto_rawDesc = nil
	file_shared_v1_shared_proto_goTypes = nil
	file_shared_v1_shared_proto_depIdxs = nil
}