package mahasiswaWeb

type MahasiswaMatkulDosenResponse struct {
	NIM           string   `json:"nim"`
	NamaMahasiswa string   `json:"nama"`
	Semester      int      `json:"semester"`
	KodeJurusan   int      `json:"kode_jurusan"`
	Jurusan       string   `json:"jurusan"`
	Angkatan      int      `json:"angkatan"`
	KodeKelas     string   `json:"kode_kelas"`
	Matkul        []string `json:"matakuliah"`
}