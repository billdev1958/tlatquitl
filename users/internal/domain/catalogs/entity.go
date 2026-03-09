package catalogs

import (
	"time"
	"user/pkg/types"
)

type Dependency struct {
	ID         types.UID
	Name       string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
