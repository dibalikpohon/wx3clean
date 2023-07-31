package dokter

import (
	"database/sql"
	"errors"
	"wx3clean/domain/entities"
	repoerr "wx3clean/domain/repo/errors"
	"wx3clean/internal/db"
)

// GetById implements repo.BasicAction.
func (r *DokterRepo) GetById(id string) (v entities.Dokter, err error) {

	db := db.MysqlGetInstance()
	query_s := `SELECT id, nama, spesialisasi, tarif_konsultasi FROM dokters WHERE id = ?`

	err = db.QueryRow(query_s, id).Scan(&v.Id, &v.Nama, &v.Spesialisasi, &v.TarifKonsultasi)
	if errors.Is(err, sql.ErrNoRows) {
		err = repoerr.NewDataNotExistsError("id", id)
	}

	return
}
