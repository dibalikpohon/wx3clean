package dokter

import "wx3clean/domain/response"

type GetDokterByIdUsecase interface {
	Execute(id string) (response response.GetDokter, err error)
}
