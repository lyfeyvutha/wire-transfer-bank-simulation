package main

import (
	"context"
	"fmt"
	"log"
	"net"

	bankpb "github.com/lyfeyvutha/grpcATCbank/gen/bank"

	commonpb "github.com/lyfeyvutha/grpcATCbank/gen/common"
	"github.com/lyfeyvutha/grpcATCbank/gen/transfer_statements"
	"google.golang.org/grpc"
)

type userID string

var (
	//tokensPerDollar         = int32(5)
	bankAddr                 = "localhost:10000"
	transferStatementsAddr   = "localhost:10001"
	transferStatementsClient transfer_statements.TransferStatementsClient
)

func main() {
	var conn *grpc.ClientConn
	transferStatementsClient, conn = setupTransferStatementsClient()
	defer conn.Close()

	log.Println("Successfully connected to transfer_statements...")

	grpcServer, lis := setupBankServer()
	grpcServer.Serve(lis)
}

// TODO: Setup the casino grpc server
func setupBankServer() (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", bankAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	bankpb.RegisterBankServer(grpcServer, newBankServer())

	return grpcServer, lis
	//panic("not implemented")
}

// TODO: Setup the payment_statements.PaymentStatementsClient
func setupTransferStatementsClient() (transfer_statements.TransferStatementsClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(transferStatementsAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return transfer_statements.NewTransferStatementsClient(conn), conn
	//panic("not implemented")
}

func newBankServer() *bankServer {
	return &bankServer{
		stockPrice:     100,
		userToCash:     map[userID]int32{},
		userToTransfer: map[userID][]int32{},
		userToStocks:   map[userID]int32{},
	}
}

// TODO: Implement the casinopb.CasinoServer interface
type bankServer struct {
	stockPrice     int32
	userToCash     map[userID]int32
	userToTransfer map[userID][]int32
	userToStocks   map[userID]int32
}

func (c *bankServer) Deposit(ctx context.Context, transfer *commonpb.Transfer) (*bankpb.Cash, error) {
	log.Printf("Deposit invoked with transfer %v\n", transfer)

	usrID := userID(transfer.User.GetId())
	cash := transfer.GetAmount() //* tokensPerDollar
	//panic("not implemented")

	c.userToTransfer[usrID] = append(c.userToTransfer[usrID],
		-transfer.Amount)
	c.userToCash[usrID] += cash

	return &bankpb.Cash{Count: cash}, nil

}

func (c *bankServer) Withdraw(ctx context.Context, withdrawReq *bankpb.WithdrawRequest) (*commonpb.Transfer, error) {
	toWithdraw := withdrawReq.GetCashCnt()
	log.Printf("Withdraw invoked with %v\n", toWithdraw)

	usrID := userID(withdrawReq.User.GetId())
	log.Println(c.userToCash[usrID])
	if !c.hasEnoughCash(usrID, toWithdraw) {
		return nil, fmt.Errorf("not enough cash to withdraw")
	}

	amount := toWithdraw // / tokensPerDollar
	c.userToCash[usrID] -= toWithdraw
	c.userToTransfer[usrID] = append(c.userToTransfer[usrID], amount)

	return &commonpb.Transfer{User: withdrawReq.User, Amount: amount}, nil
}

func (c *bankServer) GetCashBalance(ctx context.Context, user *commonpb.User) (*bankpb.Cash, error) {
	log.Printf("GetCashBalance invoked with user %v\n", user)

	usrID := userID(user.GetId())
	return &bankpb.Cash{Count: c.userToCash[usrID]}, nil
}

func (c *bankServer) GetTransfers(user *commonpb.User, stream bankpb.Bank_GetTransfersServer) error {
	panic("not implemented")
}

func (c *bankServer) GetTransferStatement(ctx context.Context, user *commonpb.User) (*commonpb.TransferStatement, error) {
	panic("not implemented")
}

func (c *bankServer) Invest(stream bankpb.Bank_InvestServer) error {
	panic("not implemented")
}
