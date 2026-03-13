package usecases

import (
	"context"
	"user/internal/domain/account"
	"user/internal/domain/user"
	"user/pkg/types"
)

type UserUsecases interface {
	Create(ctx context.Context, u *user.User, a *account.Account) (user.User, error)
	Update(ctx context.Context, u *user.User) (user.User, error)
	FindByID(ctx context.Context, id types.UID) (user.User, error)
	FindAll(ctx context.Context, pagination types.Pagination, filters types.UserFilters) ([]user.User, error)
	// TODO ChangePassword(ctx context.Context, password string) error
}

type useCases struct {
	userRepo user.UserRepository
}

func NewUsecases(userRepo user.UserRepository) UserUsecases {
	return &useCases{
		userRepo: userRepo,
	}
}
