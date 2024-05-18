package Service

import (
	"context"
	web "kampus/model/web/dosenKelasMatkulWeb"
)

type DosenKelasMatkulService interface {
	AjarMatkul(ctx context.Context, request web.DosenKelasMatkulCreateRequest) web.DosenKelasMatkulResponse
}