// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/voter/v1/voter-ip_posts.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IpPostsAPIClient is the client API for IpPostsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IpPostsAPIClient interface {
	ListIpPosts(ctx context.Context, in *ListIpPostsRequest, opts ...grpc.CallOption) (*ListIpPostsResponse, error)
	GetIpPost(ctx context.Context, in *GetIpPostRequest, opts ...grpc.CallOption) (*IpPost, error)
	CreateIpPost(ctx context.Context, in *CreateIpPostRequest, opts ...grpc.CallOption) (*IpPost, error)
	UpdateIpPost(ctx context.Context, in *UpdateIpPostRequest, opts ...grpc.CallOption) (*IpPost, error)
	DeleteIpPost(ctx context.Context, in *DeleteIpPostRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type ipPostsAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewIpPostsAPIClient(cc grpc.ClientConnInterface) IpPostsAPIClient {
	return &ipPostsAPIClient{cc}
}

func (c *ipPostsAPIClient) ListIpPosts(ctx context.Context, in *ListIpPostsRequest, opts ...grpc.CallOption) (*ListIpPostsResponse, error) {
	out := new(ListIpPostsResponse)
	err := c.cc.Invoke(ctx, "/postr.ip_posts.v1.IpPostsAPI/ListIpPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipPostsAPIClient) GetIpPost(ctx context.Context, in *GetIpPostRequest, opts ...grpc.CallOption) (*IpPost, error) {
	out := new(IpPost)
	err := c.cc.Invoke(ctx, "/postr.ip_posts.v1.IpPostsAPI/GetIpPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipPostsAPIClient) CreateIpPost(ctx context.Context, in *CreateIpPostRequest, opts ...grpc.CallOption) (*IpPost, error) {
	out := new(IpPost)
	err := c.cc.Invoke(ctx, "/postr.ip_posts.v1.IpPostsAPI/CreateIpPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipPostsAPIClient) UpdateIpPost(ctx context.Context, in *UpdateIpPostRequest, opts ...grpc.CallOption) (*IpPost, error) {
	out := new(IpPost)
	err := c.cc.Invoke(ctx, "/postr.ip_posts.v1.IpPostsAPI/UpdateIpPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipPostsAPIClient) DeleteIpPost(ctx context.Context, in *DeleteIpPostRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/postr.ip_posts.v1.IpPostsAPI/DeleteIpPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IpPostsAPIServer is the server API for IpPostsAPI service.
// All implementations must embed UnimplementedIpPostsAPIServer
// for forward compatibility
type IpPostsAPIServer interface {
	ListIpPosts(context.Context, *ListIpPostsRequest) (*ListIpPostsResponse, error)
	GetIpPost(context.Context, *GetIpPostRequest) (*IpPost, error)
	CreateIpPost(context.Context, *CreateIpPostRequest) (*IpPost, error)
	UpdateIpPost(context.Context, *UpdateIpPostRequest) (*IpPost, error)
	DeleteIpPost(context.Context, *DeleteIpPostRequest) (*empty.Empty, error)
	mustEmbedUnimplementedIpPostsAPIServer()
}

// UnimplementedIpPostsAPIServer must be embedded to have forward compatible implementations.
type UnimplementedIpPostsAPIServer struct {
}

func (UnimplementedIpPostsAPIServer) ListIpPosts(context.Context, *ListIpPostsRequest) (*ListIpPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIpPosts not implemented")
}
func (UnimplementedIpPostsAPIServer) GetIpPost(context.Context, *GetIpPostRequest) (*IpPost, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIpPost not implemented")
}
func (UnimplementedIpPostsAPIServer) CreateIpPost(context.Context, *CreateIpPostRequest) (*IpPost, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIpPost not implemented")
}
func (UnimplementedIpPostsAPIServer) UpdateIpPost(context.Context, *UpdateIpPostRequest) (*IpPost, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIpPost not implemented")
}
func (UnimplementedIpPostsAPIServer) DeleteIpPost(context.Context, *DeleteIpPostRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIpPost not implemented")
}
func (UnimplementedIpPostsAPIServer) mustEmbedUnimplementedIpPostsAPIServer() {}

// UnsafeIpPostsAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IpPostsAPIServer will
// result in compilation errors.
type UnsafeIpPostsAPIServer interface {
	mustEmbedUnimplementedIpPostsAPIServer()
}

func RegisterIpPostsAPIServer(s grpc.ServiceRegistrar, srv IpPostsAPIServer) {
	s.RegisterService(&IpPostsAPI_ServiceDesc, srv)
}

func _IpPostsAPI_ListIpPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIpPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpPostsAPIServer).ListIpPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postr.ip_posts.v1.IpPostsAPI/ListIpPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpPostsAPIServer).ListIpPosts(ctx, req.(*ListIpPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IpPostsAPI_GetIpPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIpPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpPostsAPIServer).GetIpPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postr.ip_posts.v1.IpPostsAPI/GetIpPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpPostsAPIServer).GetIpPost(ctx, req.(*GetIpPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IpPostsAPI_CreateIpPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIpPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpPostsAPIServer).CreateIpPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postr.ip_posts.v1.IpPostsAPI/CreateIpPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpPostsAPIServer).CreateIpPost(ctx, req.(*CreateIpPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IpPostsAPI_UpdateIpPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIpPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpPostsAPIServer).UpdateIpPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postr.ip_posts.v1.IpPostsAPI/UpdateIpPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpPostsAPIServer).UpdateIpPost(ctx, req.(*UpdateIpPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IpPostsAPI_DeleteIpPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIpPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpPostsAPIServer).DeleteIpPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postr.ip_posts.v1.IpPostsAPI/DeleteIpPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpPostsAPIServer).DeleteIpPost(ctx, req.(*DeleteIpPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IpPostsAPI_ServiceDesc is the grpc.ServiceDesc for IpPostsAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IpPostsAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "postr.ip_posts.v1.IpPostsAPI",
	HandlerType: (*IpPostsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListIpPosts",
			Handler:    _IpPostsAPI_ListIpPosts_Handler,
		},
		{
			MethodName: "GetIpPost",
			Handler:    _IpPostsAPI_GetIpPost_Handler,
		},
		{
			MethodName: "CreateIpPost",
			Handler:    _IpPostsAPI_CreateIpPost_Handler,
		},
		{
			MethodName: "UpdateIpPost",
			Handler:    _IpPostsAPI_UpdateIpPost_Handler,
		},
		{
			MethodName: "DeleteIpPost",
			Handler:    _IpPostsAPI_DeleteIpPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/voter/v1/voter-ip_posts.proto",
}
