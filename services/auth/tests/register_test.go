package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"plate-connect/services/auth/handlers"
)

func TestRegisterHandler_Success(t *testing.T) {
	body := []byte(`{"kennzeichen":"B-AB123", "email":"test@example.com"}`)
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handlers.RegisterHandler(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status 201 Created, got %d", rr.Code)
	}
}

func TestRegisterHandler_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer([]byte("{invalid json")))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handlers.RegisterHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %d", rr.Code)
	}
}
