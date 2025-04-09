package model

type PromoCode struct {
	Id          string `db:"id"`
	PromoId     string `db:"promo_id"`
	Code        string `db:"code"`
	IsActivated bool   `db:"is_activated"`
	MaxCount    int64  `db:"max_count"`
}
