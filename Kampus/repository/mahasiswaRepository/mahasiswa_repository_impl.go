package mahasiswaRepository

import (
	"context"
	"database/sql"
	"errors"
	"kampus/helper"
	"kampus/model/domain"
)

type MahasiswaRepositoryImpl struct{

}

func NewMahasiswaRepositoryImpl()MahasiswaRepository{
	return &MahasiswaRepositoryImpl{}
}

func(Repository *MahasiswaRepositoryImpl)Create(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa) domain.Mahasiswa{
	SQL := "insert into mahasiswa (nim,nama,gender,tanggal_lahir,semester,kode_jurusan,angkatan,kode_kelas) values (?,?,?,?,?,?,?,?)"
	_,err := tx.ExecContext(ctx,SQL,mahasiswa.NIM,mahasiswa.Nama,mahasiswa.Gender,mahasiswa.TanggalLahir,mahasiswa.Semester,mahasiswa.KodeJurusan,mahasiswa.Angkatan,mahasiswa.KodeKelas)
	helper.PanicIfError(err)

	return mahasiswa
}

func(Repository *MahasiswaRepositoryImpl)Update(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa) domain.Mahasiswa{
	SQL := "update mahasiswa set semester = ?,kode_kelas = ? where nim = ?"
	_,err := tx.ExecContext(ctx,SQL,mahasiswa.Semester,mahasiswa.KodeKelas,mahasiswa.NIM)
	helper.PanicIfError(err)

	return mahasiswa
}

func(Repository *MahasiswaRepositoryImpl)Delete(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa){
	SQL := "delete from mahasiswa where nim = ?"
	_,err := tx.ExecContext(ctx,SQL,mahasiswa.NIM)
	helper.PanicIfError(err)
}

func(Repository *MahasiswaRepositoryImpl)FindByNim(ctx context.Context, tx *sql.Tx, mahasiswaNIM string) (domain.Mahasiswa, error){
	SQL := "select nim,nama,gender,tanggal_lahir,semester,kode_jurusan,angkatan,kode_kelas from mahasiswa where nim = ?"
	rows,err := tx.QueryContext(ctx,SQL,mahasiswaNIM)
	helper.PanicIfError(err)

	defer rows.Close()

	mahasiswa := domain.Mahasiswa{}
	var kodeKelas sql.NullString
	if rows.Next(){
		err := rows.Scan(&mahasiswa.NIM,&mahasiswa.Nama,&mahasiswa.Gender,&mahasiswa.TanggalLahir,&mahasiswa.Semester,&mahasiswa.KodeJurusan,&mahasiswa.Angkatan,&kodeKelas)
		helper.PanicIfError(err)
		if kodeKelas.Valid{
			mahasiswa.KodeKelas = kodeKelas.String
		}
		return mahasiswa,nil
	}else{
		return mahasiswa,errors.New("NIM tidak ditemukan")
	}
}

func(Repository *MahasiswaRepositoryImpl)FindAll(ctx context.Context, tx *sql.Tx) []domain.Mahasiswa{
	SQL := "select nim,nama,gender,tanggal_lahir,semester,kode_jurusan,angkatan,kode_kelas from mahasiswa"
	rows,err := tx.QueryContext(ctx,SQL)
	helper.PanicIfError(err)
	defer rows.Close()
	
	var mahasiswa []domain.Mahasiswa
	
	for rows.Next(){
		mhs := domain.Mahasiswa{}
		var kodeKelas sql.NullString
		err := rows.Scan(&mhs.NIM,&mhs.Nama,&mhs.Gender,&mhs.TanggalLahir,&mhs.Semester,&mhs.KodeJurusan,&mhs.Angkatan,&kodeKelas)
		helper.PanicIfError(err)
		if kodeKelas.Valid{
			mhs.KodeKelas = kodeKelas.String
		}
		mahasiswa = append(mahasiswa, mhs)
	}
	return mahasiswa
}

func(Repository *MahasiswaRepositoryImpl)FindMatkulMahasiswa(ctx context.Context, tx *sql.Tx, mahasiswaNIM string) (domain.Mahasiswa, error){
	SQL := "select nim,nama,gender,tanggal_lahir,semester,kode_jurusan,angkatan from mahasiswa where nim = ?"
	rows,err := tx.QueryContext(ctx,SQL,mahasiswaNIM)
	helper.PanicIfError(err)

	defer rows.Close()

	mahasiswa := domain.Mahasiswa{}
	if rows.Next(){
		err := rows.Scan(&mahasiswa.NIM,&mahasiswa.Nama,&mahasiswa.Gender,&mahasiswa.TanggalLahir,&mahasiswa.Semester,&mahasiswa.KodeJurusan,&mahasiswa.Angkatan)
		helper.PanicIfError(err)
		return mahasiswa,nil
	}else{
		return mahasiswa,errors.New("NIM tidak ditemukan")
	}
}

