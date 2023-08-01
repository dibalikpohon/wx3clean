package dokter

import (
	"net/http"
	handler "wx3clean/domain/handler"
	usecases "wx3clean/domain/usecases/dokter"
)

type dokterHandler struct {
	addDokter     usecases.AddDokterUsecase
	deleteDokter  usecases.DeleteDokterUsecase
	getAllDokter  usecases.GetAllDokterUsecase
	getDokterById usecases.GetDokterByIdUsecase
	updateDokter  usecases.UpdateDokterUsecase
}

// ServeHTTP implements handler.HandlerInterface.
func (h *dokterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, v := range h.GetHandlers() {

	}
}

// GetHandlers implements handler.HandlerInterface.
func (h *dokterHandler) GetHandlers() (hs []handler.HandlerStruct) {
	hs = append(hs,
		handler.HandlerStruct{Method: http.MethodGet, Path: "/dokter", HandlerFunc: h.get}, // get All and get by id
		handler.HandlerStruct{Method: http.MethodPost, Path: "/dokter", HandlerFunc: h.post},
		handler.HandlerStruct{Method: http.MethodPut, Path: "/dokter", HandlerFunc: h.put},
		handler.HandlerStruct{Method: http.MethodDelete, Path: "/dokter", HandlerFunc: h.delete},
	)
	return
}

func NewDokterHandler(
	addDokterUsecase usecases.AddDokterUsecase,
	deleteDokterUsecase usecases.DeleteDokterUsecase,
	getAllDokterUsecase usecases.GetAllDokterUsecase,
	getDokterByIdUsecase usecases.GetDokterByIdUsecase,
	updateDokterUsecase usecases.UpdateDokterUsecase,
) handler.HandlerInterface {
	return &dokterHandler{
		addDokter:     addDokterUsecase,
		deleteDokter:  deleteDokterUsecase,
		getAllDokter:  getAllDokterUsecase,
		getDokterById: getDokterByIdUsecase,
		updateDokter:  updateDokterUsecase,
	}
}
