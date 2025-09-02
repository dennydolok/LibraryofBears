package repository

import (
	"BearLibrary/Models"
	"BearLibrary/internal/domain/repository"
	"gorm.io/gorm"
)

type artsRepository struct {
	db *gorm.DB
}

func NewArtsRepository(db *gorm.DB) repository.ArtsRepository {
	return &artsRepository{db: db}
}

func (r *artsRepository) Create(arts *Models.Arts) error {
	return r.db.Create(arts).Error
}

func (r *artsRepository) GetByID(id uint) (*Models.Arts, error) {
	var arts Models.Arts
	err := r.db.Preload("Author").Preload("Series").First(&arts, id).Error
	if err != nil {
		return nil, err
	}
	return &arts, nil
}

func (r *artsRepository) GetAll() ([]Models.Arts, error) {
	var arts []Models.Arts
	err := r.db.Preload("Author").Preload("Series").Find(&arts).Error
	return arts, err
}

func (r *artsRepository) GetByAuthorID(authorID uint) ([]Models.Arts, error) {
	var arts []Models.Arts
	err := r.db.Preload("Author").Preload("Series").Where("author_id = ?", authorID).Find(&arts).Error
	return arts, err
}

func (r *artsRepository) GetBySeriesID(seriesID uint) ([]Models.Arts, error) {
	var arts []Models.Arts
	err := r.db.Preload("Author").Preload("Series").Where("series_id = ?", seriesID).Find(&arts).Error
	return arts, err
}

func (r *artsRepository) Update(arts *Models.Arts) error {
	return r.db.Save(arts).Error
}

func (r *artsRepository) Delete(id uint) error {
	return r.db.Delete(&Models.Arts{}, id).Error
}