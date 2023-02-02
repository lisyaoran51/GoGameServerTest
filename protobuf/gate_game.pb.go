// This is Proto v3 syntax。

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protobuf/gate_game.proto

package protobuf

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

type GateLoginRequest_User_Level int32

const (
	GateLoginRequest_PLAYER GateLoginRequest_User_Level = 0
	GateLoginRequest_GUEST  GateLoginRequest_User_Level = 1
)

// Enum value maps for GateLoginRequest_User_Level.
var (
	GateLoginRequest_User_Level_name = map[int32]string{
		0: "PLAYER",
		1: "GUEST",
	}
	GateLoginRequest_User_Level_value = map[string]int32{
		"PLAYER": 0,
		"GUEST":  1,
	}
)

func (x GateLoginRequest_User_Level) Enum() *GateLoginRequest_User_Level {
	p := new(GateLoginRequest_User_Level)
	*p = x
	return p
}

func (x GateLoginRequest_User_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GateLoginRequest_User_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_gate_game_proto_enumTypes[0].Descriptor()
}

func (GateLoginRequest_User_Level) Type() protoreflect.EnumType {
	return &file_protobuf_gate_game_proto_enumTypes[0]
}

func (x GateLoginRequest_User_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GateLoginRequest_User_Level.Descriptor instead.
func (GateLoginRequest_User_Level) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_gate_game_proto_rawDescGZIP(), []int{0, 0}
}

//// GateLoginRequest Gate登入協議
type GateLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string                           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                                                       // 用戶帳號
	User_Level      GateLoginRequest_User_Level      `protobuf:"varint,2,opt,name=user_Level,json=userLevel,proto3,enum=protobuf.GateLoginRequest_User_Level" json:"user_Level,omitempty"` // 用户類型
	Productid       string                           `protobuf:"bytes,3,opt,name=productid,proto3" json:"productid,omitempty"`                                                             // 用戶產品編號
	FiatCurrency    string                           `protobuf:"bytes,4,opt,name=FiatCurrency,proto3" json:"FiatCurrency,omitempty"`                                                       // 用戶用的法幣
	VirtualCurrency string                           `protobuf:"bytes,5,opt,name=VirtualCurrency,proto3" json:"VirtualCurrency,omitempty"`                                                 // 用戶用的虛擬幣
	Currency        []*GateLoginRequest_CurrencyData `protobuf:"bytes,6,rep,name=currency,proto3" json:"currency,omitempty"`                                                               // 用戶身上的貨幣類別 可為空
	Token           string                           `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`                                                                     // 用戶Token
}

func (x *GateLoginRequest) Reset() {
	*x = GateLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_gate_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateLoginRequest) ProtoMessage() {}

