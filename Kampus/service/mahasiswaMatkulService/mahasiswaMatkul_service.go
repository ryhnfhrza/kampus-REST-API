package Service

import (
	"context"
	web "kampus/model/web/mahasiswaMatakuliahWeb"
)

type MahasiswaMatkulService interface {
	AmbilMatkul(ctx context.Context, request web.MahasiswaMatkulCreateRequest) web.MahasiswaMatkulResponse
}