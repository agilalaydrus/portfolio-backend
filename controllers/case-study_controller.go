package controllers

import (
	"net/http"
	"portfolio-backend/database"
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetCaseStudies(c *gin.Context) {
	var studies []models.CaseStudy
	if err := database.DB.Order("created_at desc").Find(&studies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch case studies"})
		return
	}
	c.JSON(http.StatusOK, studies)
}

func CreateCaseStudy(c *gin.Context) {
	var study models.CaseStudy
	if err := c.ShouldBindJSON(&study); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&study).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create case study"})
		return
	}

	c.JSON(http.StatusCreated, study)
}

func UpdateCaseStudy(c *gin.Context) {
	id := c.Param("id")
	var study models.CaseStudy

	if err := database.DB.First(&study, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Case study not found"})
		return
	}

	if err := c.ShouldBindJSON(&study); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&study)
	c.JSON(http.StatusOK, study)
}

func DeleteCaseStudy(c *gin.Context) {
	id := c.Param("id")
	var study models.CaseStudy

	if err := database.DB.First(&study, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Case study not found"})
		return
	}

	database.DB.Delete(&study)
	c.JSON(http.StatusOK, gin.H{"message": "Case study deleted"})
}
