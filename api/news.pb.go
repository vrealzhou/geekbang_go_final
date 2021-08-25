// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/news.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Limit  uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetCollectionRequest) Reset() {
	*x = GetCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_news_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionRequest) ProtoMessage() {}

func (x *GetCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_news_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionRequest.ProtoReflect.Descriptor instead.
func (*GetCollectionRequest) Descriptor() ([]byte, []int) {
	return file_api_news_proto_rawDescGZIP(), []int{0}
}

func (x *GetCollectionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetCollectionRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetCollectionRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type CollectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Articles []string `protobuf:"bytes,2,rep,name=articles,proto3" json:"articles,omitempty"`
	Limit    uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset   uint32   `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Total    uint32   `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *CollectionResponse) Reset() {
	*x = CollectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_news_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionResponse) ProtoMessage() {}

func (x *CollectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_news_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionResponse.ProtoReflect.Descriptor instead.
func (*CollectionResponse) Descriptor() ([]byte, []int) {
	return file_api_news_proto_rawDescGZIP(), []int{1}
}

func (x *CollectionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CollectionResponse) GetArticles() []string {
	if x != nil {
		return x.Articles
	}
	return nil
}

func (x *CollectionResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *CollectionResponse) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *CollectionResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text    string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Related []string `protobuf:"bytes,3,rep,name=related,proto3" json:"related,omitempty"`
	Images  []string `protobuf:"bytes,4,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_news_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_api_news_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_api_news_proto_rawDescGZIP(), []int{2}
}

func (x *Article) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Article) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Article) GetRelated() []string {
	if x != nil {
		return x.Related
	}
	return nil
}

func (x *Article) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string       `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Crops []*ImageCrop `protobuf:"bytes,3,rep,name=crops,proto3" json:"crops,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_news_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_api_news_proto_msgTypes[3]
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
	return file_api_news_proto_rawDescGZIP(), []int{3}
}

func (x *Image) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Image) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Image) GetCrops() []*ImageCrop {
	if x != nil {
		return x.Crops
	}
	return nil
}

type ImageCrop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Width  uint32 `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height uint32 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *ImageCrop) Reset() {
	*x = ImageCrop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_news_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageCrop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageCrop) ProtoMessage() {}

func (x *ImageCrop) ProtoReflect() protoreflect.Message {
	mi := &file_api_news_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageCrop.ProtoReflect.Descriptor instead.
func (*ImageCrop) Descriptor() ([]byte, []int) {
	return file_api_news_proto_rawDescGZIP(), []int{4}
}

func (x *ImageCrop) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ImageCrop) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ImageCrop) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_api_news_proto protoreflect.FileDescriptor

var file_api_news_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x12,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x5f, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x63, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x72, 0x6f,
	0x70, 0x52, 0x05, 0x63, 0x72, 0x6f, 0x70, 0x73, 0x22, 0x4b, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x43, 0x72, 0x6f, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0xc4, 0x01, 0x0a, 0x04, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x39,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x63, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x53, 0x0a, 0x17,
	0x76, 0x72, 0x65, 0x61, 0x6c, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x62, 0x61, 0x6e, 0x67, 0x5f, 0x67,
	0x6f, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x42, 0x0a, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x76, 0x72, 0x65, 0x61, 0x6c, 0x7a, 0x68, 0x6f, 0x75, 0x2f, 0x67, 0x65, 0x65, 0x6b,
	0x62, 0x61, 0x6e, 0x67, 0x5f, 0x67, 0x6f, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_news_proto_rawDescOnce sync.Once
	file_api_news_proto_rawDescData = file_api_news_proto_rawDesc
)

func file_api_news_proto_rawDescGZIP() []byte {
	file_api_news_proto_rawDescOnce.Do(func() {
		file_api_news_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_news_proto_rawDescData)
	})
	return file_api_news_proto_rawDescData
}

var file_api_news_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_news_proto_goTypes = []interface{}{
	(*GetCollectionRequest)(nil),   // 0: api.GetCollectionRequest
	(*CollectionResponse)(nil),     // 1: api.CollectionResponse
	(*Article)(nil),                // 2: api.Article
	(*Image)(nil),                  // 3: api.Image
	(*ImageCrop)(nil),              // 4: api.ImageCrop
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
}
var file_api_news_proto_depIdxs = []int32{
	4, // 0: api.Image.crops:type_name -> api.ImageCrop
	5, // 1: api.News.GetArtcle:input_type -> google.protobuf.StringValue
	0, // 2: api.News.GetCollection:input_type -> api.GetCollectionRequest
	5, // 3: api.News.GetImage:input_type -> google.protobuf.StringValue
	2, // 4: api.News.GetArtcle:output_type -> api.Article
	1, // 5: api.News.GetCollection:output_type -> api.CollectionResponse
	3, // 6: api.News.GetImage:output_type -> api.Image
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_news_proto_init() }
func file_api_news_proto_init() {
	if File_api_news_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_news_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionRequest); i {
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
		file_api_news_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionResponse); i {
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
		file_api_news_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_api_news_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_news_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageCrop); i {
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
			RawDescriptor: file_api_news_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_news_proto_goTypes,
		DependencyIndexes: file_api_news_proto_depIdxs,
		MessageInfos:      file_api_news_proto_msgTypes,
	}.Build()
	File_api_news_proto = out.File
	file_api_news_proto_rawDesc = nil
	file_api_news_proto_goTypes = nil
	file_api_news_proto_depIdxs = nil
}
