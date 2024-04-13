package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestHello tests the hello function.
func TestHello(t *testing.T) {
	// Create a new Echo instance for testing
	e := echo.New()

	// Create a new HTTP request to pass to the handler.
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a new HTTP recorder to record the response.
	rec := httptest.NewRecorder()

	// Create a new Echo context.
	c := e.NewContext(req, rec)

	// Invoke the hello handler.
	if assert.NoError(t, hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello, World!", rec.Body.String())
	}
}

// TestIsValid tests the isValid function.
func TestIsValid(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"Valid Input", "hello", false},
		{"Empty Input", "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := isValid(test.input)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
