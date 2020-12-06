package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

func IDStringToUint64(id string) uint64 {
	n, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Println("[Error] IDStringToUint64", err)
		return 0
	}
	return n
}

func CheckBCrypt(hashed, raw string) error {
	rawByte := []byte(raw)
	hashedByte := []byte(hashed)
	return bcrypt.CompareHashAndPassword(hashedByte, rawByte)
}

func HashBCrypt(raw string) (string, error) {
	rawByte := []byte(raw)
	hashed, err := bcrypt.GenerateFromPassword(rawByte, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
