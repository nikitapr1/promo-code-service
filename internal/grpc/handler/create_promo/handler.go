package create_promo

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"promo-code-service/internal/domain/dto/promo"
	promopb "promo-code-service/pkg/api/pb"
	"time"
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

	activeFrom, err := time.Parse(time.DateOnly, r.GetActiveFrom())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid active_from")
	}
	activeUntil, err := time.Parse(time.DateOnly, r.GetActiveUntil())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid active_until")
	}

	dto := promo.CreatePromoDTO{
		Mode:        promo.Mode(r.GetMode()),
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
		ActiveFrom:  activeFrom,
		ActiveUntil: activeUntil,
	}

	promoId, err := h.promoService.Create(ctx, r.GetCompanyId(), dto)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	resp := &promopb.CreatePromoResponse{
		Id: promoId,
	}
	return resp, nil
}
