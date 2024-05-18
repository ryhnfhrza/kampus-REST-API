package mahasiswaService

import (
	"context"
	"kampus/model/web/mahasiswaWeb"
)

type MahasiswaService interface {
	Create(ctx context.Context, request mahasiswaWeb.MahasiswaCreateRequest)mahasiswaWeb.MahasiswaResponse
	Update(ctx context.Context,request mahasiswaWeb.MahasiswaUpdateRequest) mahasiswaWeb.MahasiswaResponse
	Delete(ctx context.Context, mahasiswaNIM string)
	FindByNIM(ctx context.Context,mahasiswaNIM string)mahasiswaWeb.MahasiswaResponse
	FindAll(ctx context.Context)[]mahasiswaWeb.MahasiswaResponse
	FindMahasiswaMatkulDosen(ctx context.Context,mahasiswaNIM string)mahasiswaWeb.MahasiswaMatkulDosenResponse
}