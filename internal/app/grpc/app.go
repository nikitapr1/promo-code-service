package grpcapp

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	promogrpc "ya-promo-code-service/internal/grpc"
	"ya-promo-code-service/internal/grpc/interceptor"
)

type App struct {
	log        *zap.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *zap.Logger,
	promoService promogrpc.PromoService,
	port int,
) *App {

	gRPCServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.LogRequestUnaryInterceptor(log)),
	)

	promogrpc.Register(gRPCServer, promoService)

	return &App{log: log, gRPCServer: gRPCServer, port: port}
}

func (a *App) MustRun() {

	if err := a.Run(); err != nil {
		panic(err)
	}

}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(zap.String("op", op))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", zap.String("Address", lis.Addr().String()))

	if err := a.gRPCServer.Serve(lis); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *App) Stop() {

	const op = "grpcapp.Stop"

	a.log.With(zap.String("op", op)).Info("Stopping grpc server")

	a.gRPCServer.GracefulStop()
}
