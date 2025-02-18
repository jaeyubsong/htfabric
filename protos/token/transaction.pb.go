// Code generated by protoc-gen-go. DO NOT EDIT.
// source: token/transaction.proto

package token

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

// TokenTransaction governs the structure of Payload.data, when
// the transaction's envelope header indicates a transaction of type
// "Token"
type TokenTransaction struct {
	// action carries the content of this transaction.
	//
	// Types that are valid to be assigned to Action:
	//	*TokenTransaction_PlainAction
	Action               isTokenTransaction_Action `protobuf_oneof:"action"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TokenTransaction) Reset()         { *m = TokenTransaction{} }
func (m *TokenTransaction) String() string { return proto.CompactTextString(m) }
func (*TokenTransaction) ProtoMessage()    {}
func (*TokenTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_fadc60fa5929c0a6, []int{0}
}

func (m *TokenTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenTransaction.Unmarshal(m, b)
}
func (m *TokenTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenTransaction.Marshal(b, m, deterministic)
}
func (m *TokenTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenTransaction.Merge(m, src)
}
func (m *TokenTransaction) XXX_Size() int {
	return xxx_messageInfo_TokenTransaction.Size(m)
}
func (m *TokenTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_TokenTransaction proto.InternalMessageInfo

type isTokenTransaction_Action interface {
	isTokenTransaction_Action()
}

type TokenTransaction_PlainAction struct {
	PlainAction *PlainTokenAction `protobuf:"bytes,1,opt,name=plain_action,json=plainAction,proto3,oneof"`
}

func (*TokenTransaction_PlainAction) isTokenTransaction_Action() {}

func (m *TokenTransaction) GetAction() isTokenTransaction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *TokenTransaction) GetPlainAction() *PlainTokenAction {
	if x, ok := m.GetAction().(*TokenTransaction_PlainAction); ok {
		return x.PlainAction
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TokenTransaction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TokenTransaction_PlainAction)(nil),
	}
}

// PlainTokenAction governs the structure of a token action that is
// subjected to no privacy restrictions
type PlainTokenAction struct {
	// Types that are valid to be assigned to Data:
	//	*PlainTokenAction_PlainImport
	//	*PlainTokenAction_PlainTransfer
	//	*PlainTokenAction_PlainRedeem
	//	*PlainTokenAction_PlainApprove
	//	*PlainTokenAction_PlainTransfer_From
	Data                 isPlainTokenAction_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PlainTokenAction) Reset()         { *m = PlainTokenAction{} }
func (m *PlainTokenAction) String() string { return proto.CompactTextString(m) }
func (*PlainTokenAction) ProtoMessage()    {}
func (*PlainTokenAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_fadc60fa5929c0a6, []int{1}
}

func (m *PlainTokenAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlainTokenAction.Unmarshal(m, b)
}
func (m *PlainTokenAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlainTokenAction.Marshal(b, m, deterministic)
}
func (m *PlainTokenAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainTokenAction.Merge(m, src)
}
func (m *PlainTokenAction) XXX_Size() int {
	return xxx_messageInfo_PlainTokenAction.Size(m)
}
func (m *PlainTokenAction) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainTokenAction.DiscardUnknown(m)
}

var xxx_messageInfo_PlainTokenAction proto.InternalMessageInfo

type isPlainTokenAction_Data interface {
	isPlainTokenAction_Data()
}

type PlainTokenAction_PlainImport struct {
	PlainImport *PlainImport `protobuf:"bytes,1,opt,name=plain_import,json=plainImport,proto3,oneof"`
}

type PlainTokenAction_PlainTransfer struct {
	PlainTransfer *PlainTransfer `protobuf:"bytes,2,opt,name=plain_transfer,json=plainTransfer,proto3,oneof"`
}

type PlainTokenAction_PlainRedeem struct {
	PlainRedeem *PlainTransfer `protobuf:"bytes,3,opt,name=plain_redeem,json=plainRedeem,proto3,oneof"`
}

type PlainTokenAction_PlainApprove struct {
	PlainApprove *PlainApprove `protobuf:"bytes,4,opt,name=plain_approve,json=plainApprove,proto3,oneof"`
}

type PlainTokenAction_PlainTransfer_From struct {
	PlainTransfer_From *PlainTransferFrom `protobuf:"bytes,5,opt,name=plain_transfer_From,json=plainTransferFrom,proto3,oneof"`
}

func (*PlainTokenAction_PlainImport) isPlainTokenAction_Data() {}

func (*PlainTokenAction_PlainTransfer) isPlainTokenAction_Data() {}

func (*PlainTokenAction_PlainRedeem) isPlainTokenAction_Data() {}

func (*PlainTokenAction_PlainApprove) isPlainTokenAction_Data() {}

func (*PlainTokenAction_PlainTransfer_From) isPlainTokenAction_Data() {}

func (m *PlainTokenAction) GetData() isPlainTokenAction_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PlainTokenAction) GetPlainImport() *PlainImport {
	if x, ok := m.GetData().(*PlainTokenAction_PlainImport); ok {
		return x.PlainImport
	}
	return nil
}

func (m *PlainTokenAction) GetPlainTransfer() *PlainTransfer {
	if x, ok := m.GetData().(*PlainTokenAction_PlainTransfer); ok {
		return x.PlainTransfer
	}
	return nil
}

func (m *PlainTokenAction) GetPlainRedeem() *PlainTransfer {
	if x, ok := m.GetData().(*PlainTokenAction_PlainRedeem); ok {
		return x.PlainRedeem
	}
	return nil
}

func (m *PlainTokenAction) GetPlainApprove() *PlainApprove {
	if x, ok := m.GetData().(*PlainTokenAction_PlainApprove); ok {
		return x.PlainApprove
	}
	return nil
}

func (m *PlainTokenAction) GetPlainTransfer_From() *PlainTransferFrom {
	if x, ok := m.GetData().(*PlainTokenAction_PlainTransfer_From); ok {
		return x.PlainTransfer_From
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PlainTokenAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PlainTokenAction_PlainImport)(nil),
		(*PlainTokenAction_PlainTransfer)(nil),
		(*PlainTokenAction_PlainRedeem)(nil),
		(*PlainTokenAction_PlainApprove)(nil),
		(*PlainTokenAction_PlainTransfer_From)(nil),
	}
}

// PlainImport specifies an import of one or more tokens in plaintext format
type PlainImport struct {
	// An import transaction may contain one or more outputs
	Outputs              []*PlainOutput `protobuf:"bytes,1,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PlainImport) Reset()         { *m = PlainImport{} }
func (m *PlainImport) String() string { return proto.CompactTextString(m) }
func (*PlainImport) ProtoMessage()    {}
func (*PlainImport) Descriptor() ([]byte, []int) {
	return fileDescriptor_fadc60fa5929c0a6, []int{2}
}

func (m *PlainImport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlainImport.Unmarshal(m, b)
}
func (m *PlainImport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlainImport.Marshal(b, m, deterministic)
}
func (m *PlainImport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainImport.Merge(m, src)
}
func (m *PlainImport) XXX_Size() int {
	return xxx_messageInfo_PlainImport.Size(m)
}
func (m *PlainImport) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainImport.DiscardUnknown(m)
}

