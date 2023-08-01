package dokter

import (
	"wx3clean/domain/entities"
	"wx3clean/domain/repo"
	"wx3clean/domain/usecases/dokter"
)

type UpdateDokterUsecaseImpl struct {
	updateAction repo.UpdateAction[entities.Dokter]
}

// Execute implements dokter.UpdateDokterUsecase.
func (u *UpdateDokterUsecaseImpl) Execute(id string, updatedDokter entities.Dokter) (err error) {
	err = u.updateAction.Update(id, updatedDokter)
	return
}

func NewUpdateDokterUsecase(updateAction repo.UpdateAction[entities.Dokter]) dokter.UpdateDokterUsecase {
	return &UpdateDokterUsecaseImpl{updateAction}
}
