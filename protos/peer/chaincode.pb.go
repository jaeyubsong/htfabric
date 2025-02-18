// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/chaincode.proto

package peer

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Confidentiality Levels
type ConfidentialityLevel int32

const (
	ConfidentialityLevel_PUBLIC       ConfidentialityLevel = 0
	ConfidentialityLevel_CONFIDENTIAL ConfidentialityLevel = 1
)

var ConfidentialityLevel_name = map[int32]string{
	0: "PUBLIC",
	1: "CONFIDENTIAL",
}

var ConfidentialityLevel_value = map[string]int32{
	"PUBLIC":       0,
	"CONFIDENTIAL": 1,
}

func (x ConfidentialityLevel) String() string {
	return proto.EnumName(ConfidentialityLevel_name, int32(x))
}

func (ConfidentialityLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{0}
}

type ChaincodeSpec_Type int32

const (
	ChaincodeSpec_UNDEFINED ChaincodeSpec_Type = 0
	ChaincodeSpec_GOLANG    ChaincodeSpec_Type = 1
	ChaincodeSpec_NODE      ChaincodeSpec_Type = 2
	ChaincodeSpec_CAR       ChaincodeSpec_Type = 3
	ChaincodeSpec_JAVA      ChaincodeSpec_Type = 4
)

var ChaincodeSpec_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "GOLANG",
	2: "NODE",
	3: "CAR",
	4: "JAVA",
}

var ChaincodeSpec_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"GOLANG":    1,
	"NODE":      2,
	"CAR":       3,
	"JAVA":      4,
}

func (x ChaincodeSpec_Type) String() string {
	return proto.EnumName(ChaincodeSpec_Type_name, int32(x))
}

func (ChaincodeSpec_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{2, 0}
}

type ChaincodeDeploymentSpec_ExecutionEnvironment int32

const (
	ChaincodeDeploymentSpec_DOCKER ChaincodeDeploymentSpec_ExecutionEnvironment = 0
	ChaincodeDeploymentSpec_SYSTEM ChaincodeDeploymentSpec_ExecutionEnvironment = 1
)

var ChaincodeDeploymentSpec_ExecutionEnvironment_name = map[int32]string{
	0: "DOCKER",
	1: "SYSTEM",
}

var ChaincodeDeploymentSpec_ExecutionEnvironment_value = map[string]int32{
	"DOCKER": 0,
	"SYSTEM": 1,
}

func (x ChaincodeDeploymentSpec_ExecutionEnvironment) String() string {
	return proto.EnumName(ChaincodeDeploymentSpec_ExecutionEnvironment_name, int32(x))
}

func (ChaincodeDeploymentSpec_ExecutionEnvironment) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{3, 0}
}

//ChaincodeID contains the path as specified by the deploy transaction
//that created it as well as the hashCode that is generated by the
//system for the path. From the user level (ie, CLI, REST API and so on)
//deploy transaction is expected to provide the path and other requests
//are expected to provide the hashCode. The other value will be ignored.
//Internally, the structure could contain both values. For instance, the
//hashCode will be set when first generated using the path
type ChaincodeID struct {
	//deploy transaction will use the path
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	//all other requests will use the name (really a hashcode) generated by
	//the deploy transaction
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//user friendly version name for the chaincode
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeID) Reset()         { *m = ChaincodeID{} }
func (m *ChaincodeID) String() string { return proto.CompactTextString(m) }
func (*ChaincodeID) ProtoMessage()    {}
func (*ChaincodeID) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{0}
}

func (m *ChaincodeID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeID.Unmarshal(m, b)
}
func (m *ChaincodeID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeID.Marshal(b, m, deterministic)
}
func (m *ChaincodeID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeID.Merge(m, src)
}
func (m *ChaincodeID) XXX_Size() int {
	return xxx_messageInfo_ChaincodeID.Size(m)
}
func (m *ChaincodeID) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeID.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeID proto.InternalMessageInfo

