// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcapi_ratelimiter.proto

package ratelimiter // import "github.com/polarismesh/specification/source/go/api/v1/traffic_manage/ratelimiter"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RateLimitGRPCV2Client is the client API for RateLimitGRPCV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitGRPCV2Client interface {
	// 限流接口
	Service(ctx context.Context, opts ...grpc.CallOption) (RateLimitGRPCV2_ServiceClient, error)
	// 时间对齐接口
	TimeAdjust(ctx context.Context, in *TimeAdjustRequest, opts ...grpc.CallOption) (*TimeAdjustResponse, error)
}

type rateLimitGRPCV2Client struct {
	cc *grpc.ClientConn
}

func NewRateLimitGRPCV2Client(cc *grpc.ClientConn) RateLimitGRPCV2Client {
	return &rateLimitGRPCV2Client{cc}
}

func (c *rateLimitGRPCV2Client) Service(ctx context.Context, opts ...grpc.CallOption) (RateLimitGRPCV2_ServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RateLimitGRPCV2_serviceDesc.Streams[0], "/polaris.metric.v2.RateLimitGRPCV2/Service", opts...)
	if err != nil {
		return nil, err
	}
	x := &rateLimitGRPCV2ServiceClient{stream}
	return x, nil
}

type RateLimitGRPCV2_ServiceClient interface {
	Send(*RateLimitRequest) error
	Recv() (*RateLimitResponse, error)
	grpc.ClientStream
}

type rateLimitGRPCV2ServiceClient struct {
	grpc.ClientStream
}

func (x *rateLimitGRPCV2ServiceClient) Send(m *RateLimitRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rateLimitGRPCV2ServiceClient) Recv() (*RateLimitResponse, error) {
	m := new(RateLimitResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rateLimitGRPCV2Client) TimeAdjust(ctx context.Context, in *TimeAdjustRequest, opts ...grpc.CallOption) (*TimeAdjustResponse, error) {
	out := new(TimeAdjustResponse)
	err := c.cc.Invoke(ctx, "/polaris.metric.v2.RateLimitGRPCV2/TimeAdjust", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitGRPCV2Server is the server API for RateLimitGRPCV2 service.
type RateLimitGRPCV2Server interface {
	// 限流接口
	Service(RateLimitGRPCV2_ServiceServer) error
	// 时间对齐接口
	TimeAdjust(context.Context, *TimeAdjustRequest) (*TimeAdjustResponse, error)
}

func RegisterRateLimitGRPCV2Server(s *grpc.Server, srv RateLimitGRPCV2Server) {
	s.RegisterService(&_RateLimitGRPCV2_serviceDesc, srv)
}

func _RateLimitGRPCV2_Service_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RateLimitGRPCV2Server).Service(&rateLimitGRPCV2ServiceServer{stream})
}

type RateLimitGRPCV2_ServiceServer interface {
	Send(*RateLimitResponse) error
	Recv() (*RateLimitRequest, error)
	grpc.ServerStream
}

type rateLimitGRPCV2ServiceServer struct {
	grpc.ServerStream
}

func (x *rateLimitGRPCV2ServiceServer) Send(m *RateLimitResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rateLimitGRPCV2ServiceServer) Recv() (*RateLimitRequest, error) {
	m := new(RateLimitRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RateLimitGRPCV2_TimeAdjust_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeAdjustRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitGRPCV2Server).TimeAdjust(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/polaris.metric.v2.RateLimitGRPCV2/TimeAdjust",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitGRPCV2Server).TimeAdjust(ctx, req.(*TimeAdjustRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitGRPCV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "polaris.metric.v2.RateLimitGRPCV2",
	HandlerType: (*RateLimitGRPCV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TimeAdjust",
			Handler:    _RateLimitGRPCV2_TimeAdjust_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Service",
			Handler:       _RateLimitGRPCV2_Service_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpcapi_ratelimiter.proto",
}

func init() {
	proto.RegisterFile("grpcapi_ratelimiter.proto", fileDescriptor_grpcapi_ratelimiter_e6f61dd264c034b3)
}

var fileDescriptor_grpcapi_ratelimiter_e6f61dd264c034b3 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x3d, 0x4e, 0x33, 0x31,
	0x10, 0x40, 0xbf, 0x6d, 0x3e, 0x24, 0x37, 0x28, 0xee, 0x48, 0xc9, 0x8f, 0x44, 0x35, 0x26, 0xcb,
	0x09, 0x48, 0x0a, 0x1a, 0x8a, 0x68, 0x41, 0x14, 0xa1, 0x88, 0x1c, 0x33, 0xbb, 0x19, 0x14, 0xff,
	0x30, 0x9e, 0xdd, 0x4b, 0x71, 0x13, 0x4e, 0x85, 0x80, 0xcd, 0x12, 0x14, 0x04, 0xad, 0xfd, 0xe6,
	0x3d, 0x6b, 0xac, 0x8e, 0x1a, 0x4e, 0xce, 0x26, 0x5a, 0xb2, 0x15, 0xdc, 0x90, 0x27, 0x41, 0x86,
	0xc4, 0x51, 0xa2, 0x1e, 0xa5, 0xb8, 0xb1, 0x4c, 0x19, 0x3c, 0x0a, 0x93, 0x83, 0xae, 0x1c, 0x8f,
	0xf6, 0xa8, 0xf2, 0xb5, 0x50, 0x87, 0x95, 0x15, 0xbc, 0x79, 0x3f, 0xbd, 0xae, 0xe6, 0xb3, 0xfb,
	0x52, 0x2f, 0xd4, 0xc1, 0x2d, 0x72, 0x47, 0x0e, 0xf5, 0x09, 0xec, 0x59, 0x60, 0xc0, 0x2b, 0x7c,
	0x6e, 0x31, 0xcb, 0xf8, 0xf4, 0x77, 0x28, 0xa7, 0x18, 0x32, 0x1e, 0xff, 0x3b, 0x2f, 0x2e, 0x0a,
	0xfd, 0xa0, 0xd4, 0x1d, 0x79, 0xbc, 0x7a, 0x7c, 0x6a, 0xb3, 0xe8, 0x9f, 0x26, 0xbf, 0xae, 0xb7,
	0xfe, 0xb3, 0x3f, 0xa8, 0x6d, 0x60, 0xfa, 0x52, 0xa8, 0x99, 0x8b, 0x1e, 0x04, 0x83, 0xc3, 0x20,
	0xc3, 0x58, 0x4e, 0xe8, 0xa8, 0x26, 0x67, 0x85, 0x62, 0x00, 0x9b, 0x08, 0xba, 0x09, 0x08, 0xdb,
	0xba, 0x26, 0x07, 0xde, 0x06, 0xdb, 0x20, 0xec, 0xac, 0x66, 0xaa, 0x87, 0xd7, 0x23, 0xf7, 0x9b,
	0x58, 0xcc, 0x1b, 0x92, 0x75, 0xbb, 0x02, 0x17, 0xbd, 0xe9, 0xbd, 0x1e, 0xf3, 0xda, 0x7c, 0x73,
	0x9b, 0x1c, 0x5b, 0x76, 0x68, 0x9a, 0x68, 0x6c, 0x22, 0xd3, 0x4d, 0x4c, 0x5f, 0x59, 0x7e, 0x56,
	0xcc, 0x4e, 0x65, 0xf5, 0xff, 0xe3, 0x07, 0x2e, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xc3,
	0x51, 0x46, 0xc4, 0x01, 0x00, 0x00,
}
