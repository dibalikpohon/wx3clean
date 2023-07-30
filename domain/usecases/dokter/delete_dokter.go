package dokter

type DeleteDokterUsecase interface {
	Execute(id string) (err error)
}
