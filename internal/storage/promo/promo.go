package promorepository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"promo-code-service/internal/domain/dto/promo"
	"time"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, companyId string, promoDto promo.CreatePromoDTO) (id string, err error) {

	const op = "storage.promo.Create"

	stmt, err := r.db.Preparex("INSERT INTO promo(id, owner_id, description, image_url, active_from, active_until, created_at, mode, target_age_from, target_age_until, target_country, target_categories) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING id")
	defer func(stmt *sqlx.Stmt) {
		_ = stmt.Close()
	}(stmt)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	err = stmt.QueryRowxContext(
		ctx,
		uuid.New().String(),
		companyId,
		promoDto.Description,
		promoDto.ImageUrl,
		promoDto.ActiveFrom,
		promoDto.ActiveUntil,
		time.Now(),
		promoDto.Mode,
	).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
