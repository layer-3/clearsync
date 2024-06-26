// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: session_key.proto

package proto

import (
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

type ParamCondition int32

const (
	ParamCondition_PARAM_CONDITION_UNSPECIFIED           ParamCondition = 0
	ParamCondition_PARAM_CONDITION_EQUAL                 ParamCondition = 1
	ParamCondition_PARAM_CONDITION_GREATER_THAN          ParamCondition = 2
	ParamCondition_PARAM_CONDITION_LESS_THAN             ParamCondition = 3
	ParamCondition_PARAM_CONDITION_GREATER_THAN_OR_EQUAL ParamCondition = 4
	ParamCondition_PARAM_CONDITION_LESS_THAN_OR_EQUAL    ParamCondition = 5
	ParamCondition_PARAM_CONDITION_NOT_EQUAL             ParamCondition = 6
)

// Enum value maps for ParamCondition.
var (
	ParamCondition_name = map[int32]string{
		0: "PARAM_CONDITION_UNSPECIFIED",
		1: "PARAM_CONDITION_EQUAL",
		2: "PARAM_CONDITION_GREATER_THAN",
		3: "PARAM_CONDITION_LESS_THAN",
		4: "PARAM_CONDITION_GREATER_THAN_OR_EQUAL",
		5: "PARAM_CONDITION_LESS_THAN_OR_EQUAL",
		6: "PARAM_CONDITION_NOT_EQUAL",
	}
	ParamCondition_value = map[string]int32{
		"PARAM_CONDITION_UNSPECIFIED":           0,
		"PARAM_CONDITION_EQUAL":                 1,
		"PARAM_CONDITION_GREATER_THAN":          2,
		"PARAM_CONDITION_LESS_THAN":             3,
		"PARAM_CONDITION_GREATER_THAN_OR_EQUAL": 4,
		"PARAM_CONDITION_LESS_THAN_OR_EQUAL":    5,
		"PARAM_CONDITION_NOT_EQUAL":             6,
	}
)

func (x ParamCondition) Enum() *ParamCondition {
	p := new(ParamCondition)
	*p = x
	return p
}

func (x ParamCondition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParamCondition) Descriptor() protoreflect.EnumDescriptor {
	return file_session_key_proto_enumTypes[0].Descriptor()
}

func (ParamCondition) Type() protoreflect.EnumType {
	return &file_session_key_proto_enumTypes[0]
}

func (x ParamCondition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParamCondition.Descriptor instead.
func (ParamCondition) EnumDescriptor() ([]byte, []int) {
	return file_session_key_proto_rawDescGZIP(), []int{0}
}

type IncompleteUserOp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender               *Address `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Nonce                *BigInt  `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	InitCode             []byte   `protobuf:"bytes,3,opt,name=init_code,json=initCode,proto3" json:"init_code,omitempty"`
	CallData             []byte   `protobuf:"bytes,4,opt,name=call_data,json=callData,proto3" json:"call_data,omitempty"`
	CallGasLimit         *BigInt  `protobuf:"bytes,5,opt,name=call_gas_limit,json=callGasLimit,proto3" json:"call_gas_limit,omitempty"`
	VerificationGasLimit *BigInt  `protobuf:"bytes,6,opt,name=verification_gas_limit,json=verificationGasLimit,proto3" json:"verification_gas_limit,omitempty"`
	PreVerificationGas   *BigInt  `protobuf:"bytes,7,opt,name=pre_verification_gas,json=preVerificationGas,proto3" json:"pre_verification_gas,omitempty"`
	MaxFeePerGas         *BigInt  `protobuf:"bytes,8,opt,name=max_fee_per_gas,json=maxFeePerGas,proto3" json:"max_fee_per_gas,omitempty"`
	MaxPriorityFeePerGas *BigInt  `protobuf:"bytes,9,opt,name=max_priority_fee_per_gas,json=maxPriorityFeePerGas,proto3" json:"max_priority_fee_per_gas,omitempty"`
	PaymasterAndData     []byte   `protobuf:"bytes,10,opt,name=paymaster_and_data,json=paymasterAndData,proto3" json:"paymaster_and_data,omitempty"`
	Signature            []byte   `protobuf:"bytes,11,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *IncompleteUserOp) Reset() {
	*x = IncompleteUserOp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncompleteUserOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncompleteUserOp) ProtoMessage() {}

func (x *IncompleteUserOp) ProtoReflect() protoreflect.Message {
	mi := &file_session_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncompleteUserOp.ProtoReflect.Descriptor instead.
func (*IncompleteUserOp) Descriptor() ([]byte, []int) {
	return file_session_key_proto_rawDescGZIP(), []int{0}
}

func (x *IncompleteUserOp) GetSender() *Address {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *IncompleteUserOp) GetNonce() *BigInt {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *IncompleteUserOp) GetInitCode() []byte {
	if x != nil {
		return x.InitCode
	}
	return nil
}

func (x *IncompleteUserOp) GetCallData() []byte {
	if x != nil {
		return x.CallData
	}
	return nil
}

func (x *IncompleteUserOp) GetCallGasLimit() *BigInt {
	if x != nil {
		return x.CallGasLimit
	}
	return nil
}

func (x *IncompleteUserOp) GetVerificationGasLimit() *BigInt {
	if x != nil {
		return x.VerificationGasLimit
	}
	return nil
}

func (x *IncompleteUserOp) GetPreVerificationGas() *BigInt {
	if x != nil {
		return x.PreVerificationGas
	}
	return nil
}

func (x *IncompleteUserOp) GetMaxFeePerGas() *BigInt {
	if x != nil {
		return x.MaxFeePerGas
	}
	return nil
}

func (x *IncompleteUserOp) GetMaxPriorityFeePerGas() *BigInt {
	if x != nil {
		return x.MaxPriorityFeePerGas
	}
	return nil
}

func (x *IncompleteUserOp) GetPaymasterAndData() []byte {
	if x != nil {
		return x.PaymasterAndData
	}
	return nil
}

func (x *IncompleteUserOp) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type SessionKeyPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target      *Address     `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	FunctionAbi []byte       `protobuf:"bytes,2,opt,name=function_abi,json=functionAbi,proto3" json:"function_abi,omitempty"`
	ValueLimit  *BigInt      `protobuf:"bytes,3,opt,name=value_limit,json=valueLimit,proto3" json:"value_limit,omitempty"`
	ParamRules  []*ParamRule `protobuf:"bytes,4,rep,name=param_rules,json=paramRules,proto3" json:"param_rules,omitempty"`
}

func (x *SessionKeyPermission) Reset() {
	*x = SessionKeyPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionKeyPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionKeyPermission) ProtoMessage() {}

func (x *SessionKeyPermission) ProtoReflect() protoreflect.Message {
	mi := &file_session_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionKeyPermission.ProtoReflect.Descriptor instead.
func (*SessionKeyPermission) Descriptor() ([]byte, []int) {
	return file_session_key_proto_rawDescGZIP(), []int{1}
}

func (x *SessionKeyPermission) GetTarget() *Address {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *SessionKeyPermission) GetFunctionAbi() []byte {
	if x != nil {
		return x.FunctionAbi
	}
	return nil
}

func (x *SessionKeyPermission) GetValueLimit() *BigInt {
	if x != nil {
		return x.ValueLimit
	}
	return nil
}

func (x *SessionKeyPermission) GetParamRules() []*ParamRule {
	if x != nil {
		return x.ParamRules
	}
	return nil
}

type ParamRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Condition ParamCondition `protobuf:"varint,1,opt,name=condition,proto3,enum=ParamCondition" json:"condition,omitempty"`
	Param     string         `protobuf:"bytes,2,opt,name=param,proto3" json:"param,omitempty"`
}

func (x *ParamRule) Reset() {
	*x = ParamRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_key_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamRule) ProtoMessage() {}

func (x *ParamRule) ProtoReflect() protoreflect.Message {
	mi := &file_session_key_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamRule.ProtoReflect.Descriptor instead.
func (*ParamRule) Descriptor() ([]byte, []int) {
	return file_session_key_proto_rawDescGZIP(), []int{2}
}

func (x *ParamRule) GetCondition() ParamCondition {
	if x != nil {
		return x.Condition
	}
	return ParamCondition_PARAM_CONDITION_UNSPECIFIED
}

func (x *ParamRule) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

var File_session_key_proto protoreflect.FileDescriptor

var file_session_key_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x96, 0x04, 0x0a, 0x10, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x4f, 0x70, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x69, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x63, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x0e, 0x63, 0x61, 0x6c,
	0x6c, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52,
	0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x47, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x42, 0x0a,
	0x16, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x61,
	0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x14, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x3e, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x12, 0x70,
	0x72, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x61,
	0x73, 0x12, 0x33, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x65, 0x72,
	0x5f, 0x67, 0x61, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x46, 0x65, 0x65,
	0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x12, 0x44, 0x0a, 0x18, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x67,
	0x61, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x12, 0x2c, 0x0a, 0x12,
	0x70, 0x61, 0x79, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x70, 0x61, 0x79, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x62, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x62, 0x69, 0x12, 0x2d, 0x0a, 0x0b, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x0a,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2b, 0x0a, 0x0b, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x50, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x52, 0x75, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x2a, 0xff, 0x01, 0x0a, 0x0e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x1b,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a,
	0x15, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x41, 0x52, 0x41,
	0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x47, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x41,
	0x52, 0x41, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45,
	0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10, 0x03, 0x12, 0x29, 0x0a, 0x25, 0x50, 0x41, 0x52,
	0x41, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x47, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x4f, 0x52, 0x5f, 0x45, 0x51, 0x55,
	0x41, 0x4c, 0x10, 0x04, 0x12, 0x26, 0x0a, 0x22, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x43, 0x4f,
	0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41,
	0x4e, 0x5f, 0x4f, 0x52, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x06, 0x42, 0x24, 0x5a, 0x22, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2d,
	0x33, 0x2f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_session_key_proto_rawDescOnce sync.Once
	file_session_key_proto_rawDescData = file_session_key_proto_rawDesc
)

func file_session_key_proto_rawDescGZIP() []byte {
	file_session_key_proto_rawDescOnce.Do(func() {
		file_session_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_session_key_proto_rawDescData)
	})
	return file_session_key_proto_rawDescData
}

