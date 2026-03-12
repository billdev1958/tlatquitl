package roles

import (
	"time"
	"user/pkg/types"
)

type Roles struct {
	ID          types.UID
	Code        string
	Name        string
	Description string
	CreatedAt   time.Time
}

type Permission struct {
	ID          types.UID
	Code        string
	Name        string
	Description string
	CreatedAt   time.Time
}

type RolePermissions struct {
	ID           types.UID
	RoleID       types.UID
	PermissionID types.UID
	CreatedAt    time.Time
}
