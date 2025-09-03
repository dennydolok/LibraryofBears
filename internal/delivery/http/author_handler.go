package http

import (
	"BearLibrary/Models"
	"BearLibrary/internal/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AuthorHandler struct {
	authorUsecase usecase.AuthorUsecase
}

func NewAuthorHandler(authorUsecase usecase.AuthorUsecase) *AuthorHandler {
	return &AuthorHandler{authorUsecase: authorUsecase}
}

func (h *AuthorHandler) Create(c echo.Context) error {
	var author Models.Author
	if err := c.Bind(&author); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body",
		})
	}

	if err := h.authorUsecase.Create(&author); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Author created successfully",
		"author":  author,
	})
}

func (h *AuthorHandler) GetByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid author ID",
		})
	}

	author, err := h.authorUsecase.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Author not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"author": author,
	})
}

func (h *AuthorHandler) GetAll(c echo.Context) error {
	authors, err := h.authorUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"authors": authors,
	})
}

func (h *AuthorHandler) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid author ID",
		})
	}

	var author Models.Author
	if err := c.Bind(&author); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body",
		})
	}

	author.ID = uint(id)
	if err := h.authorUsecase.Update(&author); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Author updated successfully",
		"author":  author,
	})
}

func (h *AuthorHandler) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid author ID",
		})
	}

	if err := h.authorUsecase.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Author deleted successfully",
	})
}