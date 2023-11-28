package dosenWeb

type DosenCreateRequest struct {
	Nama   string `Validate:"required,max=100,min=1" json:"nama"`
	Gender string `validate:"required,oneof=pria wanita prefer_not_to" json:"gender"`
	Umur   int    `Validate:"required,max=100,min=18" json:"umur"`
}