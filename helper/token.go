package helper

import (
	"crypto/rand"
	"fmt"
)

func MakeToken() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", b)
}
