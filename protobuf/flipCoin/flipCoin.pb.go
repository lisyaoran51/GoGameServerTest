// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: flipCoin/flipCoin.proto

package flipCoin

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

// 共通標頭檔
//
// 傳送與接收的前置標頭檔，用以告知後方內容該如何解析。
type GameMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*GameMessage_BetReq
	//	*GameMessage_BetRes
	Payload isGameMessage_Payload `protobuf_oneof:"payload"`
}

func (x *GameMessage) Reset() {
	*x = GameMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flipCoin_flipCoin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameMessage) ProtoMessage() {}

func (x *GameMessage) ProtoReflect() protoreflect.Message {
	mi := &file_flipCoin_flipCoin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameMessage.ProtoReflect.Descriptor instead.
func (*GameMessage) Descriptor() ([]byte, []int) {
	return file_flipCoin_flipCoin_proto_rawDescGZIP(), []int{0}
}

func (m *GameMessage) GetPayload() isGameMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *GameMessage) GetBetReq() *BetReq {
	if x, ok := x.GetPayload().(*GameMessage_BetReq); ok {
		return x.BetReq
	}
	return nil
}

func (x *GameMessage) GetBetRes() *BetRes {
	if x, ok := x.GetPayload().(*GameMessage_BetRes); ok {
		return x.BetRes
	}
	return nil
}

type isGameMessage_Payload interface {
	isGameMessage_Payload()
}

type GameMessage_BetReq struct {
	BetReq *BetReq `protobuf:"bytes,1,opt,name=betReq,proto3,oneof"`
}

type GameMessage_BetRes struct {
	BetRes *BetRes `protobuf:"bytes,2,opt,name=betRes,proto3,oneof"`
}

func (*GameMessage_BetReq) isGameMessage_Payload() {}

func (*GameMessage_BetRes) isGameMessage_Payload() {}

