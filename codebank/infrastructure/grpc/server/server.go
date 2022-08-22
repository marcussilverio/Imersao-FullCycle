package server

import (
	"github.com/codeedu/marcussilverio/infrastructure/grpc/pb"
	"github.com/codeedu/marcussilverio/infrastructure/grpc/service"
	"github.com/codeedu/marcussilverio/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GRPCServer struct {
	ProcesTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer() GRPCServer {
	return GRPCServer {}
}

func (g GRPCServer) Serve(){
	lis, err := net.Listen(network: "tcp", address: "0.0.0.0:50051")
	if err != nil{
		log.Fatalf("Could not listen tcp port")
	}
	transactionService := service.NewTransactionService()
	transactionService.ProcesTransactionUseCase = g.ProcesTransactionUseCase

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterPaymentServiceServer(grpcServer, transactionService)
	grpcServer.Serve(lis)
}