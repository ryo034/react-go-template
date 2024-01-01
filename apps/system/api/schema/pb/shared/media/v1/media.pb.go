// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: shared/media/v1/media.proto

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

type MediaType int32

const (
	MediaType_MEDIA_TYPE_UNKNOWN MediaType = 0
	MediaType_MEDIA_TYPE_JPG     MediaType = 1
	MediaType_MEDIA_TYPE_PNG     MediaType = 2
	MediaType_MEDIA_TYPE_HIEC    MediaType = 3
	MediaType_MEDIA_TYPE_WEBP    MediaType = 4
	MediaType_MEDIA_TYPE_MOV     MediaType = 5
	MediaType_MEDIA_TYPE_MP4     MediaType = 6
)

// Enum value maps for MediaType.
var (
	MediaType_name = map[int32]string{
		0: "MEDIA_TYPE_UNKNOWN",
		1: "MEDIA_TYPE_JPG",
		2: "MEDIA_TYPE_PNG",
		3: "MEDIA_TYPE_HIEC",
		4: "MEDIA_TYPE_WEBP",
		5: "MEDIA_TYPE_MOV",
		6: "MEDIA_TYPE_MP4",
	}
	MediaType_value = map[string]int32{
		"MEDIA_TYPE_UNKNOWN": 0,
		"MEDIA_TYPE_JPG":     1,
		"MEDIA_TYPE_PNG":     2,
		"MEDIA_TYPE_HIEC":    3,
		"MEDIA_TYPE_WEBP":    4,
		"MEDIA_TYPE_MOV":     5,
		"MEDIA_TYPE_MP4":     6,
	}
)

func (x MediaType) Enum() *MediaType {
	p := new(MediaType)
	*p = x
	return p
}

func (x MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_shared_media_v1_media_proto_enumTypes[0].Descriptor()
}

func (MediaType) Type() protoreflect.EnumType {
	return &file_shared_media_v1_media_proto_enumTypes[0]
}

func (x MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaType.Descriptor instead.
func (MediaType) EnumDescriptor() ([]byte, []int) {
	return file_shared_media_v1_media_proto_rawDescGZIP(), []int{0}
}

type MediaBaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaType MediaType `protobuf:"varint,1,opt,name=media_type,json=mediaType,proto3,enum=media.v1.MediaType" json:"media_type,omitempty"`
	Name      string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Order     uint32    `protobuf:"varint,3,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *MediaBaseInfo) Reset() {
	*x = MediaBaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_media_v1_media_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaBaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaBaseInfo) ProtoMessage() {}

func (x *MediaBaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_shared_media_v1_media_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaBaseInfo.ProtoReflect.Descriptor instead.
func (*MediaBaseInfo) Descriptor() ([]byte, []int) {
	return file_shared_media_v1_media_proto_rawDescGZIP(), []int{0}
}

func (x *MediaBaseInfo) GetMediaType() MediaType {
	if x != nil {
		return x.MediaType
	}
	return MediaType_MEDIA_TYPE_UNKNOWN
}

func (x *MediaBaseInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MediaBaseInfo) GetOrder() uint32 {
	if x != nil {
		return x.Order
	}
	return 0
}

type MediaContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Base *MediaBaseInfo `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	Url  string         `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *MediaContent) Reset() {
	*x = MediaContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_media_v1_media_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaContent) ProtoMessage() {}

func (x *MediaContent) ProtoReflect() protoreflect.Message {
	mi := &file_shared_media_v1_media_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaContent.ProtoReflect.Descriptor instead.
func (*MediaContent) Descriptor() ([]byte, []int) {
	return file_shared_media_v1_media_proto_rawDescGZIP(), []int{1}
}

func (x *MediaContent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MediaContent) GetBase() *MediaBaseInfo {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *MediaContent) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content *MediaContent `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_media_v1_media_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_shared_media_v1_media_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_shared_media_v1_media_proto_rawDescGZIP(), []int{2}
}

func (x *Image) GetContent() *MediaContent {
	if x != nil {
		return x.Content
	}
	return nil
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content   *MediaContent `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Thumbnail *MediaContent `protobuf:"bytes,2,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_media_v1_media_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_shared_media_v1_media_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_shared_media_v1_media_proto_rawDescGZIP(), []int{3}
}

