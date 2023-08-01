package response

import (
	"net/http"
	"wx3clean/domain/entities"
)

type dokterResponse struct {
	Id              string  `json:"id"`
	Nama            string  `json:"nama"`
	Spesialisasi    string  `json:"spesialisasi"`
	TarifKonsultasi float32 `json:"tarif_konsultasi"`
}

type GetDokterResponse struct {
	BasicResponse
	Data dokterResponse `json:"data"`
}

type GetDokterArrayResponse struct {
	BasicResponse
	Data []dokterResponse `json:"data"`
}

func NewGetDokterResponse(e entities.Dokter) GetDokterResponse {
	return GetDokterResponse{
		BasicResponse: BasicResponse{
			Status: http.StatusText(http.StatusOK),
			Code:   http.StatusOK,
		},
		Data: DokterResponseFromEntity(e),
	}
}

func NewGetDokterArrayResponse(es []entities.Dokter) GetDokterArrayResponse {
	responseArray := make([]dokterResponse, 0)
	for _, v := range es {
		responseArray = append(responseArray, DokterResponseFromEntity(v))
	}
	return GetDokterArrayResponse{
		BasicResponse: BasicResponse{
			Status: http.StatusText(http.StatusOK),
			Code:   http.StatusOK,
		},
		Data: []dokterResponse{},
	}
}

func DokterResponseFromEntity(e entities.Dokter) dokterResponse {
	return dokterResponse{
		Id:              e.Id,
		Nama:            e.Nama,
		Spesialisasi:    e.Spesialisasi,
		TarifKonsultasi: e.TarifKonsultasi,
	}
}
