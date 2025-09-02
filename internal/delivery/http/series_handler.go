package http

import (
	"BearLibrary/Models"
	"BearLibrary/internal/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SeriesHandler struct {
	seriesUsecase usecase.SeriesUsecase
}

func NewSeriesHandler(seriesUsecase usecase.SeriesUsecase) *SeriesHandler {
	return &SeriesHandler{seriesUsecase: seriesUsecase}
}

func (h *SeriesHandler) Create(c echo.Context) error {
	var series Models.Series
	if err := c.Bind(&series); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body",
		})
	}

	if err := h.seriesUsecase.Create(&series); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Series created successfully",
		"series":  series,
	})
}

func (h *SeriesHandler) GetByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid series ID",
		})
	}

	series, err := h.seriesUsecase.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Series not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"series": series,
	})
}

func (h *SeriesHandler) GetAll(c echo.Context) error {
	series, err := h.seriesUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"series": series,
	})
}

func (h *SeriesHandler) GetByAuthorID(c echo.Context) error {
	idParam := c.Param("authorId")
	authorID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid author ID",
		})
	}

	series, err := h.seriesUsecase.GetByAuthorID(uint(authorID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"series": series,
	})
}

func (h *SeriesHandler) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid series ID",
		})
	}

	var series Models.Series
	if err := c.Bind(&series); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body",
		})
	}

	series.ID = uint(id)
	if err := h.seriesUsecase.Update(&series); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Series updated successfully",
		"series":  series,
	})
}

func (h *SeriesHandler) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid series ID",
		})
	}

	if err := h.seriesUsecase.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Series deleted successfully",
	})
}