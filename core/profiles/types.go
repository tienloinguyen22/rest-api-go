package profiles

import (
	"time"
)

type UpdateUserProfilePayload struct {
	FullName string `json:"fullName" validate:"max=50"`
	PhoneNo string `json:"phoneNo" validate:"max=10"`
	AvatarUrl string `json:"avatarUrl"`
	Dob time.Time `json:"dob"`
	Address string `json:"address" validate:"max=1000"`
	Grade int64 `json:"grade" validate:"max=12"`
	School string `json:"school" validate:"max=1000"`
	Gender string `json:"gender" validate:"max=100"`
	OwnerType string `json:"ownerType" validate:"max=100"`
}