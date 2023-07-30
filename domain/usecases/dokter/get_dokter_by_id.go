package dokter

import "karuntos_wx3/domain/response"

type GetDokterByIdUsecase interface {
	Execute(id string) (response response.GetDokter, err error)
}
