package web

type DosenKelasMatkulCreateRequest struct {
	IdDosen    int      `validate:"required" json:"id_dosen"`
	KodeKelas  string   `validate:"required" json:"kode_kelas"`
	KodeMatkul []string `validate:"required,dive,max=20,min=1" json:"kode_matkul"`
}