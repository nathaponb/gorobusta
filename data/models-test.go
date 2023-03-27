package data

import "gorm.io/gorm"

type TestPostgresRepository struct {
	Conn *gorm.DB
}

func NewTestPostgresRepository(db *gorm.DB) *TestPostgresRepository {
	return &TestPostgresRepository{
		Conn: db,
	}
}

func (g *TestPostgresRepository) GetByUsername(username string) (*User, error) {
	user := &User{
		FirstName:   "user",
		LastName:    "test",
		DateOfBirth: "2023-03-27",
		UserName:    "test01",
		Email:       "test@dev.com",
		Password:    "12345",
	}
	return user, nil
}

func (g *TestPostgresRepository) Register(user *User) error {
	return nil
}
