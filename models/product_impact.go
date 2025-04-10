package models

import "gorm.io/gorm"

type ProductImpact struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Metric      string  `json:"metric"`
	ImpactType  string  `json:"impact_type"`
	ProjectID   uint    `json:"project_id"`
	Project     Project `json:"project" gorm:"foreignKey:ProjectID"`
}
