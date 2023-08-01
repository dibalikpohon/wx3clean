package dokter

import (
	"wx3clean/domain/entities"
	"wx3clean/internal/db"
	"wx3clean/internal/helper"
)

// Insert implements repo.BasicAction.
func (r *DokterRepo) Insert(dao entities.Dokter) (newId string, err error) {

	db := db.MysqlGetInstance()
	query_s := `
	  INSERT INTO dokters (id, nama, spesialisasi, tarif_konsultasi)
	  VALUES (?, ?, ?, ?)
	`

	newId = helper.RandomChar(15)

	_, err = db.Exec(query_s, newId, dao.Nama, dao.Spesialisasi, dao.TarifKonsultasi)
	if err != nil {
		panic(err)
	}

	return
}
