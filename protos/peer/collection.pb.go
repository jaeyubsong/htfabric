// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/collection.proto

package peer

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/hyperledger/fabric/protos/common"
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

// CollectionConfigPackage represents an array of CollectionConfig
// messages; the extra struct is required because repeated oneof is
// forbidden by the protobuf syntax
type CollectionConfigPackage struct {
	Config               []*CollectionConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CollectionConfigPackage) Reset()         { *m = CollectionConfigPackage{} }
func (m *CollectionConfigPackage) String() string { return proto.CompactTextString(m) }
func (*CollectionConfigPackage) ProtoMessage()    {}
func (*CollectionConfigPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8182e05ac5917d8, []int{0}
}

func (m *CollectionConfigPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionConfigPackage.Unmarshal(m, b)
}
func (m *CollectionConfigPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionConfigPackage.Marshal(b, m, deterministic)
}
func (m *CollectionConfigPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionConfigPackage.Merge(m, src)
}
func (m *CollectionConfigPackage) XXX_Size() int {
	return xxx_messageInfo_CollectionConfigPackage.Size(m)
}
func (m *CollectionConfigPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionConfigPackage.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionConfigPackage proto.InternalMessageInfo

func (m *CollectionConfigPackage) GetConfig() []*CollectionConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// CollectionConfig defines the configuration of a collection object;
// it currently contains a single, static type.
// Dynamic collections are deferred.
type CollectionConfig struct {
	// Types that are valid to be assigned to Payload:
	//	*CollectionConfig_StaticCollectionConfig
	Payload              isCollectionConfig_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CollectionConfig) Reset()         { *m = CollectionConfig{} }
func (m *CollectionConfig) String() string { return proto.CompactTextString(m) }
func (*CollectionConfig) ProtoMessage()    {}
func (*CollectionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8182e05ac5917d8, []int{1}
}

func (m *CollectionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionConfig.Unmarshal(m, b)
}
func (m *CollectionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionConfig.Marshal(b, m, deterministic)
}
func (m *CollectionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionConfig.Merge(m, src)
}
func (m *CollectionConfig) XXX_Size() int {
	return xxx_messageInfo_CollectionConfig.Size(m)
}
func (m *CollectionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionConfig proto.InternalMessageInfo

type isCollectionConfig_Payload interface {
	isCollectionConfig_Payload()
}

type CollectionConfig_StaticCollectionConfig struct {
	StaticCollectionConfig *StaticCollectionConfig `protobuf:"bytes,1,opt,name=static_collection_config,json=staticCollectionConfig,proto3,oneof"`
}

func (*CollectionConfig_StaticCollectionConfig) isCollectionConfig_Payload() {}

func (m *CollectionConfig) GetPayload() isCollectionConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CollectionConfig) GetStaticCollectionConfig() *StaticCollectionConfig {
	if x, ok := m.GetPayload().(*CollectionConfig_StaticCollectionConfig); ok {
		return x.StaticCollectionConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CollectionConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CollectionConfig_StaticCollectionConfig)(nil),
	}
}

