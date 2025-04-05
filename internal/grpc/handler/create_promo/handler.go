package create_promo

import (
	"context"
	promopb "ya-promo-code-service/pkg/api/pb"
)

type Handler struct {
	promoService promoService
}

func New(promoService promoService) *Handler {
	return &Handler{
		promoService: promoService,
	}
}

func (h *Handler) CreatePromo(ctx context.Context, request *promopb.CreatePromoRequest) (*promopb.CreatePromoResponse, error) {
	return nil, nil
}