// 玩家請求下注
type BetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tableid   string `protobuf:"bytes,1,opt,name=tableid,proto3" json:"tableid,omitempty"`     // table 資訊
	Currency  string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`   // 幣種
	Betamount string `protobuf:"bytes,3,opt,name=betamount,proto3" json:"betamount,omitempty"` // 下注金額 小數最小只能到0.00000001
	Uuid      string `protobuf:"bytes,4,opt,name=uuid,proto3" json:"uuid,omitempty"`           // UUID
}

func (x *BetReq) Reset() {
	*x = BetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flipCoin_flipCoin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BetReq) ProtoMessage() {}

func (x *BetReq) ProtoReflect() protoreflect.Message {
	mi := &file_flipCoin_flipCoin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BetReq.ProtoReflect.Descriptor instead.
func (*BetReq) Descriptor() ([]byte, []int) {
	return file_flipCoin_flipCoin_proto_rawDescGZIP(), []int{1}
}

func (x *BetReq) GetTableid() string {
	if x != nil {
		return x.Tableid
	}
	return ""
}

func (x *BetReq) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *BetReq) GetBetamount() string {
	if x != nil {
		return x.Betamount
	}
	return ""
}

func (x *BetReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 下注回應
type BetRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tableid   string `protobuf:"bytes,1,opt,name=tableid,proto3" json:"tableid,omitempty"`     // table 資訊
	Betamount string `protobuf:"bytes,2,opt,name=betamount,proto3" json:"betamount,omitempty"` // 下注金額 小數最小只能到0.00000001
	Code      uint32 `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`          // 下注結果
	Bettime   uint64 `protobuf:"varint,4,opt,name=bettime,proto3" json:"bettime,omitempty"`    // 下注時間
	Win       string `protobuf:"bytes,5,opt,name=win,proto3" json:"win,omitempty"`             // 派彩
	Ratio     string `protobuf:"bytes,6,opt,name=ratio,proto3" json:"ratio,omitempty"`         // 派彩倍率
	Res       string `protobuf:"bytes,7,opt,name=res,proto3" json:"res,omitempty"`             // 寶石結果
	Uuid      string `protobuf:"bytes,8,opt,name=uuid,proto3" json:"uuid,omitempty"`           // UUID
}

func (x *BetRes) Reset() {
	*x = BetRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flipCoin_flipCoin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BetRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BetRes) ProtoMessage() {}

func (x *BetRes) ProtoReflect() protoreflect.Message {
	mi := &file_flipCoin_flipCoin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BetRes.ProtoReflect.Descriptor instead.
func (*BetRes) Descriptor() ([]byte, []int) {
	return file_flipCoin_flipCoin_proto_rawDescGZIP(), []int{2}
}

func (x *BetRes) GetTableid() string {
	if x != nil {
		return x.Tableid
	}
	return ""
}

func (x *BetRes) GetBetamount() string {
	if x != nil {
		return x.Betamount
	}
	return ""
}

func (x *BetRes) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BetRes) GetBettime() uint64 {
	if x != nil {
		return x.Bettime
	}
	return 0
}

func (x *BetRes) GetWin() string {
	if x != nil {
		return x.Win
	}
	return ""
}

func (x *BetRes) GetRatio() string {
	if x != nil {
		return x.Ratio
	}
	return ""
}

func (x *BetRes) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

func (x *BetRes) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_flipCoin_flipCoin_proto protoreflect.FileDescriptor

var file_flipCoin_flipCoin_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x6c, 0x69, 0x70, 0x43, 0x6f, 0x69, 0x6e, 0x2f, 0x66, 0x6c, 0x69, 0x70, 0x43,
	0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x66, 0x6c, 0x69, 0x70, 0x43,
	0x6f, 0x69, 0x6e, 0x22, 0x70, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x62, 0x65, 0x74, 0x52, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x43, 0x6f, 0x69, 0x6e, 0x2e, 0x42, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x48, 0x00, 0x52, 0x06, 0x62, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2a,
	0x0a, 0x06, 0x62, 0x65, 0x74, 0x52, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x66, 0x6c, 0x69, 0x70, 0x43, 0x6f, 0x69, 0x6e, 0x2e, 0x42, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x48, 0x00, 0x52, 0x06, 0x62, 0x65, 0x74, 0x52, 0x65, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x70, 0x0a, 0x06, 0x42, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x65, 0x74, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x65, 0x74, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0xbc, 0x01, 0x0a, 0x06, 0x42, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x62, 0x65, 0x74, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x65, 0x74, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x65, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x62, 0x65, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x69, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x73, 0x79, 0x61, 0x6f, 0x72, 0x61, 0x6e, 0x35, 0x31,
	0x2f, 0x47, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x65, 0x73,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6c, 0x69, 0x70, 0x43,
	0x6f, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flipCoin_flipCoin_proto_rawDescOnce sync.Once
	file_flipCoin_flipCoin_proto_rawDescData = file_flipCoin_flipCoin_proto_rawDesc
)

func file_flipCoin_flipCoin_proto_rawDescGZIP() []byte {
	file_flipCoin_flipCoin_proto_rawDescOnce.Do(func() {
		file_flipCoin_flipCoin_proto_rawDescData = protoimpl.X.CompressGZIP(file_flipCoin_flipCoin_proto_rawDescData)
	})
	return file_flipCoin_flipCoin_proto_rawDescData
}

var file_flipCoin_flipCoin_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_flipCoin_flipCoin_proto_goTypes = []interface{}{
	(*GameMessage)(nil), // 0: flipCoin.GameMessage
	(*BetReq)(nil),      // 1: flipCoin.BetReq
	(*BetRes)(nil),      // 2: flipCoin.BetRes
}
var file_flipCoin_flipCoin_proto_depIdxs = []int32{
	1, // 0: flipCoin.GameMessage.betReq:type_name -> flipCoin.BetReq
	2, // 1: flipCoin.GameMessage.betRes:type_name -> flipCoin.BetRes
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_flipCoin_flipCoin_proto_init() }
func file_flipCoin_flipCoin_proto_init() {
	if File_flipCoin_flipCoin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flipCoin_flipCoin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameMessage); i {
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
		file_flipCoin_flipCoin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BetReq); i {
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
		file_flipCoin_flipCoin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BetRes); i {
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
	file_flipCoin_flipCoin_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GameMessage_BetReq)(nil),
		(*GameMessage_BetRes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flipCoin_flipCoin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flipCoin_flipCoin_proto_goTypes,
		DependencyIndexes: file_flipCoin_flipCoin_proto_depIdxs,
		MessageInfos:      file_flipCoin_flipCoin_proto_msgTypes,
	}.Build()
	File_flipCoin_flipCoin_proto = out.File
	file_flipCoin_flipCoin_proto_rawDesc = nil
	file_flipCoin_flipCoin_proto_goTypes = nil
	file_flipCoin_flipCoin_proto_depIdxs = nil
}