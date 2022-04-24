package utils

import "math/rand"

func GetBankTransferCode() string {
	length := 8
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := ""
	for i := 0; i < length; i += 1 {
		position := rand.Intn(len(characters) - 1)
		result += string(characters[position])
	}

	return result;
}