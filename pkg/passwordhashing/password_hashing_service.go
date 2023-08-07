package passwordhashing

type PasswordHashingService interface {
	Hash(string) (string, error)
	Verify(hashedPassword, password string) error
}
