package book

import "gorm.io/gorm"

type Book struct {
	ID        string `json:"bid,omitempty" gorm:"column:first_name" binding:"required"`
	Title     string `json:"title,omitempty" gorm:"column:last_name" binding:"required"`
	Desc      string `json:"desc,omitempty" gorm:"column:last_name" binding:"required"`
	ISBN      string `json:"isbn,omitempty" gorm:"column:date_of_birth" binding:"required"`
	Ganre     string `json:"ganre,omitempty" gorm:"column:username" binding:"required"`
	Author    string `json:"author,omitempty" gorm:"column:email" binding:"required"`
	Published string `json:"published,omitempty" gorm:"column:password" binding:"required"`
}

// by default gorm plurallizes table name from struct model name
// to override it, create a method called tableName returns string.
func (Book) TableName() string {
	return "book"
}

type PostgresRepository struct {
	Conn *gorm.DB
}

func NewBookRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		Conn: db,
	}
}

func (g *PostgresRepository) GetByID(username string) (*Book, error) {
	var user Book

	err := g.Conn.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
