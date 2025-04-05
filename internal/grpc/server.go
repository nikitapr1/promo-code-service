package promogrpc

import (
	"context"
	"ya-promo-code-service/internal/grpc/handler/create_promo"
	promoservice "ya-promo-code-service/internal/service/promo"
	promopb "ya-promo-code-service/pkg/api/pb"
)

type ServerAPI struct {
	promopb.UnimplementedPromoServiceServer
	promoService       *promoservice.Service
	createPromoHandler *create_promo.Handler
}

func New(srv *promoservice.Service, createPromoHandler *create_promo.Handler) *ServerAPI {
	return &ServerAPI{
		UnimplementedPromoServiceServer: promopb.UnimplementedPromoServiceServer{},
		promoService:                    srv,
		createPromoHandler:              createPromoHandler,
	}
}

//
//func Register(gRPC *grpc.Server, promoService PromoService) {
//	promopb.RegisterPromoServiceServer(gRPC, &ServerAPI{
//		UnimplementedPromoServiceServer: promopb.UnimplementedPromoServiceServer{},
//		promoService:                    promoService,
//	})
//}

func (s *ServerAPI) CreatePromo(ctx context.Context, request *promopb.CreatePromoRequest) (*promopb.CreatePromoResponse, error) {
	return s.createPromoHandler.CreatePromo(ctx, request)
}

func (s *ServerAPI) GetPromo(ctx context.Context, request *promopb.GetPromoRequest) (*promopb.GetPromoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerAPI) ListPromo(ctx context.Context, request *promopb.ListPromoRequest) (*promopb.ListPromoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerAPI) UpdatePromo(ctx context.Context, request *promopb.UpdatePromoRequest) (*promopb.UpdatePromoResponse, error) {
	//TODO implement me
	panic("implement me")
}
