// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.3
// source: grpcapi.proto

package service_manage

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_grpcapi_proto protoreflect.FileDescriptor

var file_grpcapi_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x1a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x87, 0x02, 0x0a, 0x0b, 0x50, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x47, 0x52, 0x50, 0x43,
	0x12, 0x2a, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x0c, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x10,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x0c,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x12, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x1a, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x13,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x29, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x0c, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x0c, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xf5, 0x01, 0x0a, 0x14, 0x50,
	0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x47,
	0x52, 0x50, 0x43, 0x12, 0x45, 0x0a, 0x0e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x4a, 0x0a, 0x11, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12,
	0x18, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x11, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44,
	0x65, 0x6c, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x18, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x32, 0x95, 0x01, 0x0a, 0x1a, 0x50, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x52, 0x50,
	0x43, 0x12, 0x3c, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x1a,
	0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x1a, 0x0c, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x93, 0x01, 0x0a, 0x37, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x61, 0x72,
	0x69, 0x73, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x42, 0x12, 0x50, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x47,
	0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_grpcapi_proto_goTypes = []interface{}{
	(*Client)(nil),                // 0: v1.Client
	(*Instance)(nil),              // 1: v1.Instance
	(*DiscoverRequest)(nil),       // 2: v1.DiscoverRequest
	(*HeartbeatsRequest)(nil),     // 3: v1.HeartbeatsRequest
	(*GetHeartbeatsRequest)(nil),  // 4: v1.GetHeartbeatsRequest
	(*DelHeartbeatsRequest)(nil),  // 5: v1.DelHeartbeatsRequest
	(*ServiceContract)(nil),       // 6: v1.ServiceContract
	(*Response)(nil),              // 7: v1.Response
	(*DiscoverResponse)(nil),      // 8: v1.DiscoverResponse
	(*HeartbeatsResponse)(nil),    // 9: v1.HeartbeatsResponse
	(*GetHeartbeatsResponse)(nil), // 10: v1.GetHeartbeatsResponse
	(*DelHeartbeatsResponse)(nil), // 11: v1.DelHeartbeatsResponse
}
var file_grpcapi_proto_depIdxs = []int32{
	0,  // 0: v1.PolarisGRPC.ReportClient:input_type -> v1.Client
	1,  // 1: v1.PolarisGRPC.RegisterInstance:input_type -> v1.Instance
	1,  // 2: v1.PolarisGRPC.DeregisterInstance:input_type -> v1.Instance
	2,  // 3: v1.PolarisGRPC.Discover:input_type -> v1.DiscoverRequest
	1,  // 4: v1.PolarisGRPC.Heartbeat:input_type -> v1.Instance
	3,  // 5: v1.PolarisHeartbeatGRPC.BatchHeartbeat:input_type -> v1.HeartbeatsRequest
	4,  // 6: v1.PolarisHeartbeatGRPC.BatchGetHeartbeat:input_type -> v1.GetHeartbeatsRequest
	5,  // 7: v1.PolarisHeartbeatGRPC.BatchDelHeartbeat:input_type -> v1.DelHeartbeatsRequest
	6,  // 8: v1.PolarisServiceContractGRPC.ReportServiceContract:input_type -> v1.ServiceContract
	6,  // 9: v1.PolarisServiceContractGRPC.GetServiceContract:input_type -> v1.ServiceContract
	7,  // 10: v1.PolarisGRPC.ReportClient:output_type -> v1.Response
	7,  // 11: v1.PolarisGRPC.RegisterInstance:output_type -> v1.Response
	7,  // 12: v1.PolarisGRPC.DeregisterInstance:output_type -> v1.Response
	8,  // 13: v1.PolarisGRPC.Discover:output_type -> v1.DiscoverResponse
	7,  // 14: v1.PolarisGRPC.Heartbeat:output_type -> v1.Response
	9,  // 15: v1.PolarisHeartbeatGRPC.BatchHeartbeat:output_type -> v1.HeartbeatsResponse
	10, // 16: v1.PolarisHeartbeatGRPC.BatchGetHeartbeat:output_type -> v1.GetHeartbeatsResponse
	11, // 17: v1.PolarisHeartbeatGRPC.BatchDelHeartbeat:output_type -> v1.DelHeartbeatsResponse
	7,  // 18: v1.PolarisServiceContractGRPC.ReportServiceContract:output_type -> v1.Response
	7,  // 19: v1.PolarisServiceContractGRPC.GetServiceContract:output_type -> v1.Response
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_grpcapi_proto_init() }
func file_grpcapi_proto_init() {
	if File_grpcapi_proto != nil {
		return
	}
	file_client_proto_init()
	file_service_proto_init()
	file_request_proto_init()
	file_response_proto_init()
	file_heartbeat_proto_init()
	file_contract_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpcapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_grpcapi_proto_goTypes,
		DependencyIndexes: file_grpcapi_proto_depIdxs,
	}.Build()
	File_grpcapi_proto = out.File
	file_grpcapi_proto_rawDesc = nil
	file_grpcapi_proto_goTypes = nil
	file_grpcapi_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PolarisGRPCClient is the client API for PolarisGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PolarisGRPCClient interface {
	// 客户端上报
	ReportClient(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Response, error)
	// 被调方注册服务实例
	RegisterInstance(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error)
	// 被调方反注册服务实例
	DeregisterInstance(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error)
	// 统一发现接口
	Discover(ctx context.Context, opts ...grpc.CallOption) (PolarisGRPC_DiscoverClient, error)
	// 被调方上报心跳
	Heartbeat(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error)
}

type polarisGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewPolarisGRPCClient(cc grpc.ClientConnInterface) PolarisGRPCClient {
	return &polarisGRPCClient{cc}
}

func (c *polarisGRPCClient) ReportClient(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.PolarisGRPC/ReportClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisGRPCClient) RegisterInstance(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.PolarisGRPC/RegisterInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisGRPCClient) DeregisterInstance(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.PolarisGRPC/DeregisterInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisGRPCClient) Discover(ctx context.Context, opts ...grpc.CallOption) (PolarisGRPC_DiscoverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PolarisGRPC_serviceDesc.Streams[0], "/v1.PolarisGRPC/Discover", opts...)
	if err != nil {
		return nil, err
	}
	x := &polarisGRPCDiscoverClient{stream}
	return x, nil
}

type PolarisGRPC_DiscoverClient interface {
	Send(*DiscoverRequest) error
	Recv() (*DiscoverResponse, error)
	grpc.ClientStream
}

type polarisGRPCDiscoverClient struct {
	grpc.ClientStream
}

func (x *polarisGRPCDiscoverClient) Send(m *DiscoverRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *polarisGRPCDiscoverClient) Recv() (*DiscoverResponse, error) {
	m := new(DiscoverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *polarisGRPCClient) Heartbeat(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.PolarisGRPC/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolarisGRPCServer is the server API for PolarisGRPC service.
type PolarisGRPCServer interface {
	// 客户端上报
	ReportClient(context.Context, *Client) (*Response, error)
	// 被调方注册服务实例
	RegisterInstance(context.Context, *Instance) (*Response, error)
	// 被调方反注册服务实例
	DeregisterInstance(context.Context, *Instance) (*Response, error)
	// 统一发现接口
	Discover(PolarisGRPC_DiscoverServer) error
	// 被调方上报心跳
	Heartbeat(context.Context, *Instance) (*Response, error)
}

// UnimplementedPolarisGRPCServer can be embedded to have forward compatible implementations.
type UnimplementedPolarisGRPCServer struct {
}

func (*UnimplementedPolarisGRPCServer) ReportClient(context.Context, *Client) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportClient not implemented")
}
func (*UnimplementedPolarisGRPCServer) RegisterInstance(context.Context, *Instance) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterInstance not implemented")
}
func (*UnimplementedPolarisGRPCServer) DeregisterInstance(context.Context, *Instance) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterInstance not implemented")
}
func (*UnimplementedPolarisGRPCServer) Discover(PolarisGRPC_DiscoverServer) error {
	return status.Errorf(codes.Unimplemented, "method Discover not implemented")
}
func (*UnimplementedPolarisGRPCServer) Heartbeat(context.Context, *Instance) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}

