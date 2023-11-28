package dosenWeb

type DosenUpdateRequest struct {
	Id   int    `validate:"required" json:"id"`
	Nama string `Validate:"required,max=100,min=1" json:"nama"`
	Umur int    `Validate:"required,max=100,min=18" json:"umur"`
}
