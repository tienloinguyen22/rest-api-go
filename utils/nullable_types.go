package utils

import (
	"database/sql"
	"encoding/json"
	"time"
)

type NullString struct {
	sql.NullString
}

func (v NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
			return json.Marshal(v.String)
	} else {
			return json.Marshal(nil)
	}
}

type NullTime struct {
	sql.NullTime
}

func (v NullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
			return json.Marshal(v.Time.Format(time.RFC3339))
	} else {
			return json.Marshal(nil)
	}
}

type NullInt64 struct {
	sql.NullInt64
}

func (v NullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
			return json.Marshal(v.Int64)
	} else {
			return json.Marshal(nil)
	}
}