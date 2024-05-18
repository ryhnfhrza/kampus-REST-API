package web

type MahasiswaMatkulResponse struct {
	NIM      string   `json:"nim"`
	NamaMhs  string   `json:"nama"`
	Matkul   []string `json:"matakuliah"`
	TotalSKS int      `json:"total_sks"`
}