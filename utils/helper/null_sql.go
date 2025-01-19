package helper

import "database/sql"

func Assign[T any](v T) sql.Null[T] {
	return sql.Null[T]{
		Valid: true,
		V:     v,
	}
}
