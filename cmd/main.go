package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"promo-code-service/internal/config"
	promogrpc "promo-code-service/internal/grpc"
	"promo-code-service/internal/grpc/handler/create_promo"
	promoservice "promo-code-service/internal/service/promo"
	"promo-code-service/internal/storage/repository"
	promopb "promo-code-service/pkg/api/pb"
	"syscall"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {

	const op = "cmd.Main"

	cfg := config.MustLoad()

	log, err := setupLogger(cfg.Env)
	if err != nil {
		panic(fmt.Errorf("%s: failed to setup logger: %s", op, err))
	}

	log.Info("Starting Promo Service", zap.Int("port", cfg.GRPCPort))

	storage, err := repository.New(cfg.DatabaseUrl)
	if err != nil {
		panic(fmt.Errorf("%s: failed to connect to database:\n%s", op, err))
	}

	promoService := promoservice.New(log, storage, storage)

	createPromoHandler := create_promo.New(promoService)

	server := grpc.NewServer()

	serverAPI := promogrpc.New(promoService, createPromoHandler)

	promopb.RegisterPromoServiceServer(server, serverAPI)

	go mustRunGRPCServer(server, cfg.GRPCPort)

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("Stopping application...\n", zap.String("signal", sign.String()))

	server.GracefulStop()

	log.Info("Application stopped")
}

func setupLogger(env string) (*zap.Logger, error) {

	var log *zap.Logger
	var err error

	switch env {
	case envProd:
		log, err = zap.NewProduction()
	case envLocal:
		log, err = zap.NewDevelopment()
	default:
		return nil, fmt.Errorf("unknown environment %q", env)
	}
	return log, err
}

func mustRunGRPCServer(server *grpc.Server, port int) {

	const op = "main.RunGRPCServer"

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Errorf("%s: failed to listen: %s", op, err))
	}

	err = server.Serve(lis)
	if err != nil {
		panic(fmt.Errorf("%s: failed to serve: %s", op, err))
	}

}
