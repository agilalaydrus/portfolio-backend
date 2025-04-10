package routes

import (
	"github.com/gin-gonic/gin"
	"portfolio-backend/controllers"
	"portfolio-backend/middlewares"
)

func SetupRouter(r *gin.Engine) {
	// Public routes (no auth)
	public := r.Group("/api")
	{
		public.GET("/projects", controllers.GetProjects)
		public.GET("/projects/:id", controllers.GetProjectByID)
		public.POST("/contact", controllers.SubmitContact)
		public.GET("/case-studies", controllers.GetCaseStudies)
		public.GET("/product-impacts", controllers.GetProductImpacts)
	}

	// Admin routes (with JWT)
	admin := r.Group("/admin/api")
	{
		// Auth endpoint (open)
		admin.POST("/login", controllers.Login)

		// Protected endpoints (JWT required)
		admin.Use(middlewares.AuthMiddleware())

		admin.GET("/me", controllers.Me)

		// Contact management
		admin.GET("/contacts", controllers.GetContacts)
		admin.GET("/contacts/:id", controllers.GetContactByID)
		admin.PATCH("/contacts/:id/read", controllers.MarkContactAsRead)
		admin.PATCH("/contacts/:id/respond", controllers.MarkContactAsResponded)

		// Projects
		admin.POST("/projects", controllers.CreateProject)
		admin.PUT("/projects/:id", controllers.UpdateProject)
		admin.DELETE("/projects/:id", controllers.DeleteProject)

		// Case studies
		admin.POST("/case-studies", controllers.CreateCaseStudy)
		admin.PUT("/case-studies/:id", controllers.UpdateCaseStudy)
		admin.DELETE("/case-studies/:id", controllers.DeleteCaseStudy)

		// Product impacts
		admin.POST("/product-impacts", controllers.CreateProductImpact)
		admin.PUT("/product-impacts/:id", controllers.UpdateProductImpact)
		admin.DELETE("/product-impacts/:id", controllers.DeleteProductImpact)
	}
}
