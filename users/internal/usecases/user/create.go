package usecases

import (
	"context"
	"user/internal/domain/user"
	"user/pkg/error"
	"user/pkg/uuid"
)

func (uc *useCases) Create(ctx context.Context, u *user.User) (user.User, error) {
	userID := uuid.NewUUID()

	if u.Name == "" {
		return user.User{}, apperror.ErrEmptyInput
	}

	if u.Lastname1 == "" {
		return user.User{}, apperror.ErrEmptyInput
	}

	if u.Lastname2 == "" {
		return user.User{}, apperror.ErrEmptyInput
	}

	u.ID = userID

	createdUser, err := uc.userRepo.Create(ctx, u)
	if err != nil {
		return user.User{}, err
	}

	return createdUser, nil

}