func (m *ChaincodeID) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ChaincodeID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeID) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Carries the chaincode function and its arguments.
// UnmarshalJSON in transaction.go converts the string-based REST/JSON input to
// the []byte-based current ChaincodeInput structure.
type ChaincodeInput struct {
	Args                 [][]byte          `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	Decorations          map[string][]byte `protobuf:"bytes,2,rep,name=decorations,proto3" json:"decorations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ChaincodeInput) Reset()         { *m = ChaincodeInput{} }
func (m *ChaincodeInput) String() string { return proto.CompactTextString(m) }
func (*ChaincodeInput) ProtoMessage()    {}
func (*ChaincodeInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{1}
}

func (m *ChaincodeInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeInput.Unmarshal(m, b)
}
func (m *ChaincodeInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeInput.Marshal(b, m, deterministic)
}
func (m *ChaincodeInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeInput.Merge(m, src)
}
func (m *ChaincodeInput) XXX_Size() int {
	return xxx_messageInfo_ChaincodeInput.Size(m)
}
func (m *ChaincodeInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeInput.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeInput proto.InternalMessageInfo

func (m *ChaincodeInput) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *ChaincodeInput) GetDecorations() map[string][]byte {
	if m != nil {
		return m.Decorations
	}
	return nil
}

// Carries the chaincode specification. This is the actual metadata required for
// defining a chaincode.
type ChaincodeSpec struct {
	Type                 ChaincodeSpec_Type `protobuf:"varint,1,opt,name=type,proto3,enum=protos.ChaincodeSpec_Type" json:"type,omitempty"`
	ChaincodeId          *ChaincodeID       `protobuf:"bytes,2,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	Input                *ChaincodeInput    `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	Timeout              int32              `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ChaincodeSpec) Reset()         { *m = ChaincodeSpec{} }
func (m *ChaincodeSpec) String() string { return proto.CompactTextString(m) }
func (*ChaincodeSpec) ProtoMessage()    {}
func (*ChaincodeSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{2}
}

func (m *ChaincodeSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeSpec.Unmarshal(m, b)
}
func (m *ChaincodeSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeSpec.Marshal(b, m, deterministic)
}
func (m *ChaincodeSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeSpec.Merge(m, src)
}
func (m *ChaincodeSpec) XXX_Size() int {
	return xxx_messageInfo_ChaincodeSpec.Size(m)
}
func (m *ChaincodeSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeSpec proto.InternalMessageInfo

func (m *ChaincodeSpec) GetType() ChaincodeSpec_Type {
	if m != nil {
		return m.Type
	}
	return ChaincodeSpec_UNDEFINED
}

func (m *ChaincodeSpec) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

func (m *ChaincodeSpec) GetInput() *ChaincodeInput {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *ChaincodeSpec) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// Specify the deployment of a chaincode.
// TODO: Define `codePackage`.
type ChaincodeDeploymentSpec struct {
	ChaincodeSpec        *ChaincodeSpec                               `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec,proto3" json:"chaincode_spec,omitempty"`
	CodePackage          []byte                                       `protobuf:"bytes,3,opt,name=code_package,json=codePackage,proto3" json:"code_package,omitempty"`
	ExecEnv              ChaincodeDeploymentSpec_ExecutionEnvironment `protobuf:"varint,4,opt,name=exec_env,json=execEnv,proto3,enum=protos.ChaincodeDeploymentSpec_ExecutionEnvironment" json:"exec_env,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *ChaincodeDeploymentSpec) Reset()         { *m = ChaincodeDeploymentSpec{} }
func (m *ChaincodeDeploymentSpec) String() string { return proto.CompactTextString(m) }
func (*ChaincodeDeploymentSpec) ProtoMessage()    {}
func (*ChaincodeDeploymentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{3}
}

func (m *ChaincodeDeploymentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeDeploymentSpec.Unmarshal(m, b)
}
func (m *ChaincodeDeploymentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeDeploymentSpec.Marshal(b, m, deterministic)
}
func (m *ChaincodeDeploymentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeDeploymentSpec.Merge(m, src)
}
func (m *ChaincodeDeploymentSpec) XXX_Size() int {
	return xxx_messageInfo_ChaincodeDeploymentSpec.Size(m)
}
func (m *ChaincodeDeploymentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeDeploymentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeDeploymentSpec proto.InternalMessageInfo

func (m *ChaincodeDeploymentSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetCodePackage() []byte {
	if m != nil {
		return m.CodePackage
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetExecEnv() ChaincodeDeploymentSpec_ExecutionEnvironment {
	if m != nil {
		return m.ExecEnv
	}
	return ChaincodeDeploymentSpec_DOCKER
}

// Carries the chaincode function and its arguments.
type ChaincodeInvocationSpec struct {
	ChaincodeSpec        *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec,proto3" json:"chaincode_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ChaincodeInvocationSpec) Reset()         { *m = ChaincodeInvocationSpec{} }
func (m *ChaincodeInvocationSpec) String() string { return proto.CompactTextString(m) }
func (*ChaincodeInvocationSpec) ProtoMessage()    {}
func (*ChaincodeInvocationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{4}
}

func (m *ChaincodeInvocationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeInvocationSpec.Unmarshal(m, b)
}
func (m *ChaincodeInvocationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeInvocationSpec.Marshal(b, m, deterministic)
}
func (m *ChaincodeInvocationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeInvocationSpec.Merge(m, src)
}
func (m *ChaincodeInvocationSpec) XXX_Size() int {
	return xxx_messageInfo_ChaincodeInvocationSpec.Size(m)
}
func (m *ChaincodeInvocationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeInvocationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeInvocationSpec proto.InternalMessageInfo

func (m *ChaincodeInvocationSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

// LifecycleEvent is used as the payload of the chaincode event emitted by LSCC
type LifecycleEvent struct {
	ChaincodeName        string   `protobuf:"bytes,1,opt,name=chaincode_name,json=chaincodeName,proto3" json:"chaincode_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LifecycleEvent) Reset()         { *m = LifecycleEvent{} }
func (m *LifecycleEvent) String() string { return proto.CompactTextString(m) }
func (*LifecycleEvent) ProtoMessage()    {}
func (*LifecycleEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_202814c635ff5fee, []int{5}
}

func (m *LifecycleEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecycleEvent.Unmarshal(m, b)
}
func (m *LifecycleEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecycleEvent.Marshal(b, m, deterministic)
}
func (m *LifecycleEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecycleEvent.Merge(m, src)
}
func (m *LifecycleEvent) XXX_Size() int {
	return xxx_messageInfo_LifecycleEvent.Size(m)
}
func (m *LifecycleEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LifecycleEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LifecycleEvent proto.InternalMessageInfo

func (m *LifecycleEvent) GetChaincodeName() string {
	if m != nil {
		return m.ChaincodeName
	}
	return ""
}

func init() {
	proto.RegisterEnum("protos.ConfidentialityLevel", ConfidentialityLevel_name, ConfidentialityLevel_value)
	proto.RegisterEnum("protos.ChaincodeSpec_Type", ChaincodeSpec_Type_name, ChaincodeSpec_Type_value)
	proto.RegisterEnum("protos.ChaincodeDeploymentSpec_ExecutionEnvironment", ChaincodeDeploymentSpec_ExecutionEnvironment_name, ChaincodeDeploymentSpec_ExecutionEnvironment_value)
	proto.RegisterType((*ChaincodeID)(nil), "protos.ChaincodeID")
	proto.RegisterType((*ChaincodeInput)(nil), "protos.ChaincodeInput")
	proto.RegisterMapType((map[string][]byte)(nil), "protos.ChaincodeInput.DecorationsEntry")
	proto.RegisterType((*ChaincodeSpec)(nil), "protos.ChaincodeSpec")
	proto.RegisterType((*ChaincodeDeploymentSpec)(nil), "protos.ChaincodeDeploymentSpec")
	proto.RegisterType((*ChaincodeInvocationSpec)(nil), "protos.ChaincodeInvocationSpec")
	proto.RegisterType((*LifecycleEvent)(nil), "protos.LifecycleEvent")
}

func init() { proto.RegisterFile("peer/chaincode.proto", fileDescriptor_202814c635ff5fee) }

var fileDescriptor_202814c635ff5fee = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x6d, 0x6f, 0xda, 0x3c,
	0x14, 0x6d, 0x80, 0xbe, 0xdd, 0x50, 0x94, 0xc7, 0x0f, 0xcf, 0x33, 0xd4, 0x4f, 0x2c, 0xd2, 0x34,
	0x36, 0x4d, 0x41, 0x62, 0xd5, 0x36, 0x4d, 0x53, 0x25, 0x4a, 0xd2, 0x2a, 0x1d, 0x83, 0xca, 0x6d,
	0x27, 0x6d, 0x5f, 0x50, 0xea, 0x5c, 0xc0, 0x6a, 0x70, 0xa2, 0x60, 0xa2, 0xe6, 0x67, 0xec, 0x97,
	0xec, 0x27, 0x6e, 0xb2, 0x53, 0x5e, 0xba, 0xf6, 0xdb, 0x3e, 0x71, 0x7d, 0x7c, 0x7c, 0xee, 0x3d,
	0x07, 0xc7, 0x50, 0x4f, 0x10, 0xd3, 0x36, 0x9b, 0x06, 0x5c, 0xb0, 0x38, 0x44, 0x27, 0x49, 0x63,
	0x19, 0x93, 0x1d, 0xfd, 0x33, 0xb7, 0x87, 0x60, 0xf6, 0x96, 0x5b, 0xbe, 0x4b, 0x08, 0x54, 0x92,
	0x40, 0x4e, 0x1b, 0x46, 0xd3, 0x68, 0xed, 0x53, 0x5d, 0x2b, 0x4c, 0x04, 0x33, 0x6c, 0x94, 0x0a,
	0x4c, 0xd5, 0xa4, 0x01, 0xbb, 0x19, 0xa6, 0x73, 0x1e, 0x8b, 0x46, 0x59, 0xc3, 0xcb, 0xa5, 0xfd,
	0xd3, 0x80, 0xda, 0x5a, 0x51, 0x24, 0x0b, 0xa9, 0x04, 0x82, 0x74, 0x32, 0x6f, 0x18, 0xcd, 0x72,
	0xab, 0x4a, 0x75, 0x4d, 0x7c, 0x30, 0x43, 0x64, 0x71, 0x1a, 0x48, 0x1e, 0x8b, 0x79, 0xa3, 0xd4,
	0x2c, 0xb7, 0xcc, 0xce, 0xcb, 0x62, 0xb8, 0xb9, 0xf3, 0x50, 0xc0, 0x71, 0xd7, 0x4c, 0x4f, 0xc8,
	0x34, 0xa7, 0x9b, 0x67, 0x0f, 0x8f, 0xc1, 0xfa, 0x93, 0x40, 0x2c, 0x28, 0xdf, 0x62, 0x7e, 0x6f,
	0x43, 0x95, 0xa4, 0x0e, 0xdb, 0x59, 0x10, 0x2d, 0x0a, 0x1b, 0x55, 0x5a, 0x2c, 0x3e, 0x96, 0x3e,
	0x18, 0xf6, 0x2f, 0x03, 0x0e, 0x56, 0x0d, 0x2f, 0x13, 0x64, 0xc4, 0x81, 0x8a, 0xcc, 0x13, 0xd4,
	0xc7, 0x6b, 0x9d, 0xc3, 0x47, 0x53, 0x29, 0x92, 0x73, 0x95, 0x27, 0x48, 0x35, 0x8f, 0xbc, 0x83,
	0xea, 0x2a, 0xdf, 0x11, 0x0f, 0x75, 0x0b, 0xb3, 0xf3, 0xef, 0x63, 0x37, 0x2e, 0x35, 0x57, 0x44,
	0x3f, 0x24, 0x6f, 0x60, 0x9b, 0x2b, 0x83, 0x3a, 0x43, 0xb3, 0xf3, 0xff, 0xd3, 0xf6, 0x69, 0x41,
	0x52, 0x99, 0x4b, 0x3e, 0xc3, 0x78, 0x21, 0x1b, 0x95, 0xa6, 0xd1, 0xda, 0xa6, 0xcb, 0xa5, 0x7d,
	0x0c, 0x15, 0x35, 0x0d, 0x39, 0x80, 0xfd, 0xeb, 0x81, 0xeb, 0x9d, 0xfa, 0x03, 0xcf, 0xb5, 0xb6,
	0x08, 0xc0, 0xce, 0xd9, 0xb0, 0xdf, 0x1d, 0x9c, 0x59, 0x06, 0xd9, 0x83, 0xca, 0x60, 0xe8, 0x7a,
	0x56, 0x89, 0xec, 0x42, 0xb9, 0xd7, 0xa5, 0x56, 0x59, 0x41, 0xe7, 0xdd, 0xaf, 0x5d, 0xab, 0x62,
	0xff, 0x28, 0xc1, 0xb3, 0x55, 0x4f, 0x17, 0x93, 0x28, 0xce, 0x67, 0x28, 0xa4, 0xce, 0xe2, 0x13,
	0xd4, 0xd6, 0xde, 0xe6, 0x09, 0x32, 0x9d, 0x8a, 0xd9, 0xf9, 0xef, 0xc9, 0x54, 0xe8, 0x01, 0x7b,
	0x90, 0xe4, 0x73, 0xa8, 0xea, 0x83, 0x49, 0xc0, 0x6e, 0x83, 0x09, 0x6a, 0xa3, 0x55, 0x6a, 0x2a,
	0xec, 0xa2, 0x80, 0xc8, 0x10, 0xf6, 0xf0, 0x0e, 0xd9, 0x08, 0x45, 0xa6, 0x7d, 0xd5, 0x3a, 0x47,
	0x8f, 0xa4, 0x1f, 0xce, 0xe4, 0x78, 0x77, 0xc8, 0x16, 0xea, 0xdf, 0xf6, 0x44, 0xc6, 0xd3, 0x58,
	0xa8, 0x0d, 0xba, 0xab, 0x54, 0x3c, 0x91, 0xd9, 0x0e, 0xd4, 0x9f, 0x22, 0xa8, 0x38, 0xdc, 0x61,
	0xef, 0xb3, 0x47, 0x8b, 0x68, 0x2e, 0xbf, 0x5d, 0x5e, 0x79, 0x5f, 0x2c, 0xe3, 0xbc, 0xb2, 0x57,
	0xb2, 0xca, 0xb4, 0x86, 0xe3, 0x31, 0x32, 0xc9, 0x33, 0x1c, 0x85, 0x81, 0x44, 0x3b, 0xd9, 0x88,
	0xc4, 0x17, 0x59, 0xcc, 0xf4, 0xf5, 0xfa, 0xfb, 0x48, 0xee, 0xdb, 0xfd, 0xc3, 0xc3, 0xd1, 0x04,
	0x05, 0x16, 0xb7, 0x76, 0x14, 0x44, 0x13, 0xfb, 0x3d, 0xd4, 0xfa, 0x7c, 0x8c, 0x2c, 0x67, 0x11,
	0x7a, 0x99, 0x9a, 0xf8, 0xc5, 0x66, 0x23, 0xfd, 0x0d, 0x16, 0x17, 0x7a, 0xad, 0x38, 0x08, 0x66,
	0xf8, 0xfa, 0x08, 0xea, 0xbd, 0x58, 0x8c, 0x79, 0x88, 0x42, 0xf2, 0x20, 0xe2, 0x32, 0xef, 0x63,
	0x86, 0x91, 0x32, 0x79, 0x71, 0x7d, 0xd2, 0xf7, 0x7b, 0xd6, 0x16, 0xb1, 0xa0, 0xda, 0x1b, 0x0e,
	0x4e, 0x7d, 0xd7, 0x1b, 0x5c, 0xf9, 0xdd, 0xbe, 0x65, 0x9c, 0x0c, 0xc1, 0x8e, 0xd3, 0x89, 0x33,
	0xcd, 0x13, 0x4c, 0x23, 0x0c, 0x27, 0x98, 0x3a, 0xe3, 0xe0, 0x26, 0xe5, 0x6c, 0xe9, 0x42, 0xbd,
	0x1b, 0xdf, 0x5f, 0x4d, 0xb8, 0x9c, 0x2e, 0x6e, 0x1c, 0x16, 0xcf, 0xda, 0x1b, 0xd4, 0x76, 0x41,
	0x6d, 0x17, 0xd4, 0xb6, 0xa2, 0xde, 0x14, 0x4f, 0xca, 0xdb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xea, 0x1f, 0x3b, 0xe4, 0x71, 0x04, 0x00, 0x00,
}
