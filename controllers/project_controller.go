package controllers

import (
	"net/http"
	"portfolio-backend/database"
	"portfolio-backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	var projects []models.Project
	err := database.DB.
		Preload("CaseStudies").
		Preload("ProductImpacts").
		Find(&projects).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

func GetProjectByID(c *gin.Context) {
	id := c.Param("id")
	var project models.Project

	if err := database.DB.Preload("CaseStudies").Preload("ProductImpacts").First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

func CreateProject(c *gin.Context) {
	var input models.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat proyek"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Project dibuat", "data": input})
}

func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	var project models.Project

	if err := database.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project tidak ditemukan"})
		return
	}

	var input models.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&project).Updates(input)
	c.JSON(http.StatusOK, gin.H{"message": "Project diupdate", "data": project})
}

func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	var project models.Project

	if err := database.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project tidak ditemukan"})
		return
	}

	database.DB.Delete(&project)
	c.JSON(http.StatusOK, gin.H{"message": "Project dihapus"})
}

func MarkContactAsResponded(c *gin.Context) {
	id := c.Param("id")

	var contact models.Contact
	if err := database.DB.First(&contact, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	now := time.Now()
	contact.RespondedAt = &now
	database.DB.Save(&contact)

	c.JSON(http.StatusOK, gin.H{"message": "Contact marked as responded"})
}