func (x *Video) GetContent() *MediaContent {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Video) GetThumbnail() *MediaContent {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

type Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//	*Media_Image
	//	*Media_Video
	Content isMedia_Content `protobuf_oneof:"content"`
}

func (x *Media) Reset() {
	*x = Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_media_v1_media_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_shared_media_v1_media_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_shared_media_v1_media_proto_rawDescGZIP(), []int{4}
}

func (m *Media) GetContent() isMedia_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Media) GetImage() *Image {
	if x, ok := x.GetContent().(*Media_Image); ok {
		return x.Image
	}
	return nil
}

func (x *Media) GetVideo() *Video {
	if x, ok := x.GetContent().(*Media_Video); ok {
		return x.Video
	}
	return nil
}

type isMedia_Content interface {
	isMedia_Content()
}

type Media_Image struct {
	Image *Image `protobuf:"bytes,1,opt,name=image,proto3,oneof"`
}

type Media_Video struct {
	Video *Video `protobuf:"bytes,2,opt,name=video,proto3,oneof"`
}

func (*Media_Image) isMedia_Content() {}

func (*Media_Video) isMedia_Content() {}

type UploadContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to PhotoContent:
	//	*UploadContent_Info
	//	*UploadContent_UploadFile
	PhotoContent isUploadContent_PhotoContent `protobuf_oneof:"photo_content"`
}

func (x *UploadContent) Reset() {
	*x = UploadContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_media_v1_media_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadContent) ProtoMessage() {}

