package promo_code_repository

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

func (r *Repository) Create(ctx context.Context, promoCodeModel *model.PromoCode) (id string, err error) {

	const op = "storage.promo_code.Create"

	query := `
		INSERT INTO promo_code(id, promo_id, code, is_activated, max_count) 
		VALUES (:id, :promo_id, :code, :is_activated, :max_count) 
		RETURNING id
		`

	stmt, err := r.db.PrepareNamedContext(ctx, query)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	err = stmt.QueryRowxContext(
		ctx, promoCodeModel,
	).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil

}
