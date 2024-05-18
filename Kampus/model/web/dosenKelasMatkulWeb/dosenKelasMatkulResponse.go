package web

type DosenKelasMatkulResponse struct {
	IdDosen   int      `json:"id_dosen"`
	NamaDosen string   `json:"nama"`
	KodeKelas string   `json:"kode_kelas"`
	Matkul    []string `json:"matakuliah"`
}