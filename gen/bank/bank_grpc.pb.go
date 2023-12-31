// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: bank/bank.proto

package bank

import (
	context "context"
	common "github.com/lyfeyvutha/grpcATCbank/gen/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Bank_Deposit_FullMethodName              = "/bank.Bank/Deposit"
	Bank_Withdraw_FullMethodName             = "/bank.Bank/Withdraw"
	Bank_GetCashBalance_FullMethodName       = "/bank.Bank/GetCashBalance"
	Bank_GetTransfers_FullMethodName         = "/bank.Bank/GetTransfers"
	Bank_GetTransferStatement_FullMethodName = "/bank.Bank/GetTransferStatement"
	Bank_Invest_FullMethodName               = "/bank.Bank/Invest"
)

// BankClient is the client API for Bank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankClient interface {
	Deposit(ctx context.Context, in *common.Transfer, opts ...grpc.CallOption) (*Cash, error)
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*common.Transfer, error)
	GetCashBalance(ctx context.Context, in *common.User, opts ...grpc.CallOption) (*Cash, error)
	GetTransfers(ctx context.Context, in *common.User, opts ...grpc.CallOption) (Bank_GetTransfersClient, error)
	GetTransferStatement(ctx context.Context, in *common.User, opts ...grpc.CallOption) (*common.TransferStatement, error)
	Invest(ctx context.Context, opts ...grpc.CallOption) (Bank_InvestClient, error)
}

type bankClient struct {
	cc grpc.ClientConnInterface
}

func NewBankClient(cc grpc.ClientConnInterface) BankClient {
	return &bankClient{cc}
}

func (c *bankClient) Deposit(ctx context.Context, in *common.Transfer, opts ...grpc.CallOption) (*Cash, error) {
	out := new(Cash)
	err := c.cc.Invoke(ctx, Bank_Deposit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*common.Transfer, error) {
	out := new(common.Transfer)
	err := c.cc.Invoke(ctx, Bank_Withdraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) GetCashBalance(ctx context.Context, in *common.User, opts ...grpc.CallOption) (*Cash, error) {
	out := new(Cash)
	err := c.cc.Invoke(ctx, Bank_GetCashBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) GetTransfers(ctx context.Context, in *common.User, opts ...grpc.CallOption) (Bank_GetTransfersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Bank_ServiceDesc.Streams[0], Bank_GetTransfers_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bankGetTransfersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Bank_GetTransfersClient interface {
	Recv() (*common.Transfer, error)
	grpc.ClientStream
}

type bankGetTransfersClient struct {
	grpc.ClientStream
}

func (x *bankGetTransfersClient) Recv() (*common.Transfer, error) {
	m := new(common.Transfer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bankClient) GetTransferStatement(ctx context.Context, in *common.User, opts ...grpc.CallOption) (*common.TransferStatement, error) {
	out := new(common.TransferStatement)
	err := c.cc.Invoke(ctx, Bank_GetTransferStatement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) Invest(ctx context.Context, opts ...grpc.CallOption) (Bank_InvestClient, error) {
	stream, err := c.cc.NewStream(ctx, &Bank_ServiceDesc.Streams[1], Bank_Invest_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bankInvestClient{stream}
	return x, nil
}

type Bank_InvestClient interface {
	Send(*Action) error
	Recv() (*InvestInfo, error)
	grpc.ClientStream
}

type bankInvestClient struct {
	grpc.ClientStream
}

func (x *bankInvestClient) Send(m *Action) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bankInvestClient) Recv() (*InvestInfo, error) {
	m := new(InvestInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BankServer is the server API for Bank service.
// All implementations should embed UnimplementedBankServer
// for forward compatibility
type BankServer interface {
	Deposit(context.Context, *common.Transfer) (*Cash, error)
	Withdraw(context.Context, *WithdrawRequest) (*common.Transfer, error)
	GetCashBalance(context.Context, *common.User) (*Cash, error)
	GetTransfers(*common.User, Bank_GetTransfersServer) error
	GetTransferStatement(context.Context, *common.User) (*common.TransferStatement, error)
	Invest(Bank_InvestServer) error
}

// UnimplementedBankServer should be embedded to have forward compatible implementations.
type UnimplementedBankServer struct {
}

func (UnimplementedBankServer) Deposit(context.Context, *common.Transfer) (*Cash, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedBankServer) Withdraw(context.Context, *WithdrawRequest) (*common.Transfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedBankServer) GetCashBalance(context.Context, *common.User) (*Cash, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCashBalance not implemented")
}
func (UnimplementedBankServer) GetTransfers(*common.User, Bank_GetTransfersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTransfers not implemented")
}
func (UnimplementedBankServer) GetTransferStatement(context.Context, *common.User) (*common.TransferStatement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransferStatement not implemented")
}
func (UnimplementedBankServer) Invest(Bank_InvestServer) error {
	return status.Errorf(codes.Unimplemented, "method Invest not implemented")
}

// UnsafeBankServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankServer will
// result in compilation errors.
type UnsafeBankServer interface {
	mustEmbedUnimplementedBankServer()
}

func RegisterBankServer(s grpc.ServiceRegistrar, srv BankServer) {
	s.RegisterService(&Bank_ServiceDesc, srv)
}

func _Bank_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Transfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_Deposit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).Deposit(ctx, req.(*common.Transfer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_Withdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).Withdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_GetCashBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).GetCashBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_GetCashBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).GetCashBalance(ctx, req.(*common.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_GetTransfers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(common.User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BankServer).GetTransfers(m, &bankGetTransfersServer{stream})
}

type Bank_GetTransfersServer interface {
	Send(*common.Transfer) error
	grpc.ServerStream
}

type bankGetTransfersServer struct {
	grpc.ServerStream
}

func (x *bankGetTransfersServer) Send(m *common.Transfer) error {
	return x.ServerStream.SendMsg(m)
}

func _Bank_GetTransferStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).GetTransferStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_GetTransferStatement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).GetTransferStatement(ctx, req.(*common.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_Invest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BankServer).Invest(&bankInvestServer{stream})
}

type Bank_InvestServer interface {
	Send(*InvestInfo) error
	Recv() (*Action, error)
	grpc.ServerStream
}

type bankInvestServer struct {
	grpc.ServerStream
}

func (x *bankInvestServer) Send(m *InvestInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bankInvestServer) Recv() (*Action, error) {
	m := new(Action)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Bank_ServiceDesc is the grpc.ServiceDesc for Bank service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bank_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bank.Bank",
	HandlerType: (*BankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deposit",
			Handler:    _Bank_Deposit_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _Bank_Withdraw_Handler,
		},
		{
			MethodName: "GetCashBalance",
			Handler:    _Bank_GetCashBalance_Handler,
		},
		{
			MethodName: "GetTransferStatement",
			Handler:    _Bank_GetTransferStatement_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTransfers",
			Handler:       _Bank_GetTransfers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Invest",
			Handler:       _Bank_Invest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bank/bank.proto",
}
