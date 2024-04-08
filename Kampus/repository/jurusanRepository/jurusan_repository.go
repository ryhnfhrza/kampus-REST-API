package jurusanRepository

import (
	"context"
	"database/sql"
	"kampus/model/domain"
)

type JurusanRepository interface{
	AddJurusan(ctx context.Context, tx *sql.Tx, jurusan domain.Jurusan) domain.Jurusan
	UpdateNamaJurusan(ctx context.Context, tx *sql.Tx, jurusan domain.Jurusan) domain.Jurusan
	DeleteJurusan(ctx context.Context, tx *sql.Tx, jurusan domain.Jurusan)
	FindJurusanByKode(ctx context.Context, tx *sql.Tx,jurusanKode int)  (domain.Jurusan, error)
	FindAllJurusan(ctx context.Context, tx *sql.Tx) []domain.Jurusan
}