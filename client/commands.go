package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	bankpb "github.com/lyfeyvutha/grpcATCbank/gen/bank"

	commonpb "github.com/lyfeyvutha/grpcATCbank/gen/common"
	"google.golang.org/grpc"
)

type command string

const bankAddr = "localhost:10000"

var errStopGambling = errors.New("user exits gambling session")

// TODO: Setup the casinopb.CasinoClient
func setupClient() (bankpb.BankClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial("localhost:10000", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return bankpb.NewBankClient(conn), conn
	//panic("not implemented")
}

func buyTokens(tokensCnt int) (string, error) {
	res, err := client.Deposit(context.Background(), &commonpb.Transfer{
		User: &commonpb.User{
			Id: username,
		},
		Amount: int32(tokensCnt),
	})
	if err != nil {
		return "", fmt.Errorf("couldn't deposit: %w", err)
	}

	return fmt.Sprintf("Successfully deposited $%v", res.GetCount()), nil
	//panic("not implemented")
}

func withdraw(tokensCnt int) (string, error) {
  res, err := client.Withdraw(context.Background(), &bankpb.WithdrawRequest{
    User:      &commonpb.User{Id: username},
    CashCnt: int32(tokensCnt),
  })
  if err != nil {
    return "", fmt.Errorf("couldn't withdraw: %w", err)
  }

  return fmt.Sprintf("Successfully withdrew $%d", res.GetAmount()), nil
}


func tokenBalance() (string, error) {
	res, err := client.GetCashBalance(context.Background(), &commonpb.User{
		Id: username,
	})
	if err != nil {
		return "", fmt.Errorf("couldn't get balance: %w", err)
	}

	return fmt.Sprintf("Your balance is $%v.", res.GetCount()), nil
}

func payments() (string, error) {
	panic("not implemented")
}

func paymentStatement() (string, error) {
	panic("not implemented")
}

func gamble() (string, error) {
	panic("not implemented")
}
