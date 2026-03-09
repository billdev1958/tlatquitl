package uuid

import (
	"github.com/google/uuid"
	"user/pkg/types"
)

func NewUUID() types.UID {
	return types.UID(uuid.New().String())
}
