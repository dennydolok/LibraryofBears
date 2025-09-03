package Route

import (
	config "BearLibrary/Config"
	"BearLibrary/internal/container"
	"github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Initialize dependency injection container
	container, err := container.NewContainer(config.DB)
	if err != nil {
		panic("Failed to initialize container: " + err.Error())
	}

	// JWT Middleware
	jwtMiddleware := echojwt.JWT([]byte(config.Secret))

	// Public routes
	// Auth routes
	e.POST("/auth/register", container.UserHandler.Register)
	e.POST("/auth/login", container.UserHandler.Login)

	// Public read-only routes
	e.GET("/authors", container.AuthorHandler.GetAll)
	e.GET("/authors/:id", container.AuthorHandler.GetByID)
	e.GET("/series", container.SeriesHandler.GetAll)
	e.GET("/series/:id", container.SeriesHandler.GetByID)
	e.GET("/series/author/:authorId", container.SeriesHandler.GetByAuthorID)
	e.GET("/arts", container.ArtsHandler.GetAll)
	e.GET("/arts/:id", container.ArtsHandler.GetByID)
	e.GET("/arts/author/:authorId", container.ArtsHandler.GetByAuthorID)
	e.GET("/arts/series/:seriesId", container.ArtsHandler.GetBySeriesID)

	// Protected routes - require JWT authentication
	protected := e.Group("")
	protected.Use(jwtMiddleware)

	// User management routes (protected)
	protected.GET("/users", container.UserHandler.GetAll)
	protected.GET("/users/:id", container.UserHandler.GetByID)
	protected.PUT("/users/:id", container.UserHandler.Update)
	protected.DELETE("/users/:id", container.UserHandler.Delete)

	// Author management routes (protected)
	protected.POST("/authors", container.AuthorHandler.Create)
	protected.PUT("/authors/:id", container.AuthorHandler.Update)
	protected.DELETE("/authors/:id", container.AuthorHandler.Delete)

	// Series management routes (protected)
	protected.POST("/series", container.SeriesHandler.Create)
	protected.PUT("/series/:id", container.SeriesHandler.Update)
	protected.DELETE("/series/:id", container.SeriesHandler.Delete)

	// Arts management routes (protected)
	protected.POST("/arts", container.ArtsHandler.Create)
	protected.POST("/arts/upload", container.ArtsHandler.CreateWithFile)
	protected.PUT("/arts/:id", container.ArtsHandler.Update)
	protected.PUT("/arts/:id/upload", container.ArtsHandler.UpdateWithFile)
	protected.DELETE("/arts/:id", container.ArtsHandler.Delete)

	// Health check
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	})

	return e
}
