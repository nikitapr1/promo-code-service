//go:generate mockgen -source deps.go -package $GOPACKAGE -typed -destination mock_deps_test.go
package create_promo

import (
	"context"
	"ya-promo-code-service/internal/domain/dto"
)

type promoService interface {
	Create(ctx context.Context, promoDto dto.PromoDTO)
}
