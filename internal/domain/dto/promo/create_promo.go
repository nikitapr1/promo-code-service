package promo

import "time"

type Mode string

const (
	COMMON Mode = "COMMON"
	UNIQUE Mode = "UNIQUE"
)

type Target struct {
	AgeFrom    uint32
	AgeUntil   uint32
	Country    string
	Categories []string
}

type CreatePromoDTO struct {
	Mode        Mode
	PromoCommon string
	PromoUnique []string
	Description string
	ImageUrl    string
	Target      Target
	MaxCount    int64
	ActiveFrom  time.Time
	ActiveUntil time.Time
}
