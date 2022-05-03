package users

import (
	"github.com/google/uuid"
	"github.com/tienloinguyen22/rest-api-go/utils"
)

type User struct {
	ID uuid.UUID `db:"id" json:"id"`
	FullName string `db:"full_name" json:"fullName"`
	Email string `db:"email" json:"email"`
	PhoneNo utils.NullString `db:"phone_no" json:"phoneNo"`
	AvatarUrl utils.NullString `db:"avatar_url" json:"avatarUrl"`
	Dob utils.NullTime `db:"dob" json:"dob"`
	Address utils.NullString `db:"address" json:"address"`
	Grade utils.NullInt64 `db:"grade" json:"grade"`
	School utils.NullString `db:"school" json:"school"`
	Gender utils.NullString `db:"gender" json:"gender"`
	OwnerType utils.NullString `db:"owner_type" json:"ownerType"`
	SignupProvider string `db:"signup_provider" json:"signupProvider"`
	BankTransferCode string `db:"bank_transfer_code" json:"bankTransferCode"`
	FirebaseID string `db:"firebase_id" json:"firebaseId"`
	IsActive bool `db:"is_active" json:"isActive"`
	utils.CommonEntityFields
}