package promo_repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"promo-code-service/internal/storage/model"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, promoModel *model.Promo) (id string, err error) {

	const op = "storage.promo.Create"

	query := `
		INSERT INTO promo(
			id, company_id, description, image_url, active_from, active_until,
			created_at, mode, target_age_from, target_age_until,
			target_country, target_categories
		) VALUES (
			:id, :company_id, :description, :image_url, :active_from, :active_until,
			:created_at, :mode, :target_age_from, :target_age_until,
			:target_country, :target_categories
		)
		RETURNING id
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return "", fmt.Errorf("%s: prepare failed: %w", op, err)
	}

	err = stmt.QueryRowxContext(ctx, promoModel).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
