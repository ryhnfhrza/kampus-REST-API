package mahasiswaWeb

type MahasiswaCreateRequest struct {
	NIM      string `Validate:"required,max=15,min=1" json:"nim"`
	Nama     string `Validate:"required,max=100,min=1" json:"nama"`
	Gender   string `validate:"required,oneof=pria wanita prefer_not_to" json:"gender"`
	Umur     int    `Validate:"required,max=100,min=10" json:"umur"`
	Semester int    `Validate:"required,max=10,min=1" json:"semester"`
}