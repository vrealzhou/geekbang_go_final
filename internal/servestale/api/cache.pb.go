// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: internal/servestale/cache.proto

package servestale

import (
	api "github.com/vrealzhou/geekbang_go_final/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_internal_servestale_cache_proto protoreflect.FileDescriptor

var file_internal_servestale_cache_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x73, 0x74, 0x61, 0x6c, 0x65, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x1a, 0x0f, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf5, 0x01, 0x0a, 0x05,
	0x4d, 0x75, 0x73, 0x69, 0x63, 0x12, 0x37, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x39,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12,
	0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3d, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x39, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x28, 0x01, 0x32, 0xc0, 0x01, 0x0a, 0x04, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x38, 0x0a, 0x0c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x63, 0x6c, 0x65, 0x12, 0x0c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x47, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12,
	0x35, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x42, 0x5a, 0x0a, 0x17, 0x76, 0x72, 0x65, 0x61, 0x6c, 0x2e,
	0x67, 0x65, 0x65, 0x6b, 0x62, 0x61, 0x6e, 0x67, 0x5f, 0x67, 0x6f, 0x5f, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x42, 0x0a, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x72, 0x65, 0x61,
	0x6c, 0x7a, 0x68, 0x6f, 0x75, 0x2f, 0x67, 0x65, 0x65, 0x6b, 0x62, 0x61, 0x6e, 0x67, 0x5f, 0x67,
	0x6f, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x74, 0x61,
	0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_internal_servestale_cache_proto_goTypes = []interface{}{
	(*api.Artist)(nil),             // 0: api.Artist
	(*api.Release)(nil),            // 1: api.Release
	(*api.Recording)(nil),          // 2: api.Recording
	(*api.Artwork)(nil),            // 3: api.Artwork
	(*api.Article)(nil),            // 4: api.Article
	(*api.CollectionResponse)(nil), // 5: api.CollectionResponse
	(*api.Image)(nil),              // 6: api.Image
	(*emptypb.Empty)(nil),          // 7: google.protobuf.Empty
}
var file_internal_servestale_cache_proto_depIdxs = []int32{
	0, // 0: servestale.Music.UpdateArtist:input_type -> api.Artist
	1, // 1: servestale.Music.UpdateRelease:input_type -> api.Release
	2, // 2: servestale.Music.UpdateRecording:input_type -> api.Recording
	3, // 3: servestale.Music.UpdateArtwork:input_type -> api.Artwork
	4, // 4: servestale.News.UpdateArtcle:input_type -> api.Article
	5, // 5: servestale.News.UpdateCollection:input_type -> api.CollectionResponse
	6, // 6: servestale.News.UpdateImage:input_type -> api.Image
	7, // 7: servestale.Music.UpdateArtist:output_type -> google.protobuf.Empty
	7, // 8: servestale.Music.UpdateRelease:output_type -> google.protobuf.Empty
	7, // 9: servestale.Music.UpdateRecording:output_type -> google.protobuf.Empty
	7, // 10: servestale.Music.UpdateArtwork:output_type -> google.protobuf.Empty
	7, // 11: servestale.News.UpdateArtcle:output_type -> google.protobuf.Empty
	7, // 12: servestale.News.UpdateCollection:output_type -> google.protobuf.Empty
	7, // 13: servestale.News.UpdateImage:output_type -> google.protobuf.Empty
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_servestale_cache_proto_init() }
func file_internal_servestale_cache_proto_init() {
	if File_internal_servestale_cache_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_servestale_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_internal_servestale_cache_proto_goTypes,
		DependencyIndexes: file_internal_servestale_cache_proto_depIdxs,
	}.Build()
	File_internal_servestale_cache_proto = out.File
	file_internal_servestale_cache_proto_rawDesc = nil
	file_internal_servestale_cache_proto_goTypes = nil
	file_internal_servestale_cache_proto_depIdxs = nil
}
