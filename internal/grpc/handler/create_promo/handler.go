package create_promo

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"promo-code-service/internal/domain/dto/promo"
	"promo-code-service/internal/service/promo"
	promopb "promo-code-service/pkg/api/pb"
)

type Handler struct {
	promoService promoService
}

func New(promoService promoService) *Handler {
	return &Handler{
		promoService: promoService,
	}
}

func (h *Handler) CreatePromo(ctx context.Context, r *promopb.CreatePromoRequest) (*promopb.CreatePromoResponse, error) {

	dto := &promo.CreatePromoDTO{
		CompanyId:   r.GetCompanyId(),
		Mode:        promo.Mode(r.GetMode().String()),
		PromoCommon: r.GetPromoCommon(),
		PromoUnique: r.GetPromoUnique(),
		Description: r.GetDescription(),
		ImageUrl:    r.GetImageUrl(),
		Target: promo.Target{
			AgeFrom:    r.Target.GetAgeFrom(),
			AgeUntil:   r.Target.GetAgeUntil(),
			Country:    r.Target.GetCountry(),
			Categories: r.Target.GetCategories(),
		},
		MaxCount:    r.GetMaxCount(),
		ActiveFrom:  r.GetActiveFrom().AsTime(),
		ActiveUntil: r.GetActiveUntil().AsTime(),
	}
	promoId, err := h.promoService.Create(ctx, dto)
	if err != nil {

		if errors.As(err, &promo_service.ValidationError{}) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	resp := &promopb.CreatePromoResponse{
		Id: promoId,
	}
	return resp, nil
}
