package hash

import (
	"mygram/commons/exceptions"

	"golang.org/x/crypto/bcrypt"
)

func GetHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		exceptions.PanicIfNeeded(err)
	}
	return string(hash)
}