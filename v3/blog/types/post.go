package types

type Post struct {
	ID    string
	Title string
	Body  string
}

type PostStorage interface {
	Create(*Post) error
	GetAll() ([]*Post, error)
	GetByID(string) (*Post, error)
}
