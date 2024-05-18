package mahasiswaWeb

type MahasiswaUpdateRequest struct {
	NIM       string `validate:"required,max=15,min=1" json:"nim"`
	Semester  int    `validate:"max=10" json:"semester"`
	KodeKelas string ` json:"kode_kelas"`
}