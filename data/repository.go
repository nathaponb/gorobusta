package data

type Repository interface {
	GetByUsername(username string) (*User, error)
	Register(user *User) error
}
