package dosenWeb

type DosenUpdateRequest struct {
	Id   int    `validate:"required"`
	Nama string `Validate:"required,max=100,min=1"`
	Umur int    `Validate:"required,max=100,min=18"`
}
