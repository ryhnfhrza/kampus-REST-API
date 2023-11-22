package helper

import (
	"kampus/model/domain"
	"kampus/model/web/dosenWeb"
	"kampus/model/web/mahasiswaWeb"
	"kampus/model/web/matakuliahWeb"
)

func ToMahasiswaResponse(mahasiswa domain.Mahasiswa)mahasiswaWeb.MahasiswaResponse{
	return mahasiswaWeb.MahasiswaResponse{
		NIM: mahasiswa.NIM,
		Nama: mahasiswa.Nama,
		Gender: mahasiswa.Gender,
		Umur: mahasiswa.Umur,
		Semester: mahasiswa.Semester,
	}
}

func ToMahasiswaResponses(mahasiswa []domain.Mahasiswa)[]mahasiswaWeb.MahasiswaResponse{
	var mahasiswaResponses []mahasiswaWeb.MahasiswaResponse
	for _,mhs := range mahasiswa{
		mahasiswaResponses = append(mahasiswaResponses, ToMahasiswaResponse(mhs))
	}
	return mahasiswaResponses
}

func ToDosenResponse(dosen domain.Dosen)dosenWeb.DosenResponse{
	return dosenWeb.DosenResponse{
		Id: dosen.Id,
		Nama: dosen.Nama,
		Gender: dosen.Gender,
		Umur: dosen.Umur,
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