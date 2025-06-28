// apps/api/handlers/auth.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad payload"})
		return
	}
	// TODO: DB 照合 & JWT 発行
	c.JSON(http.StatusOK, gin.H{"token": "dummy"})
}
