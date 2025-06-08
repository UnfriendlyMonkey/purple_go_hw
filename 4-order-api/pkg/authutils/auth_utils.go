package authutils

import (
	"fmt"
	"math/rand"
)

func GenerateCode() string {
	code := fmt.Sprintf("%04d", rand.Intn(9999))
	return code
}

func GenerateSessionID() string {
	const (
		letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numberBytes = "0123456789"
		length      = 12
	)

	b := make([]byte, length)
	for i := range b {
		if i%2 == 0 {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		} else {
			b[i] = numberBytes[rand.Intn(len(numberBytes))]
		}
	}
	return string(b)
}

func SendCode(phone, code string) (bool, error) {
	// imitates sending code to phone
	fmt.Printf("Sending code %s to %s\n", code, phone)
	return true, nil
}
