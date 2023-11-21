package matakuliah

import (
	"context"
	"database/sql"
	"kampus/model/domain"
)

type Matakuliah interface{
Create(ctx context.Context, tx *sql.Tx, matakuliah domain.Matakuliah) domain.Matakuliah
Update(ctx context.Context, tx *sql.Tx, matakuliah domain.Matakuliah) domain.Matakuliah
Delete(ctx context.Context, tx *sql.Tx, matakuliah domain.Matakuliah)
FindByKode(ctx context.Context, tx *sql.Tx,kodeMatakuliah int) (domain.Matakuliah, error)
FindAll(ctx context.Context, tx *sql.Tx) []domain.Matakuliah
}