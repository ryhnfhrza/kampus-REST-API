package mahasiswaWeb

type MahasiswaUpdateRequest struct {
	NIM      string `Validate:"required,max=15,min=1"`
	Nama     string `Validate:"required,max=100,min=1"`
	Umur     int    `Validate:"required,max=100,min=10"`
	Semester int    `Validate:"required,max=10,min=1"`
}