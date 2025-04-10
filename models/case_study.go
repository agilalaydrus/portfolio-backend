package models

import "gorm.io/gorm"

type CaseStudy struct {
	gorm.Model
	Title           string `json:"title"`
	Description     string `json:"description"`
	Role            string `json:"role"`
	TeamSize        int    `json:"team_size"`
	TechStack       string `json:"tech_stack"`
	Duration        string `json:"duration"`
	Impact          string `json:"impact"`
	DecisionProcess string `json:"decision_process"`
	Link            string `json:"link"`

	ProjectID uint    `json:"project_id"` // <-- Tambahkan ini
	Project   Project `json:"project" gorm:"foreignKey:ProjectID"`
}
