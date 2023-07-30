package response

type GetDokter struct {
	Id              string  `json:"id"`
	Nama            string  `json:"nama"`
	Spesialisasi    string  `json:"spesialisasi"`
	TarifKonsultasi float32 `json:"tarif_konsultasi"`
}

type GetDokterArray []GetDokter
