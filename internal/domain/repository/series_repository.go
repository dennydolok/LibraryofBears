package repository

import (
	"BearLibrary/Models"
)

type SeriesRepository interface {
	Create(series *Models.Series) error
	GetByID(id uint) (*Models.Series, error)
	GetAll() ([]Models.Series, error)
	GetByAuthorID(authorID uint) ([]Models.Series, error)
	Update(series *Models.Series) error
	Delete(id uint) error
}