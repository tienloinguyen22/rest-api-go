package utils

import "time"

type CommonEntityFields struct {
	CreatedAt time.Time `db:"created_at"`
	CreatedBy string `db:"created_by"`
	UpdatedAt time.Time `db:"updated_at"`
	UpdatedBy string `db:"updated_by"`
}