package repository

import (
	"BearLibrary/Models"
)

type UserRepository interface {
	Create(user *Models.User) error
	GetByID(id uint) (*Models.User, error)
	GetByEmail(email string) (*Models.User, error)
	GetAll() ([]Models.User, error)
	Update(user *Models.User) error
	Delete(id uint) error
}