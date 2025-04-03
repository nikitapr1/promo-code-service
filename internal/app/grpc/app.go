package grpcapp

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	log        *zap.Logger
	port       int
	gRPCServer *grpc.Server
}
