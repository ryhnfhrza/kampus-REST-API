package mahasiswaRepository

import (
	"context"
	"database/sql"
	"kampus/model/domain"
)

type MahasiswaRepository interface {
	Create(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa) domain.Mahasiswa
	Update(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa) domain.Mahasiswa
	Delete(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa)
	FindByNim(ctx context.Context, tx *sql.Tx,mahasiswaNIM string)  (domain.Mahasiswa, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Mahasiswa
	FindMatkulMahasiswa(ctx context.Context, tx *sql.Tx,mahasiswaNIM string)(domain.Mahasiswa, error)
	FindMatkuldanDosen (ctx context.Context,tx *sql.Tx,mahasiswaNIM string)([]domain.MahasiswaMatkulDosen, error)
	NIMTerakhirMahasiswaPadaJurusan(ctx context.Context,tx *sql.Tx,kodeJurusan int,angkatan int)string
	CheckKelas(ctx context.Context,tx *sql.Tx,kodeKelas string,kodeJurusan int)bool
	JumlahMhsPadaKelas(ctx context.Context, tx *sql.Tx,kodeKelas string,kodeJurusan int)  (int, error)
	CreateKelas(ctx context.Context,tx *sql.Tx,kodeKelas string,kodeJurusan int)
}