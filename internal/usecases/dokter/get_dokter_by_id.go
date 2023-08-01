package dokter

import (
	"wx3clean/domain/entities"
	"wx3clean/domain/repo"
	"wx3clean/domain/usecases/dokter"
)

type GetDokterByIdUsecaseImpl struct {
	getByIdAction repo.GetByIdAction[entities.Dokter]
}

// Execute implements dokter.GetDokterByIdUsecase.
func (u *GetDokterByIdUsecaseImpl) Execute(id string) (result entities.Dokter, err error) {
	result, err = u.getByIdAction.GetById(id)
	return
}

func NewGetDokterByIdUsecase(getByIdAction repo.GetByIdAction[entities.Dokter]) dokter.GetDokterByIdUsecase {
	return &GetDokterByIdUsecaseImpl{getByIdAction}
}
