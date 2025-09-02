package repository

import (
	"BearLibrary/Models"
	"BearLibrary/internal/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *Models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uint) (*Models.User, error) {
	var user Models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*Models.User, error) {
	var user Models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll() ([]Models.User, error) {
	var users []Models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) Update(user *Models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&Models.User{}, id).Error
}