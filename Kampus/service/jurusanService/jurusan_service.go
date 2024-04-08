package jurusanService

import (
	"context"
	"kampus/model/web/jurusanWeb"
)

type JurusanService interface{
	AddJurusan(ctx context.Context, request jurusanWeb.JurusanCreateRequest) jurusanWeb.JurusanWebResponse
	UpdateNamaJurusan(ctx context.Context, request jurusanWeb.JurusanUpdateRequest) jurusanWeb.JurusanWebResponse
	DeleteJurusan(ctx context.Context, jurusanKode int)
	FindJurusanByKode(ctx context.Context, jurusanKode int) jurusanWeb.JurusanWebResponse
	FindAllJurusan(ctx context.Context) []jurusanWeb.JurusanWebResponse	
}