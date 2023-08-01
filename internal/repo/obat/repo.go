package obat

import (
	"database/sql"
)

type ObatRepo struct {
	db *sql.DB
}

func NewObatRepo(db *sql.DB) *ObatRepo {
	return &ObatRepo{db}
}
