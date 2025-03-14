// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: stock.proto

package stock

import (
	context "context"
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

type GetStocksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductIds []uint32 `protobuf:"varint,1,rep,packed,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
}

func (x *GetStocksReq) Reset() {
	*x = GetStocksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStocksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStocksReq) ProtoMessage() {}

func (x *GetStocksReq) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStocksReq.ProtoReflect.Descriptor instead.
func (*GetStocksReq) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{0}
}

func (x *GetStocksReq) GetProductIds() []uint32 {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

type Stock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint32 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  uint32 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Stock) Reset() {
	*x = Stock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stock) ProtoMessage() {}

func (x *Stock) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stock.ProtoReflect.Descriptor instead.
func (*Stock) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{1}
}

func (x *Stock) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Stock) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type GetStocksResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stocks map[uint32]*Stock `protobuf:"bytes,1,rep,name=stocks,proto3" json:"stocks,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetStocksResp) Reset() {
	*x = GetStocksResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStocksResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStocksResp) ProtoMessage() {}

func (x *GetStocksResp) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStocksResp.ProtoReflect.Descriptor instead.
func (*GetStocksResp) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{2}
}

func (x *GetStocksResp) GetStocks() map[uint32]*Stock {
	if x != nil {
		return x.Stocks
	}
	return nil
}

type DeductStocksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stocks  []*Stock `protobuf:"bytes,1,rep,name=stocks,proto3" json:"stocks,omitempty"` // 需要扣减的商品列表
	OrderId string   `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *DeductStocksReq) Reset() {
	*x = DeductStocksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeductStocksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeductStocksReq) ProtoMessage() {}

func (x *DeductStocksReq) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeductStocksReq.ProtoReflect.Descriptor instead.
func (*DeductStocksReq) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{3}
}

func (x *DeductStocksReq) GetStocks() []*Stock {
	if x != nil {
		return x.Stocks
	}
	return nil
}

func (x *DeductStocksReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type DeductStocksResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeductStocksResp) Reset() {
	*x = DeductStocksResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeductStocksResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeductStocksResp) ProtoMessage() {}

func (x *DeductStocksResp) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeductStocksResp.ProtoReflect.Descriptor instead.
func (*DeductStocksResp) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{4}
}

func (x *DeductStocksResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type OccupyStocksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stocks  []*Stock `protobuf:"bytes,1,rep,name=stocks,proto3" json:"stocks,omitempty"` // 需要扣减的商品列表
	OrderId string   `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *OccupyStocksReq) Reset() {
	*x = OccupyStocksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupyStocksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupyStocksReq) ProtoMessage() {}

func (x *OccupyStocksReq) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupyStocksReq.ProtoReflect.Descriptor instead.
func (*OccupyStocksReq) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{5}
}

func (x *OccupyStocksReq) GetStocks() []*Stock {
	if x != nil {
		return x.Stocks
	}
	return nil
}

func (x *OccupyStocksReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type OccupyStocksResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *OccupyStocksResp) Reset() {
	*x = OccupyStocksResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OccupyStocksResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OccupyStocksResp) ProtoMessage() {}

func (x *OccupyStocksResp) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OccupyStocksResp.ProtoReflect.Descriptor instead.
func (*OccupyStocksResp) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{6}
}

