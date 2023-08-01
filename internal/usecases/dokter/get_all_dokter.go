package dokter

import (
	"wx3clean/domain/entities"
	"wx3clean/domain/repo"
	"wx3clean/domain/usecases/dokter"
)

type GetAllDokterUsecaseImpl struct {
	getAllAction repo.GetAllAction[entities.Dokter]
}

// Execute implements dokter.GetAllDokterUsecase.
func (u *GetAllDokterUsecaseImpl) Execute() (result []entities.Dokter, err error) {
	result, err = u.getAllAction.GetAll()
	return
}

func NewGetAllDokterUsecase(getAllAction repo.GetAllAction[entities.Dokter]) dokter.GetAllDokterUsecase {
	return &GetAllDokterUsecaseImpl{getAllAction}
}
