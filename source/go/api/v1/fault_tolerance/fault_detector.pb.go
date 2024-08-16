// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.3
// source: fault_detector.proto

package fault_tolerance

import (
	model "github.com/polarismesh/specification/source/go/api/v1/model"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// detect protocol
type FaultDetectRule_Protocol int32

const (
	FaultDetectRule_UNKNOWN FaultDetectRule_Protocol = 0
	FaultDetectRule_HTTP    FaultDetectRule_Protocol = 1
	FaultDetectRule_TCP     FaultDetectRule_Protocol = 2
	FaultDetectRule_UDP     FaultDetectRule_Protocol = 3
)

// Enum value maps for FaultDetectRule_Protocol.
var (
	FaultDetectRule_Protocol_name = map[int32]string{
		0: "UNKNOWN",
		1: "HTTP",
		2: "TCP",
		3: "UDP",
	}
	FaultDetectRule_Protocol_value = map[string]int32{
		"UNKNOWN": 0,
		"HTTP":    1,
		"TCP":     2,
		"UDP":     3,
	}
)

func (x FaultDetectRule_Protocol) Enum() *FaultDetectRule_Protocol {
	p := new(FaultDetectRule_Protocol)
	*p = x
	return p
}

func (x FaultDetectRule_Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FaultDetectRule_Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_fault_detector_proto_enumTypes[0].Descriptor()
}

func (FaultDetectRule_Protocol) Type() protoreflect.EnumType {
	return &file_fault_detector_proto_enumTypes[0]
}

func (x FaultDetectRule_Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FaultDetectRule_Protocol.Descriptor instead.
func (FaultDetectRule_Protocol) EnumDescriptor() ([]byte, []int) {
	return file_fault_detector_proto_rawDescGZIP(), []int{1, 0}
}

type FaultDetector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// fault detect rules for current service
	Rules []*FaultDetectRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	// total revision for the fault detect rules
	Revision string `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *FaultDetector) Reset() {
	*x = FaultDetector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fault_detector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultDetector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultDetector) ProtoMessage() {}

func (x *FaultDetector) ProtoReflect() protoreflect.Message {
	mi := &file_fault_detector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultDetector.ProtoReflect.Descriptor instead.
func (*FaultDetector) Descriptor() ([]byte, []int) {
	return file_fault_detector_proto_rawDescGZIP(), []int{0}
}

func (x *FaultDetector) GetRules() []*FaultDetectRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *FaultDetector) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

type FaultDetectRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// rule name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// namespace of rule
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// revision routing version
	Revision string `protobuf:"bytes,4,opt,name=revision,proto3" json:"revision,omitempty"`
	// ctime create time of the rules
	Ctime string `protobuf:"bytes,5,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// mtime modify time of the rules
	Mtime string `protobuf:"bytes,6,opt,name=mtime,proto3" json:"mtime,omitempty"`
	// description simple description rules
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	// detect target
	TargetService *FaultDetectRule_DestinationService `protobuf:"bytes,21,opt,name=target_service,json=targetService,proto3" json:"target_service,omitempty"`
	// detect interval
	Interval uint32 `protobuf:"varint,22,opt,name=interval,proto3" json:"interval,omitempty"`
	// detect timeout
	Timeout uint32 `protobuf:"varint,23,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// detect port
	Port     uint32                   `protobuf:"varint,24,opt,name=port,proto3" json:"port,omitempty"`
	Protocol FaultDetectRule_Protocol `protobuf:"varint,25,opt,name=protocol,proto3,enum=v1.FaultDetectRule_Protocol" json:"protocol,omitempty"`
	// http detect config
	HttpConfig *HttpProtocolConfig `protobuf:"bytes,26,opt,name=http_config,json=httpConfig,proto3" json:"http_config,omitempty"`
	// tcp detect config
	TcpConfig *TcpProtocolConfig `protobuf:"bytes,27,opt,name=tcp_config,json=tcpConfig,proto3" json:"tcp_config,omitempty"`
	// udp detect config
	UdpConfig *UdpProtocolConfig `protobuf:"bytes,28,opt,name=udp_config,json=udpConfig,proto3" json:"udp_config,omitempty"`
	// priority rules priority
	Priority uint32 `protobuf:"varint,29,opt,name=priority,proto3" json:"priority,omitempty"`
	// 探测规则标签数据
	Metadata map[string]string `protobuf:"bytes,30,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FaultDetectRule) Reset() {
	*x = FaultDetectRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fault_detector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultDetectRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultDetectRule) ProtoMessage() {}

func (x *FaultDetectRule) ProtoReflect() protoreflect.Message {
	mi := &file_fault_detector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultDetectRule.ProtoReflect.Descriptor instead.
func (*FaultDetectRule) Descriptor() ([]byte, []int) {
	return file_fault_detector_proto_rawDescGZIP(), []int{1}
}

func (x *FaultDetectRule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FaultDetectRule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FaultDetectRule) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *FaultDetectRule) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *FaultDetectRule) GetCtime() string {
	if x != nil {
		return x.Ctime
	}
	return ""
}

