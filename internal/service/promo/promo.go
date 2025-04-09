package promo_service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"promo-code-service/internal/domain/dto/promo"
	"promo-code-service/internal/storage/model"
	"time"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field '%s': %s", e.Field, e.Message)
}

type Service struct {
	log *zap.Logger
	promoRepository
	promoCodeRepository
}

func New(
	log *zap.Logger,
	promoRepository promoRepository,
	promoCodeRepository promoCodeRepository,
) *Service {
	return &Service{
		log:                 log,
		promoRepository:     promoRepository,
		promoCodeRepository: promoCodeRepository,
	}
}

func (s *Service) Create(ctx context.Context, promoDto *promo.CreatePromoDTO) (id string, err error) {
	const op = "service.promo.Create"

	if promoDto.PromoCommon != "" && len(promoDto.PromoUnique) > 0 {
		return "", ValidationError{Field: "promo_common", Message: "promo_common and promo_unique can't be set together"}
	}

	promoModel := &model.Promo{
		Id:               uuid.New().String(),
		CompanyId:        promoDto.CompanyId,
		Description:      promoDto.Description,
		ImageUrl:         promoDto.ImageUrl,
		ActiveFrom:       promoDto.ActiveFrom,
		ActiveUntil:      promoDto.ActiveUntil,
		CreatedAt:        time.Now(),
		Mode:             promoDto.Mode,
		TargetAgeFrom:    int(promoDto.Target.AgeFrom),
		TargetAgeUntil:   int(promoDto.Target.AgeUntil),
		TargetCountry:    promoDto.Target.Country,
		TargetCategories: promoDto.Target.Categories,
	}

	id, err = s.promoRepository.Create(ctx, promoModel)

	if err != nil {
		s.log.Error("Failed to save promo to db", zap.Error(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if promoDto.Mode == promo.COMMON {
		promoCodeModel := &model.PromoCode{
			Id:          uuid.New().String(),
			PromoId:     id,
			Code:        promoDto.PromoCommon,
			IsActivated: false,
			MaxCount:    promoDto.MaxCount,
		}

		_, err = s.promoCodeRepository.Create(ctx, promoCodeModel)
		if err != nil {
			s.log.Error("Failed to save promo code to db", zap.Error(err))
			return "", fmt.Errorf("%s: %w", op, err)
		}
	} else if promoDto.Mode == promo.UNIQUE {
		for _, code := range promoDto.PromoUnique {
			promoCodeModel := &model.PromoCode{
				Id:          uuid.New().String(),
				PromoId:     id,
				Code:        code,
				IsActivated: false,
				MaxCount:    1,
			}

			_, err = s.promoCodeRepository.Create(ctx, promoCodeModel)
			if err != nil {
				s.log.Error("Failed to save unique promo code to db", zap.Error(err))
			}
		}
	} else {
		s.log.Error("Invalid promo mode", zap.String("mode", string(promoDto.Mode)))
		return "", ValidationError{Field: "mode", Message: "Invalid promo mode"}
	}

	return id, nil
}
