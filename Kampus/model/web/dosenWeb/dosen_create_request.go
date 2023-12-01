package dosenWeb

type DosenCreateRequest struct {
	Nama   string `validate:"required,max=100,min=1" json:"nama"`
	Gender string `validate:"required,oneof=pria wanita prefer_not_to" json:"gender"`
	Umur   int    `validate:"required,max=100,min=18" json:"umur"`
}