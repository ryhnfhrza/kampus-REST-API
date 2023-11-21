package dosen

import (
	"context"
	"database/sql"
	"kampus/model/domain"
)

type DosenRepository interface{
	Create(ctx context.Context, tx *sql.Tx, dosen domain.Dosen) domain.Dosen
	Update(ctx context.Context, tx *sql.Tx, dosen domain.Dosen) domain.Dosen
	Delete(ctx context.Context, tx *sql.Tx, dosen domain.Dosen)
	FindById(ctx context.Context, tx *sql.Tx,dosenId int)  (domain.Dosen, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Dosen
}