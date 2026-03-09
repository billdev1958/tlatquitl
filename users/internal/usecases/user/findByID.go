package usecases

import (
	"context"
	"fmt"
	"user/internal/domain/user"
	"user/pkg/error"
	"user/pkg/types"
)

func (uc *useCases) FindByID(ctx context.Context, id types.UID) (user.User, error) {
	if id == "" {
		return user.User{}, apperror.ErrEmptyInput
	}

	u, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return user.User{}, fmt.Errorf("usuario no encontrado con ID %s: %w", id, err)
	}

	return u, nil
}