func (x *UploadContent) ProtoReflect() protoreflect.Message {
	mi := &file_shared_media_v1_media_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadContent.ProtoReflect.Descriptor instead.
func (*UploadContent) Descriptor() ([]byte, []int) {
	return file_shared_media_v1_media_proto_rawDescGZIP(), []int{5}
}

func (m *UploadContent) GetPhotoContent() isUploadContent_PhotoContent {
	if m != nil {
		return m.PhotoContent
	}
	return nil
}

func (x *UploadContent) GetInfo() *MediaBaseInfo {
	if x, ok := x.GetPhotoContent().(*UploadContent_Info); ok {
		return x.Info
	}
	return nil
}

func (x *UploadContent) GetUploadFile() []byte {
	if x, ok := x.GetPhotoContent().(*UploadContent_UploadFile); ok {
		return x.UploadFile
	}
	return nil
}

type isUploadContent_PhotoContent interface {
	isUploadContent_PhotoContent()
}

type UploadContent_Info struct {
	Info *MediaBaseInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadContent_UploadFile struct {
	UploadFile []byte `protobuf:"bytes,2,opt,name=upload_file,json=uploadFile,proto3,oneof"`
}

func (*UploadContent_Info) isUploadContent_PhotoContent() {}

func (*UploadContent_UploadFile) isUploadContent_PhotoContent() {}

type UploadImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content *UploadContent `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UploadImage) Reset() {
	*x = UploadImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_media_v1_media_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImage) ProtoMessage() {}

func (x *UploadImage) ProtoReflect() protoreflect.Message {
	mi := &file_shared_media_v1_media_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImage.ProtoReflect.Descriptor instead.
func (*UploadImage) Descriptor() ([]byte, []int) {
	return file_shared_media_v1_media_proto_rawDescGZIP(), []int{6}
}

func (x *UploadImage) GetContent() *UploadContent {
	if x != nil {
		return x.Content
	}
	return nil
}

type UploadVideo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to VideoContent:
	//	*UploadVideo_Thumbnail
	//	*UploadVideo_MainContent
	VideoContent isUploadVideo_VideoContent `protobuf_oneof:"video_content"`
}

func (x *UploadVideo) Reset() {
	*x = UploadVideo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_media_v1_media_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideo) ProtoMessage() {}

func (x *UploadVideo) ProtoReflect() protoreflect.Message {
	mi := &file_shared_media_v1_media_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideo.ProtoReflect.Descriptor instead.
func (*UploadVideo) Descriptor() ([]byte, []int) {
	return file_shared_media_v1_media_proto_rawDescGZIP(), []int{7}
}

func (m *UploadVideo) GetVideoContent() isUploadVideo_VideoContent {
	if m != nil {
		return m.VideoContent
	}
	return nil
}

func (x *UploadVideo) GetThumbnail() *UploadContent {
	if x, ok := x.GetVideoContent().(*UploadVideo_Thumbnail); ok {
		return x.Thumbnail
	}
	return nil
}

func (x *UploadVideo) GetMainContent() *UploadContent {
	if x, ok := x.GetVideoContent().(*UploadVideo_MainContent); ok {
		return x.MainContent
	}
	return nil
}

type isUploadVideo_VideoContent interface {
	isUploadVideo_VideoContent()
}

type UploadVideo_Thumbnail struct {
	Thumbnail *UploadContent `protobuf:"bytes,1,opt,name=thumbnail,proto3,oneof"`
}

type UploadVideo_MainContent struct {
	MainContent *UploadContent `protobuf:"bytes,2,opt,name=main_content,json=mainContent,proto3,oneof"`
}

func (*UploadVideo_Thumbnail) isUploadVideo_VideoContent() {}

func (*UploadVideo_MainContent) isUploadVideo_VideoContent() {}

type MixedUploadContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MixedContent:
	//	*MixedUploadContent_Image
	//	*MixedUploadContent_Video
	MixedContent isMixedUploadContent_MixedContent `protobuf_oneof:"mixed_content"`
}

func (x *MixedUploadContent) Reset() {
	*x = MixedUploadContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_media_v1_media_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MixedUploadContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MixedUploadContent) ProtoMessage() {}

func (x *MixedUploadContent) ProtoReflect() protoreflect.Message {
	mi := &file_shared_media_v1_media_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MixedUploadContent.ProtoReflect.Descriptor instead.
func (*MixedUploadContent) Descriptor() ([]byte, []int) {
	return file_shared_media_v1_media_proto_rawDescGZIP(), []int{8}
}

func (m *MixedUploadContent) GetMixedContent() isMixedUploadContent_MixedContent {
	if m != nil {
		return m.MixedContent
	}
	return nil
}

func (x *MixedUploadContent) GetImage() *UploadImage {
	if x, ok := x.GetMixedContent().(*MixedUploadContent_Image); ok {
		return x.Image
	}
	return nil
}

func (x *MixedUploadContent) GetVideo() *UploadVideo {
	if x, ok := x.GetMixedContent().(*MixedUploadContent_Video); ok {
		return x.Video
	}
	return nil
}

type isMixedUploadContent_MixedContent interface {
	isMixedUploadContent_MixedContent()
}

type MixedUploadContent_Image struct {
	Image *UploadImage `protobuf:"bytes,1,opt,name=image,proto3,oneof"`
}

type MixedUploadContent_Video struct {
	Video *UploadVideo `protobuf:"bytes,2,opt,name=video,proto3,oneof"`
}

func (*MixedUploadContent_Image) isMixedUploadContent_MixedContent() {}

func (*MixedUploadContent_Video) isMixedUploadContent_MixedContent() {}

var File_shared_media_v1_media_proto protoreflect.FileDescriptor

var file_shared_media_v1_media_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x22, 0x6d, 0x0a, 0x0d, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x5d, 0x0a, 0x0c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x39, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x30,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x6f, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x22, 0x64, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x09, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x72, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x40, 0x0a, 0x0b, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x95, 0x01,
	0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x37, 0x0a,
	0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x09, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x3c, 0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x4d, 0x69, 0x78, 0x65, 0x64, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x48, 0x00, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x42, 0x0f, 0x0a, 0x0d, 0x6d, 0x69,
	0x78, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2a, 0x9d, 0x01, 0x0a, 0x09,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x45, 0x44,
	0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4a, 0x50, 0x47, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x44,
	0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x49, 0x45, 0x43, 0x10, 0x03, 0x12, 0x13,
	0x0a, 0x0f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x45, 0x42,
	0x50, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x4f, 0x56, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45, 0x44, 0x49, 0x41,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x50, 0x34, 0x10, 0x06, 0x42, 0x4f, 0x5a, 0x4d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x79, 0x6f, 0x30, 0x33, 0x34,
	0x2f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x2d, 0x67, 0x6f, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shared_media_v1_media_proto_rawDescOnce sync.Once
	file_shared_media_v1_media_proto_rawDescData = file_shared_media_v1_media_proto_rawDesc
)

func file_shared_media_v1_media_proto_rawDescGZIP() []byte {
	file_shared_media_v1_media_proto_rawDescOnce.Do(func() {
		file_shared_media_v1_media_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_media_v1_media_proto_rawDescData)
	})
	return file_shared_media_v1_media_proto_rawDescData
}

var file_shared_media_v1_media_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shared_media_v1_media_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_shared_media_v1_media_proto_goTypes = []interface{}{
	(MediaType)(0),             // 0: media.v1.MediaType
	(*MediaBaseInfo)(nil),      // 1: media.v1.MediaBaseInfo
	(*MediaContent)(nil),       // 2: media.v1.MediaContent
	(*Image)(nil),              // 3: media.v1.Image
	(*Video)(nil),              // 4: media.v1.Video
	(*Media)(nil),              // 5: media.v1.Media
	(*UploadContent)(nil),      // 6: media.v1.UploadContent
	(*UploadImage)(nil),        // 7: media.v1.UploadImage
	(*UploadVideo)(nil),        // 8: media.v1.UploadVideo
	(*MixedUploadContent)(nil), // 9: media.v1.MixedUploadContent
}
var file_shared_media_v1_media_proto_depIdxs = []int32{
	0,  // 0: media.v1.MediaBaseInfo.media_type:type_name -> media.v1.MediaType
	1,  // 1: media.v1.MediaContent.base:type_name -> media.v1.MediaBaseInfo
	2,  // 2: media.v1.Image.content:type_name -> media.v1.MediaContent
	2,  // 3: media.v1.Video.content:type_name -> media.v1.MediaContent
	2,  // 4: media.v1.Video.thumbnail:type_name -> media.v1.MediaContent
	3,  // 5: media.v1.Media.image:type_name -> media.v1.Image
	4,  // 6: media.v1.Media.video:type_name -> media.v1.Video
	1,  // 7: media.v1.UploadContent.info:type_name -> media.v1.MediaBaseInfo
	6,  // 8: media.v1.UploadImage.content:type_name -> media.v1.UploadContent
	6,  // 9: media.v1.UploadVideo.thumbnail:type_name -> media.v1.UploadContent
	6,  // 10: media.v1.UploadVideo.main_content:type_name -> media.v1.UploadContent
	7,  // 11: media.v1.MixedUploadContent.image:type_name -> media.v1.UploadImage
	8,  // 12: media.v1.MixedUploadContent.video:type_name -> media.v1.UploadVideo
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_shared_media_v1_media_proto_init() }
func file_shared_media_v1_media_proto_init() {
	if File_shared_media_v1_media_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shared_media_v1_media_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaBaseInfo); i {
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
		file_shared_media_v1_media_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaContent); i {
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
		file_shared_media_v1_media_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_shared_media_v1_media_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_shared_media_v1_media_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Media); i {
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
		file_shared_media_v1_media_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadContent); i {
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
		file_shared_media_v1_media_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImage); i {
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
		file_shared_media_v1_media_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideo); i {
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
		file_shared_media_v1_media_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MixedUploadContent); i {
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
	file_shared_media_v1_media_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Media_Image)(nil),
		(*Media_Video)(nil),
	}
	file_shared_media_v1_media_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*UploadContent_Info)(nil),
		(*UploadContent_UploadFile)(nil),
	}
	file_shared_media_v1_media_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*UploadVideo_Thumbnail)(nil),
		(*UploadVideo_MainContent)(nil),
	}
	file_shared_media_v1_media_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*MixedUploadContent_Image)(nil),
		(*MixedUploadContent_Video)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shared_media_v1_media_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_media_v1_media_proto_goTypes,
		DependencyIndexes: file_shared_media_v1_media_proto_depIdxs,
		EnumInfos:         file_shared_media_v1_media_proto_enumTypes,
		MessageInfos:      file_shared_media_v1_media_proto_msgTypes,
	}.Build()
	File_shared_media_v1_media_proto = out.File
	file_shared_media_v1_media_proto_rawDesc = nil
	file_shared_media_v1_media_proto_goTypes = nil
	file_shared_media_v1_media_proto_depIdxs = nil
}