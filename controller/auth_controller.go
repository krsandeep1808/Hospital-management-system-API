package controller

import (
    "clinic-portal/model"
    "clinic-portal/utils"
    "net/http"
    "github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
    var user model.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if user.Email == "doctor@demo.com" && user.Password == "password" {
        token := utils.GenerateToken("doctor")
        c.JSON(http.StatusOK, gin.H{"token": token})
    } else if user.Email == "reception@demo.com" && user.Password == "password" {
        token := utils.GenerateToken("receptionist")
        c.JSON(http.StatusOK, gin.H{"token": token})
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    }
}