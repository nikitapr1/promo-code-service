package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"ya-promo-code-service/internal/app"
	"ya-promo-code-service/internal/config"
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

	//application := app.New(log, cfg.GRPCPort, cfg.DatabaseUrl)
	//
	//go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Info("Stopping application...\n", zap.String("signal", sign.String()))

	application.GRPCSrv.Stop()

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

func runGRPCServer(log *zap.Logger, server *grpc.Server) error {
	const op = "grpcapp.Run"

	log = log.With(zap.String("op", op))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", zap.String("Address", lis.Addr().String()))

	if err := server.Serve(lis); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
