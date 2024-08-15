// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.3
// source: configrelease.proto

package service_manage

import (
	fault_tolerance "github.com/polarismesh/specification/source/go/api/v1/fault_tolerance"
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

type ConfigRelease struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service        *Service                        `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Ctime          *wrapperspb.StringValue         `protobuf:"bytes,2,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime          *wrapperspb.StringValue         `protobuf:"bytes,3,opt,name=mtime,proto3" json:"mtime,omitempty"`
	CircuitBreaker *fault_tolerance.CircuitBreaker `protobuf:"bytes,4,opt,name=circuitBreaker,proto3" json:"circuitBreaker,omitempty"`
}

func (x *ConfigRelease) Reset() {
	*x = ConfigRelease{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configrelease_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRelease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRelease) ProtoMessage() {}

func (x *ConfigRelease) ProtoReflect() protoreflect.Message {
	mi := &file_configrelease_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRelease.ProtoReflect.Descriptor instead.
func (*ConfigRelease) Descriptor() ([]byte, []int) {
	return file_configrelease_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigRelease) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *ConfigRelease) GetCtime() *wrapperspb.StringValue {
	if x != nil {
		return x.Ctime
	}
	return nil
}

func (x *ConfigRelease) GetMtime() *wrapperspb.StringValue {
	if x != nil {
		return x.Mtime
	}
	return nil
}

func (x *ConfigRelease) GetCircuitBreaker() *fault_tolerance.CircuitBreaker {
	if x != nil {
		return x.CircuitBreaker
	}
	return nil
}

type ConfigWithService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services       []*Service                      `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	CircuitBreaker *fault_tolerance.CircuitBreaker `protobuf:"bytes,2,opt,name=circuitBreaker,proto3" json:"circuitBreaker,omitempty"`
}

func (x *ConfigWithService) Reset() {
	*x = ConfigWithService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configrelease_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigWithService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigWithService) ProtoMessage() {}

func (x *ConfigWithService) ProtoReflect() protoreflect.Message {
	mi := &file_configrelease_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigWithService.ProtoReflect.Descriptor instead.
func (*ConfigWithService) Descriptor() ([]byte, []int) {
	return file_configrelease_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigWithService) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *ConfigWithService) GetCircuitBreaker() *fault_tolerance.CircuitBreaker {
	if x != nil {
		return x.CircuitBreaker
	}
	return nil
}

var File_configrelease_proto protoreflect.FileDescriptor

var file_configrelease_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69,
	0x74, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda,
	0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x6d,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x3a, 0x0a, 0x0e, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x69, 0x72,
	0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x0e, 0x63, 0x69, 0x72,
	0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x22, 0x78, 0x0a, 0x11, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x57, 0x69, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x27, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0e, 0x63, 0x69, 0x72,
	0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72,
	0x65, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x0e, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72,
	0x65, 0x61, 0x6b, 0x65, 0x72, 0x42, 0x93, 0x01, 0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65,
	0x6e, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x42, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_configrelease_proto_rawDescOnce sync.Once
	file_configrelease_proto_rawDescData = file_configrelease_proto_rawDesc
)

func file_configrelease_proto_rawDescGZIP() []byte {
	file_configrelease_proto_rawDescOnce.Do(func() {
		file_configrelease_proto_rawDescData = protoimpl.X.CompressGZIP(file_configrelease_proto_rawDescData)
	})
	return file_configrelease_proto_rawDescData
}

var file_configrelease_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_configrelease_proto_goTypes = []interface{}{
	(*ConfigRelease)(nil),                  // 0: v1.ConfigRelease
	(*ConfigWithService)(nil),              // 1: v1.ConfigWithService
	(*Service)(nil),                        // 2: v1.Service
	(*wrapperspb.StringValue)(nil),         // 3: google.protobuf.StringValue
	(*fault_tolerance.CircuitBreaker)(nil), // 4: v1.CircuitBreaker
}
var file_configrelease_proto_depIdxs = []int32{
	2, // 0: v1.ConfigRelease.service:type_name -> v1.Service
	3, // 1: v1.ConfigRelease.ctime:type_name -> google.protobuf.StringValue
	3, // 2: v1.ConfigRelease.mtime:type_name -> google.protobuf.StringValue
	4, // 3: v1.ConfigRelease.circuitBreaker:type_name -> v1.CircuitBreaker
	2, // 4: v1.ConfigWithService.services:type_name -> v1.Service
	4, // 5: v1.ConfigWithService.circuitBreaker:type_name -> v1.CircuitBreaker
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_configrelease_proto_init() }
func file_configrelease_proto_init() {
	if File_configrelease_proto != nil {
		return
	}
	file_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_configrelease_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRelease); i {
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
		file_configrelease_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigWithService); i {
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
			RawDescriptor: file_configrelease_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_configrelease_proto_goTypes,
		DependencyIndexes: file_configrelease_proto_depIdxs,
		MessageInfos:      file_configrelease_proto_msgTypes,
	}.Build()
	File_configrelease_proto = out.File
	file_configrelease_proto_rawDesc = nil
	file_configrelease_proto_goTypes = nil
	file_configrelease_proto_depIdxs = nil
}
