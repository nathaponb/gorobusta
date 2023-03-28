package user

type UserRepository interface {
	GetByUsername(username string) (*User, error)
	Register(user *User) error
}
