package request

type AddDokter struct {
	Nama            string  `json:"nama"`
	Spesialisasi    string  `json:"spesialisasi"`
	TarifKonsultasi float32 `json:"tarif_konsultasi"`
}

type UpdateDokter AddDokter
