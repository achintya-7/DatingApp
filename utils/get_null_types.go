package utils

import "database/sql"

func GetNullInt(i *int32) sql.NullInt32 {
	if i != nil {
		return sql.NullInt32{Int32: *i, Valid: true}
	} else {
		return sql.NullInt32{Valid: false}
	}
}

func GetNullString(s string) sql.NullString {
	if s != "" {
		return sql.NullString{String: s, Valid: true}
	} else {
		return sql.NullString{Valid: false}
	}
}
