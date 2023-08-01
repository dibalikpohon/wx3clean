package dokter

import (
	"wx3clean/domain/entities"
	"wx3clean/domain/repo"
	"wx3clean/domain/usecases/dokter"
)

type AddDokterUsecaseImpl struct {
	insertAction repo.InsertAction[entities.Dokter]
}

// Execute implements dokter.AddDokterUsecase.
func (u *AddDokterUsecaseImpl) Execute(dokter entities.Dokter) (newId string, err error) {
	newId, err = u.insertAction.Insert(dokter)
	return
}

func NewAddDokterUsecase(insertAction repo.InsertAction[entities.Dokter]) dokter.AddDokterUsecase {
	return &AddDokterUsecaseImpl{insertAction}
}
