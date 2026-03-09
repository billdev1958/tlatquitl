package account

import (
	"time"
	"user/pkg/types"
)

type Account struct {
	ID               types.UID
	UserID           types.UID
	DependencyID     int
	Email            string
	Password         string
	RoleID           int
	IsVerified       bool
	Version          int
	Created_at       time.Time
	Updated_at       time.Time
	PasswordChangeAt time.Time
	Deleted_at       time.Time
}
