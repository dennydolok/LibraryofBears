package repository

import (
	"BearLibrary/Models"
	"BearLibrary/internal/domain/repository"
	"gorm.io/gorm"
)

type seriesRepository struct {
	db *gorm.DB
}

func NewSeriesRepository(db *gorm.DB) repository.SeriesRepository {
	return &seriesRepository{db: db}
}

func (r *seriesRepository) Create(series *Models.Series) error {
	return r.db.Create(series).Error
}

func (r *seriesRepository) GetByID(id uint) (*Models.Series, error) {
	var series Models.Series
	err := r.db.Preload("Author").First(&series, id).Error
	if err != nil {
		return nil, err
	}
	return &series, nil
}

func (r *seriesRepository) GetAll() ([]Models.Series, error) {
	var series []Models.Series
	err := r.db.Preload("Author").Find(&series).Error
	return series, err
}

func (r *seriesRepository) GetByAuthorID(authorID uint) ([]Models.Series, error) {
	var series []Models.Series
	err := r.db.Preload("Author").Where("author_id = ?", authorID).Find(&series).Error
	return series, err
}

func (r *seriesRepository) Update(series *Models.Series) error {
	return r.db.Save(series).Error
}

func (r *seriesRepository) Delete(id uint) error {
	return r.db.Delete(&Models.Series{}, id).Error
}