func (x *FaultDetectRule) GetMtime() string {
	if x != nil {
		return x.Mtime
	}
	return ""
}

func (x *FaultDetectRule) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FaultDetectRule) GetTargetService() *FaultDetectRule_DestinationService {
	if x != nil {
		return x.TargetService
	}
	return nil
}

func (x *FaultDetectRule) GetInterval() uint32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *FaultDetectRule) GetTimeout() uint32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *FaultDetectRule) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *FaultDetectRule) GetProtocol() FaultDetectRule_Protocol {
	if x != nil {
		return x.Protocol
	}
	return FaultDetectRule_UNKNOWN
}

func (x *FaultDetectRule) GetHttpConfig() *HttpProtocolConfig {
	if x != nil {
		return x.HttpConfig
	}
	return nil
}

func (x *FaultDetectRule) GetTcpConfig() *TcpProtocolConfig {
	if x != nil {
		return x.TcpConfig
	}
	return nil
}

func (x *FaultDetectRule) GetUdpConfig() *UdpProtocolConfig {
	if x != nil {
		return x.UdpConfig
	}
	return nil
}

func (x *FaultDetectRule) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *FaultDetectRule) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type HttpProtocolConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method  string                              `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Url     string                              `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Headers []*HttpProtocolConfig_MessageHeader `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
	Body    string                              `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *HttpProtocolConfig) Reset() {
	*x = HttpProtocolConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fault_detector_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpProtocolConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpProtocolConfig) ProtoMessage() {}

func (x *HttpProtocolConfig) ProtoReflect() protoreflect.Message {
	mi := &file_fault_detector_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpProtocolConfig.ProtoReflect.Descriptor instead.
func (*HttpProtocolConfig) Descriptor() ([]byte, []int) {
	return file_fault_detector_proto_rawDescGZIP(), []int{2}
}

func (x *HttpProtocolConfig) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *HttpProtocolConfig) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HttpProtocolConfig) GetHeaders() []*HttpProtocolConfig_MessageHeader {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *HttpProtocolConfig) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type TcpProtocolConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Send    string   `protobuf:"bytes,1,opt,name=send,proto3" json:"send,omitempty"`
	Receive []string `protobuf:"bytes,2,rep,name=receive,proto3" json:"receive,omitempty"`
}

func (x *TcpProtocolConfig) Reset() {
	*x = TcpProtocolConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fault_detector_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TcpProtocolConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TcpProtocolConfig) ProtoMessage() {}

func (x *TcpProtocolConfig) ProtoReflect() protoreflect.Message {
	mi := &file_fault_detector_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TcpProtocolConfig.ProtoReflect.Descriptor instead.
func (*TcpProtocolConfig) Descriptor() ([]byte, []int) {
	return file_fault_detector_proto_rawDescGZIP(), []int{3}
}

func (x *TcpProtocolConfig) GetSend() string {
	if x != nil {
		return x.Send
	}
	return ""
}

func (x *TcpProtocolConfig) GetReceive() []string {
	if x != nil {
		return x.Receive
	}
	return nil
}

type UdpProtocolConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Send    string   `protobuf:"bytes,1,opt,name=send,proto3" json:"send,omitempty"`
	Receive []string `protobuf:"bytes,2,rep,name=receive,proto3" json:"receive,omitempty"`
}

