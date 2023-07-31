package dokter

import (
	"wx3clean/domain/entities"
	"wx3clean/internal/db"
)

// Insert implements repo.BasicAction.
func (r *DokterRepo) Insert(dao entities.Dokter) (err error) {

	db := db.MysqlGetInstance()
	query_s := `
	  INSERT INTO dokters (id, nama, spesialisasi, tarif_konsultasi)
	  VALUES (?, ?, ?, ?)
	`
	_, err = db.Exec(query_s, dao.Id, dao.Nama, dao.Spesialisasi, dao.TarifKonsultasi)
	if err != nil {
		panic(err)
	}

	return
}
