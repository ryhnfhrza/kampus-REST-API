package matakuliahService

import (
	"context"
	"kampus/model/web/matakuliahWeb"
)

type MatakuliahService interface {
	Create(ctx context.Context, request matakuliahWeb.MatakuliahCreateRequest) matakuliahWeb.MatakuliahResponse
	Update(ctx context.Context, request matakuliahWeb.MatakuliahUpdateRequest) matakuliahWeb.MatakuliahResponse
	Delete(ctx context.Context, matakuliahKode string)
	FindByKode(ctx context.Context, matakuliahKode string) matakuliahWeb.MatakuliahResponse
	FindAll(ctx context.Context) []matakuliahWeb.MatakuliahResponse	
}