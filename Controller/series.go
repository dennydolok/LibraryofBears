package Controller

import (
	"BearLibrary/Config"
	model "BearLibrary/Models"

	"github.com/labstack/echo/v4"
	"net/http"
)

var series []model.Series

func GetSeries(c echo.Context) error {
	err := Config.DB.Find(&series).Preload("author")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"series":  series,
	})
}

func GetSeriesByAuthor(c echo.Context) error {
	err := Config.DB.Find(&series, "AuthorID = ?", c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"series":  series,
	})
}

func AddSeries(c echo.Context) error {
	series := model.Series{}
	err := c.Bind(&series)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	err = Config.DB.Create(&series).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"series":  series,
	})
}
