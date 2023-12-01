package mahasiswaWeb

type MahasiswaUpdateRequest struct {
	NIM      string `validate:"required,max=15,min=1" json:"nim"`
	Nama     string `validate:"required,max=100,min=1" json:"nama"`
	Umur     int    `validate:"required,max=100,min=10" json:"umur"`
	Semester int    `validate:"required,max=10,min=1" json:"semester"`
}