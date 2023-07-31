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
func (*DeleteDokterUsecaseImpl) Execute(id string) (err error) {
	panic("unimplemented")
}

func NewDeleteDokterUsecase(deleteAction repo.DeleteAction[entities.Dokter]) dokter.DeleteDokterUsecase {
	return &DeleteDokterUsecaseImpl{deleteAction}
}
