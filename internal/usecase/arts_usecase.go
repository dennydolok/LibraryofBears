package usecase

import (
	"BearLibrary/Models"
	"BearLibrary/internal/domain/repository"
	"BearLibrary/internal/infrastructure/service"
	"mime/multipart"
)

type ArtsUsecase interface {
	Create(arts *Models.Arts) error
	CreateWithFile(arts *Models.Arts, file multipart.File, fileHeader *multipart.FileHeader) error
	GetByID(id uint) (*Models.Arts, error)
	GetAll() ([]Models.Arts, error)
	GetByAuthorID(authorID uint) ([]Models.Arts, error)
	GetBySeriesID(seriesID uint) ([]Models.Arts, error)
	Update(arts *Models.Arts) error
	UpdateWithFile(arts *Models.Arts, file multipart.File, fileHeader *multipart.FileHeader) error
	Delete(id uint) error
}

type artsUsecase struct {
	artsRepo    repository.ArtsRepository
	fileService service.FileService
}

func NewArtsUsecase(artsRepo repository.ArtsRepository, fileService service.FileService) ArtsUsecase {
	return &artsUsecase{
		artsRepo:    artsRepo,
		fileService: fileService,
	}
}

func (u *artsUsecase) Create(arts *Models.Arts) error {
	return u.artsRepo.Create(arts)
}

func (u *artsUsecase) CreateWithFile(arts *Models.Arts, file multipart.File, fileHeader *multipart.FileHeader) error {
	// Upload file to MinIO
	fileURL, err := u.fileService.UploadFile(file, fileHeader, "arts")
	if err != nil {
		return err
	}
	
	arts.File = fileURL
	return u.artsRepo.Create(arts)
}

func (u *artsUsecase) GetByID(id uint) (*Models.Arts, error) {
	return u.artsRepo.GetByID(id)
}

func (u *artsUsecase) GetAll() ([]Models.Arts, error) {
	return u.artsRepo.GetAll()
}

func (u *artsUsecase) GetByAuthorID(authorID uint) ([]Models.Arts, error) {
	return u.artsRepo.GetByAuthorID(authorID)
}

func (u *artsUsecase) GetBySeriesID(seriesID uint) ([]Models.Arts, error) {
	return u.artsRepo.GetBySeriesID(seriesID)
}

func (u *artsUsecase) Update(arts *Models.Arts) error {
	return u.artsRepo.Update(arts)
}

func (u *artsUsecase) UpdateWithFile(arts *Models.Arts, file multipart.File, fileHeader *multipart.FileHeader) error {
	// Upload new file to MinIO
	fileURL, err := u.fileService.UploadFile(file, fileHeader, "arts")
	if err != nil {
		return err
	}
	
	arts.File = fileURL
	return u.artsRepo.Update(arts)
}

func (u *artsUsecase) Delete(id uint) error {
	return u.artsRepo.Delete(id)
}