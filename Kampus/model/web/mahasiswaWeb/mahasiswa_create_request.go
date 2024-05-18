package mahasiswaWeb

type MahasiswaCreateRequest struct {
	Nama         string `validate:"required,max=100,min=1" json:"nama"`
	Gender       string `validate:"required,oneof=pria wanita prefer_not_to" json:"gender"`
	TanggalLahir string `validate:"required" json:"tanggal_lahir"`
	KodeJurusan  int    `validate:"required" json:"kode_jurusan"`
}