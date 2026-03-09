package usecases

import (
	"context"
	"fmt"
	"user/internal/domain/user"
	"user/pkg/types"
)

func (uc *useCases) FindAll(ctx context.Context, pagination types.Pagination, filters types.UserFilters) ([]user.User, error) {
	if pagination.Limit <= 0 {
		pagination.Limit = 10
	}

	if pagination.Offset < 0 {
		pagination.Offset = 0
	}

	users, err := uc.userRepo.FindAll(ctx, pagination, filters)
	if err != nil {
		return []user.User{}, fmt.Errorf("error en la búsqueda filtrada: %w", err)
	}

	return users, nil
}
