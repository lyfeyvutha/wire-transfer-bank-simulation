#!/bin/bash

mkdir -p gen/common
mkdir -p gen/bank
mkdir -p gen/transfer_statements

protoc -I idl --go_out=gen --go-grpc_out=require_unimplemented_servers=false:gen --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative idl/common/types.proto
protoc -I idl --go_out=gen --go-grpc_out=require_unimplemented_servers=false:gen --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative idl/bank/bank.proto
protoc -I idl --go_out=gen --go-grpc_out=require_unimplemented_servers=false:gen --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative idl/transfer_statements/transfer_statements.proto

