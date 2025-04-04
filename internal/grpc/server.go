package promogrpc

import (
	"context"
	"google.golang.org/grpc"
	"ya-promo-code-service/internal/domain/dto"
	promopb "ya-promo-code-service/pkg/api/pb"
)

type PromoService interface {
	Create(
		ctx context.Context,
		companyId string,
		promoDto dto.PromoDTO,
	)
	// TODO: описать остальные методы сервиса
}

type serverAPI struct {
	promopb.UnimplementedPromoServiceServer
	promoService PromoService
}

func Register(gRPC *grpc.Server, promoService PromoService) {
	promopb.RegisterPromoServiceServer(gRPC, &serverAPI{
		UnimplementedPromoServiceServer: promopb.UnimplementedPromoServiceServer{},
		promoService:                    promoService,
	})
}

func (s serverAPI) CreatePromo(ctx context.Context, request *promopb.CreatePromoRequest) (*promopb.CreatePromoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s serverAPI) ListPromo(ctx context.Context, request *promopb.ListPromoRequest) (*promopb.ListPromoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s serverAPI) GetPromo(ctx context.Context, request *promopb.GetPromoRequest) (*promopb.GetPromoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s serverAPI) UpdatePromo(ctx context.Context, request *promopb.UpdatePromoRequest) (*promopb.UpdatePromoResponse, error) {
	//TODO implement me
	panic("implement me")
}
