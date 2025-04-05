package promoservice

import (
	"context"
	"go.uber.org/zap"
	"promo-code-service/internal/domain/dto/promo"
)

type Service struct {
	log           *zap.Logger
	promoProvider PromoProvider
	promoSaver    PromoSaver
}

type PromoProvider interface {
	//TODO: описать
}

type PromoSaver interface {
	// TODO: описать

}

func New(
	log *zap.Logger,
	promoProvider PromoProvider,
	promoSaver PromoSaver,
) *Service {
	return &Service{
		log:           log,
		promoProvider: promoProvider,
		promoSaver:    promoSaver,
	}
}

func (p *Service) Create(ctx context.Context, companyId string, promoDto promo.CreatePromoDTO) {
	//TODO implement me
	panic("implement me")
}
