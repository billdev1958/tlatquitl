package account

import (
	"time"
	"user/pkg/types"
)

type Account struct {
	ID               types.UID
	UserID           types.UID
	Email            string
	Password         string
	IsVerified       bool
	Version          int
	Created_at       time.Time
	Updated_at       time.Time
	PasswordChangeAt time.Time
	Deleted_at       time.Time
}
