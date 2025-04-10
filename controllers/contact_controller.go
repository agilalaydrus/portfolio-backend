package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portfolio-backend/database"
	"portfolio-backend/models"
)

import "portfolio-backend/utils"

func SubmitContact(c *gin.Context) {
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&contact)

	go utils.SendResendContact(
		contact.Name,
		contact.Email,
		contact.Phone,
		contact.LinkedIn,
		contact.Message,
	)

	c.JSON(http.StatusOK, gin.H{"message": "Contact submitted successfully"})
}

func GetContacts(c *gin.Context) {
	var contacts []models.Contact

	if err := database.DB.Order("created_at desc").Find(&contacts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch contacts"})
		return
	}

	c.JSON(http.StatusOK, contacts)
}

func GetContactByID(c *gin.Context) {
	id := c.Param("id")

	var contact models.Contact
	if err := database.DB.First(&contact, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	c.JSON(http.StatusOK, contact)
}

func MarkContactAsRead(c *gin.Context) {
	id := c.Param("id")

	var contact models.Contact
	if err := database.DB.First(&contact, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	contact.IsRead = true
	database.DB.Save(&contact)

	c.JSON(http.StatusOK, gin.H{"message": "Contact marked as read"})
}
