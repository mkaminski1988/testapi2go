package model

type Book struct {
	Id        string
	Name      string
	CreatedAt int64 `bson:"created_at"`
	UpdatedAt int64 `bson:"updated_at"`
}

func (book Book) GetID() string {
	return book.Id
}

func (book *Book) SetID(id string) error {
	book.Id = id
	return nil
}
