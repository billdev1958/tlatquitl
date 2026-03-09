package user

import (
	"time"
	"user/pkg/types"
)

type User struct {
	ID         types.UID
	Name       string
	Lastname1  string
	Lastname2  string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