func (x *UdpProtocolConfig) Reset() {
	*x = UdpProtocolConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fault_detector_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UdpProtocolConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UdpProtocolConfig) ProtoMessage() {}

func (x *UdpProtocolConfig) ProtoReflect() protoreflect.Message {
	mi := &file_fault_detector_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UdpProtocolConfig.ProtoReflect.Descriptor instead.
func (*UdpProtocolConfig) Descriptor() ([]byte, []int) {
	return file_fault_detector_proto_rawDescGZIP(), []int{4}
}

func (x *UdpProtocolConfig) GetSend() string {
	if x != nil {
		return x.Send
	}
	return ""
}

func (x *UdpProtocolConfig) GetReceive() []string {
	if x != nil {
		return x.Receive
	}
	return nil
}

type FaultDetectRule_DestinationService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service   string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// deprecated_filed  use api.path instead
	//
	// Deprecated: Marked as deprecated in fault_detector.proto.
	Method *model.MatchString `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Api    *model.API         `protobuf:"bytes,4,opt,name=api,proto3" json:"api,omitempty"`
}

func (x *FaultDetectRule_DestinationService) Reset() {
	*x = FaultDetectRule_DestinationService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fault_detector_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultDetectRule_DestinationService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultDetectRule_DestinationService) ProtoMessage() {}

func (x *FaultDetectRule_DestinationService) ProtoReflect() protoreflect.Message {
	mi := &file_fault_detector_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultDetectRule_DestinationService.ProtoReflect.Descriptor instead.
func (*FaultDetectRule_DestinationService) Descriptor() ([]byte, []int) {
	return file_fault_detector_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FaultDetectRule_DestinationService) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *FaultDetectRule_DestinationService) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// Deprecated: Marked as deprecated in fault_detector.proto.
func (x *FaultDetectRule_DestinationService) GetMethod() *model.MatchString {
	if x != nil {
		return x.Method
	}
	return nil
}

func (x *FaultDetectRule_DestinationService) GetApi() *model.API {
	if x != nil {
		return x.Api
	}
	return nil
}

type HttpProtocolConfig_MessageHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *HttpProtocolConfig_MessageHeader) Reset() {
	*x = HttpProtocolConfig_MessageHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fault_detector_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpProtocolConfig_MessageHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpProtocolConfig_MessageHeader) ProtoMessage() {}

func (x *HttpProtocolConfig_MessageHeader) ProtoReflect() protoreflect.Message {
	mi := &file_fault_detector_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpProtocolConfig_MessageHeader.ProtoReflect.Descriptor instead.
func (*HttpProtocolConfig_MessageHeader) Descriptor() ([]byte, []int) {
	return file_fault_detector_proto_rawDescGZIP(), []int{2, 0}
}

func (x *HttpProtocolConfig_MessageHeader) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *HttpProtocolConfig_MessageHeader) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_fault_detector_proto protoreflect.FileDescriptor

var file_fault_detector_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0d, 0x46, 0x61, 0x75, 0x6c, 0x74,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x75,
	0x6c, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x9f, 0x07, 0x0a, 0x0f, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x4d, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x75,
	0x6c, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x12, 0x37, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x74, 0x74,
	0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0a, 0x68, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34, 0x0a, 0x0a, 0x74,
	0x63, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x63, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x74, 0x63, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x34, 0x0a, 0x0a, 0x75, 0x64, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x1c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x64, 0x70, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x75, 0x64,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x1e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x94, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x2b, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x19,
	0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x50, 0x49, 0x52, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x33, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x43, 0x50,
	0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x44, 0x50, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x08, 0x10,
	0x15, 0x22, 0xcb, 0x01, 0x0a, 0x12, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x3e, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x37, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x41, 0x0a, 0x11, 0x54, 0x63, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x22, 0x41, 0x0a, 0x11, 0x55, 0x64, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x42, 0x95, 0x01, 0x0a, 0x38, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65,
	0x6e, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x74, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x6e,
	0x63, 0x65, 0x42, 0x12, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x74, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fault_detector_proto_rawDescOnce sync.Once
	file_fault_detector_proto_rawDescData = file_fault_detector_proto_rawDesc
)

func file_fault_detector_proto_rawDescGZIP() []byte {
	file_fault_detector_proto_rawDescOnce.Do(func() {
		file_fault_detector_proto_rawDescData = protoimpl.X.CompressGZIP(file_fault_detector_proto_rawDescData)
	})
	return file_fault_detector_proto_rawDescData
}

var file_fault_detector_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_fault_detector_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_fault_detector_proto_goTypes = []interface{}{
	(FaultDetectRule_Protocol)(0),              // 0: v1.FaultDetectRule.Protocol
	(*FaultDetector)(nil),                      // 1: v1.FaultDetector
	(*FaultDetectRule)(nil),                    // 2: v1.FaultDetectRule
	(*HttpProtocolConfig)(nil),                 // 3: v1.HttpProtocolConfig
	(*TcpProtocolConfig)(nil),                  // 4: v1.TcpProtocolConfig
	(*UdpProtocolConfig)(nil),                  // 5: v1.UdpProtocolConfig
	(*FaultDetectRule_DestinationService)(nil), // 6: v1.FaultDetectRule.DestinationService
	nil,                                      // 7: v1.FaultDetectRule.MetadataEntry
	(*HttpProtocolConfig_MessageHeader)(nil), // 8: v1.HttpProtocolConfig.MessageHeader
	(*model.MatchString)(nil),                // 9: v1.MatchString
	(*model.API)(nil),                        // 10: v1.API
}
var file_fault_detector_proto_depIdxs = []int32{
	2,  // 0: v1.FaultDetector.rules:type_name -> v1.FaultDetectRule
	6,  // 1: v1.FaultDetectRule.target_service:type_name -> v1.FaultDetectRule.DestinationService
	0,  // 2: v1.FaultDetectRule.protocol:type_name -> v1.FaultDetectRule.Protocol
	3,  // 3: v1.FaultDetectRule.http_config:type_name -> v1.HttpProtocolConfig
	4,  // 4: v1.FaultDetectRule.tcp_config:type_name -> v1.TcpProtocolConfig
	5,  // 5: v1.FaultDetectRule.udp_config:type_name -> v1.UdpProtocolConfig
	7,  // 6: v1.FaultDetectRule.metadata:type_name -> v1.FaultDetectRule.MetadataEntry
	8,  // 7: v1.HttpProtocolConfig.headers:type_name -> v1.HttpProtocolConfig.MessageHeader
	9,  // 8: v1.FaultDetectRule.DestinationService.method:type_name -> v1.MatchString
	10, // 9: v1.FaultDetectRule.DestinationService.api:type_name -> v1.API
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_fault_detector_proto_init() }
func file_fault_detector_proto_init() {
	if File_fault_detector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fault_detector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultDetector); i {
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
		file_fault_detector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultDetectRule); i {
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
		file_fault_detector_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpProtocolConfig); i {
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
		file_fault_detector_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TcpProtocolConfig); i {
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
		file_fault_detector_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UdpProtocolConfig); i {
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
		file_fault_detector_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultDetectRule_DestinationService); i {
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
		file_fault_detector_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpProtocolConfig_MessageHeader); i {
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
			RawDescriptor: file_fault_detector_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fault_detector_proto_goTypes,
		DependencyIndexes: file_fault_detector_proto_depIdxs,
		EnumInfos:         file_fault_detector_proto_enumTypes,
		MessageInfos:      file_fault_detector_proto_msgTypes,
	}.Build()
	File_fault_detector_proto = out.File
	file_fault_detector_proto_rawDesc = nil
	file_fault_detector_proto_goTypes = nil
	file_fault_detector_proto_depIdxs = nil
}
