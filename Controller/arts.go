package Controller

import (
	config "BearLibrary/Config"
	model "BearLibrary/Models"
	"github.com/labstack/echo/v4"
	"net/http"
)

var art model.Arts

func AddArts(c echo.Context) error {
	c.Bind(&art)
	err := config.DB.Save(&art).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    art,
	})
}
