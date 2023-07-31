package dokter

import (
	"wx3clean/domain/entities"
	"wx3clean/domain/repo"
	"wx3clean/domain/request"
	"wx3clean/domain/usecases/dokter"
)

type UpdateDokterUsecaseImpl struct {
	updateAction repo.UpdateAction[entities.Dokter]
}

// Execute implements dokter.UpdateDokterUsecase.
func (*UpdateDokterUsecaseImpl) Execute(id string, request request.UpdateDokter) (err error) {
	panic("unimplemented")
}

func NewUpdateDokterUsecase(updateAction repo.UpdateAction[entities.Dokter]) dokter.UpdateDokterUsecase {
	return &UpdateDokterUsecaseImpl{updateAction}
}
