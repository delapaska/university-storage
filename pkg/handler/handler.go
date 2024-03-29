package handler

import (
	"github.com/delapaska/university-storage/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{

		projects := api.Group("/projects")
		{

			projects.POST("/", h.createProject)
			projects.GET("/", h.getAllLists)
			projects.GET("/:id", h.getProjectById)
		}
	}
	return router
}
