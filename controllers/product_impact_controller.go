package controllers

import (
	"net/http"
	"portfolio-backend/database"
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetProductImpacts(c *gin.Context) {
	var impacts []models.ProductImpact

	// Cek apakah ada query ?project_id=xxx
	projectID := c.Query("project_id")
	query := database.DB.Preload("Project").Order("created_at desc")

	if projectID != "" {
		query = query.Where("project_id = ?", projectID)
	}

	// Eksekusi query (yang bisa pakai filter atau tidak)
	if err := query.Find(&impacts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product impacts"})
		return
	}

	c.JSON(http.StatusOK, impacts)

}

func CreateProductImpact(c *gin.Context) {
	var impact models.ProductImpact
	if err := c.ShouldBindJSON(&impact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&impact)
	c.JSON(http.StatusCreated, impact)
}

func UpdateProductImpact(c *gin.Context) {
	id := c.Param("id")
	var impact models.ProductImpact

	if err := database.DB.First(&impact, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product impact not found"})
		return
	}

	if err := c.ShouldBindJSON(&impact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&impact)
	c.JSON(http.StatusOK, impact)
}

func DeleteProductImpact(c *gin.Context) {
	id := c.Param("id")
	var impact models.ProductImpact

	if err := database.DB.First(&impact, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product impact not found"})
		return
	}

	database.DB.Delete(&impact)
	c.JSON(http.StatusOK, gin.H{"message": "Product impact deleted"})
}
