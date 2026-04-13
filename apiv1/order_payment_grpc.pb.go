package apiv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	PaymentService_ProcessPayment_FullMethodName = "/orderpayment.v1.PaymentService/ProcessPayment"
)

type PaymentServiceClient interface {
	ProcessPayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) ProcessPayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_ProcessPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type PaymentServiceServer interface {
	ProcessPayment(context.Context, *PaymentRequest) (*PaymentResponse, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

type UnimplementedPaymentServiceServer struct{}

func (UnimplementedPaymentServiceServer) ProcessPayment(context.Context, *PaymentRequest) (*PaymentResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method ProcessPayment not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}
func (UnimplementedPaymentServiceServer) testEmbeddedByValue()                        {}

type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_ProcessPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).ProcessPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_ProcessPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).ProcessPayment(ctx, req.(*PaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orderpayment.v1.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessPayment",
			Handler:    _PaymentService_ProcessPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orderpayment/v1/order_payment.proto",
}

const (
	OrderUpdateService_SubscribeToOrderUpdates_FullMethodName = "/orderpayment.v1.OrderUpdateService/SubscribeToOrderUpdates"
)

type OrderUpdateServiceClient interface {
	SubscribeToOrderUpdates(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OrderStatusUpdate], error)
}

type orderUpdateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderUpdateServiceClient(cc grpc.ClientConnInterface) OrderUpdateServiceClient {
	return &orderUpdateServiceClient{cc}
}

func (c *orderUpdateServiceClient) SubscribeToOrderUpdates(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OrderStatusUpdate], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OrderUpdateService_ServiceDesc.Streams[0], OrderUpdateService_SubscribeToOrderUpdates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[OrderRequest, OrderStatusUpdate]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderUpdateService_SubscribeToOrderUpdatesClient = grpc.ServerStreamingClient[OrderStatusUpdate]

type OrderUpdateServiceServer interface {
	SubscribeToOrderUpdates(*OrderRequest, grpc.ServerStreamingServer[OrderStatusUpdate]) error
	mustEmbedUnimplementedOrderUpdateServiceServer()
}

type UnimplementedOrderUpdateServiceServer struct{}

func (UnimplementedOrderUpdateServiceServer) SubscribeToOrderUpdates(*OrderRequest, grpc.ServerStreamingServer[OrderStatusUpdate]) error {
	return status.Error(codes.Unimplemented, "method SubscribeToOrderUpdates not implemented")
}
func (UnimplementedOrderUpdateServiceServer) mustEmbedUnimplementedOrderUpdateServiceServer() {}
func (UnimplementedOrderUpdateServiceServer) testEmbeddedByValue()                            {}

type UnsafeOrderUpdateServiceServer interface {
	mustEmbedUnimplementedOrderUpdateServiceServer()
}

func RegisterOrderUpdateServiceServer(s grpc.ServiceRegistrar, srv OrderUpdateServiceServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderUpdateService_ServiceDesc, srv)
}

func _OrderUpdateService_SubscribeToOrderUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrderRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderUpdateServiceServer).SubscribeToOrderUpdates(m, &grpc.GenericServerStream[OrderRequest, OrderStatusUpdate]{ServerStream: stream})
}

type OrderUpdateService_SubscribeToOrderUpdatesServer = grpc.ServerStreamingServer[OrderStatusUpdate]

var OrderUpdateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orderpayment.v1.OrderUpdateService",
	HandlerType: (*OrderUpdateServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToOrderUpdates",
			Handler:       _OrderUpdateService_SubscribeToOrderUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "orderpayment/v1/order_payment.proto",
}
