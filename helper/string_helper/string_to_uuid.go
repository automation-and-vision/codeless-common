package string_helper

import (
	"github.com/google/uuid"
	"strings"
)

func StringToUUID(text string) (uuid.UUID, error) {
	id := strings.Trim(text, "/")

	return uuid.Parse(id)
}
