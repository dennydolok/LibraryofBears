package repository

import (
	"BearLibrary/Models"
	"BearLibrary/internal/domain/repository"
	"gorm.io/gorm"
)

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) repository.AuthorRepository {
	return &authorRepository{db: db}
}

func (r *authorRepository) Create(author *Models.Author) error {
	return r.db.Create(author).Error
}

func (r *authorRepository) GetByID(id uint) (*Models.Author, error) {
	var author Models.Author
	err := r.db.First(&author, id).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *authorRepository) GetAll() ([]Models.Author, error) {
	var authors []Models.Author
	err := r.db.Find(&authors).Error
	return authors, err
}

func (r *authorRepository) Update(author *Models.Author) error {
	return r.db.Save(author).Error
}

func (r *authorRepository) Delete(id uint) error {
	return r.db.Delete(&Models.Author{}, id).Error
}