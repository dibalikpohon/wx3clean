package request

type AddDokter struct {
	Nama            string `json:"nama"`
	Spesialisasi    string `json:"spesialisasi"`
	TarifKonsultasi string `json:"tarif_konsultasi"`
}

type UpdateDokter AddDokter
