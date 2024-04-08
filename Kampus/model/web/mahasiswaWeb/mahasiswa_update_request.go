package mahasiswaWeb

type MahasiswaUpdateRequest struct {
	NIM      string `validate:"required,max=15,min=1" json:"nim"`
	Semester int    `validate:"required,max=10,min=1" json:"semester"`
}