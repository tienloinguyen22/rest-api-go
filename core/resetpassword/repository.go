package resetpassword

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ResetPasswordTokenRepository struct {
	DB *sqlx.DB
}

func NewResetPasswordTokenRepository(db *sqlx.DB) *ResetPasswordTokenRepository {
	return &ResetPasswordTokenRepository{
		DB: db,
	}
}

func (r ResetPasswordTokenRepository) FindNonExpiredByUserID(ctx context.Context, userID uuid.UUID) (*ResetPasswordToken, error) {
	var resetPasswordToken ResetPasswordToken
	query := `SELECT * FROM reset_password_tokens WHERE user_id=$1 AND expired_at>$2 AND completed=false ORDER BY expired_at DESC LIMIT 1`
	if err := r.DB.GetContext(ctx, &resetPasswordToken, query, userID.String(), time.Now()); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &resetPasswordToken, nil
}

func (r ResetPasswordTokenRepository) Create(ctx context.Context, resetPasswordToken *ResetPasswordToken) (*ResetPasswordToken, error) {
	query := `
		INSERT INTO reset_password_tokens(
			user_id, expired_at, created_by
		) VALUES ($1, $2, $3) RETURNING id, expired_at;
	`
	row := r.DB.QueryRowxContext(
		ctx,
		query,
		resetPasswordToken.ID,
		resetPasswordToken.ExpiredAt,
	)
	if err := row.Scan(&resetPasswordToken.ID, &resetPasswordToken.ExpiredAt); err != nil {
		return resetPasswordToken, err
	}
	return resetPasswordToken, nil
}