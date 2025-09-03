package usecase

import (
	"BearLibrary/Models"
	"BearLibrary/internal/domain/repository"
)

type AuthorUsecase interface {
	Create(author *Models.Author) error
	GetByID(id uint) (*Models.Author, error)
	GetAll() ([]Models.Author, error)
	Update(author *Models.Author) error
	Delete(id uint) error
}

type authorUsecase struct {
	authorRepo repository.AuthorRepository
}

func NewAuthorUsecase(authorRepo repository.AuthorRepository) AuthorUsecase {
	return &authorUsecase{authorRepo: authorRepo}
}

func (u *authorUsecase) Create(author *Models.Author) error {
	return u.authorRepo.Create(author)
}

func (u *authorUsecase) GetByID(id uint) (*Models.Author, error) {
	return u.authorRepo.GetByID(id)
}

func (u *authorUsecase) GetAll() ([]Models.Author, error) {
	return u.authorRepo.GetAll()
}

func (u *authorUsecase) Update(author *Models.Author) error {
	return u.authorRepo.Update(author)
}

func (u *authorUsecase) Delete(id uint) error {
	return u.authorRepo.Delete(id)
}