func RegisterPolarisGRPCServer(s *grpc.Server, srv PolarisGRPCServer) {
	s.RegisterService(&_PolarisGRPC_serviceDesc, srv)
}

func _PolarisGRPC_ReportClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Client)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisGRPCServer).ReportClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisGRPC/ReportClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisGRPCServer).ReportClient(ctx, req.(*Client))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisGRPC_RegisterInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Instance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisGRPCServer).RegisterInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisGRPC/RegisterInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisGRPCServer).RegisterInstance(ctx, req.(*Instance))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisGRPC_DeregisterInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Instance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisGRPCServer).DeregisterInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisGRPC/DeregisterInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisGRPCServer).DeregisterInstance(ctx, req.(*Instance))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisGRPC_Discover_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PolarisGRPCServer).Discover(&polarisGRPCDiscoverServer{stream})
}

type PolarisGRPC_DiscoverServer interface {
	Send(*DiscoverResponse) error
	Recv() (*DiscoverRequest, error)
	grpc.ServerStream
}

type polarisGRPCDiscoverServer struct {
	grpc.ServerStream
}

func (x *polarisGRPCDiscoverServer) Send(m *DiscoverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *polarisGRPCDiscoverServer) Recv() (*DiscoverRequest, error) {
	m := new(DiscoverRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PolarisGRPC_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Instance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisGRPCServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisGRPC/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisGRPCServer).Heartbeat(ctx, req.(*Instance))
	}
	return interceptor(ctx, in, info, handler)
}

var _PolarisGRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PolarisGRPC",
	HandlerType: (*PolarisGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportClient",
			Handler:    _PolarisGRPC_ReportClient_Handler,
		},
		{
			MethodName: "RegisterInstance",
			Handler:    _PolarisGRPC_RegisterInstance_Handler,
		},
		{
			MethodName: "DeregisterInstance",
			Handler:    _PolarisGRPC_DeregisterInstance_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _PolarisGRPC_Heartbeat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Discover",
			Handler:       _PolarisGRPC_Discover_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpcapi.proto",
}

// PolarisHeartbeatGRPCClient is the client API for PolarisHeartbeatGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PolarisHeartbeatGRPCClient interface {
	// 被调方批量上报心跳
	BatchHeartbeat(ctx context.Context, opts ...grpc.CallOption) (PolarisHeartbeatGRPC_BatchHeartbeatClient, error)
	// 批量获取心跳记录
	BatchGetHeartbeat(ctx context.Context, in *GetHeartbeatsRequest, opts ...grpc.CallOption) (*GetHeartbeatsResponse, error)
	// 批量删除心跳记录
	BatchDelHeartbeat(ctx context.Context, in *DelHeartbeatsRequest, opts ...grpc.CallOption) (*DelHeartbeatsResponse, error)
}

type polarisHeartbeatGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewPolarisHeartbeatGRPCClient(cc grpc.ClientConnInterface) PolarisHeartbeatGRPCClient {
	return &polarisHeartbeatGRPCClient{cc}
}

func (c *polarisHeartbeatGRPCClient) BatchHeartbeat(ctx context.Context, opts ...grpc.CallOption) (PolarisHeartbeatGRPC_BatchHeartbeatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PolarisHeartbeatGRPC_serviceDesc.Streams[0], "/v1.PolarisHeartbeatGRPC/BatchHeartbeat", opts...)
	if err != nil {
		return nil, err
	}
	x := &polarisHeartbeatGRPCBatchHeartbeatClient{stream}
	return x, nil
}

