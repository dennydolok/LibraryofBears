package repository

import (
	"BearLibrary/Models"
)

type AuthorRepository interface {
	Create(author *Models.Author) error
	GetByID(id uint) (*Models.Author, error)
	GetAll() ([]Models.Author, error)
	Update(author *Models.Author) error
	Delete(id uint) error
}