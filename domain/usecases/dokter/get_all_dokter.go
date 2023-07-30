package dokter

import "karuntos_wx3/domain/response"

type GetAllDokterUsecase interface {
	Execute() (response response.GetDokterArray, err error)
}
