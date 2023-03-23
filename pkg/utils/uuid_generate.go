package utils

import (
	"strings"

	"github.com/google/uuid"
)

// GenerateSecurityToken - func generate uuid4 based security token without "-"
func GenerateUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
