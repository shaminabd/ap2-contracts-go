package apiv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Amount        int64                  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	mi := &file_orderpayment_v1_order_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderpayment_v1_order_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_orderpayment_v1_order_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *PaymentRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type PaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId       string                 `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	TransactionId string                 `protobuf:"bytes,3,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Amount        int64                  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentResponse) Reset() {
	*x = PaymentResponse{}
	mi := &file_orderpayment_v1_order_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentResponse) ProtoMessage() {}

func (x *PaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderpayment_v1_order_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return file_orderpayment_v1_order_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PaymentResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *PaymentResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *PaymentResponse) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PaymentResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PaymentResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type OrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	mi := &file_orderpayment_v1_order_payment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderpayment_v1_order_payment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_orderpayment_v1_order_payment_proto_rawDescGZIP(), []int{2}
}

func (x *OrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type OrderStatusUpdate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderStatusUpdate) Reset() {
	*x = OrderStatusUpdate{}
	mi := &file_orderpayment_v1_order_payment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderStatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatusUpdate) ProtoMessage() {}

func (x *OrderStatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_orderpayment_v1_order_payment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*OrderStatusUpdate) Descriptor() ([]byte, []int) {
	return file_orderpayment_v1_order_payment_proto_rawDescGZIP(), []int{3}
}

func (x *OrderStatusUpdate) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderStatusUpdate) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderStatusUpdate) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_orderpayment_v1_order_payment_proto protoreflect.FileDescriptor

const file_orderpayment_v1_order_payment_proto_rawDesc = "" +
	"\n" +
	"#orderpayment/v1/order_payment.proto\x12\x0forderpayment.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"C\n" +
	"\x0ePaymentRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x16\n" +
	"\x06amount\x18\x02 \x01(\x03R\x06amount\"\xce\x01\n" +
	"\x0fPaymentResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x19\n" +
	"\border_id\x18\x02 \x01(\tR\aorderId\x12%\n" +
	"\x0etransaction_id\x18\x03 \x01(\tR\rtransactionId\x12\x16\n" +
	"\x06amount\x18\x04 \x01(\x03R\x06amount\x12\x16\n" +
	"\x06status\x18\x05 \x01(\tR\x06status\x129\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\")\n" +
	"\fOrderRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\"\x81\x01\n" +
	"\x11OrderStatusUpdate\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\x129\n" +
	"\n" +
	"updated_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt2e\n" +
	"\x0ePaymentService\x12S\n" +
	"\x0eProcessPayment\x12\x1f.orderpayment.v1.PaymentRequest\x1a .orderpayment.v1.PaymentResponse2t\n" +
	"\x12OrderUpdateService\x12^\n" +
	"\x17SubscribeToOrderUpdates\x12\x1d.orderpayment.v1.OrderRequest\x1a\".orderpayment.v1.OrderStatusUpdate0\x01B3Z1github.com/shaminabd/ap2-contracts-go/apiv1;apiv1b\x06proto3"

var (
	file_orderpayment_v1_order_payment_proto_rawDescOnce sync.Once
	file_orderpayment_v1_order_payment_proto_rawDescData []byte
)

func file_orderpayment_v1_order_payment_proto_rawDescGZIP() []byte {
	file_orderpayment_v1_order_payment_proto_rawDescOnce.Do(func() {
		file_orderpayment_v1_order_payment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_orderpayment_v1_order_payment_proto_rawDesc), len(file_orderpayment_v1_order_payment_proto_rawDesc)))
	})
	return file_orderpayment_v1_order_payment_proto_rawDescData
}

var file_orderpayment_v1_order_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_orderpayment_v1_order_payment_proto_goTypes = []any{
	(*PaymentRequest)(nil),
	(*PaymentResponse)(nil),
	(*OrderRequest)(nil),
	(*OrderStatusUpdate)(nil),
	(*timestamppb.Timestamp)(nil),
}
var file_orderpayment_v1_order_payment_proto_depIdxs = []int32{
	4,
	4,
	0,
	2,
	1,
	3,
	4,
	2,
	2,
	2,
	0,
}

func init() { file_orderpayment_v1_order_payment_proto_init() }
func file_orderpayment_v1_order_payment_proto_init() {
	if File_orderpayment_v1_order_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_orderpayment_v1_order_payment_proto_rawDesc), len(file_orderpayment_v1_order_payment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_orderpayment_v1_order_payment_proto_goTypes,
		DependencyIndexes: file_orderpayment_v1_order_payment_proto_depIdxs,
		MessageInfos:      file_orderpayment_v1_order_payment_proto_msgTypes,
	}.Build()
	File_orderpayment_v1_order_payment_proto = out.File
	file_orderpayment_v1_order_payment_proto_goTypes = nil
	file_orderpayment_v1_order_payment_proto_depIdxs = nil
}
