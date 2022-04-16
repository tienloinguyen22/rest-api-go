package users

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/tienloinguyen22/edwork-api-go/utils"
)

type User struct {
	ID uuid.UUID `db:"id"`
	FullName string `db:"full_name"`
	Email string `db:"email"`
	PhoneNo string `db:"phone_no"`
	AvatarUrl string `db:"avatar_url"`
	Dob sql.NullTime `db:"dob"`
	Address string `db:"address"`
	Grade int `db:"grade"`
	School string `db:"school"`
	Gender string `db:"gender"`
	OwnerType string `db:"owner_type"`
	SignupProvider string `db:"signup_provider"`
	BankTransferCode string `db:"bank_transfer_code"`
	FirebaseID string `db:"firebase_id"`
	IsActive string `db:"is_active"`
	utils.CommonEntityFields
}