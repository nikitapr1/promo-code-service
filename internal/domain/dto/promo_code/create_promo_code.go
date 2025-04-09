package promo_code

type CreatePromoCodeDTO struct {
	PromoId     string
	Code        string
	IsActivated bool
	MaxCount    int64
}
