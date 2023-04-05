package models

// Response adalah struktur untuk menyimpan respons dari server.
type Cuaca struct {
	Water float64 `json:"water"`
	Wind  float64 `json:"wind"`
}
