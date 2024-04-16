// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.23.4
// source: trade.proto

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

// Message Trading Channel state definition
type TradeState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId     []byte            `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	State         *State            `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	StateHash     []byte            `protobuf:"bytes,3,opt,name=state_hash,json=stateHash,proto3" json:"state_hash,omitempty"`
	StateHashSigs []*Signature      `protobuf:"bytes,4,rep,name=state_hash_sigs,json=stateHashSigs,proto3" json:"state_hash_sigs,omitempty"`
	Safety        *Decimal          `protobuf:"bytes,5,opt,name=safety,proto3" json:"safety,omitempty"`
	MarginCall    *SignedMarginCall `protobuf:"bytes,6,opt,name=margin_call,json=marginCall,proto3" json:"margin_call,omitempty"`
}

func (x *TradeState) Reset() {
	*x = TradeState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeState) ProtoMessage() {}

func (x *TradeState) ProtoReflect() protoreflect.Message {
	mi := &file_trade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeState.ProtoReflect.Descriptor instead.
func (*TradeState) Descriptor() ([]byte, []int) {
	return file_trade_proto_rawDescGZIP(), []int{0}
}

func (x *TradeState) GetChannelId() []byte {
	if x != nil {
		return x.ChannelId
	}
	return nil
}

func (x *TradeState) GetState() *State {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *TradeState) GetStateHash() []byte {
	if x != nil {
		return x.StateHash
	}
	return nil
}

func (x *TradeState) GetStateHashSigs() []*Signature {
	if x != nil {
		return x.StateHashSigs
	}
	return nil
}

func (x *TradeState) GetSafety() *Decimal {
	if x != nil {
		return x.Safety
	}
	return nil
}

func (x *TradeState) GetMarginCall() *SignedMarginCall {
	if x != nil {
		return x.MarginCall
	}
	return nil
}

var File_trade_proto protoreflect.FileDescriptor

var file_trade_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x02, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x37, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x5f, 0x73, 0x69, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x53, 0x69, 0x67, 0x73, 0x12,
	0x25, 0x0a, 0x06, 0x73, 0x61, 0x66, 0x65, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x06,
	0x73, 0x61, 0x66, 0x65, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x0b, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e,
	0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x0a, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x32,
	0xee, 0x01, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x23, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x0b, 0x2e, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x0b, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x66, 0x75, 0x6e,
	0x64, 0x12, 0x0b, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x0b,
	0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x43,
	0x61, 0x6c, 0x6c, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x2d, 0x0a, 0x11, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0b, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x0b, 0x2e,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x46, 0x69,
	0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x12, 0x0b, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x1a, 0x0b, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2d, 0x33, 0x2f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trade_proto_rawDescOnce sync.Once
	file_trade_proto_rawDescData = file_trade_proto_rawDesc
)

func file_trade_proto_rawDescGZIP() []byte {
	file_trade_proto_rawDescOnce.Do(func() {
		file_trade_proto_rawDescData = protoimpl.X.CompressGZIP(file_trade_proto_rawDescData)
	})
	return file_trade_proto_rawDescData
}

var file_trade_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_trade_proto_goTypes = []interface{}{
	(*TradeState)(nil),       // 0: TradeState
	(*State)(nil),            // 1: core.State
	(*Signature)(nil),        // 2: core.Signature
	(*Decimal)(nil),          // 3: core.Decimal
	(*SignedMarginCall)(nil), // 4: core.SignedMarginCall
}
var file_trade_proto_depIdxs = []int32{
	1, // 0: TradeState.state:type_name -> core.State
	2, // 1: TradeState.state_hash_sigs:type_name -> core.Signature
	3, // 2: TradeState.safety:type_name -> core.Decimal
	4, // 3: TradeState.margin_call:type_name -> core.SignedMarginCall
	0, // 4: TradeChannel.Prefund:input_type -> TradeState
	0, // 5: TradeChannel.Postfund:input_type -> TradeState
	4, // 6: TradeChannel.UpdateMargin:input_type -> core.SignedMarginCall
	0, // 7: TradeChannel.RequestSettlement:input_type -> TradeState
	0, // 8: TradeChannel.Finalize:input_type -> TradeState
	0, // 9: TradeChannel.Prefund:output_type -> TradeState
	0, // 10: TradeChannel.Postfund:output_type -> TradeState
	4, // 11: TradeChannel.UpdateMargin:output_type -> core.SignedMarginCall
	0, // 12: TradeChannel.RequestSettlement:output_type -> TradeState
	0, // 13: TradeChannel.Finalize:output_type -> TradeState
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_trade_proto_init() }
func file_trade_proto_init() {
	if File_trade_proto != nil {
		return
	}
	file_core_proto_init()
	file_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_trade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeState); i {
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
			RawDescriptor: file_trade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trade_proto_goTypes,
		DependencyIndexes: file_trade_proto_depIdxs,
		MessageInfos:      file_trade_proto_msgTypes,
	}.Build()
	File_trade_proto = out.File
	file_trade_proto_rawDesc = nil
	file_trade_proto_goTypes = nil
	file_trade_proto_depIdxs = nil
}
