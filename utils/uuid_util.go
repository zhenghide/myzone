package utils

import (
	"github.com/satori/go.uuid"
)

func GetUUID() string {
	id, _ := uuid.NewV4()
	return id.String()
}
