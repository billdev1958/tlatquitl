package usecases

import (
	"context"
	"user/internal/domain/user"
	"user/pkg/error"
)

func (uc *useCases) Update(ctx context.Context, u *user.User) (user.User, error) {
	oldUser, err := uc.userRepo.FindByID(ctx, u.ID)
	if err != nil {
		return user.User{}, err
	}

	if u.Name == oldUser.Name &&
		u.Lastname1 == oldUser.Lastname1 &&
		u.Lastname2 == oldUser.Lastname2 {
		return user.User{}, apperror.ErrNoChanges
	}

	updatedUser, err := uc.userRepo.Update(ctx, u)
	if err != nil {
		return user.User{}, nil
	}

	return updatedUser, nil
}
