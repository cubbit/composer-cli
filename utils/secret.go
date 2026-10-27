package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSecret() (string, error) {
	secretBytes := make([]byte, 64)
	_, err := rand.Read(secretBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(secretBytes), nil
}
