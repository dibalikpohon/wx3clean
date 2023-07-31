package dokter

import (
	"errors"
	repoerr "wx3clean/domain/repo/errors"
	"wx3clean/internal/db"
)

// Delete implements repo.BasicAction.
func (r *DokterRepo) Delete(id string) error {

	_, err := r.GetById(id)
	var dataNotExists *repoerr.DataNotExistsError
	if errors.As(err, &dataNotExists) {
		return err
	}

	db := db.MysqlGetInstance()
	query_s := `DELETE FROM dokters WHERE id = ?`
	_, err = db.Exec(query_s, id)
	if err != nil {
		panic(err)
	}

	return nil
}
