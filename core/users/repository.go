package users

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tienloinguyen22/edwork-api-go/utils"
)

type User struct {
	ID uuid.UUID `db:"id"`
	FullName string `db:"full_name"`
	Email string `db:"email"`
	PhoneNo sql.NullString `db:"phone_no"`
	AvatarUrl sql.NullString `db:"avatar_url"`
	Dob sql.NullTime `db:"dob"`
	Address sql.NullString `db:"address"`
	Grade sql.NullInt64 `db:"grade"`
	School sql.NullString `db:"school"`
	Gender sql.NullString `db:"gender"`
	OwnerType sql.NullString `db:"owner_type"`
	SignupProvider string `db:"signup_provider"`
	BankTransferCode string `db:"bank_transfer_code"`
	FirebaseID string `db:"firebase_id"`
	IsActive bool `db:"is_active"`
	utils.CommonEntityFields
}

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r UserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE email=$1`
	if err := r.DB.GetContext(ctx, &user, query, email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r UserRepository) Create(ctx context.Context, user *User) (*User, error) {
	query := `
		INSERT INTO users(
			full_name, email, phone_no, avatar_url, dob, address, grade, school, gender, owner_type, signup_provider, bank_transfer_code, firebase_id, created_by, updated_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id;
	`
	row := r.DB.QueryRowxContext(
		ctx,
		query,
		user.FullName,
		user.Email,
		user.PhoneNo,
		user.AvatarUrl,
		user.Dob,
		user.Address,
		user.Grade,
		user.School,
		user.Gender,
		user.OwnerType,
		user.SignupProvider,
		user.BankTransferCode,
		user.FirebaseID,
		"self",
		"self",
	)
	if err := row.Scan(&user.ID); err != nil {
		return user, err
	}
	return user, nil
}

func (r UserRepository) UpdateFirebaseInfoByID(ctx context.Context, id uuid.UUID, user *User) (*User, error) {
	query := `
		UPDATE users SET firebase_id=$1, signup_provider=$2 WHERE id=$3
	`
	if err := r.DB.GetContext(ctx, &user, query, user.FirebaseID, user.SignupProvider, user.ID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}