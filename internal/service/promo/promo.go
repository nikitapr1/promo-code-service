package promoservice

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"promo-code-service/internal/domain/dto/promo"
)

type Service struct {
	log           *zap.Logger
	promoProvider promoProvider
	promoSaver    promoSaver
}

func New(
	log *zap.Logger,
	promoProvider promoProvider,
	promoSaver promoSaver,
) *Service {
	return &Service{
		log:           log,
		promoProvider: promoProvider,
		promoSaver:    promoSaver,
	}
}

func (s *Service) Create(ctx context.Context, companyId string, promoDto promo.CreatePromoDTO) (id string, err error) {
	const op = "service.promo.Create"

	id, err = s.promoSaver.Create(ctx, companyId, promoDto)

	if err != nil {
		s.log.Error("Failed to save promo to db", zap.Error(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
