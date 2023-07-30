package dokter

import "wx3clean/domain/response"

type GetAllDokterUsecase interface {
	Execute() (response response.GetDokterArray, err error)
}
