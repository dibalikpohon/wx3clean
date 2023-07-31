package dokter

import (
	"wx3clean/domain/entities"
	"wx3clean/internal/db"
)

// GetAll implements repo.BasicAction.
func (r *DokterRepo) GetAll() (arr []entities.Dokter, err error) {

	db := db.MysqlGetInstance()
	sql := `SELECT id, nama, spesialisasi, tarif_konsultasi FROM dokters`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var e entities.Dokter
		rows.Scan(&e.Id, &e.Nama, &e.Spesialisasi, &e.TarifKonsultasi)

		arr = append(arr, e)
	}

	return
}
