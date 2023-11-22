package dosenService

import (
	"context"
	"kampus/model/web/dosenWeb"
)

type DosenService interface {
	Create(ctx context.Context, request dosenWeb.DosenCreateRequest) dosenWeb.DosenResponse
	Update(ctx context.Context, request dosenWeb.DosenUpdateRequest) dosenWeb.DosenResponse
	Delete(ctx context.Context, dosenId int)
	FindById(ctx context.Context, dosenId int) dosenWeb.DosenResponse
	FindAll(ctx context.Context) []dosenWeb.DosenResponse	
}