package providers

import (
	"bytes"
	"crypto/sha512"
	"errors"

	error "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	errCreateHash = errors.New("error to create hash")
)

type (
	HashProvider interface {
		GenerateHash(string) string
		CompareHash(string, string) bool
	}

	hashProviderContainer struct{}
)

func NewHashProvider() HashProvider {
	return hashProviderContainer{}
}

func (pro hashProviderContainer) GenerateHash(payload string) string {
	hashedInput := sha512.Sum512_256([]byte(payload))

	trimmedHash := bytes.Trim(hashedInput[:], "\x00")

	preparedPassword := string(trimmedHash)

	passwordHashInBytes, err := bcrypt.GenerateFromPassword([]byte(preparedPassword), bcrypt.DefaultCost)
	if err != nil {
		error.Wrap(err, errCreateHash.Error())
	}

	return string(passwordHashInBytes)
}

func (pro hashProviderContainer) CompareHash(payload string, hashed string) bool {
	hashedInput := sha512.Sum512_256([]byte(payload))
	trimmedHash := bytes.Trim(hashedInput[:], "\x00")
	preparedPassword := string(trimmedHash)

	plainTextInBytes := []byte(preparedPassword)
	hashTextInBytes := []byte(hashed)

	err := bcrypt.CompareHashAndPassword(hashTextInBytes, plainTextInBytes)
	if err != nil {
		error.Wrap(err, errCreateHash.Error())
		return false
	}

	return true
}
