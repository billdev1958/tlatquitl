package usecases

import (
	"context"
	"user/internal/domain/catalogs"
	"user/pkg/error"
	"user/pkg/uuid"
)

func (uc *useCases) Create(ctx context.Context, d *catalogs.Dependency) (catalogs.Dependency, error) {
	dependencyID := uuid.NewUUID()

	if d.Name == "" {
		return catalogs.Dependency{}, apperror.ErrEmptyInput
	}

	d.ID = dependencyID

	createdDependency, err := uc.catalogsRepo.Create(ctx, d)
	if err != nil {
		return catalogs.Dependency{}, err
	}

	return createdDependency, nil

}
