package user

import "errors"

type User struct {
	username       string
	hashedPassword string
}

func (u User) Username() string {
	return u.username
}

func (u User) HashedPassword() string {
	return u.hashedPassword
}

func NewUser(username, hashedPassword string) (*User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if hashedPassword == "" {
		return nil, errors.New("hashed password cannot be empty")
	}
	return &User{
		username,
		hashedPassword,
	}, nil
}
