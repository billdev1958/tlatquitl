package usecases

import (
	"context"
	"user/internal/domain/catalogs"
	"user/pkg/types"
)

type CatalogsUsecases interface {
	Create(ctx context.Context, d *catalogs.Dependency) (catalogs.Dependency, error)
	FindByID(ctx context.Context, id types.UID) (catalogs.Dependency, error)
}

type useCases struct {
	catalogsRepo catalogs.CatalogsRepository
}

func NewUsecases(catalogsRepo catalogs.CatalogsRepository) CatalogsUsecases {
	return &useCases{
		catalogsRepo: catalogsRepo,
	}
}
