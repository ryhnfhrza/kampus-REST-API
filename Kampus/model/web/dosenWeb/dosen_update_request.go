package dosenWeb

type DosenUpdateRequest struct {
	Id   int    `validate:"required" json:"id"`
	Nama string `validate:"required,max=100,min=1" json:"nama"`
	Umur int    `validate:"required,max=100,min=18" json:"umur"`
}
