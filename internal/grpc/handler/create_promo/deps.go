//go:generate mockgen -source deps.go -package $GOPACKAGE -typed -destination mock_deps_test.go
package create_promo

import (
	"context"
	"promo-code-service/internal/domain/dto"
)

type promoService interface {
	Create(ctx context.Context, companyId string, promoDto dto.PromoDTO)
}
