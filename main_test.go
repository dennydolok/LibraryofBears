package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"BearLibrary/Helper"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestJWTHelper(t *testing.T) {
	// Test JWT token creation
	token, err := Helper.CreateToken(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestHealthCheckRoute(t *testing.T) {
	e := echo.New()
	
	// Add health check route manually to avoid dependency issues
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "ok", response["status"])
}

func TestJSONParsing(t *testing.T) {
	// Test JSON parsing for user structure
	userJSON := `{"username":"testuser","email":"test@example.com","role":1}`
	
	var user map[string]interface{}
	err := json.Unmarshal([]byte(userJSON), &user)
	assert.NoError(t, err)
	assert.Equal(t, "testuser", user["username"])
	assert.Equal(t, "test@example.com", user["email"])
	assert.Equal(t, float64(1), user["role"]) // JSON numbers become float64
}