package util

import (
	"uuid"
)

func UuidGenerateV4() string {
	return uuid.NewV4().String()
}

func UuidGenerateV7() string {
	return uuid.NewV7().String()
}
