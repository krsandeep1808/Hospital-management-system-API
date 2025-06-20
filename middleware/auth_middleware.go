package middleware

import (
    "clinic-portal/utils"
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware(role string) gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        token = strings.TrimPrefix(token, "Bearer ")
        if utils.ValidateToken(token, role) {
            c.Next()
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
        }
    }
}