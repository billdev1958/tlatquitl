package account

import (
	"context"
	"user/pkg/types"
)

type AccountRepository interface {
	Create(ctx context.Context, a Account) error
	Update(ctx context.Context, a Account) error
	FindByID(ctx context.Context, id types.UID) (Account, error)
}
