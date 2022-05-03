package resetpassword

import (
	"time"

	"github.com/google/uuid"
	"github.com/tienloinguyen22/rest-api-go/utils"
)

type ResetPasswordToken struct {
	ID uuid.UUID `db:"id"`
	UserID uuid.UUID `db:"user_id"`
	Completed bool `db:"completed"`
	ExpiredAt time.Time `db:"expired_at"`
	utils.CommonEntityFields
}

type RequestResetPasswordTokenPayload struct {
	Email string `json:"email" validate:"nonzero"`
}