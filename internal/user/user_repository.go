package user

type UserRepository interface {
	Save(*User) error
	Get(username string) (*User, error)
}
