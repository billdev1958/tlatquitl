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

type UserRoles struct {
	ID            types.UID
	UserID        types.UID
	RoleID        types.UID
	BranchScopeID types.UID
	CreatedAt     time.Time
}

type RolePermissions struct {
	ID           types.UID
	RoleID       types.UID
	PermissionID types.UID
	CreatedAt    time.Time
}

type UserRoleAssignments struct {
	ID         types.UID
	UserID     types.UID
	RoleID     types.UID
	ScopeType  string
	ScopeID    types.UID
	Effect     string
	ValidFrom  time.Time
	ValidUntil time.Time
	AssignedBy types.UID
	CreatedAt  time.Time
}
