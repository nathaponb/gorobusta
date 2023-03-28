package book

type BookRepository interface {
	GetByID(id string) (*Book, error)
}
