// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/music.proto

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

type Artwork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arid   string `protobuf:"bytes,1,opt,name=arid,proto3" json:"arid,omitempty"`
	Url    string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Width  uint32 `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`
	Height uint32 `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
	Size   uint32 `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Artwork) Reset() {
	*x = Artwork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_music_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artwork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artwork) ProtoMessage() {}

func (x *Artwork) ProtoReflect() protoreflect.Message {
	mi := &file_api_music_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artwork.ProtoReflect.Descriptor instead.
func (*Artwork) Descriptor() ([]byte, []int) {
	return file_api_music_proto_rawDescGZIP(), []int{0}
}

func (x *Artwork) GetArid() string {
	if x != nil {
		return x.Arid
	}
	return ""
}

func (x *Artwork) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Artwork) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Artwork) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Artwork) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Artwork) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Artist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arid     string   `protobuf:"bytes,1,opt,name=arid,proto3" json:"arid,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type     string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Artworks []string `protobuf:"bytes,4,rep,name=artworks,proto3" json:"artworks,omitempty"`
}

func (x *Artist) Reset() {
	*x = Artist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_music_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artist) ProtoMessage() {}

func (x *Artist) ProtoReflect() protoreflect.Message {
	mi := &file_api_music_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artist.ProtoReflect.Descriptor instead.
func (*Artist) Descriptor() ([]byte, []int) {
	return file_api_music_proto_rawDescGZIP(), []int{1}
}

func (x *Artist) GetArid() string {
	if x != nil {
		return x.Arid
	}
	return ""
}

func (x *Artist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Artist) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Artist) GetArtworks() []string {
	if x != nil {
		return x.Artworks
	}
	return nil
}

type Recording struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arid     string   `protobuf:"bytes,1,opt,name=arid,proto3" json:"arid,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Duration uint32   `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Metadata string   `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Artists  []string `protobuf:"bytes,5,rep,name=artists,proto3" json:"artists,omitempty"`
	Releases []string `protobuf:"bytes,6,rep,name=releases,proto3" json:"releases,omitempty"`
	Artworks []string `protobuf:"bytes,7,rep,name=artworks,proto3" json:"artworks,omitempty"`
}

func (x *Recording) Reset() {
	*x = Recording{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_music_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recording) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recording) ProtoMessage() {}

func (x *Recording) ProtoReflect() protoreflect.Message {
	mi := &file_api_music_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recording.ProtoReflect.Descriptor instead.
func (*Recording) Descriptor() ([]byte, []int) {
	return file_api_music_proto_rawDescGZIP(), []int{2}
}

func (x *Recording) GetArid() string {
	if x != nil {
		return x.Arid
	}
	return ""
}

func (x *Recording) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Recording) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Recording) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *Recording) GetArtists() []string {
	if x != nil {
		return x.Artists
	}
	return nil
}

func (x *Recording) GetReleases() []string {
	if x != nil {
		return x.Releases
	}
	return nil
}

func (x *Recording) GetArtworks() []string {
	if x != nil {
		return x.Artworks
	}
	return nil
}

type Release struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arid        string   `protobuf:"bytes,1,opt,name=arid,proto3" json:"arid,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Year        uint32   `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	RecordLabel string   `protobuf:"bytes,4,opt,name=recordLabel,proto3" json:"recordLabel,omitempty"`
	Artists     []string `protobuf:"bytes,5,rep,name=artists,proto3" json:"artists,omitempty"`
	Recordings  []string `protobuf:"bytes,6,rep,name=recordings,proto3" json:"recordings,omitempty"`
	Artworks    []string `protobuf:"bytes,7,rep,name=artworks,proto3" json:"artworks,omitempty"`
}

func (x *Release) Reset() {
	*x = Release{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_music_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Release) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release) ProtoMessage() {}

func (x *Release) ProtoReflect() protoreflect.Message {
	mi := &file_api_music_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Release.ProtoReflect.Descriptor instead.
func (*Release) Descriptor() ([]byte, []int) {
	return file_api_music_proto_rawDescGZIP(), []int{3}
}

func (x *Release) GetArid() string {
	if x != nil {
		return x.Arid
	}
	return ""
}

func (x *Release) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Release) GetYear() uint32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Release) GetRecordLabel() string {
	if x != nil {
		return x.RecordLabel
	}
	return ""
}

func (x *Release) GetArtists() []string {
	if x != nil {
		return x.Artists
	}
	return nil
}

func (x *Release) GetRecordings() []string {
	if x != nil {
		return x.Recordings
	}
	return nil
}

func (x *Release) GetArtworks() []string {
	if x != nil {
		return x.Artworks
	}
	return nil
}

var File_api_music_proto protoreflect.FileDescriptor

var file_api_music_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x72, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x22, 0x60, 0x0a, 0x06, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x72, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x72, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x32, 0xfd, 0x01, 0x0a, 0x05, 0x4d, 0x75, 0x73, 0x69, 0x63,
	0x12, 0x38, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x53, 0x0a, 0x17, 0x76, 0x72, 0x65, 0x61, 0x6c, 0x2e,
	0x67, 0x65, 0x65, 0x6b, 0x62, 0x61, 0x6e, 0x67, 0x5f, 0x67, 0x6f, 0x5f, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x42, 0x0a, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x72, 0x65, 0x61,
	0x6c, 0x7a, 0x68, 0x6f, 0x75, 0x2f, 0x67, 0x65, 0x65, 0x6b, 0x62, 0x61, 0x6e, 0x67, 0x5f, 0x67,
	0x6f, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_music_proto_rawDescOnce sync.Once
	file_api_music_proto_rawDescData = file_api_music_proto_rawDesc
)

func file_api_music_proto_rawDescGZIP() []byte {
	file_api_music_proto_rawDescOnce.Do(func() {
		file_api_music_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_music_proto_rawDescData)
	})
	return file_api_music_proto_rawDescData
}

var file_api_music_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_music_proto_goTypes = []interface{}{
	(*Artwork)(nil),                // 0: api.Artwork
	(*Artist)(nil),                 // 1: api.Artist
	(*Recording)(nil),              // 2: api.Recording
	(*Release)(nil),                // 3: api.Release
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
}
var file_api_music_proto_depIdxs = []int32{
	4, // 0: api.Music.GetArtist:input_type -> google.protobuf.StringValue
	4, // 1: api.Music.GetRelease:input_type -> google.protobuf.StringValue
	4, // 2: api.Music.GetRecording:input_type -> google.protobuf.StringValue
	4, // 3: api.Music.GetArtwork:input_type -> google.protobuf.StringValue
	1, // 4: api.Music.GetArtist:output_type -> api.Artist
	3, // 5: api.Music.GetRelease:output_type -> api.Release
	2, // 6: api.Music.GetRecording:output_type -> api.Recording
	0, // 7: api.Music.GetArtwork:output_type -> api.Artwork
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_music_proto_init() }
func file_api_music_proto_init() {
	if File_api_music_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_music_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artwork); i {
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
		file_api_music_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artist); i {
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
		file_api_music_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recording); i {
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
		file_api_music_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Release); i {
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
			RawDescriptor: file_api_music_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_music_proto_goTypes,
		DependencyIndexes: file_api_music_proto_depIdxs,
		MessageInfos:      file_api_music_proto_msgTypes,
	}.Build()
	File_api_music_proto = out.File
	file_api_music_proto_rawDesc = nil
	file_api_music_proto_goTypes = nil
	file_api_music_proto_depIdxs = nil
}
