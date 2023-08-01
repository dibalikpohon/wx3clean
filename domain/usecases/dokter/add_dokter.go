package dokter

import (
	"wx3clean/domain/entities"
)

type AddDokterUsecase interface {
	Execute(dokter entities.Dokter) (newId string, err error)
}
