package utils

import (
	"strings"

	"github.com/google/uuid"
)

// generate a new uuid without "-"
func NewUUID() string {
	u2 := uuid.New().String()
	u3 := strings.ReplaceAll(u2, "-", "")
	return u3
}

func NewRandomUUID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		id = uuid.New()
	}
	return strings.ReplaceAll(id.String(), "-", "")
}
