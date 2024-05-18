package helper

import (
	"fmt"
	"kampus/model/domain"
	webDosenMatkul "kampus/model/web/dosenKelasMatkulWeb"
	"kampus/model/web/dosenWeb"
	"kampus/model/web/jurusanWeb"
	webMahasiswaMatkul "kampus/model/web/mahasiswaMatakuliahWeb"
	"kampus/model/web/mahasiswaWeb"
	"kampus/model/web/matakuliahWeb"
	"strings"
)

func ToMahasiswaResponse(mahasiswa domain.Mahasiswa)mahasiswaWeb.MahasiswaResponse{
	return mahasiswaWeb.MahasiswaResponse{
		NIM: mahasiswa.NIM,
		Nama: mahasiswa.Nama,
		Gender: mahasiswa.Gender,
		TanggalLahir: mahasiswa.TanggalLahir,
		Semester: mahasiswa.Semester,
		KodeJurusan: mahasiswa.KodeJurusan,
		Angkatan: mahasiswa.Angkatan,
		KodeKelas: mahasiswa.KodeKelas,
	}
}

func ToMahasiswaResponses(mahasiswa []domain.Mahasiswa)[]mahasiswaWeb.MahasiswaResponse{
	var mahasiswaResponses []mahasiswaWeb.MahasiswaResponse
	for _,mhs := range mahasiswa{
		mahasiswaResponses = append(mahasiswaResponses, ToMahasiswaResponse(mhs))
	}
	return mahasiswaResponses
}

func ToMahasiswaMatkulDosenResponse(mahasiswa []domain.MahasiswaMatkulDosen) mahasiswaWeb.MahasiswaMatkulDosenResponse {
    uniqueEntries := make(map[string]bool)
    var combinedEntries []string

    for _, m := range mahasiswa {
        key := m.KodeMatakuliah + m.Matakuliah + " - " + m.NamaDosen
        if _, ok := uniqueEntries[key]; !ok {
            uniqueEntries[key] = true
            combinedEntries = append(combinedEntries, "["+m.KodeMatakuliah+"] "+m.Matakuliah+" - "+m.NamaDosen)
        }
    }

    firstEntry := mahasiswa[0]

    // Menggabungkan semua string dalam slice menjadi satu string dengan pemisah ", "
    matkulString := strings.Join(combinedEntries, ", ")

    return mahasiswaWeb.MahasiswaMatkulDosenResponse{
        NIM:           firstEntry.NIM,
        NamaMahasiswa: firstEntry.NamaMahasiswa,
        Semester:      firstEntry.Semester,
        KodeJurusan:   firstEntry.KodeJurusan,
        Jurusan:       firstEntry.Jurusan,
        Angkatan:      firstEntry.Angkatan,
        KodeKelas:     firstEntry.KodeKelas,
        Matkul:        []string{matkulString}, // Menggunakan string yang telah digabungkan dalam slice string tunggal
    }
}



func ToDosenResponse(dosen domain.Dosen)dosenWeb.DosenResponse{
	return dosenWeb.DosenResponse{
		Id: dosen.Id,
		Nama: dosen.Nama,
		Gender: dosen.Gender,
		TanggalLahir: dosen.TanggalLahir,
	}

}

func ToDosenResponses(dosen []domain.Dosen)[]dosenWeb.DosenResponse{
	var dosenResponses []dosenWeb.DosenResponse
	for _,dsn := range dosen{
		dosenResponses = append(dosenResponses, ToDosenResponse(dsn))
	}
	return dosenResponses
}

func ToMatakuliahResponse(matakuliah domain.Matakuliah)matakuliahWeb.MatakuliahResponse{
	return matakuliahWeb.MatakuliahResponse{
		Kode: matakuliah.Kode,
		Matakuliah: matakuliah.Mata_kuliah,
		SKS: matakuliah.SKS,
	}

}

func ToMatakuliahResponses(matakuliah []domain.Matakuliah)[]matakuliahWeb.MatakuliahResponse{
	var matakuliahResponses []matakuliahWeb.MatakuliahResponse
	for _,mk := range matakuliah{
		matakuliahResponses = append(matakuliahResponses, ToMatakuliahResponse(mk))
	}
	return matakuliahResponses
}

func ToJurusanResponse(jurusan domain.Jurusan)jurusanWeb.JurusanWebResponse{
	return jurusanWeb.JurusanWebResponse{
		KodeJurusan: jurusan.Kode,
		NamaJurusan: jurusan.NamaJurusan,
	}

}

func ToJurusanResponses(jurusan []domain.Jurusan)[]jurusanWeb.JurusanWebResponse{
	var jurusanResponses []jurusanWeb.JurusanWebResponse
	for _,jsn := range jurusan{
		jurusanResponses = append(jurusanResponses, ToJurusanResponse(jsn))
	}
	return jurusanResponses
}

func ToMahasiswaMatkulResponse(mahasiswaMatkuls []domain.MahasiswaMatkul, matakuliah []domain.Matakuliah, mahasiswa domain.Mahasiswa) webMahasiswaMatkul.MahasiswaMatkulResponse {
	var matkulDetails []string
	var totalSKS int

	for i:= range mahasiswaMatkuls {
			matkulDetails = append(matkulDetails, fmt.Sprintf("%s - %s - %d", matakuliah[i].Kode, matakuliah[i].Mata_kuliah, matakuliah[i].SKS))
			
			totalSKS += matakuliah[i].SKS
	}

	return webMahasiswaMatkul.MahasiswaMatkulResponse{
			NIM:      mahasiswa.NIM,
			NamaMhs:  mahasiswa.Nama,
			Matkul:   matkulDetails,
			TotalSKS: totalSKS,
	}
}

func ToDosenMatkulResponse(dosenMatkuls []domain.DosenKelasMatkul, matakuliah []domain.Matakuliah, dosen domain.Dosen ,dosenKelasMatkul domain.DosenKelasMatkul) webDosenMatkul.DosenKelasMatkulResponse {
	var matkulDetails []string

	for i:= range dosenMatkuls {
			matkulDetails = append(matkulDetails, fmt.Sprintf("%s - %s - %d", matakuliah[i].Kode, matakuliah[i].Mata_kuliah, matakuliah[i].SKS))
			
	}

	return webDosenMatkul.DosenKelasMatkulResponse{
			IdDosen: dosen.Id,
			NamaDosen: dosen.Nama,
			Matkul:   matkulDetails,
			KodeKelas: dosenKelasMatkul.KodeKelas,
	}
}