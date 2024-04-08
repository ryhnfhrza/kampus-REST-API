package jurusanWeb

type JurusanUpdateRequest struct {
	KodeJurusan int    `validate:"required,max=999,min=100,numeric" json:"kode_jurusan"`
	NamaJurusan string `validate:"required,max=50,min=1,alpha" json:"jurusan"`
}