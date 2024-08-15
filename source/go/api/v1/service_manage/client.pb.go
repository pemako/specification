// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.3
// source: client.proto

package service_manage

import (
	model "github.com/polarismesh/specification/source/go/api/v1/model"
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

type Client_ClientType int32

const (
	Client_UNKNOWN Client_ClientType = 0
	Client_SDK     Client_ClientType = 1
	Client_AGENT   Client_ClientType = 2
)

// Enum value maps for Client_ClientType.
var (
	Client_ClientType_name = map[int32]string{
		0: "UNKNOWN",
		1: "SDK",
		2: "AGENT",
	}
	Client_ClientType_value = map[string]int32{
		"UNKNOWN": 0,
		"SDK":     1,
		"AGENT":   2,
	}
)

func (x Client_ClientType) Enum() *Client_ClientType {
	p := new(Client_ClientType)
	*p = x
	return p
}

func (x Client_ClientType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Client_ClientType) Descriptor() protoreflect.EnumDescriptor {
	return file_client_proto_enumTypes[0].Descriptor()
}

func (Client_ClientType) Type() protoreflect.EnumType {
	return &file_client_proto_enumTypes[0]
}

func (x Client_ClientType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Client_ClientType.Descriptor instead.
func (Client_ClientType) EnumDescriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0, 0}
}

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Type     Client_ClientType       `protobuf:"varint,2,opt,name=type,proto3,enum=v1.Client_ClientType" json:"type,omitempty"`
	Version  *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Location *model.Location         `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Id       *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	Stat     []*StatInfo             `protobuf:"bytes,6,rep,name=stat,proto3" json:"stat,omitempty"`
	Ctime    *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime    *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=mtime,proto3" json:"mtime,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

func (x *Client) GetHost() *wrapperspb.StringValue {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *Client) GetType() Client_ClientType {
	if x != nil {
		return x.Type
	}
	return Client_UNKNOWN
}

func (x *Client) GetVersion() *wrapperspb.StringValue {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *Client) GetLocation() *model.Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Client) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Client) GetStat() []*StatInfo {
	if x != nil {
		return x.Stat
	}
	return nil
}

func (x *Client) GetCtime() *wrapperspb.StringValue {
	if x != nil {
		return x.Ctime
	}
	return nil
}

func (x *Client) GetMtime() *wrapperspb.StringValue {
	if x != nil {
		return x.Mtime
	}
	return nil
}

type StatInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target   *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Port     *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Path     *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Protocol *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *StatInfo) Reset() {
	*x = StatInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatInfo) ProtoMessage() {}

func (x *StatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatInfo.ProtoReflect.Descriptor instead.
func (*StatInfo) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{1}
}

func (x *StatInfo) GetTarget() *wrapperspb.StringValue {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *StatInfo) GetPort() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *StatInfo) GetPath() *wrapperspb.StringValue {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *StatInfo) GetProtocol() *wrapperspb.StringValue {
	if x != nil {
		return x.Protocol
	}
	return nil
}

var File_client_proto protoreflect.FileDescriptor

var file_client_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xae, 0x03, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x28, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x63, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a,
	0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6d, 0x74, 0x69, 0x6d,
	0x65, 0x22, 0x2d, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x53, 0x44, 0x4b, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10, 0x02,
	0x22, 0xde, 0x01, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x42, 0x8c, 0x01, 0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x42, 0x0b, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_rawDescOnce sync.Once
	file_client_proto_rawDescData = file_client_proto_rawDesc
)

func file_client_proto_rawDescGZIP() []byte {
	file_client_proto_rawDescOnce.Do(func() {
		file_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_rawDescData)
	})
	return file_client_proto_rawDescData
}

var file_client_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_client_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_proto_goTypes = []interface{}{
	(Client_ClientType)(0),         // 0: v1.Client.ClientType
	(*Client)(nil),                 // 1: v1.Client
	(*StatInfo)(nil),               // 2: v1.StatInfo
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
	(*model.Location)(nil),         // 4: v1.Location
	(*wrapperspb.UInt32Value)(nil), // 5: google.protobuf.UInt32Value
}
var file_client_proto_depIdxs = []int32{
	3,  // 0: v1.Client.host:type_name -> google.protobuf.StringValue
	0,  // 1: v1.Client.type:type_name -> v1.Client.ClientType
	3,  // 2: v1.Client.version:type_name -> google.protobuf.StringValue
	4,  // 3: v1.Client.location:type_name -> v1.Location
	3,  // 4: v1.Client.id:type_name -> google.protobuf.StringValue
	2,  // 5: v1.Client.stat:type_name -> v1.StatInfo
	3,  // 6: v1.Client.ctime:type_name -> google.protobuf.StringValue
	3,  // 7: v1.Client.mtime:type_name -> google.protobuf.StringValue
	3,  // 8: v1.StatInfo.target:type_name -> google.protobuf.StringValue
	5,  // 9: v1.StatInfo.port:type_name -> google.protobuf.UInt32Value
	3,  // 10: v1.StatInfo.path:type_name -> google.protobuf.StringValue
	3,  // 11: v1.StatInfo.protocol:type_name -> google.protobuf.StringValue
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_client_proto_init() }
func file_client_proto_init() {
	if File_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatInfo); i {
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
			RawDescriptor: file_client_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_proto_goTypes,
		DependencyIndexes: file_client_proto_depIdxs,
		EnumInfos:         file_client_proto_enumTypes,
		MessageInfos:      file_client_proto_msgTypes,
	}.Build()
	File_client_proto = out.File
	file_client_proto_rawDesc = nil
	file_client_proto_goTypes = nil
	file_client_proto_depIdxs = nil
}
