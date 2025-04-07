//go:generate mockgen -source deps.go -package $GOPACKAGE -typed -destination mock_deps_test.go
package promoservice

import (
	"context"
	"promo-code-service/internal/domain/dto/promo"
)

type promoProvider interface {
}

type promoSaver interface {
	Create(ctx context.Context, companyId string, promoDto promo.CreatePromoDTO) (id string, err error)
}
