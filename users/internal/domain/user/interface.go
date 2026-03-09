package user

import (
	"context"
	"user/pkg/types"
)

type UserRepository interface {
	Create(ctx context.Context, u *User) (User, error)
	Update(ctx context.Context, u *User) (User, error)
	FindByID(ctx context.Context, id types.UID) (User, error)
	FindAll(ctx context.Context, p types.Pagination, f types.UserFilters) ([]User, error)
	// TODO ChangePassword(ctx context.Context, password string) error
}
