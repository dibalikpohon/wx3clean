package dokter

import (
	"wx3clean/domain/entities"
	"wx3clean/domain/repo"
	"wx3clean/domain/request"
	"wx3clean/domain/usecases/dokter"
)

type AddDokterUsecaseImpl struct {
	insertAction repo.InsertAction[entities.Dokter]
}

// Execute implements dokter.AddDokterUsecase.
func (*AddDokterUsecaseImpl) Execute(request request.AddDokter) (err error) {
	panic("unimplemented")
}

func NewAddDokterUsecase(insertAction repo.InsertAction[entities.Dokter]) dokter.AddDokterUsecase {
	return &AddDokterUsecaseImpl{insertAction}
}
