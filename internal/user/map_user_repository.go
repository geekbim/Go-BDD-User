package user

import "errors"

type mapUserRepository struct {
	users map[string]string
}

func (r *mapUserRepository) Save(u *User) error {
	if _, ok := r.users[u.Username()]; ok {
		return errors.New("user is already exists")
	}

	r.users[u.Username()] = u.HashedPassword()

	return nil
}

func (r *mapUserRepository) Get(username string) (*User, error) {
	if hashedPassword, ok := r.users[username]; ok {
		return &User{
			username,
			hashedPassword,
		}, nil
	}

	return nil, errors.New("user is not exists")
}

func NewMapUserRepository() UserRepository {
	return &mapUserRepository{
		make(map[string]string),
	}
}
