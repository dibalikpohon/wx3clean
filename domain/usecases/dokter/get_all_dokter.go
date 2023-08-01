package dokter

import (
	"wx3clean/domain/entities"
)

type GetAllDokterUsecase interface {
	Execute() (result []entities.Dokter, err error)
}
