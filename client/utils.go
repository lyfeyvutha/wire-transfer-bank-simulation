package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	bankpb "github.com/lyfeyvutha/grpcATCbank/gen/bank"

	commonpb "github.com/lyfeyvutha/grpcATCbank/gen/common"
)

func promptUserForAction() (*bankpb.Action, bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.Trim(input, "\n ")
	tokens := regexp.MustCompile("[ ]+").Split(input, -1)
	cmd, args := tokens[0], tokens[1:]

	var actionType bankpb.ActionType
	switch cmd {
	case "":
		return nil, false
	case "exit":
		fallthrough
	case "stop":
		return nil, true
	case "buy":
		actionType = bankpb.ActionType_BUY
	case "sell":
		actionType = bankpb.ActionType_SELL

	default:
		fmt.Println("unknown action")
		return nil, false
	}

	if len(args) != 1 {
		fmt.Println("bad arguments")
		return nil, false
	}

	amount, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("error while reading input: %v\n", err)
		return nil, false
	}

	return &bankpb.Action{
		User:        &commonpb.User{Id: username},
		Type:        actionType,
		StocksCount: int32(amount),
	}, false
}

func paymentHistoryString(payments []*commonpb.Transfer) string {
	res := ""
	for i, payment := range payments {
		res += fmt.Sprintf("\tPayment %d: %d$\n", i, payment.Amount)
	}

	return res
}
