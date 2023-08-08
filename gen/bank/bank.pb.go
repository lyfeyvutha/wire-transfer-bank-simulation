// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: bank/bank.proto

package bank

import (
	common "github.com/lyfeyvutha/grpcATCbank/gen/common"
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

type ActionType int32

const (
	ActionType_BUY  ActionType = 0
	ActionType_SELL ActionType = 1
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0: "BUY",
		1: "SELL",
	}
	ActionType_value = map[string]int32{
		"BUY":  0,
		"SELL": 1,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_bank_bank_proto_enumTypes[0].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_bank_bank_proto_enumTypes[0]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{0}
}

type InvestType int32

const (
	InvestType_STOCK_INFO    InvestType = 0
	InvestType_ACTION_RESULT InvestType = 1
)

// Enum value maps for InvestType.
var (
	InvestType_name = map[int32]string{
		0: "STOCK_INFO",
		1: "ACTION_RESULT",
	}
	InvestType_value = map[string]int32{
		"STOCK_INFO":    0,
		"ACTION_RESULT": 1,
	}
)

func (x InvestType) Enum() *InvestType {
	p := new(InvestType)
	*p = x
	return p
}

func (x InvestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvestType) Descriptor() protoreflect.EnumDescriptor {
	return file_bank_bank_proto_enumTypes[1].Descriptor()
}

func (InvestType) Type() protoreflect.EnumType {
	return &file_bank_bank_proto_enumTypes[1]
}

func (x InvestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvestType.Descriptor instead.
func (InvestType) EnumDescriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{1}
}

type WithdrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    *common.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	CashCnt int32        `protobuf:"varint,2,opt,name=cashCnt,proto3" json:"cashCnt,omitempty"`
}

func (x *WithdrawRequest) Reset() {
	*x = WithdrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawRequest) ProtoMessage() {}

func (x *WithdrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawRequest.ProtoReflect.Descriptor instead.
func (*WithdrawRequest) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{0}
}

func (x *WithdrawRequest) GetUser() *common.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *WithdrawRequest) GetCashCnt() int32 {
	if x != nil {
		return x.CashCnt
	}
	return 0
}

type Cash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Cash) Reset() {
	*x = Cash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cash) ProtoMessage() {}

func (x *Cash) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cash.ProtoReflect.Descriptor instead.
func (*Cash) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{1}
}

