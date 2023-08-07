package user

import "go-bdd-user/pkg/passwordhashing"

type UserService struct {
	userRepository         UserRepository
	passwordHashingService passwordhashing.PasswordHashingService
}

func (s *UserService) Register(username, password string) error {
	hashedPassword, err := s.passwordHashingService.Hash(password)
	if err != nil {
		return err
	}

	user, err := NewUser(username, hashedPassword)
	if err != nil {
		return err
	}

	return s.userRepository.Save(user)
}

func (s *UserService) Login(username, password string) error {
	user, err := s.userRepository.Get(username)
	if err != nil {
		return err
	}

	return s.passwordHashingService.Verify(user.HashedPassword(), password)
}

func NewUserService(
	userRepository UserRepository,
	passwordHashingService passwordhashing.PasswordHashingService,
) *UserService {
	return &UserService{userRepository, passwordHashingService}
}
