package app

import (
	"fmt"
	"go.uber.org/zap"
	grpcapp "ya-promo-code-service/internal/app/grpc"
	"ya-promo-code-service/internal/service/promo"
	"ya-promo-code-service/internal/storage/postgres"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *zap.Logger,
	grpcPort int,
	storagePath string,
) *App {
	// TODO: postgres storage
	// TODO: promo service

	const op = "app.New"

	storage, err := postgres.New(storagePath)
	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

	promoService := promoservice.New(log, storage, storage)

	grpcSrv := grpcapp.New(log, promoService, grpcPort)

	return &App{
		GRPCSrv: grpcSrv,
	}
}
