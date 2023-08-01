package dokter

import (
	"wx3clean/domain/entities"
)

type GetDokterByIdUsecase interface {
	Execute(id string) (result entities.Dokter, err error)
}
