package repository

import (
	"context"
	"database/sql"
	"kampus/model/domain"
)

type DosenKelasMatkulRepository interface{
	AjarMatkul(ctx context.Context, tx *sql.Tx, ajarMatkul domain.DosenKelasMatkul) domain.DosenKelasMatkul
}