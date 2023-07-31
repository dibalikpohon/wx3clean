package dokter

import (
	"wx3clean/domain/entities"
	"wx3clean/domain/repo"
	"wx3clean/domain/response"
	"wx3clean/domain/usecases/dokter"
)

type GetDokterByIdUsecaseImpl struct {
	getByIdAction repo.GetByIdAction[entities.Dokter]
}

// Execute implements dokter.GetDokterByIdUsecase.
func (*GetDokterByIdUsecaseImpl) Execute(id string) (response response.GetDokter, err error) {
	panic("unimplemented")
}

func NewGetDokterByIdUsecase(getByIdAction repo.GetByIdAction[entities.Dokter]) dokter.GetDokterByIdUsecase {
	return &GetDokterByIdUsecaseImpl{getByIdAction}
}
