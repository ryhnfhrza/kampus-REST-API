package mahasiswaWeb

type MahasiswaCreateRequest struct {
	NIM      string `Validate:"required,max=15,min=1"`
	Nama     string `Validate:"required,max=100,min=1"`
	Gender   string `validate:"required,oneof=pria wanita prefer_not_to"`
	Umur     int    `Validate:"required,max=100,min=10"`
	Semester int    `Validate:"required,max=10,min=1"`
}