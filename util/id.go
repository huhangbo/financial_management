package util

import (
	"github.com/google/uuid"
)

func GenID() int {
	return int(uuid.New().ID())
}
