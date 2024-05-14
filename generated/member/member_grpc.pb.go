// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: member.proto

package member

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MemberService_JoinCommunity_FullMethodName = "/member.MemberService/JoinCommunity"
)

// MemberServiceClient is the client API for MemberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemberServiceClient interface {
	JoinCommunity(ctx context.Context, in *CreateMemberInput, opts ...grpc.CallOption) (*Message, error)
}

type memberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberServiceClient(cc grpc.ClientConnInterface) MemberServiceClient {
	return &memberServiceClient{cc}
}

func (c *memberServiceClient) JoinCommunity(ctx context.Context, in *CreateMemberInput, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, MemberService_JoinCommunity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServiceServer is the server API for MemberService service.
// All implementations must embed UnimplementedMemberServiceServer
// for forward compatibility
type MemberServiceServer interface {
	JoinCommunity(context.Context, *CreateMemberInput) (*Message, error)
	mustEmbedUnimplementedMemberServiceServer()
}

// UnimplementedMemberServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMemberServiceServer struct {
}

func (UnimplementedMemberServiceServer) JoinCommunity(context.Context, *CreateMemberInput) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinCommunity not implemented")
}
func (UnimplementedMemberServiceServer) mustEmbedUnimplementedMemberServiceServer() {}

// UnsafeMemberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemberServiceServer will
// result in compilation errors.
type UnsafeMemberServiceServer interface {
	mustEmbedUnimplementedMemberServiceServer()
}

func RegisterMemberServiceServer(s grpc.ServiceRegistrar, srv MemberServiceServer) {
	s.RegisterService(&MemberService_ServiceDesc, srv)
}

func _MemberService_JoinCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMemberInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).JoinCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberService_JoinCommunity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).JoinCommunity(ctx, req.(*CreateMemberInput))
	}
	return interceptor(ctx, in, info, handler)
}

// MemberService_ServiceDesc is the grpc.ServiceDesc for MemberService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemberService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "member.MemberService",
	HandlerType: (*MemberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinCommunity",
			Handler:    _MemberService_JoinCommunity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "member.proto",
}
