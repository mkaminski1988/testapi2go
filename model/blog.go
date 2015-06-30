package model

type Blog struct {
	Id          string
	Description string
	CreatedAt   int64 `bson:"created_at"`
	UpdatedAt   int64 `bson:"updated_at"`
}

func (comment Blog) GetID() string {
	return comment.Id
}

func (comment *Blog) SetID(id string) error {
	comment.Id = id
	return nil
}
