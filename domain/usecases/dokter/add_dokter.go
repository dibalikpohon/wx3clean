package dokter

import (
	"karuntos_wx3/domain/request"
)

type AddDokterUsecase interface {
	Execute(request request.AddDokter) (err error)
}
