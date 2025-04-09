//go:generate mockgen -source deps.go -package $GOPACKAGE -typed -destination mock_deps_test.go
package create_promo

import (
	"context"
	"promo-code-service/internal/domain/dto/promo"
)

type promoService interface {
	Create(ctx context.Context, promoDto *promo.CreatePromoDTO) (id string, err error)
}