func (x *GateLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_gate_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateLoginRequest.ProtoReflect.Descriptor instead.
func (*GateLoginRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_gate_game_proto_rawDescGZIP(), []int{0}
}

func (x *GateLoginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GateLoginRequest) GetUser_Level() GateLoginRequest_User_Level {
	if x != nil {
		return x.User_Level
	}
	return GateLoginRequest_PLAYER
}

func (x *GateLoginRequest) GetProductid() string {
	if x != nil {
		return x.Productid
	}
	return ""
}

func (x *GateLoginRequest) GetFiatCurrency() string {
	if x != nil {
		return x.FiatCurrency
	}
	return ""
}

func (x *GateLoginRequest) GetVirtualCurrency() string {
	if x != nil {
		return x.VirtualCurrency
	}
	return ""
}

func (x *GateLoginRequest) GetCurrency() []*GateLoginRequest_CurrencyData {
	if x != nil {
		return x.Currency
	}
	return nil
}

func (x *GateLoginRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// gateLoginResponse Gate登入協議回包
type GateLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code           uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`                    // 結果
	Username       string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`             // 使用者名稱
	CryptoCurrency string `protobuf:"bytes,3,opt,name=cryptoCurrency,proto3" json:"cryptoCurrency,omitempty"` // 使用虛擬貨幣
	FiatCurrency   string `protobuf:"bytes,4,opt,name=fiatCurrency,proto3" json:"fiatCurrency,omitempty"`     // 使用法幣 沒使用會是空字串
}

func (x *GateLoginResponse) Reset() {
	*x = GateLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_gate_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateLoginResponse) ProtoMessage() {}

func (x *GateLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_gate_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateLoginResponse.ProtoReflect.Descriptor instead.
func (*GateLoginResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_gate_game_proto_rawDescGZIP(), []int{1}
}

func (x *GateLoginResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GateLoginResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GateLoginResponse) GetCryptoCurrency() string {
	if x != nil {
		return x.CryptoCurrency
	}
	return ""
}

func (x *GateLoginResponse) GetFiatCurrency() string {
	if x != nil {
		return x.FiatCurrency
	}
	return ""
}

type GateLoginRequest_CurrencyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyType string `protobuf:"bytes,1,opt,name=currencyType,proto3" json:"currencyType,omitempty"` //貨幣種類
	Amount       string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`             //貨幣數量
}

func (x *GateLoginRequest_CurrencyData) Reset() {
	*x = GateLoginRequest_CurrencyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_gate_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateLoginRequest_CurrencyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateLoginRequest_CurrencyData) ProtoMessage() {}

func (x *GateLoginRequest_CurrencyData) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_gate_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateLoginRequest_CurrencyData.ProtoReflect.Descriptor instead.
func (*GateLoginRequest_CurrencyData) Descriptor() ([]byte, []int) {
	return file_protobuf_gate_game_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GateLoginRequest_CurrencyData) GetCurrencyType() string {
	if x != nil {
		return x.CurrencyType
	}
	return ""
}

func (x *GateLoginRequest_CurrencyData) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_protobuf_gate_game_proto protoreflect.FileDescriptor

var file_protobuf_gate_game_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x5f,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x22, 0xa4, 0x03, 0x0a, 0x10, 0x47, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x5f, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x46, 0x69, 0x61, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x46, 0x69, 0x61, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x43, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x4a, 0x0a, 0x0c, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x23, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x47, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x22, 0x8f, 0x01, 0x0a, 0x11,
	0x47, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x69, 0x61,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x66, 0x69, 0x61, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x32, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x73, 0x79,
	0x61, 0x6f, 0x72, 0x61, 0x6e, 0x35, 0x31, 0x2f, 0x47, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_gate_game_proto_rawDescOnce sync.Once
	file_protobuf_gate_game_proto_rawDescData = file_protobuf_gate_game_proto_rawDesc
)

func file_protobuf_gate_game_proto_rawDescGZIP() []byte {
	file_protobuf_gate_game_proto_rawDescOnce.Do(func() {
		file_protobuf_gate_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_gate_game_proto_rawDescData)
	})
	return file_protobuf_gate_game_proto_rawDescData
}

var file_protobuf_gate_game_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_gate_game_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protobuf_gate_game_proto_goTypes = []interface{}{
	(GateLoginRequest_User_Level)(0),      // 0: protobuf.GateLoginRequest.User_Level
	(*GateLoginRequest)(nil),              // 1: protobuf.GateLoginRequest
	(*GateLoginResponse)(nil),             // 2: protobuf.GateLoginResponse
	(*GateLoginRequest_CurrencyData)(nil), // 3: protobuf.GateLoginRequest.CurrencyData
}
var file_protobuf_gate_game_proto_depIdxs = []int32{
	0, // 0: protobuf.GateLoginRequest.user_Level:type_name -> protobuf.GateLoginRequest.User_Level
	3, // 1: protobuf.GateLoginRequest.currency:type_name -> protobuf.GateLoginRequest.CurrencyData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protobuf_gate_game_proto_init() }
func file_protobuf_gate_game_proto_init() {
	if File_protobuf_gate_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_gate_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateLoginRequest); i {
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
		file_protobuf_gate_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateLoginResponse); i {
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
		file_protobuf_gate_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateLoginRequest_CurrencyData); i {
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
			RawDescriptor: file_protobuf_gate_game_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_gate_game_proto_goTypes,
		DependencyIndexes: file_protobuf_gate_game_proto_depIdxs,
		EnumInfos:         file_protobuf_gate_game_proto_enumTypes,
		MessageInfos:      file_protobuf_gate_game_proto_msgTypes,
	}.Build()
	File_protobuf_gate_game_proto = out.File
	file_protobuf_gate_game_proto_rawDesc = nil
	file_protobuf_gate_game_proto_goTypes = nil
	file_protobuf_gate_game_proto_depIdxs = nil
}
