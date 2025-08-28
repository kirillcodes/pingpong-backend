package auth

import (
	"crypto/rand"
	"fmt"
)

func GenerateDigitCode() (string, error) {
	var b [4]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}

	n := (int(b[0])<<24 | int(b[1])<<16 | int(b[2])<<8 | int(b[3])) & 0xFFFFF
	code := fmt.Sprintf("%06d", n%1000000)
	return code, nil
}
