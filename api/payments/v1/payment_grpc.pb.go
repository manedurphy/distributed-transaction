// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	AddCreditCard(ctx context.Context, in *AddCreditCardRequest, opts ...grpc.CallOption) (*AddCreditCardResponse, error)
	GetCreditCards(ctx context.Context, in *GetCreditCardsRequest, opts ...grpc.CallOption) (*GetCreditCardsResponse, error)
	AddFunds(ctx context.Context, in *AddFundsRequest, opts ...grpc.CallOption) (*AddFundsResponse, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) AddCreditCard(ctx context.Context, in *AddCreditCardRequest, opts ...grpc.CallOption) (*AddCreditCardResponse, error) {
	out := new(AddCreditCardResponse)
	err := c.cc.Invoke(ctx, "/payment.PaymentService/AddCreditCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetCreditCards(ctx context.Context, in *GetCreditCardsRequest, opts ...grpc.CallOption) (*GetCreditCardsResponse, error) {
	out := new(GetCreditCardsResponse)
	err := c.cc.Invoke(ctx, "/payment.PaymentService/GetCreditCards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) AddFunds(ctx context.Context, in *AddFundsRequest, opts ...grpc.CallOption) (*AddFundsResponse, error) {
	out := new(AddFundsResponse)
	err := c.cc.Invoke(ctx, "/payment.PaymentService/AddFunds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility
type PaymentServiceServer interface {
	AddCreditCard(context.Context, *AddCreditCardRequest) (*AddCreditCardResponse, error)
	GetCreditCards(context.Context, *GetCreditCardsRequest) (*GetCreditCardsResponse, error)
	AddFunds(context.Context, *AddFundsRequest) (*AddFundsResponse, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (UnimplementedPaymentServiceServer) AddCreditCard(context.Context, *AddCreditCardRequest) (*AddCreditCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCreditCard not implemented")
}
func (UnimplementedPaymentServiceServer) GetCreditCards(context.Context, *GetCreditCardsRequest) (*GetCreditCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreditCards not implemented")
}
func (UnimplementedPaymentServiceServer) AddFunds(context.Context, *AddFundsRequest) (*AddFundsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFunds not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_AddCreditCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCreditCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).AddCreditCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.PaymentService/AddCreditCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).AddCreditCard(ctx, req.(*AddCreditCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetCreditCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCreditCardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetCreditCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.PaymentService/GetCreditCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetCreditCards(ctx, req.(*GetCreditCardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_AddFunds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFundsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).AddFunds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.PaymentService/AddFunds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).AddFunds(ctx, req.(*AddFundsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCreditCard",
			Handler:    _PaymentService_AddCreditCard_Handler,
		},
		{
			MethodName: "GetCreditCards",
			Handler:    _PaymentService_GetCreditCards_Handler,
		},
		{
			MethodName: "AddFunds",
			Handler:    _PaymentService_AddFunds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment.proto",
}
