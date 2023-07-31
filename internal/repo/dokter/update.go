package dokter

import (
	"errors"
	"wx3clean/domain/entities"
	repoerr "wx3clean/domain/repo/errors"
	"wx3clean/internal/db"
)

// Update implements repo.BasicAction.
func (r *DokterRepo) Update(id string, dao entities.Dokter) (err error) {

	_, err = r.GetById(id)
	var dataNotExists *repoerr.DataNotExistsError
	if errors.As(err, &dataNotExists) {
		return err
	}

	db := db.MysqlGetInstance()
	query_s := `
	  UPDATE dokters 
	  SET 
	    nama=?, 
	    spesialisasi = ?, 
	    tarif_konsultasi = ? 
	  WHERE id = ?
	`
	_, err = db.Exec(query_s, dao.Nama, dao.Spesialisasi, dao.TarifKonsultasi, id)
	if err != nil {
		panic(err)
	}

	return nil
}
