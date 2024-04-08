package jurusanWeb

type JurusanCreateRequest struct {
	KodeJurusan int    `validate:"required,max=999,min=100,number" json:"kode_jurusan"`
	NamaJurusan string `validate:"required,max=50,min=1" json:"jurusan"`
}