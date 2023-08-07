package passwordhashing

import (
	"golang.org/x/crypto/bcrypt"
)

type bcryptPasswordHashingService struct {
}

func (s *bcryptPasswordHashingService) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *bcryptPasswordHashingService) Verify(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func NewBcryptPasswordHashingService() PasswordHashingService {
	return new(bcryptPasswordHashingService)
}
