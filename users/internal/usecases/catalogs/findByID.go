package usecases

import (
	"context"
	"user/internal/domain/catalogs"
	"user/pkg/error"
	"user/pkg/types"
)

func (uc *useCases) FindByID(ctx context.Context, id types.UID) (catalogs.Dependency, error) {
	if id == "" {
		return catalogs.Dependency{}, apperror.ErrEmptyInput
	}

	catalog, err := uc.catalogsRepo.FindByID(ctx, id)
	if err != nil {
		return catalogs.Dependency{}, err
	}

	return catalog, nil
}
