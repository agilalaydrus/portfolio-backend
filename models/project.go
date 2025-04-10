package models

type Project struct {
	ID             uint            `json:"ID"`
	Title          string          `json:"title"`
	Description    string          `json:"description"`
	TechStack      string          `json:"tech_stack"`
	RepositoryURL  string          `json:"repo_url"`
	DemoURL        string          `json:"demo_url"`
	Tags           string          `json:"tags"`
	CaseStudies    []CaseStudy     `json:"case_studies"`    // relasi
	ProductImpacts []ProductImpact `json:"product_impacts"` // relasi
}
