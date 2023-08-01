package dokter

import (
	"encoding/json"
	"errors"
	"net/http"
	repoerror "wx3clean/domain/repo/errors"
	"wx3clean/domain/response"
)

func (h *dokterHandler) get(w http.ResponseWriter, r *http.Request) {
	// Get by id
	id := r.URL.Query().Get("id")
	if len(id) > 0 {
		result, err := h.getDokterById.Execute(id)
		var dataNotFoundError *repoerror.DataNotExistsError
		if errors.As(err, &dataNotFoundError) {
			// return error response
			response := response.NewDataNotFoundErrorResponse("id", id)
			w.WriteHeader(int(response.Code))
			err = json.NewEncoder(w).Encode(&response)
			if err != nil {
				panic(err)
			}
		}
		response := response.NewGetDokterResponse(result)
		w.WriteHeader(int(response.Code))
		err = json.NewEncoder(w).Encode(&response)
		if err != nil {
			panic(err)
		}
	}

	result, err := h.getAllDokter.Execute()
	response := response.NewGetDokterArrayResponse(result)
	w.WriteHeader(response.Code)
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		panic(err)
	}
}
