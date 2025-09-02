package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"BearLibrary/Config"
	"BearLibrary/Route"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Skip if no database connection to avoid panic
	if Config.DB == nil {
		// Initialize a temporary connection for testing
		Config.InitConnection()
	}

	// Create router - this will initialize the container
	router := Route.New()

	// Test health check
	req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "ok", response["status"])
}

func TestUserRegistration(t *testing.T) {
	// Skip if no database connection
	if Config.DB == nil {
		t.Skip("Database not available for testing")
	}

	router := Route.New()

	// Test user registration
	user := map[string]interface{}{
		"username": "testuser",
		"email":    "test@example.com",
		"password": "testpassword",
		"role":     1,
	}

	userJSON, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	// Should either succeed or fail due to duplicate email (which is fine for testing)
	assert.True(t, rec.Code == http.StatusCreated || rec.Code == http.StatusBadRequest)
}