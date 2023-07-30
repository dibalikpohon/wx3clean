package dokter

import "wx3clean/domain/request"

type UpdateDokterUsecase interface {
	Execute(id string, request request.UpdateDokter) (err error)
}