func (Repository *MahasiswaRepositoryImpl) FindMatkuldanDosen(ctx context.Context, tx *sql.Tx, mahasiswaNIM string) ([]domain.MahasiswaMatkulDosen, error) {
	SQL := "SELECT m.nim, m.nama, m.semester, m.kode_jurusan, j.nama AS nama_jurusan, m.angkatan, m.kode_kelas, mm.kode_matakuliah, mk.matakuliah AS nama_matakuliah, (SELECT d.nama FROM dosen_kelas_matakuliah dkm INNER JOIN dosen d ON dkm.id_dosen = d.id WHERE dkm.kode_kelas = m.kode_kelas AND dkm.kode_matakuliah = mm.kode_matakuliah LIMIT 1) AS nama_dosen FROM mahasiswa m INNER JOIN jurusan j ON m.kode_jurusan = j.kode INNER JOIN mahasiswa_matakuliah mm ON m.nim = mm.nim INNER JOIN matakuliah mk ON mm.kode_matakuliah = mk.kode WHERE m.nim = ?"
	rows, err := tx.QueryContext(ctx, SQL, mahasiswaNIM)
	helper.PanicIfError(err)
	defer rows.Close()

	var mahasiswas []domain.MahasiswaMatkulDosen
	for rows.Next() {
			var mahasiswa domain.MahasiswaMatkulDosen
			err := rows.Scan(&mahasiswa.NIM, &mahasiswa.NamaMahasiswa, &mahasiswa.Semester, &mahasiswa.KodeJurusan, &mahasiswa.Jurusan, &mahasiswa.Angkatan, &mahasiswa.KodeKelas, &mahasiswa.KodeMatakuliah, &mahasiswa.Matakuliah, &mahasiswa.NamaDosen)
			helper.PanicIfError(err)
			mahasiswas = append(mahasiswas, mahasiswa)
	}

	if len(mahasiswas) == 0 {
			return nil, errors.New("NIM tidak ditemukan")
	}

	return mahasiswas, nil
}


func(repository *MahasiswaRepositoryImpl)NIMTerakhirMahasiswaPadaJurusan(ctx context.Context,tx *sql.Tx,kodeJurusan int , angkatan int)string{
	SQL := "SELECT nim FROM mahasiswa WHERE kode_jurusan = ? AND angkatan = ? ORDER BY nim DESC LIMIT 1;"
	rows , err := tx.QueryContext(ctx,SQL,kodeJurusan,angkatan)
	helper.PanicIfError(err)
	defer rows.Close()

	var LastInsertNim string
	if rows.Next(){
		err := rows.Scan(&LastInsertNim)
		helper.PanicIfError(err)
		return LastInsertNim
	}else{
		return ""
	}

}

func(repository *MahasiswaRepositoryImpl)CheckKelas(ctx context.Context,tx *sql.Tx,kodeKelas string,kodeJurusan int)bool{
	SQL := "SELECT EXISTS (SELECT 1 FROM mahasiswa WHERE kode_kelas = ? and kode_jurusan = ? LIMIT 1);"
	rows,err := tx.QueryContext(ctx,SQL,kodeKelas,kodeJurusan)
	helper.PanicIfError(err)

	defer rows.Close()

	var exists bool
	if rows.Next() {
		err := rows.Scan(&exists)
		helper.PanicIfError(err)
	}

	return exists
}


func(repository *MahasiswaRepositoryImpl)JumlahMhsPadaKelas(ctx context.Context, tx *sql.Tx,kodeKelas string, kodeJurusan int)  (int, error){
	SQL := "select count(*) from mahasiswa where kode_kelas = ? and kode_jurusan = ?"
	rows,err := tx.QueryContext(ctx,SQL,kodeKelas,kodeJurusan)
	helper.PanicIfError(err)

	defer rows.Close()

	var mhsKelas int
	if rows.Next(){
		err := rows.Scan(&mhsKelas)
		helper.PanicIfError(err)
		return mhsKelas,nil
	}else{
		return mhsKelas,errors.New("kelas tidak ditemukan")
	}
}

func(repository *MahasiswaRepositoryImpl)CreateKelas(ctx context.Context,tx *sql.Tx,kodeKelas string,kodeJurusan int){
	SQL := "insert into kelas (kode_kelas,kode_jurusan) values (?,?)"
	_,err := tx.ExecContext(ctx,SQL,kodeKelas,kodeJurusan)
	helper.PanicIfError(err)
}


