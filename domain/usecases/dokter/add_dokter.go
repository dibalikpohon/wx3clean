package dokter

import (
	"wx3clean/domain/request"
)

type AddDokterUsecase interface {
	Execute(request request.AddDokter) (err error)
}
