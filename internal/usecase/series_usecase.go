package usecase

import (
	"BearLibrary/Models"
	"BearLibrary/internal/domain/repository"
)

type SeriesUsecase interface {
	Create(series *Models.Series) error
	GetByID(id uint) (*Models.Series, error)
	GetAll() ([]Models.Series, error)
	GetByAuthorID(authorID uint) ([]Models.Series, error)
	Update(series *Models.Series) error
	Delete(id uint) error
}

type seriesUsecase struct {
	seriesRepo repository.SeriesRepository
}

func NewSeriesUsecase(seriesRepo repository.SeriesRepository) SeriesUsecase {
	return &seriesUsecase{seriesRepo: seriesRepo}
}

func (u *seriesUsecase) Create(series *Models.Series) error {
	return u.seriesRepo.Create(series)
}

func (u *seriesUsecase) GetByID(id uint) (*Models.Series, error) {
	return u.seriesRepo.GetByID(id)
}

func (u *seriesUsecase) GetAll() ([]Models.Series, error) {
	return u.seriesRepo.GetAll()
}

func (u *seriesUsecase) GetByAuthorID(authorID uint) ([]Models.Series, error) {
	return u.seriesRepo.GetByAuthorID(authorID)
}

func (u *seriesUsecase) Update(series *Models.Series) error {
	return u.seriesRepo.Update(series)
}

func (u *seriesUsecase) Delete(id uint) error {
	return u.seriesRepo.Delete(id)
}