// StaticCollectionConfig constitutes the configuration parameters of a
// static collection object. Static collections are collections that are
// known at chaincode instantiation time, and that cannot be changed.
// Dynamic collections are deferred.
type StaticCollectionConfig struct {
	// the name of the collection inside the denoted chaincode
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// a reference to a policy residing / managed in the config block
	// to define which orgs have access to this collection’s private data
	MemberOrgsPolicy *CollectionPolicyConfig `protobuf:"bytes,2,opt,name=member_orgs_policy,json=memberOrgsPolicy,proto3" json:"member_orgs_policy,omitempty"`
	// The minimum number of peers private data will be sent to upon
	// endorsement. The endorsement would fail if dissemination to at least
	// this number of peers is not achieved.
	RequiredPeerCount int32 `protobuf:"varint,3,opt,name=required_peer_count,json=requiredPeerCount,proto3" json:"required_peer_count,omitempty"`
	// The maximum number of peers that private data will be sent to
	// upon endorsement. This number has to be bigger than required_peer_count.
	MaximumPeerCount int32 `protobuf:"varint,4,opt,name=maximum_peer_count,json=maximumPeerCount,proto3" json:"maximum_peer_count,omitempty"`
	// The number of blocks after which the collection data expires.
	// For instance if the value is set to 10, a key last modified by block number 100
	// will be purged at block number 111. A zero value is treated same as MaxUint64
	BlockToLive uint64 `protobuf:"varint,5,opt,name=block_to_live,json=blockToLive,proto3" json:"block_to_live,omitempty"`
	// The member only read access denotes whether only collection member clients
	// can read the private data (if set to true), or even non members can
	// read the data (if set to false, for example if you want to implement more granular
	// access logic in the chaincode)
	MemberOnlyRead bool `protobuf:"varint,6,opt,name=member_only_read,json=memberOnlyRead,proto3" json:"member_only_read,omitempty"`
	// The member only write access denotes whether only collection member clients
	// can write the private data (if set to true), or even non members can
	// write the data (if set to false, for example if you want to implement more granular
	// access logic in the chaincode)
	MemberOnlyWrite bool `protobuf:"varint,7,opt,name=member_only_write,json=memberOnlyWrite,proto3" json:"member_only_write,omitempty"`
	// a reference to a policy residing / managed in the config block
	// to define the endorsement policy for this collection
	EndorsementPolicy    *ApplicationPolicy `protobuf:"bytes,8,opt,name=endorsement_policy,json=endorsementPolicy,proto3" json:"endorsement_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StaticCollectionConfig) Reset()         { *m = StaticCollectionConfig{} }
func (m *StaticCollectionConfig) String() string { return proto.CompactTextString(m) }
func (*StaticCollectionConfig) ProtoMessage()    {}
func (*StaticCollectionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8182e05ac5917d8, []int{2}
}

func (m *StaticCollectionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticCollectionConfig.Unmarshal(m, b)
}
func (m *StaticCollectionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticCollectionConfig.Marshal(b, m, deterministic)
}
func (m *StaticCollectionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticCollectionConfig.Merge(m, src)
}
func (m *StaticCollectionConfig) XXX_Size() int {
	return xxx_messageInfo_StaticCollectionConfig.Size(m)
}
func (m *StaticCollectionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticCollectionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StaticCollectionConfig proto.InternalMessageInfo

func (m *StaticCollectionConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StaticCollectionConfig) GetMemberOrgsPolicy() *CollectionPolicyConfig {
	if m != nil {
		return m.MemberOrgsPolicy
	}
	return nil
}

func (m *StaticCollectionConfig) GetRequiredPeerCount() int32 {
	if m != nil {
		return m.RequiredPeerCount
	}
	return 0
}

func (m *StaticCollectionConfig) GetMaximumPeerCount() int32 {
	if m != nil {
		return m.MaximumPeerCount
	}
	return 0
}

func (m *StaticCollectionConfig) GetBlockToLive() uint64 {
	if m != nil {
		return m.BlockToLive
	}
	return 0
}

func (m *StaticCollectionConfig) GetMemberOnlyRead() bool {
	if m != nil {
		return m.MemberOnlyRead
	}
	return false
}

func (m *StaticCollectionConfig) GetMemberOnlyWrite() bool {
	if m != nil {
		return m.MemberOnlyWrite
	}
	return false
}

func (m *StaticCollectionConfig) GetEndorsementPolicy() *ApplicationPolicy {
	if m != nil {
		return m.EndorsementPolicy
	}
	return nil
}

// Collection policy configuration. Initially, the configuration can only
// contain a SignaturePolicy. In the future, the SignaturePolicy may be a
// more general Policy. Instead of containing the actual policy, the
// configuration may in the future contain a string reference to a policy.
type CollectionPolicyConfig struct {
	// Types that are valid to be assigned to Payload:
	//	*CollectionPolicyConfig_SignaturePolicy
	Payload              isCollectionPolicyConfig_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CollectionPolicyConfig) Reset()         { *m = CollectionPolicyConfig{} }
func (m *CollectionPolicyConfig) String() string { return proto.CompactTextString(m) }
func (*CollectionPolicyConfig) ProtoMessage()    {}
func (*CollectionPolicyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8182e05ac5917d8, []int{3}
}

func (m *CollectionPolicyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionPolicyConfig.Unmarshal(m, b)
}
func (m *CollectionPolicyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionPolicyConfig.Marshal(b, m, deterministic)
}
func (m *CollectionPolicyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionPolicyConfig.Merge(m, src)
}
func (m *CollectionPolicyConfig) XXX_Size() int {
	return xxx_messageInfo_CollectionPolicyConfig.Size(m)
}
func (m *CollectionPolicyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionPolicyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionPolicyConfig proto.InternalMessageInfo

type isCollectionPolicyConfig_Payload interface {
	isCollectionPolicyConfig_Payload()
}

type CollectionPolicyConfig_SignaturePolicy struct {
	SignaturePolicy *common.SignaturePolicyEnvelope `protobuf:"bytes,1,opt,name=signature_policy,json=signaturePolicy,proto3,oneof"`
}

func (*CollectionPolicyConfig_SignaturePolicy) isCollectionPolicyConfig_Payload() {}

func (m *CollectionPolicyConfig) GetPayload() isCollectionPolicyConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CollectionPolicyConfig) GetSignaturePolicy() *common.SignaturePolicyEnvelope {
	if x, ok := m.GetPayload().(*CollectionPolicyConfig_SignaturePolicy); ok {
		return x.SignaturePolicy
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CollectionPolicyConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CollectionPolicyConfig_SignaturePolicy)(nil),
	}
}

func init() {
	proto.RegisterType((*CollectionConfigPackage)(nil), "protos.CollectionConfigPackage")
	proto.RegisterType((*CollectionConfig)(nil), "protos.CollectionConfig")
	proto.RegisterType((*StaticCollectionConfig)(nil), "protos.StaticCollectionConfig")
	proto.RegisterType((*CollectionPolicyConfig)(nil), "protos.CollectionPolicyConfig")
}

func init() { proto.RegisterFile("peer/collection.proto", fileDescriptor_d8182e05ac5917d8) }

var fileDescriptor_d8182e05ac5917d8 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xd1, 0x6a, 0xdb, 0x30,
	0x18, 0x85, 0x9b, 0x25, 0x4d, 0x5b, 0x85, 0xad, 0x89, 0x46, 0x33, 0xaf, 0x17, 0x5b, 0xf0, 0x95,
	0x37, 0x86, 0x3d, 0xba, 0x27, 0x58, 0xc3, 0x20, 0xb0, 0x40, 0x83, 0x3b, 0x18, 0xf4, 0xc6, 0xc8,
	0xf2, 0x5f, 0x57, 0x54, 0x96, 0x5c, 0x49, 0xce, 0xe6, 0x77, 0xd9, 0xc3, 0x0e, 0x4b, 0x71, 0xed,
	0x86, 0x5c, 0x25, 0x9c, 0xf3, 0x9d, 0x5f, 0xbf, 0x0e, 0x32, 0xba, 0x28, 0x01, 0x54, 0x44, 0x25,
	0xe7, 0x40, 0x0d, 0x93, 0x22, 0x2c, 0x95, 0x34, 0x12, 0x8f, 0xed, 0x8f, 0xbe, 0xbc, 0xa0, 0xb2,
	0x28, 0xa4, 0x88, 0x4a, 0xc9, 0x19, 0x65, 0xa0, 0x9d, 0x7d, 0x39, 0xb3, 0x29, 0x2b, 0xd6, 0x4e,
	0xf2, 0x7f, 0xa2, 0x77, 0xcb, 0xe7, 0x29, 0x4b, 0x29, 0xee, 0x59, 0xbe, 0x21, 0xf4, 0x91, 0xe4,
	0x80, 0xbf, 0xa2, 0x31, 0xb5, 0x82, 0x37, 0x58, 0x0c, 0x83, 0xc9, 0x95, 0xe7, 0x22, 0x3a, 0xdc,
	0x0f, 0xc4, 0x3b, 0xce, 0xaf, 0xd1, 0x74, 0xdf, 0xc3, 0x77, 0xc8, 0xd3, 0x86, 0x18, 0x46, 0x93,
	0x6e, 0xdb, 0xe4, 0x79, 0xee, 0x20, 0x98, 0x5c, 0x7d, 0x68, 0xe7, 0xde, 0x5a, 0x6e, 0x7f, 0xc2,
	0xea, 0x28, 0x9e, 0xeb, 0x83, 0xce, 0xf5, 0x19, 0x3a, 0x29, 0x49, 0xcd, 0x25, 0xc9, 0xfc, 0x7f,
	0x43, 0x34, 0x3f, 0x9c, 0xc7, 0x18, 0x8d, 0x04, 0x29, 0xc0, 0x9e, 0x76, 0x16, 0xdb, 0xff, 0x78,
	0x8d, 0x70, 0x01, 0x45, 0x0a, 0x2a, 0x91, 0x2a, 0xd7, 0x89, 0xab, 0xc4, 0x7b, 0xf5, 0x72, 0x9f,
	0x6e, 0xd2, 0xc6, 0xfa, 0xbb, 0xdb, 0x4e, 0x5d, 0xf2, 0x46, 0xe5, 0xda, 0xe9, 0x38, 0x44, 0x6f,
	0x15, 0x3c, 0x55, 0x4c, 0x41, 0x96, 0x34, 0x15, 0x27, 0x54, 0x56, 0xc2, 0x78, 0xc3, 0xc5, 0x20,
	0x38, 0x8e, 0x67, 0xad, 0xb5, 0x01, 0x50, 0xcb, 0xc6, 0xc0, 0x5f, 0x10, 0x2e, 0xc8, 0x5f, 0x56,
	0x54, 0x45, 0x1f, 0x1f, 0x59, 0x7c, 0xba, 0x73, 0x3a, 0xda, 0x47, 0xaf, 0x53, 0x2e, 0xe9, 0x63,
	0x62, 0x64, 0xc2, 0xd9, 0x16, 0xbc, 0xe3, 0xc5, 0x20, 0x18, 0xc5, 0x13, 0x2b, 0xfe, 0x92, 0x6b,
	0xb6, 0x05, 0x1c, 0xa0, 0x69, 0x7b, 0x1f, 0xc1, 0xeb, 0x44, 0x01, 0xc9, 0xbc, 0xf1, 0x62, 0x10,
	0x9c, 0xc6, 0x6f, 0x76, 0xdb, 0x0a, 0x5e, 0xc7, 0x40, 0x32, 0xfc, 0x19, 0xcd, 0xfa, 0xe4, 0x1f,
	0xc5, 0x0c, 0x78, 0x27, 0x16, 0x3d, 0xef, 0xd0, 0xdf, 0x8d, 0x8c, 0x57, 0x08, 0x83, 0xc8, 0xa4,
	0xd2, 0x50, 0x80, 0x30, 0x6d, 0x4b, 0xa7, 0xb6, 0xa5, 0xf7, 0x6d, 0x4b, 0xdf, 0xcb, 0x92, 0x33,
	0x4a, 0xba, 0x9a, 0xe2, 0x59, 0x2f, 0xe4, 0x24, 0xff, 0x09, 0xcd, 0x0f, 0xb7, 0x89, 0xd7, 0x68,
	0xaa, 0x59, 0x2e, 0x88, 0xa9, 0x14, 0xb4, 0x27, 0xb8, 0x77, 0xf1, 0x31, 0x74, 0xaf, 0x38, 0xbc,
	0x6d, 0x7d, 0x17, 0xfc, 0x21, 0xb6, 0xc0, 0x65, 0x09, 0xab, 0xa3, 0xf8, 0x5c, 0xbf, 0xb4, 0x7a,
	0x2f, 0xe2, 0xfa, 0x06, 0xf9, 0x52, 0xe5, 0xe1, 0x43, 0x5d, 0x82, 0xe2, 0x90, 0xe5, 0xa0, 0xc2,
	0x7b, 0x92, 0x2a, 0x46, 0xdb, 0xc5, 0x9b, 0xea, 0xef, 0x3e, 0xe5, 0xcc, 0x3c, 0x54, 0x69, 0x73,
	0x54, 0xd4, 0x43, 0x23, 0x87, 0x46, 0x0e, 0x8d, 0x1a, 0x34, 0x75, 0x1f, 0xd7, 0xb7, 0xff, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x04, 0xb9, 0xf7, 0xb4, 0x7c, 0x03, 0x00, 0x00,
}
