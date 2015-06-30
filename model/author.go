package model

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
)

type Author struct {
	Id        string
	Name      string
	CreatedAt int64 `bson:"created_at"`
	UpdatedAt int64 `bson:"updated_at"`

	Book    []Book   `json:"-" bson:"-"`
	BookIDs []string `json:"-" bson:"books"`

	Blogs   []Blog   `json:"-" bson:"-"`
	BlogIDs []string `json:"-" bson:"blogs"`
}

func (author Author) GetID() string {
	return author.Id
}

func (author *Author) SetID(id string) error {
	author.Id = id
	return nil
}

func (author Author) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "books",
			Name: "books",
		},
		{
			Type: "blogs",
			Name: "blogs",
		},
	}
}

func (author Author) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, appid := range author.BookIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   appid,
			Type: "books",
			Name: "books",
		})
	}

	for _, id := range author.BlogIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   id,
			Type: "blogs",
			Name: "blogs",
		})
	}
	return result
}

func (author Author) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range author.Book {
		result = append(result, author.Book[key])
	}
	for key := range author.Blogs {
		result = append(result, author.Blogs[key])
	}

	return result
}

func (author *Author) SetReferencedIDs(references []jsonapi.ReferenceID) error {
	if len(references) == 0 {
		author.BlogIDs = author.BlogIDs[:0]
		author.BookIDs = author.BookIDs[:0]
		return nil
	}
	for _, reference := range references {
		if reference.Name == "books" {
			author.BookIDs = append(author.BookIDs, reference.ID)
		}
		if reference.Name == "blogs" {
			author.BlogIDs = append(author.BlogIDs, reference.ID)
		}
	}

	return nil
}

func (author *Author) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "books" {
		author.BookIDs = IDs
		return nil
	}
	if name == "blogs" {
		author.BlogIDs = IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (author *Author) AddToManyIDs(name string, IDs []string) error {
	if name == "books" {
		author.BookIDs = append(author.BookIDs, IDs...)
		return nil
	}
	if name == "blogs" {
		author.BlogIDs = append(author.BlogIDs, IDs...)
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

//Comment this function for the code to work.
func (author *Author) DeleteToManyIDs(name string, IDs []string) error {
	if name == "books" {
		for _, ID := range IDs {
			for pos, oldID := range author.BookIDs {
				if ID == oldID {
					author.BookIDs = append(author.BookIDs[:pos], author.BookIDs[pos+1:]...)
				}
			}
		}
		return nil
	}
	if name == "blogs" {
		for _, ID := range IDs {
			for pos, oldID := range author.BlogIDs {
				if ID == oldID {
					author.BlogIDs = append(author.BlogIDs[:pos], author.BlogIDs[pos+1:]...)
				}
			}
		}
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}
