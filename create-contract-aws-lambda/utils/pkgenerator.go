package utils

import "github.com/gofrs/uuid"

func UUIDStringGenerator() string {
	pk, _ := uuid.NewV4()
	return pk.String()
}
