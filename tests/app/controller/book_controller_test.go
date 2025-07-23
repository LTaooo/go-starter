package controller_test

import (
	"encoding/json"
	"fmt"
	"go-starter/core/response"
	"go-starter/core/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBook_Success(t *testing.T) {
	router := server.Init()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/book?id=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	response := response.Response{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	fmt.Println(response)
	assert.Equal(t, response.Code, 200)
}
