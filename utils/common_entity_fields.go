package utils

import "time"

type CommonEntityFields struct {
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	CreatedBy string `db:"created_by" json:"createdBy"`
	UpdatedAt time.Time `db:"updated_at" json:"updpatedAt"`
	UpdatedBy string `db:"updated_by" json:"updatedBy"`
}