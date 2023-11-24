package dosenWeb

type DosenCreateRequest struct {
	Nama   string `Validate:"required,max=100,min=1"`
	Gender string `validate:"required,oneof=pria wanita prefer_not_to"`
	Umur   int    `Validate:"required,max=100,min=18"`
}