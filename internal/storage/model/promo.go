package model

import (
	"promo-code-service/internal/domain/dto/promo"
	"time"
)

type Promo struct {
	Id               string     `db:"id"`
	CompanyId        string     `db:"company_id"`
	Description      string     `db:"description"`
	ImageUrl         string     `db:"image_url"`
	ActiveFrom       time.Time  `db:"active_from"`
	ActiveUntil      time.Time  `db:"active_until"`
	CreatedAt        time.Time  `db:"created_at"`
	Mode             promo.Mode `db:"mode"`
	TargetAgeFrom    int        `db:"target_age_from"`
	TargetAgeUntil   int        `db:"target_age_until"`
	TargetCountry    string     `db:"target_country"`
	TargetCategories []string   `db:"target_categories"`
}
