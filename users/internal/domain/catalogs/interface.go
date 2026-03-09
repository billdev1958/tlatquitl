package catalogs

import (
	"context"
	"user/pkg/types"
)

type CatalogsRepository interface {
	Create(ctx context.Context, d *Dependency) (Dependency, error)
	FindByID(ctx context.Context, id types.UID) (Dependency, error)
	// TODO FindAll(ctx context.Context) ([]Dependency, error)
}
