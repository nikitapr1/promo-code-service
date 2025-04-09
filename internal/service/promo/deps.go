//go:generate mockgen -source deps.go -package $GOPACKAGE -typed -destination mock_deps_test.go
package promo_service

import (
	"context"
	"promo-code-service/internal/storage/model"
)

type promoRepository interface {
	Create(ctx context.Context, promoModel *model.Promo) (id string, err error)
}

type promoCodeRepository interface {
	Create(ctx context.Context, promoCodeModel *model.PromoCode) (id string, err error)
}
