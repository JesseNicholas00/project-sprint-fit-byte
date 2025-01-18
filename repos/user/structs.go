package user

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID        `db:"user_id"`
	Name       sql.Null[string] `db:"name"`
	Email      string           `db:"email"`
	Password   string           `db:"password"`
	Preference sql.Null[string] `db:"preference"`
	WeightUnit sql.Null[string] `db:"weight_unit"`
	HeightUnit sql.Null[string] `db:"height_unit"`
	Weight     int              `db:"weight"`
	Height     int              `db:"height"`
	ImageURI   sql.Null[string] `db:"image_uri"`
}
