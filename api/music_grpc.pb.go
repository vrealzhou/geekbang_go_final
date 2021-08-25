// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MusicClient is the client API for Music service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MusicClient interface {
	// Get artist by arid
	GetArtist(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Artist, error)
	// Get release information by arid
	GetRelease(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Release, error)
	// Get recording information by arid
	GetRecording(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Recording, error)
	// Get artwork by arid
	GetArtwork(ctx context.Context, opts ...grpc.CallOption) (Music_GetArtworkClient, error)
}

type musicClient struct {
	cc grpc.ClientConnInterface
}

func NewMusicClient(cc grpc.ClientConnInterface) MusicClient {
	return &musicClient{cc}
}

func (c *musicClient) GetArtist(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Artist, error) {
	out := new(Artist)
	err := c.cc.Invoke(ctx, "/api.Music/GetArtist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicClient) GetRelease(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Release, error) {
	out := new(Release)
	err := c.cc.Invoke(ctx, "/api.Music/GetRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicClient) GetRecording(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Recording, error) {
	out := new(Recording)
	err := c.cc.Invoke(ctx, "/api.Music/GetRecording", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicClient) GetArtwork(ctx context.Context, opts ...grpc.CallOption) (Music_GetArtworkClient, error) {
	stream, err := c.cc.NewStream(ctx, &Music_ServiceDesc.Streams[0], "/api.Music/GetArtwork", opts...)
	if err != nil {
		return nil, err
	}
	x := &musicGetArtworkClient{stream}
	return x, nil
}

type Music_GetArtworkClient interface {
	Send(*wrapperspb.StringValue) error
	Recv() (*Artwork, error)
	grpc.ClientStream
}

type musicGetArtworkClient struct {
	grpc.ClientStream
}

func (x *musicGetArtworkClient) Send(m *wrapperspb.StringValue) error {
	return x.ClientStream.SendMsg(m)
}

func (x *musicGetArtworkClient) Recv() (*Artwork, error) {
	m := new(Artwork)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MusicServer is the server API for Music service.
// All implementations must embed UnimplementedMusicServer
// for forward compatibility
type MusicServer interface {
	// Get artist by arid
	GetArtist(context.Context, *wrapperspb.StringValue) (*Artist, error)
	// Get release information by arid
	GetRelease(context.Context, *wrapperspb.StringValue) (*Release, error)
	// Get recording information by arid
	GetRecording(context.Context, *wrapperspb.StringValue) (*Recording, error)
	// Get artwork by arid
	GetArtwork(Music_GetArtworkServer) error
	mustEmbedUnimplementedMusicServer()
}

// UnimplementedMusicServer must be embedded to have forward compatible implementations.
type UnimplementedMusicServer struct {
}

func (UnimplementedMusicServer) GetArtist(context.Context, *wrapperspb.StringValue) (*Artist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtist not implemented")
}
func (UnimplementedMusicServer) GetRelease(context.Context, *wrapperspb.StringValue) (*Release, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelease not implemented")
}
func (UnimplementedMusicServer) GetRecording(context.Context, *wrapperspb.StringValue) (*Recording, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecording not implemented")
}
func (UnimplementedMusicServer) GetArtwork(Music_GetArtworkServer) error {
	return status.Errorf(codes.Unimplemented, "method GetArtwork not implemented")
}
func (UnimplementedMusicServer) mustEmbedUnimplementedMusicServer() {}

// UnsafeMusicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MusicServer will
// result in compilation errors.
type UnsafeMusicServer interface {
	mustEmbedUnimplementedMusicServer()
}

func RegisterMusicServer(s grpc.ServiceRegistrar, srv MusicServer) {
	s.RegisterService(&Music_ServiceDesc, srv)
}

func _Music_GetArtist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServer).GetArtist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Music/GetArtist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServer).GetArtist(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Music_GetRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServer).GetRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Music/GetRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServer).GetRelease(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Music_GetRecording_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServer).GetRecording(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Music/GetRecording",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServer).GetRecording(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Music_GetArtwork_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MusicServer).GetArtwork(&musicGetArtworkServer{stream})
}

type Music_GetArtworkServer interface {
	Send(*Artwork) error
	Recv() (*wrapperspb.StringValue, error)
	grpc.ServerStream
}

type musicGetArtworkServer struct {
	grpc.ServerStream
}

func (x *musicGetArtworkServer) Send(m *Artwork) error {
	return x.ServerStream.SendMsg(m)
}

func (x *musicGetArtworkServer) Recv() (*wrapperspb.StringValue, error) {
	m := new(wrapperspb.StringValue)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Music_ServiceDesc is the grpc.ServiceDesc for Music service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Music_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Music",
	HandlerType: (*MusicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArtist",
			Handler:    _Music_GetArtist_Handler,
		},
		{
			MethodName: "GetRelease",
			Handler:    _Music_GetRelease_Handler,
		},
		{
			MethodName: "GetRecording",
			Handler:    _Music_GetRecording_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetArtwork",
			Handler:       _Music_GetArtwork_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/music.proto",
}
