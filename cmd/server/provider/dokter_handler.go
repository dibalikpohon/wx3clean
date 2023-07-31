package provider

import (
	iface "wx3clean/domain/handler"

	"wx3clean/internal/db"
	handler "wx3clean/internal/handler/dokter"
	repo "wx3clean/internal/repo/dokter"
	usecases "wx3clean/internal/usecases/dokter"
)

func NewDokterHandler() iface.HandlerInterface {
	db := db.MysqlGetInstance()

	repo := repo.NewDokterRepo(db)

	addDokterUsecase := usecases.NewAddDokterUsecase(repo)
	deleteDokterUsecase := usecases.NewDeleteDokterUsecase(repo)
	getAllDokterUsecase := usecases.NewGetAllDokterUsecase(repo)
	getDokterByIdUsecase := usecases.NewGetDokterByIdUsecase(repo)
	updateDokterUsecase := usecases.NewUpdateDokterUsecase(repo)

	return handler.NewDokterHandler(
		addDokterUsecase,
		deleteDokterUsecase,
		getAllDokterUsecase,
		getDokterByIdUsecase,
		updateDokterUsecase)
}