var file_session_key_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_session_key_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_session_key_proto_goTypes = []interface{}{
	(ParamCondition)(0),          // 0: ParamCondition
	(*IncompleteUserOp)(nil),     // 1: IncompleteUserOp
	(*SessionKeyPermission)(nil), // 2: SessionKeyPermission
	(*ParamRule)(nil),            // 3: ParamRule
	(*Address)(nil),              // 4: core.Address
	(*BigInt)(nil),               // 5: core.BigInt
}
var file_session_key_proto_depIdxs = []int32{
	4,  // 0: IncompleteUserOp.sender:type_name -> core.Address
	5,  // 1: IncompleteUserOp.nonce:type_name -> core.BigInt
	5,  // 2: IncompleteUserOp.call_gas_limit:type_name -> core.BigInt
	5,  // 3: IncompleteUserOp.verification_gas_limit:type_name -> core.BigInt
	5,  // 4: IncompleteUserOp.pre_verification_gas:type_name -> core.BigInt
	5,  // 5: IncompleteUserOp.max_fee_per_gas:type_name -> core.BigInt
	5,  // 6: IncompleteUserOp.max_priority_fee_per_gas:type_name -> core.BigInt
	4,  // 7: SessionKeyPermission.target:type_name -> core.Address
	5,  // 8: SessionKeyPermission.value_limit:type_name -> core.BigInt
	3,  // 9: SessionKeyPermission.param_rules:type_name -> ParamRule
	0,  // 10: ParamRule.condition:type_name -> ParamCondition
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_session_key_proto_init() }
func file_session_key_proto_init() {
	if File_session_key_proto != nil {
		return
	}
	file_core_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_session_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncompleteUserOp); i {
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
		file_session_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionKeyPermission); i {
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
		file_session_key_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamRule); i {
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
			RawDescriptor: file_session_key_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_session_key_proto_goTypes,
		DependencyIndexes: file_session_key_proto_depIdxs,
		EnumInfos:         file_session_key_proto_enumTypes,
		MessageInfos:      file_session_key_proto_msgTypes,
	}.Build()
	File_session_key_proto = out.File
	file_session_key_proto_rawDesc = nil
	file_session_key_proto_goTypes = nil
	file_session_key_proto_depIdxs = nil
}