var xxx_messageInfo_PlainImport proto.InternalMessageInfo

func (m *PlainImport) GetOutputs() []*PlainOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// PlainTransfer specifies a transfer of one or more plaintext tokens to one or more outputs
type PlainTransfer struct {
	// The inputs to the transfer transaction are specified by their ID
	Inputs []*InputId `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// A transfer transaction may contain one or more outputs
	Outputs              []*PlainOutput `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PlainTransfer) Reset()         { *m = PlainTransfer{} }
func (m *PlainTransfer) String() string { return proto.CompactTextString(m) }
func (*PlainTransfer) ProtoMessage()    {}
func (*PlainTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_fadc60fa5929c0a6, []int{3}
}

func (m *PlainTransfer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlainTransfer.Unmarshal(m, b)
}
func (m *PlainTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlainTransfer.Marshal(b, m, deterministic)
}
func (m *PlainTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainTransfer.Merge(m, src)
}
func (m *PlainTransfer) XXX_Size() int {
	return xxx_messageInfo_PlainTransfer.Size(m)
}
func (m *PlainTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_PlainTransfer proto.InternalMessageInfo

func (m *PlainTransfer) GetInputs() []*InputId {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PlainTransfer) GetOutputs() []*PlainOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// PlainApprove specifies an approve of one or more tokens in plaintext format
type PlainApprove struct {
	// The inputs to the transfer transaction are specified by their ID
	Inputs []*InputId `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// An approve transaction contains one or more plain delegated outputs
	DelegatedOutputs []*PlainDelegatedOutput `protobuf:"bytes,2,rep,name=delegated_outputs,json=delegatedOutputs,proto3" json:"delegated_outputs,omitempty"`
	// An approve transaction contains one plain output
	Output               *PlainOutput `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PlainApprove) Reset()         { *m = PlainApprove{} }
func (m *PlainApprove) String() string { return proto.CompactTextString(m) }
func (*PlainApprove) ProtoMessage()    {}
func (*PlainApprove) Descriptor() ([]byte, []int) {
	return fileDescriptor_fadc60fa5929c0a6, []int{4}
}

func (m *PlainApprove) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlainApprove.Unmarshal(m, b)
}
func (m *PlainApprove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlainApprove.Marshal(b, m, deterministic)
}
func (m *PlainApprove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainApprove.Merge(m, src)
}
func (m *PlainApprove) XXX_Size() int {
	return xxx_messageInfo_PlainApprove.Size(m)
}
func (m *PlainApprove) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainApprove.DiscardUnknown(m)
}

