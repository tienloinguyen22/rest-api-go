package users

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r UserRepository) FindByID(ctx context.Context, userID uuid.UUID) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE id=$1`
	if err := r.DB.GetContext(ctx, &user, query, userID.String()); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
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

func (r UserRepository) FindByFirebaseID(ctx context.Context, firebaseId string) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE firebase_id=$1`
	if err := r.DB.GetContext(ctx, &user, query, firebaseId); err != nil {
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
			id, full_name, email, phone_no, avatar_url, dob, address, grade, school, gender, owner_type, signup_provider, bank_transfer_code, firebase_id, created_by, updated_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id, is_active;
	`
	row := r.DB.QueryRowxContext(
		ctx,
		query,
		user.ID,
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
		user.ID.String(),
		user.ID.String(),
	)
	if err := row.Scan(&user.ID, &user.IsActive); err != nil {
		return user, err
	}
	return user, nil
}

func (r UserRepository) UpdateFirebaseInfoByID(ctx context.Context, user *User) (*User, error) {
	query := `
		UPDATE users SET firebase_id=$1, signup_provider=$2, updated_at=$3, updated_by=$4 WHERE id=$5 RETURNING firebase_id, signup_provider, updated_by, updated_at
	`
	row := r.DB.QueryRowxContext(ctx, query, user.FirebaseID, user.SignupProvider, time.Now(), user.ID, user.ID)
	if err := row.Scan(&user.FirebaseID, &user.SignupProvider, &user.UpdatedBy, &user.UpdatedAt); err != nil {
		return user, err
	}
	return user, nil
}