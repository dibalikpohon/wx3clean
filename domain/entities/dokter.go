package entities

type Dokter struct {
	Id              string  `json:"id"`
	Nama            string  `json:"nama"`
	Spesialisasi    string  `json:"spesialisasi"`
	TarifKonsultasi float32 `json:"tarif_konsultasi"`
}