var xxx_messageInfo_PlainApprove proto.InternalMessageInfo

func (m *PlainApprove) GetInputs() []*InputId {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PlainApprove) GetDelegatedOutputs() []*PlainDelegatedOutput {
	if m != nil {
		return m.DelegatedOutputs
	}
	return nil
}

func (m *PlainApprove) GetOutput() *PlainOutput {
	if m != nil {
		return m.Output
	}
	return nil
}

// PlainTransferFrom specifies a transfer of one or more plaintext delegated tokens to one or more outputs
// an to a delegated output
type PlainTransferFrom struct {
	// The inputs to the transfer transaction are specified by their ID
	Inputs []*InputId `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// A transferFrom transaction contains multiple outputs
	Outputs []*PlainOutput `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// A transferFrom transaction may contain one delegatable output
	DelegatedOutput      *PlainDelegatedOutput `protobuf:"bytes,3,opt,name=delegated_output,json=delegatedOutput,proto3" json:"delegated_output,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PlainTransferFrom) Reset()         { *m = PlainTransferFrom{} }
func (m *PlainTransferFrom) String() string { return proto.CompactTextString(m) }
func (*PlainTransferFrom) ProtoMessage()    {}
func (*PlainTransferFrom) Descriptor() ([]byte, []int) {
	return fileDescriptor_fadc60fa5929c0a6, []int{5}
}

func (m *PlainTransferFrom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlainTransferFrom.Unmarshal(m, b)
}
func (m *PlainTransferFrom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlainTransferFrom.Marshal(b, m, deterministic)
}
func (m *PlainTransferFrom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainTransferFrom.Merge(m, src)
}
func (m *PlainTransferFrom) XXX_Size() int {
	return xxx_messageInfo_PlainTransferFrom.Size(m)
}
func (m *PlainTransferFrom) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainTransferFrom.DiscardUnknown(m)
}

var xxx_messageInfo_PlainTransferFrom proto.InternalMessageInfo

func (m *PlainTransferFrom) GetInputs() []*InputId {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PlainTransferFrom) GetOutputs() []*PlainOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *PlainTransferFrom) GetDelegatedOutput() *PlainDelegatedOutput {
	if m != nil {
		return m.DelegatedOutput
	}
	return nil
}

// A PlainOutput is the result of import and transfer transactions using plaintext tokens
type PlainOutput struct {
	// The owner is the serialization of a SerializedIdentity struct
	Owner []byte `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// The token type
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// The quantity of tokens
	Quantity             uint64   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlainOutput) Reset()         { *m = PlainOutput{} }