func (x *Cash) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        *common.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Type        ActionType   `protobuf:"varint,2,opt,name=type,proto3,enum=bank.ActionType" json:"type,omitempty"`
	StocksCount int32        `protobuf:"varint,3,opt,name=stocksCount,proto3" json:"stocksCount,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{2}
}

func (x *Action) GetUser() *common.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Action) GetType() ActionType {
	if x != nil {
		return x.Type
	}
	return ActionType_BUY
}

func (x *Action) GetStocksCount() int32 {
	if x != nil {
		return x.StocksCount
	}
	return 0
}

type StockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *StockInfo) Reset() {
	*x = StockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockInfo) ProtoMessage() {}

func (x *StockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockInfo.ProtoReflect.Descriptor instead.
func (*StockInfo) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{3}
}

func (x *StockInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StockInfo) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ActionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ActionResult) Reset() {
	*x = ActionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionResult) ProtoMessage() {}

func (x *ActionResult) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionResult.ProtoReflect.Descriptor instead.
func (*ActionResult) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{4}
}

func (x *ActionResult) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type InvestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   InvestType    `protobuf:"varint,1,opt,name=type,proto3,enum=bank.InvestType" json:"type,omitempty"`
	Info   *StockInfo    `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Result *ActionResult `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *InvestInfo) Reset() {
	*x = InvestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_bank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvestInfo) ProtoMessage() {}

func (x *InvestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bank_bank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvestInfo.ProtoReflect.Descriptor instead.
func (*InvestInfo) Descriptor() ([]byte, []int) {
	return file_bank_bank_proto_rawDescGZIP(), []int{5}
}

func (x *InvestInfo) GetType() InvestType {
	if x != nil {
		return x.Type
	}
	return InvestType_STOCK_INFO
}

func (x *InvestInfo) GetInfo() *StockInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *InvestInfo) GetResult() *ActionResult {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_bank_bank_proto protoreflect.FileDescriptor

var file_bank_bank_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x0f, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x61, 0x73, 0x68, 0x43, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x63, 0x61, 0x73, 0x68, 0x43, 0x6e, 0x74, 0x22, 0x1c, 0x0a, 0x04, 0x43, 0x61, 0x73,
	0x68, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x71, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x09, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x22, 0x20, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x83, 0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x1f, 0x0a, 0x0a, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x55, 0x59, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x01, 0x2a, 0x2f, 0x0a, 0x0a, 0x49, 0x6e,
	0x76, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x4f, 0x43,
	0x4b, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x01, 0x32, 0xb6, 0x02, 0x0a, 0x04,
	0x42, 0x61, 0x6e, 0x6b, 0x12, 0x28, 0x0a, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12,
	0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x1a, 0x0a, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x08, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x15, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x68, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x0a, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x73, 0x12, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0f,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0b, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x12, 0x0c,
	0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x65, 0x79, 0x76, 0x75, 0x74, 0x68, 0x61, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x41, 0x54, 0x43, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x61,
	0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bank_bank_proto_rawDescOnce sync.Once
	file_bank_bank_proto_rawDescData = file_bank_bank_proto_rawDesc
)

func file_bank_bank_proto_rawDescGZIP() []byte {
	file_bank_bank_proto_rawDescOnce.Do(func() {
		file_bank_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_bank_bank_proto_rawDescData)
	})
	return file_bank_bank_proto_rawDescData
}

var file_bank_bank_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bank_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bank_bank_proto_goTypes = []interface{}{
	(ActionType)(0),                  // 0: bank.ActionType
	(InvestType)(0),                  // 1: bank.InvestType
	(*WithdrawRequest)(nil),          // 2: bank.WithdrawRequest
	(*Cash)(nil),                     // 3: bank.Cash
	(*Action)(nil),                   // 4: bank.Action
	(*StockInfo)(nil),                // 5: bank.StockInfo
	(*ActionResult)(nil),             // 6: bank.ActionResult
	(*InvestInfo)(nil),               // 7: bank.InvestInfo
	(*common.User)(nil),              // 8: types.User
	(*common.Transfer)(nil),          // 9: types.Transfer
	(*common.TransferStatement)(nil), // 10: types.TransferStatement
}
var file_bank_bank_proto_depIdxs = []int32{
	8,  // 0: bank.WithdrawRequest.user:type_name -> types.User
	8,  // 1: bank.Action.user:type_name -> types.User
	0,  // 2: bank.Action.type:type_name -> bank.ActionType
	1,  // 3: bank.InvestInfo.type:type_name -> bank.InvestType
	5,  // 4: bank.InvestInfo.info:type_name -> bank.StockInfo
	6,  // 5: bank.InvestInfo.result:type_name -> bank.ActionResult
	9,  // 6: bank.Bank.Deposit:input_type -> types.Transfer
	2,  // 7: bank.Bank.Withdraw:input_type -> bank.WithdrawRequest
	8,  // 8: bank.Bank.GetCashBalance:input_type -> types.User
	8,  // 9: bank.Bank.GetTransfers:input_type -> types.User
	8,  // 10: bank.Bank.GetTransferStatement:input_type -> types.User
	4,  // 11: bank.Bank.Invest:input_type -> bank.Action
	3,  // 12: bank.Bank.Deposit:output_type -> bank.Cash
	9,  // 13: bank.Bank.Withdraw:output_type -> types.Transfer
	3,  // 14: bank.Bank.GetCashBalance:output_type -> bank.Cash
	9,  // 15: bank.Bank.GetTransfers:output_type -> types.Transfer
	10, // 16: bank.Bank.GetTransferStatement:output_type -> types.TransferStatement
	7,  // 17: bank.Bank.Invest:output_type -> bank.InvestInfo
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_bank_bank_proto_init() }
func file_bank_bank_proto_init() {
	if File_bank_bank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bank_bank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawRequest); i {
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
		file_bank_bank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cash); i {
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
		file_bank_bank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_bank_bank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockInfo); i {
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
		file_bank_bank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionResult); i {
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
		file_bank_bank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvestInfo); i {
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
			RawDescriptor: file_bank_bank_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bank_bank_proto_goTypes,
		DependencyIndexes: file_bank_bank_proto_depIdxs,
		EnumInfos:         file_bank_bank_proto_enumTypes,
		MessageInfos:      file_bank_bank_proto_msgTypes,
	}.Build()
	File_bank_bank_proto = out.File
	file_bank_bank_proto_rawDesc = nil
	file_bank_bank_proto_goTypes = nil
	file_bank_bank_proto_depIdxs = nil
}