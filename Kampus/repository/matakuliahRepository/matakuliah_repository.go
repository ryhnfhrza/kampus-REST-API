package matakuliahRepository

import (
	"context"
	"database/sql"
	"kampus/model/domain"
)

type MatakuliahRepository interface{
Create(ctx context.Context, tx *sql.Tx, matakuliah domain.Matakuliah) domain.Matakuliah
Update(ctx context.Context, tx *sql.Tx, matakuliah domain.Matakuliah) domain.Matakuliah
Delete(ctx context.Context, tx *sql.Tx, matakuliah domain.Matakuliah)
FindByKode(ctx context.Context, tx *sql.Tx,kodeMatakuliah string) (domain.Matakuliah, error)
FindAll(ctx context.Context, tx *sql.Tx) []domain.Matakuliah
}