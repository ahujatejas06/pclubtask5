package controllers

import (
	"net/http"

	"github.com/ahujatejas06/pclubtask5/models"
	"github.com/gin-gonic/gin"
)

type CreateAccountInput struct {
	Name       string `json:"name" binding:"required"`
	Rollnumber int    `json:"roll_number" binding:"required"`
	Branch     string `json:"branch" binding:"required"`
	UserID     string `json:"userid" binding:"required"`
}

type UpdateAccountInput struct {
	Name       string `json:"name"`
	Rollnumber int    `json:"roll_number"`
	Branch     string `json:"branch"`
	UserID     string `json:"userid"`
}

func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func FindUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
	var input CreateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Rollnumber: input.Rollnumber, Branch: input.Branch, UserID: input.UserID}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	update := models.User{Name: input.Name, Rollnumber: input.Rollnumber, Branch: input.Branch, UserID: input.UserID}
	models.DB.Model(&user).Updates(update)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
