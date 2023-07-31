package dokter

import (
	"wx3clean/domain/entities"
	"wx3clean/domain/repo"
	"wx3clean/domain/response"
	"wx3clean/domain/usecases/dokter"
)

type GetAllDokterUsecaseImpl struct {
	getAllAction repo.GetAllAction[entities.Dokter]
}

// Execute implements dokter.GetAllDokterUsecase.
func (*GetAllDokterUsecaseImpl) Execute() (response response.GetDokterArray, err error) {
	panic("unimplemented")
}

func NewGetAllDokterUsecase(getAllAction repo.GetAllAction[entities.Dokter]) dokter.GetAllDokterUsecase {
	return &GetAllDokterUsecaseImpl{getAllAction}
}
