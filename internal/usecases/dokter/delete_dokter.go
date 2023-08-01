package dokter

import (
	"wx3clean/domain/entities"
	"wx3clean/domain/repo"
	"wx3clean/domain/usecases/dokter"
)

type DeleteDokterUsecaseImpl struct {
	deleteAction repo.DeleteAction[entities.Dokter]
}

// Execute implements dokter.DeleteDokterUsecase.
func (u *DeleteDokterUsecaseImpl) Execute(id string) (err error) {
	err = u.deleteAction.Delete(id)
	return
}

func NewDeleteDokterUsecase(deleteAction repo.DeleteAction[entities.Dokter]) dokter.DeleteDokterUsecase {
	return &DeleteDokterUsecaseImpl{deleteAction}
}