type PolarisHeartbeatGRPC_BatchHeartbeatClient interface {
	Send(*HeartbeatsRequest) error
	Recv() (*HeartbeatsResponse, error)
	grpc.ClientStream
}

type polarisHeartbeatGRPCBatchHeartbeatClient struct {
	grpc.ClientStream
}

func (x *polarisHeartbeatGRPCBatchHeartbeatClient) Send(m *HeartbeatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *polarisHeartbeatGRPCBatchHeartbeatClient) Recv() (*HeartbeatsResponse, error) {
	m := new(HeartbeatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *polarisHeartbeatGRPCClient) BatchGetHeartbeat(ctx context.Context, in *GetHeartbeatsRequest, opts ...grpc.CallOption) (*GetHeartbeatsResponse, error) {
	out := new(GetHeartbeatsResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisHeartbeatGRPC/BatchGetHeartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisHeartbeatGRPCClient) BatchDelHeartbeat(ctx context.Context, in *DelHeartbeatsRequest, opts ...grpc.CallOption) (*DelHeartbeatsResponse, error) {
	out := new(DelHeartbeatsResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisHeartbeatGRPC/BatchDelHeartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolarisHeartbeatGRPCServer is the server API for PolarisHeartbeatGRPC service.
type PolarisHeartbeatGRPCServer interface {
	// 被调方批量上报心跳
	BatchHeartbeat(PolarisHeartbeatGRPC_BatchHeartbeatServer) error
	// 批量获取心跳记录
	BatchGetHeartbeat(context.Context, *GetHeartbeatsRequest) (*GetHeartbeatsResponse, error)
	// 批量删除心跳记录
	BatchDelHeartbeat(context.Context, *DelHeartbeatsRequest) (*DelHeartbeatsResponse, error)
}

// UnimplementedPolarisHeartbeatGRPCServer can be embedded to have forward compatible implementations.
type UnimplementedPolarisHeartbeatGRPCServer struct {
}

func (*UnimplementedPolarisHeartbeatGRPCServer) BatchHeartbeat(PolarisHeartbeatGRPC_BatchHeartbeatServer) error {
	return status.Errorf(codes.Unimplemented, "method BatchHeartbeat not implemented")
}
func (*UnimplementedPolarisHeartbeatGRPCServer) BatchGetHeartbeat(context.Context, *GetHeartbeatsRequest) (*GetHeartbeatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetHeartbeat not implemented")
}
func (*UnimplementedPolarisHeartbeatGRPCServer) BatchDelHeartbeat(context.Context, *DelHeartbeatsRequest) (*DelHeartbeatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDelHeartbeat not implemented")
}

func RegisterPolarisHeartbeatGRPCServer(s *grpc.Server, srv PolarisHeartbeatGRPCServer) {
	s.RegisterService(&_PolarisHeartbeatGRPC_serviceDesc, srv)
}

func _PolarisHeartbeatGRPC_BatchHeartbeat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PolarisHeartbeatGRPCServer).BatchHeartbeat(&polarisHeartbeatGRPCBatchHeartbeatServer{stream})
}

type PolarisHeartbeatGRPC_BatchHeartbeatServer interface {
	Send(*HeartbeatsResponse) error
	Recv() (*HeartbeatsRequest, error)
	grpc.ServerStream
}

type polarisHeartbeatGRPCBatchHeartbeatServer struct {
	grpc.ServerStream
}

func (x *polarisHeartbeatGRPCBatchHeartbeatServer) Send(m *HeartbeatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *polarisHeartbeatGRPCBatchHeartbeatServer) Recv() (*HeartbeatsRequest, error) {
	m := new(HeartbeatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PolarisHeartbeatGRPC_BatchGetHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeartbeatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisHeartbeatGRPCServer).BatchGetHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisHeartbeatGRPC/BatchGetHeartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisHeartbeatGRPCServer).BatchGetHeartbeat(ctx, req.(*GetHeartbeatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisHeartbeatGRPC_BatchDelHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelHeartbeatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisHeartbeatGRPCServer).BatchDelHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisHeartbeatGRPC/BatchDelHeartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisHeartbeatGRPCServer).BatchDelHeartbeat(ctx, req.(*DelHeartbeatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PolarisHeartbeatGRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PolarisHeartbeatGRPC",
	HandlerType: (*PolarisHeartbeatGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BatchGetHeartbeat",
			Handler:    _PolarisHeartbeatGRPC_BatchGetHeartbeat_Handler,
		},
		{
			MethodName: "BatchDelHeartbeat",
			Handler:    _PolarisHeartbeatGRPC_BatchDelHeartbeat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BatchHeartbeat",
			Handler:       _PolarisHeartbeatGRPC_BatchHeartbeat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpcapi.proto",
}

// PolarisServiceContractGRPCClient is the client API for PolarisServiceContractGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PolarisServiceContractGRPCClient interface {
	// 上报服务契约
	ReportServiceContract(ctx context.Context, in *ServiceContract, opts ...grpc.CallOption) (*Response, error)
	// 查询服务契约
	GetServiceContract(ctx context.Context, in *ServiceContract, opts ...grpc.CallOption) (*Response, error)
}

type polarisServiceContractGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewPolarisServiceContractGRPCClient(cc grpc.ClientConnInterface) PolarisServiceContractGRPCClient {
	return &polarisServiceContractGRPCClient{cc}
}

func (c *polarisServiceContractGRPCClient) ReportServiceContract(ctx context.Context, in *ServiceContract, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.PolarisServiceContractGRPC/ReportServiceContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisServiceContractGRPCClient) GetServiceContract(ctx context.Context, in *ServiceContract, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.PolarisServiceContractGRPC/GetServiceContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolarisServiceContractGRPCServer is the server API for PolarisServiceContractGRPC service.
type PolarisServiceContractGRPCServer interface {
	// 上报服务契约
	ReportServiceContract(context.Context, *ServiceContract) (*Response, error)
	// 查询服务契约
	GetServiceContract(context.Context, *ServiceContract) (*Response, error)
}

// UnimplementedPolarisServiceContractGRPCServer can be embedded to have forward compatible implementations.
type UnimplementedPolarisServiceContractGRPCServer struct {
}

func (*UnimplementedPolarisServiceContractGRPCServer) ReportServiceContract(context.Context, *ServiceContract) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportServiceContract not implemented")
}
func (*UnimplementedPolarisServiceContractGRPCServer) GetServiceContract(context.Context, *ServiceContract) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceContract not implemented")
}

func RegisterPolarisServiceContractGRPCServer(s *grpc.Server, srv PolarisServiceContractGRPCServer) {
	s.RegisterService(&_PolarisServiceContractGRPC_serviceDesc, srv)
}

func _PolarisServiceContractGRPC_ReportServiceContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisServiceContractGRPCServer).ReportServiceContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisServiceContractGRPC/ReportServiceContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisServiceContractGRPCServer).ReportServiceContract(ctx, req.(*ServiceContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisServiceContractGRPC_GetServiceContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisServiceContractGRPCServer).GetServiceContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisServiceContractGRPC/GetServiceContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisServiceContractGRPCServer).GetServiceContract(ctx, req.(*ServiceContract))
	}
	return interceptor(ctx, in, info, handler)
}

var _PolarisServiceContractGRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PolarisServiceContractGRPC",
	HandlerType: (*PolarisServiceContractGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportServiceContract",
			Handler:    _PolarisServiceContractGRPC_ReportServiceContract_Handler,
		},
		{
			MethodName: "GetServiceContract",
			Handler:    _PolarisServiceContractGRPC_GetServiceContract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcapi.proto",
}
