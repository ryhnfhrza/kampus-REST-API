package mahasiswa

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
}