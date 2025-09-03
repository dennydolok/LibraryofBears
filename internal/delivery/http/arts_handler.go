package http

import (
	"BearLibrary/Models"
	"BearLibrary/internal/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArtsHandler struct {
	artsUsecase usecase.ArtsUsecase
}

func NewArtsHandler(artsUsecase usecase.ArtsUsecase) *ArtsHandler {
	return &ArtsHandler{artsUsecase: artsUsecase}
}

func (h *ArtsHandler) Create(c echo.Context) error {
	var arts Models.Arts
	if err := c.Bind(&arts); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body",
		})
	}

	if err := h.artsUsecase.Create(&arts); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Arts created successfully",
		"arts":    arts,
	})
}

func (h *ArtsHandler) CreateWithFile(c echo.Context) error {
	var arts Models.Arts

	// Parse form data
	arts.Title = c.FormValue("title")
	arts.Description = c.FormValue("description")
	arts.Source = c.FormValue("source")
	arts.Website = c.FormValue("website")
	
	if authorIDStr := c.FormValue("author_id"); authorIDStr != "" {
		authorID, err := strconv.ParseUint(authorIDStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid author ID",
			})
		}
		arts.AuthorID = uint(authorID)
	}

	if seriesIDStr := c.FormValue("series_id"); seriesIDStr != "" {
		seriesID, err := strconv.ParseUint(seriesIDStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid series ID",
			})
		}
		arts.SeriesID = uint(seriesID)
	}

	// Handle file upload
	file, fileHeader, err := c.Request().FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "No file provided",
		})
	}
	defer file.Close()

	if err := h.artsUsecase.CreateWithFile(&arts, file, fileHeader); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Arts created successfully with file",
		"arts":    arts,
	})
}

func (h *ArtsHandler) GetByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid arts ID",
		})
	}

	arts, err := h.artsUsecase.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Arts not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"arts": arts,
	})
}

func (h *ArtsHandler) GetAll(c echo.Context) error {
	arts, err := h.artsUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"arts": arts,
	})
}

func (h *ArtsHandler) GetByAuthorID(c echo.Context) error {
	idParam := c.Param("authorId")
	authorID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid author ID",
		})
	}

	arts, err := h.artsUsecase.GetByAuthorID(uint(authorID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"arts": arts,
	})
}

func (h *ArtsHandler) GetBySeriesID(c echo.Context) error {
	idParam := c.Param("seriesId")
	seriesID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid series ID",
		})
	}

	arts, err := h.artsUsecase.GetBySeriesID(uint(seriesID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"arts": arts,
	})
}

func (h *ArtsHandler) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid arts ID",
		})
	}

	var arts Models.Arts
	if err := c.Bind(&arts); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body",
		})
	}

	arts.ID = uint(id)
	if err := h.artsUsecase.Update(&arts); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Arts updated successfully",
		"arts":    arts,
	})
}

func (h *ArtsHandler) UpdateWithFile(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid arts ID",
		})
	}

	var arts Models.Arts
	arts.ID = uint(id)

	// Parse form data
	arts.Title = c.FormValue("title")
	arts.Description = c.FormValue("description")
	arts.Source = c.FormValue("source")
	arts.Website = c.FormValue("website")
	
	if authorIDStr := c.FormValue("author_id"); authorIDStr != "" {
		authorID, err := strconv.ParseUint(authorIDStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid author ID",
			})
		}
		arts.AuthorID = uint(authorID)
	}

	if seriesIDStr := c.FormValue("series_id"); seriesIDStr != "" {
		seriesID, err := strconv.ParseUint(seriesIDStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid series ID",
			})
		}
		arts.SeriesID = uint(seriesID)
	}

	// Handle file upload
	file, fileHeader, err := c.Request().FormFile("file")
	if err != nil {
		// No file provided, just update without file
		if err := h.artsUsecase.Update(&arts); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
	} else {
		defer file.Close()
		if err := h.artsUsecase.UpdateWithFile(&arts, file, fileHeader); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Arts updated successfully",
		"arts":    arts,
	})
}

func (h *ArtsHandler) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid arts ID",
		})
	}

	if err := h.artsUsecase.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Arts deleted successfully",
	})
}