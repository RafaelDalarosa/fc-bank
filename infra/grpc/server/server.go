package server

import (
	"log"
	"net"

	"github.com/RafaelDalarosa/fc-bank/domain/usecase"
	"github.com/RafaelDalarosa/fc-bank/infra/grpc/pb"
	"github.com/RafaelDalarosa/fc-bank/infra/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

func (g GRPCServer) Server() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("could not listen tcp port")
	}
	transactionService := service.NewTransactionService()
	transactionService.ProcessTransactinUseCase = g.ProcessTransactionUseCase
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterPaymentServiceServer(grpcServer, transactionService)
	grpcServer.Serve(lis)
}
