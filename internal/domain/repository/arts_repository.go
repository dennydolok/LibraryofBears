package repository

import (
	"BearLibrary/Models"
)

type ArtsRepository interface {
	Create(arts *Models.Arts) error
	GetByID(id uint) (*Models.Arts, error)
	GetAll() ([]Models.Arts, error)
	GetByAuthorID(authorID uint) ([]Models.Arts, error)
	GetBySeriesID(seriesID uint) ([]Models.Arts, error)
	Update(arts *Models.Arts) error
	Delete(id uint) error
}