func (m *PlainOutput) String() string { return proto.CompactTextString(m) }
func (*PlainOutput) ProtoMessage()    {}
func (*PlainOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_fadc60fa5929c0a6, []int{6}
}

func (m *PlainOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlainOutput.Unmarshal(m, b)
}
func (m *PlainOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlainOutput.Marshal(b, m, deterministic)
}
func (m *PlainOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainOutput.Merge(m, src)
}
func (m *PlainOutput) XXX_Size() int {
	return xxx_messageInfo_PlainOutput.Size(m)
}
func (m *PlainOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainOutput.DiscardUnknown(m)
}

var xxx_messageInfo_PlainOutput proto.InternalMessageInfo

func (m *PlainOutput) GetOwner() []byte {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PlainOutput) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PlainOutput) GetQuantity() uint64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

// An InputId specifies an output using the transaction ID and the index of the output in the transaction
type InputId struct {
	// The transaction ID
	TxId string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// The index of the output in the transaction
	Index                uint32   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputId) Reset()         { *m = InputId{} }
func (m *InputId) String() string { return proto.CompactTextString(m) }
func (*InputId) ProtoMessage()    {}
func (*InputId) Descriptor() ([]byte, []int) {
	return fileDescriptor_fadc60fa5929c0a6, []int{7}
}

func (m *InputId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputId.Unmarshal(m, b)
}
func (m *InputId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputId.Marshal(b, m, deterministic)
}
func (m *InputId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputId.Merge(m, src)
}
func (m *InputId) XXX_Size() int {
	return xxx_messageInfo_InputId.Size(m)
}
func (m *InputId) XXX_DiscardUnknown() {
	xxx_messageInfo_InputId.DiscardUnknown(m)
}

var xxx_messageInfo_InputId proto.InternalMessageInfo

func (m *InputId) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *InputId) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

// A PlainDelegatedOutput is the result of approve transactions using plaintext tokens
type PlainDelegatedOutput struct {
	// The owner is the serialization of a SerializedIdentity struct
	Owner []byte `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// The delegatees is an arrary of the serialized identities that can spend the output on behalf
	// the owner
	Delegatees [][]byte `protobuf:"bytes,2,rep,name=delegatees,proto3" json:"delegatees,omitempty"`
	// The token type
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// The quantity of tokens
	Quantity             uint64   `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlainDelegatedOutput) Reset()         { *m = PlainDelegatedOutput{} }
func (m *PlainDelegatedOutput) String() string { return proto.CompactTextString(m) }
func (*PlainDelegatedOutput) ProtoMessage()    {}
func (*PlainDelegatedOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_fadc60fa5929c0a6, []int{8}
}

func (m *PlainDelegatedOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlainDelegatedOutput.Unmarshal(m, b)
}
func (m *PlainDelegatedOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlainDelegatedOutput.Marshal(b, m, deterministic)
}
func (m *PlainDelegatedOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainDelegatedOutput.Merge(m, src)
}
func (m *PlainDelegatedOutput) XXX_Size() int {
	return xxx_messageInfo_PlainDelegatedOutput.Size(m)
}
func (m *PlainDelegatedOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainDelegatedOutput.DiscardUnknown(m)
}

var xxx_messageInfo_PlainDelegatedOutput proto.InternalMessageInfo

func (m *PlainDelegatedOutput) GetOwner() []byte {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PlainDelegatedOutput) GetDelegatees() [][]byte {
	if m != nil {
		return m.Delegatees
	}
	return nil
}

