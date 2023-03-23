package data

import "gorm.io/gorm"

type User struct {
	FirstName   string `json:"first_name,omitempty" gorm:"column:first_name" binding:"required"`
	LastName    string `json:"last_name,omitempty" gorm:"column:last_name" binding:"required"`
	DateOfBirth string `json:"date_of_birth,omitempty" gorm:"column:date_of_birth" binding:"required"`
	UserName    string `json:"username,omitempty" gorm:"column:username" binding:"required"`
	Email       string `json:"email,omitempty" gorm:"column:email" binding:"required"`
	Password    string `json:"password,omitempty" gorm:"column:password" binding:"required"`
}

type PostgresRepository struct {
	Conn *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		Conn: db,
	}
}

func (g *PostgresRepository) GetByUsername(username string) (*User, error) {
	var user User

	err := g.Conn.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (g *PostgresRepository) Register(user *User) error {
	err := g.Conn.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
