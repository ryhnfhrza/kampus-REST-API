package dosenWeb

type DosenResponse struct {
	Id     int    `json:"id"`
	Nama   string `json:"nama"`
	Gender string `json:"gender"`
	Umur   int    `json:"umur"`
}