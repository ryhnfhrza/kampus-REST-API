package mahasiswaWeb

type MahasiswaResponse struct {
	NIM      string `json:"nim"`
	Nama     string `json:"nama"`
	Gender   string `json:"gender"`
	Umur     int    `json:"umur"`
	Semester int    `json:"semester"`
}