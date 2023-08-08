package main

import (
	"log"
	"net"

	// commonpb "github.com/preslavmihaylov/go-grpc-crash-course/gen/common"
	"github.com/lyfeyvutha/grpcATCbank/gen/transfer_statements"
	"google.golang.org/grpc"
)

var (
	paymentStatementsAddr = "localhost:10001"
)

func main() {
	grpcServer, lis := setupPaymentStatementsServer()

	log.Println("Successfully started transfer_statements grpc server...")
	grpcServer.Serve(lis)
}

// TODO: Setup the payment statements grpc server
func setupPaymentStatementsServer() (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", paymentStatementsAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	transfer_statements.RegisterTransferStatementsServer(grpcServer, &server{})

	return grpcServer, lis
	//panic("not implemented")
}

// TODO: Implement the payment_statements.PaymentStatementsServer interface
type server struct{}

func (s *server) CreateStatement(stream transfer_statements.TransferStatements_CreateStatementServer) error {
	panic("not implemented")
}
