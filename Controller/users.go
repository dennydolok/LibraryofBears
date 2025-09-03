package Controller

import (
	"BearLibrary/Config"
	"BearLibrary/Helper"
	"BearLibrary/Models"
	"github.com/labstack/echo/v4"
	"net/http"
)

var users Models.User

func Login(c echo.Context) error {
	loginInfo := Models.User{}
	err := c.Bind(&loginInfo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	err = Config.DB.Find(&users).Where("email = ?", loginInfo.Email).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "username not found",
		})
	}
	if users.Password != loginInfo.Password {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "password incorrect",
		})
	}
	token, err := Helper.CreateToken(users.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"message": "failed to generate token",
		})
	}
	users.Password = ""
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"user":    users,
		"token":   token,
	})
}
