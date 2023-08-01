package dokter

import (
	"wx3clean/domain/entities"
)

type UpdateDokterUsecase interface {
	Execute(id string, updatedDokter entities.Dokter) (err error)
}
