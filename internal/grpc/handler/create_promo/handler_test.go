package create_promo

import (
	"context"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"promo-code-service/internal/domain/dto/promo"
	"promo-code-service/internal/pkg/pointer"
	promopb "promo-code-service/pkg/api/pb"
	"testing"
	"time"
)

func TestHandler_CreatePromo(t *testing.T) {
	type fields struct {
		promoService *MockpromoService
	}

	rq := &promopb.CreatePromoRequest{
		CompanyId:   "company_id",
		Mode:        promopb.Mode(promo.COMMON),
		PromoCommon: pointer.To("promo"),
		Description: "description",
		ImageUrl:    pointer.To("https://google.com"),
		Target: pointer.To(promopb.Target{
			AgeFrom:    pointer.ToUint32(10),
			AgeUntil:   pointer.ToUint32(30),
			Country:    pointer.To("RU"),
			Categories: []string{"pen", "city"},
		}),
		MaxCount:    5,
		ActiveFrom:  "2025-01-01",
		ActiveUntil: "2026-01-01",
	}

	createdId := "someId"

	tests := []struct {
		name    string
		prepare func(f *fields)
		want    *promopb.CreatePromoResponse
		wantErr bool
	}{
		{
			name: "success",
			prepare: func(f *fields) {

				activeFrom, err := time.Parse(time.DateOnly, rq.GetActiveFrom())

				if err != nil {
					t.Errorf("failed to parse activeFrom from time: %v", err)
				}

				activeUntil, err := time.Parse(time.DateOnly, rq.GetActiveUntil())

				if err != nil {
					t.Errorf("failed to parse activeUntil from time: %v", err)
				}

				f.promoService.EXPECT().Create(gomock.Any(), rq.GetCompanyId(), promo.CreatePromoDTO{
					Mode:        promo.Mode(rq.Mode),
					PromoCommon: rq.GetPromoCommon(),
					Description: rq.GetDescription(),
					ImageUrl:    rq.GetImageUrl(),
					Target: promo.Target{
						AgeFrom:    rq.Target.GetAgeFrom(),
						AgeUntil:   rq.Target.GetAgeUntil(),
						Country:    rq.Target.GetCountry(),
						Categories: rq.Target.GetCategories(),
					},
					MaxCount:    rq.GetMaxCount(),
					ActiveFrom:  activeFrom,
					ActiveUntil: activeUntil,
				}).Return(createdId, nil) //TODO: DoAndReturn

			},
			want: &promopb.CreatePromoResponse{
				Id: createdId,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)

			f := &fields{
				promoService: NewMockpromoService(ctrl),
			}

			tt.prepare(f)

			h := &Handler{
				promoService: f.promoService,
			}
			got, err := h.CreatePromo(context.Background(), rq)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePromo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			expected := &promopb.CreatePromoResponse{Id: createdId}

			require.Equal(t, expected, got)
		})
	}
}
