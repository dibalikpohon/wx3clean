package dokter

import (
	"database/sql"
)

type DokterRepo struct {
	db *sql.DB
}

func NewDokterRepo(db *sql.DB) *DokterRepo {
	return &DokterRepo{db}
}
