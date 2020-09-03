package utils

import "golang.org/x/crypto/bcrypt"

func CreateBcryptHashString(plaintext string) (hashtext string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), 14)
	if err != nil {
		return "", err
	}

	return string(hash[:]), nil
}

// CheckUserSecret checks a plaintext password against a password hash. nil error on success.
func CheckUserSecret(hashedSecret, plainSecret string) (bool, error) {
	byteSecret := []byte(plainSecret)
	byteHash := []byte(hashedSecret)
	err := bcrypt.CompareHashAndPassword(byteHash, byteSecret)
	if err != nil {
		return false, err
	}

	return true, nil
}
