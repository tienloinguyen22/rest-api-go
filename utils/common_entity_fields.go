package utils

import "time"

type CommonEntityFields struct {
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	CreatedBy string `db:"created_by" json:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updpated_at"`
	UpdatedBy string `db:"updated_by" json:"updated_by"`
}