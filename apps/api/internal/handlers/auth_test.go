// apps/api/handlers/auth_test.go
package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/auth/login", Login)
	return r
}

func TestLoginOK(t *testing.T) {
	router := setupRouter()

	body, _ := json.Marshal(LoginReq{
		Email:    "alice@example.com",
		Password: "secret",
	})
	req := httptest.NewRequest("POST", "/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req) // ★ Gin は ServeHTTP で実行 :contentReference[oaicite:1]{index=1}

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), `"token":"dummy"`)
}

func TestLoginBadPayload(t *testing.T) {
	router := setupRouter()
	req := httptest.NewRequest("POST", "/auth/login", nil) // JSONなし

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
