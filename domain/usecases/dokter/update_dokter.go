package dokter

import "karuntos_wx3/domain/request"

type UpdateDokterUsecase interface {
	Execute(id string, request request.UpdateDokter) (err error)
}