func (x *OccupyStocksResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RecoverStocksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stocks  []*Stock `protobuf:"bytes,1,rep,name=stocks,proto3" json:"stocks,omitempty"`
	OrderId string   `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *RecoverStocksReq) Reset() {
	*x = RecoverStocksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecoverStocksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecoverStocksReq) ProtoMessage() {}

func (x *RecoverStocksReq) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecoverStocksReq.ProtoReflect.Descriptor instead.
func (*RecoverStocksReq) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{7}
}

func (x *RecoverStocksReq) GetStocks() []*Stock {
	if x != nil {
		return x.Stocks
	}
	return nil
}

func (x *RecoverStocksReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type RecoverStocksResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RecoverStocksResp) Reset() {
	*x = RecoverStocksResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecoverStocksResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecoverStocksResp) ProtoMessage() {}

func (x *RecoverStocksResp) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecoverStocksResp.ProtoReflect.Descriptor instead.
func (*RecoverStocksResp) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{8}
}

func (x *RecoverStocksResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_stock_proto protoreflect.FileDescriptor

var file_stock_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x22, 0x2f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x73, 0x22, 0x42, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x92, 0x01, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x06, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x73, 0x1a, 0x47, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x52,
	0x0a, 0x0f, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x2c, 0x0a, 0x10, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x52, 0x0a, 0x0f, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x79, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x10, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x79, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x53, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x99, 0x02, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0c, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x44,
	0x65, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x17,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0c, 0x4f, 0x63, 0x63,
	0x75, 0x70, 0x79, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x79, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x79,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d,
	0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x17, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x75, 0x75, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x6e,
	0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stock_proto_rawDescOnce sync.Once
	file_stock_proto_rawDescData = file_stock_proto_rawDesc
)

func file_stock_proto_rawDescGZIP() []byte {
	file_stock_proto_rawDescOnce.Do(func() {
		file_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_stock_proto_rawDescData)
	})
	return file_stock_proto_rawDescData
}

var file_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_stock_proto_goTypes = []interface{}{
	(*GetStocksReq)(nil),      // 0: stock.GetStocksReq
	(*Stock)(nil),             // 1: stock.Stock
	(*GetStocksResp)(nil),     // 2: stock.GetStocksResp
	(*DeductStocksReq)(nil),   // 3: stock.DeductStocksReq
	(*DeductStocksResp)(nil),  // 4: stock.DeductStocksResp
	(*OccupyStocksReq)(nil),   // 5: stock.OccupyStocksReq
	(*OccupyStocksResp)(nil),  // 6: stock.OccupyStocksResp
	(*RecoverStocksReq)(nil),  // 7: stock.RecoverStocksReq
	(*RecoverStocksResp)(nil), // 8: stock.RecoverStocksResp
	nil,                       // 9: stock.GetStocksResp.StocksEntry
}
var file_stock_proto_depIdxs = []int32{
	9, // 0: stock.GetStocksResp.stocks:type_name -> stock.GetStocksResp.StocksEntry
	1, // 1: stock.DeductStocksReq.stocks:type_name -> stock.Stock
	1, // 2: stock.OccupyStocksReq.stocks:type_name -> stock.Stock
	1, // 3: stock.RecoverStocksReq.stocks:type_name -> stock.Stock
	1, // 4: stock.GetStocksResp.StocksEntry.value:type_name -> stock.Stock
	0, // 5: stock.StockService.BatchGetStocks:input_type -> stock.GetStocksReq
	3, // 6: stock.StockService.DeductStocks:input_type -> stock.DeductStocksReq
	5, // 7: stock.StockService.OccupyStocks:input_type -> stock.OccupyStocksReq
	7, // 8: stock.StockService.RecoverStocks:input_type -> stock.RecoverStocksReq
	2, // 9: stock.StockService.BatchGetStocks:output_type -> stock.GetStocksResp
	4, // 10: stock.StockService.DeductStocks:output_type -> stock.DeductStocksResp
	6, // 11: stock.StockService.OccupyStocks:output_type -> stock.OccupyStocksResp
	8, // 12: stock.StockService.RecoverStocks:output_type -> stock.RecoverStocksResp
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_stock_proto_init() }
func file_stock_proto_init() {
	if File_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStocksReq); i {
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
		file_stock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stock); i {
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
		file_stock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStocksResp); i {
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
		file_stock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeductStocksReq); i {
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
		file_stock_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeductStocksResp); i {
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
		file_stock_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupyStocksReq); i {
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
		file_stock_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OccupyStocksResp); i {
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
		file_stock_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoverStocksReq); i {
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
		file_stock_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoverStocksResp); i {
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
			RawDescriptor: file_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stock_proto_goTypes,
		DependencyIndexes: file_stock_proto_depIdxs,
		MessageInfos:      file_stock_proto_msgTypes,
	}.Build()
	File_stock_proto = out.File
	file_stock_proto_rawDesc = nil
	file_stock_proto_goTypes = nil
	file_stock_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.1. DO NOT EDIT.

type StockService interface {
	BatchGetStocks(ctx context.Context, req *GetStocksReq) (res *GetStocksResp, err error)
	DeductStocks(ctx context.Context, req *DeductStocksReq) (res *DeductStocksResp, err error)
	OccupyStocks(ctx context.Context, req *OccupyStocksReq) (res *OccupyStocksResp, err error)
	RecoverStocks(ctx context.Context, req *RecoverStocksReq) (res *RecoverStocksResp, err error)
}