func (m *PlainDelegatedOutput) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PlainDelegatedOutput) GetQuantity() uint64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func init() {
	proto.RegisterType((*TokenTransaction)(nil), "TokenTransaction")
	proto.RegisterType((*PlainTokenAction)(nil), "PlainTokenAction")
	proto.RegisterType((*PlainImport)(nil), "PlainImport")
	proto.RegisterType((*PlainTransfer)(nil), "PlainTransfer")
	proto.RegisterType((*PlainApprove)(nil), "PlainApprove")
	proto.RegisterType((*PlainTransferFrom)(nil), "PlainTransferFrom")
	proto.RegisterType((*PlainOutput)(nil), "PlainOutput")
	proto.RegisterType((*InputId)(nil), "InputId")
	proto.RegisterType((*PlainDelegatedOutput)(nil), "PlainDelegatedOutput")
}

func init() { proto.RegisterFile("token/transaction.proto", fileDescriptor_fadc60fa5929c0a6) }

var fileDescriptor_fadc60fa5929c0a6 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x8f, 0xd2, 0x40,
	0x14, 0xa7, 0x50, 0xba, 0xf8, 0x68, 0x57, 0x98, 0x5d, 0x63, 0xe3, 0xc1, 0x90, 0xc6, 0x18, 0x63,
	0x0c, 0x8d, 0xee, 0xaa, 0x57, 0x97, 0x6c, 0x0c, 0x9c, 0x34, 0x23, 0x17, 0xbd, 0x90, 0xc2, 0xcc,
	0xb2, 0x8d, 0xd0, 0x19, 0x87, 0x41, 0x21, 0xf1, 0x93, 0x78, 0xf1, 0x63, 0xf8, 0xf5, 0x4c, 0xdf,
	0x4c, 0xd9, 0x16, 0xc5, 0x78, 0xd8, 0x5b, 0x7f, 0xbf, 0x37, 0xbf, 0x3f, 0x6f, 0xda, 0x14, 0xee,
	0x6b, 0xf1, 0x99, 0x67, 0xb1, 0x56, 0x49, 0xb6, 0x4a, 0x66, 0x3a, 0x15, 0x59, 0x5f, 0x2a, 0xa1,
	0x45, 0x34, 0x86, 0xce, 0x38, 0x1f, 0x8d, 0x6f, 0x26, 0xe4, 0x15, 0xf8, 0x72, 0x91, 0xa4, 0xd9,
	0xc4, 0xe0, 0xd0, 0xe9, 0x39, 0x4f, 0xda, 0x2f, 0xba, 0xfd, 0xf7, 0x39, 0x89, 0xa7, 0x2f, 0x70,
	0x30, 0xac, 0xd1, 0x36, 0x1e, 0x34, 0x70, 0xd0, 0x02, 0xcf, 0x28, 0xa2, 0x5f, 0x75, 0xe8, 0xec,
	0x9f, 0x26, 0xcf, 0x0b, 0xdb, 0x74, 0x29, 0x85, 0xd2, 0xd6, 0xd6, 0x37, 0xb6, 0x23, 0xe4, 0x76,
	0x8e, 0x06, 0x92, 0xd7, 0x70, 0x6c, 0x24, 0x58, 0xfc, 0x8a, 0xab, 0xb0, 0x8e, 0xa2, 0x63, 0xdb,
	0xc5, 0xb2, 0xc3, 0x1a, 0x0d, 0x64, 0x99, 0x20, 0x67, 0x45, 0x96, 0xe2, 0x8c, 0xf3, 0x65, 0xd8,
	0x38, 0x20, 0x33, 0x69, 0x14, 0x0f, 0x91, 0x73, 0x08, 0xec, 0xde, 0x52, 0x2a, 0xf1, 0x95, 0x87,
	0x2e, 0xaa, 0x02, 0xa3, 0xba, 0x30, 0xe4, 0xb0, 0x46, 0x8d, 0xb5, 0xc5, 0xe4, 0x12, 0x4e, 0xaa,
	0x1d, 0x27, 0x6f, 0x95, 0x58, 0x86, 0x4d, 0xd4, 0x92, 0x6a, 0x62, 0x3e, 0x19, 0xd6, 0x68, 0x57,
	0xee, 0x93, 0x03, 0x0f, 0x5c, 0x96, 0xe8, 0x24, 0x7a, 0x09, 0xed, 0xd2, 0x7d, 0x90, 0xc7, 0x70,
	0x24, 0xd6, 0x5a, 0xae, 0xf5, 0x2a, 0x74, 0x7a, 0x8d, 0x9b, 0xeb, 0x7a, 0x87, 0x24, 0x2d, 0x86,
	0xd1, 0x47, 0x08, 0x2a, 0x41, 0xa4, 0x07, 0x5e, 0x9a, 0x95, 0x74, 0xad, 0xfe, 0x28, 0x87, 0x23,
	0x46, 0x2d, 0x5f, 0xb6, 0xae, 0xff, 0xcb, 0xfa, 0x87, 0x03, 0x7e, 0xf9, 0x02, 0xfe, 0xc3, 0x7a,
	0x00, 0x5d, 0xc6, 0x17, 0x7c, 0x9e, 0x68, 0xce, 0x26, 0xd5, 0x90, 0x7b, 0x26, 0xe4, 0xb2, 0x18,
	0xdb, 0xb4, 0x0e, 0xab, 0x12, 0x2b, 0xf2, 0x08, 0x3c, 0xa3, 0xb4, 0xef, 0xae, 0xda, 0xce, 0xce,
	0xa2, 0x9f, 0x0e, 0x74, 0xff, 0xb8, 0xe1, 0xdb, 0x5b, 0x9e, 0xbc, 0x81, 0xce, 0xfe, 0x26, 0xb6,
	0xcf, 0x81, 0x45, 0xee, 0xee, 0x2d, 0x12, 0x7d, 0xb0, 0x2f, 0xd4, 0x40, 0x72, 0x0a, 0x4d, 0xf1,
	0x2d, 0xe3, 0x0a, 0xbf, 0x7e, 0x9f, 0x1a, 0x40, 0x08, 0xb8, 0x7a, 0x2b, 0x39, 0x7e, 0xdd, 0x77,
	0x28, 0x3e, 0x93, 0x07, 0xd0, 0xfa, 0xb2, 0x4e, 0x32, 0x9d, 0xea, 0x2d, 0x46, 0xba, 0x74, 0x87,
	0xa3, 0x73, 0x38, 0xb2, 0x1b, 0x91, 0x13, 0x68, 0xea, 0xcd, 0x24, 0x65, 0x68, 0x98, 0x6b, 0x37,
	0x23, 0x96, 0xa7, 0xa4, 0x19, 0xe3, 0x1b, 0x34, 0x0c, 0xa8, 0x01, 0xd1, 0x77, 0x38, 0xfd, 0x5b,
	0xe7, 0x03, 0x9d, 0x1e, 0x02, 0x14, 0xbb, 0x70, 0x73, 0x4b, 0x3e, 0x2d, 0x31, 0xbb, 0xce, 0x8d,
	0x03, 0x9d, 0xdd, 0x6a, 0xe7, 0xc1, 0xb3, 0x4f, 0x4f, 0xe7, 0xa9, 0xbe, 0x5e, 0x4f, 0xfb, 0x33,
	0xb1, 0x8c, 0xaf, 0xb7, 0x92, 0xab, 0x05, 0x67, 0x73, 0xae, 0xe2, 0xab, 0x64, 0xaa, 0xd2, 0x59,
	0x8c, 0x3f, 0xa4, 0x55, 0x8c, 0x7f, 0xaa, 0xa9, 0x87, 0xe8, 0xec, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x10, 0xe3, 0xb3, 0x61, 0xb9, 0x04, 0x00, 0x00,
}
