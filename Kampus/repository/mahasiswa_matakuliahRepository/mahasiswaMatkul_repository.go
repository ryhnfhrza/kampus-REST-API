package repository

import (
	"context"
	"database/sql"
	"kampus/model/domain"
)

type MahasiswaMatkulRepository interface{
	AmbilMatkul(ctx context.Context, tx *sql.Tx, ambilMatkul domain.MahasiswaMatkul) domain.MahasiswaMatkul
}