package container

import (
	config "BearLibrary/Config"
	"BearLibrary/internal/delivery/http"
	"BearLibrary/internal/infrastructure/repository"
	"BearLibrary/internal/infrastructure/service"
	"BearLibrary/internal/usecase"
	"gorm.io/gorm"
)

type Container struct {
	UserHandler   *http.UserHandler
	AuthorHandler *http.AuthorHandler
	SeriesHandler *http.SeriesHandler
	ArtsHandler   *http.ArtsHandler
}

func NewContainer(db *gorm.DB) (*Container, error) {
	// Initialize File Service
	endpoint, accessKey, secretKey, useSSL := config.GetMinIOConfig()
	fileService, err := service.NewFileService(endpoint, accessKey, secretKey, useSSL)
	if err != nil {
		return nil, err
	}

	// Initialize Repositories
	userRepo := repository.NewUserRepository(db)
	authorRepo := repository.NewAuthorRepository(db)
	seriesRepo := repository.NewSeriesRepository(db)
	artsRepo := repository.NewArtsRepository(db)

	// Initialize Usecases
	userUsecase := usecase.NewUserUsecase(userRepo)
	authorUsecase := usecase.NewAuthorUsecase(authorRepo)
	seriesUsecase := usecase.NewSeriesUsecase(seriesRepo)
	artsUsecase := usecase.NewArtsUsecase(artsRepo, fileService)

	// Initialize Handlers
	userHandler := http.NewUserHandler(userUsecase)
	authorHandler := http.NewAuthorHandler(authorUsecase)
	seriesHandler := http.NewSeriesHandler(seriesUsecase)
	artsHandler := http.NewArtsHandler(artsUsecase)

	return &Container{
		UserHandler:   userHandler,
		AuthorHandler: authorHandler,
		SeriesHandler: seriesHandler,
		ArtsHandler:   artsHandler,
	}, nil
}