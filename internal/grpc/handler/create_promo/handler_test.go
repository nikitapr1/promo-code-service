package create_promo

import (
	"context"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
	"promo-code-service/internal/domain/dto/promo"
	promopb "promo-code-service/pkg/api/pb"
	"promo-code-service/pkg/pointer"
	"testing"
	"time"
)

func TestHandler_CreatePromo(t *testing.T) {
	type fields struct {
		promoService *MockpromoService
	}

	activeFrom, err := time.Parse(time.DateOnly, "2025-01-01")
	require.NoError(t, err)
	activeUntil, err := time.Parse(time.DateOnly, "2026-01-01")
	require.NoError(t, err)

	rq := &promopb.CreatePromoRequest{
		CompanyId: "company_id",
		Mode:      promopb.Mode_COMMON,

		PromoCommon: pointer.To("promo"),
		Description: "description",
		ImageUrl:    pointer.To("https://google.com"),
		Target: pointer.To(promopb.Target{
			AgeFrom:    pointer.ToInt64(10),
			AgeUntil:   pointer.ToInt64(30),
			Country:    pointer.To("RU"),
			Categories: []string{"pen", "city"},
		}),
		MaxCount:    5,
		ActiveFrom:  timestamppb.New(activeFrom),
		ActiveUntil: timestamppb.New(activeUntil),
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

				rqActiveFrom := rq.GetActiveFrom().AsTime()

				rqActiveUntil := rq.GetActiveUntil().AsTime()

				dto := promo.CreatePromoDTO{
					Mode:        promo.Mode(rq.GetMode().String()),
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
					ActiveFrom:  rqActiveFrom,
					ActiveUntil: rqActiveUntil,
				}

				f.promoService.EXPECT().Create(gomock.Any(), dto).Return(createdId, nil)

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
