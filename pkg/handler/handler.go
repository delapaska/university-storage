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
	router.LoadHTMLGlob("templates/*")

	auth := router.Group("/auth")
	{
		auth.GET("/sign-up", h.signUp)
		auth.POST("/sign-up", h.signUp)
		auth.GET("/sign-in", h.signIn)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		api.GET("/connect", h.connectToProject)
		api.POST("/connect", h.connectToProject)
		api.GET("/main", h.loadMainPage)
		projects := api.Group("/projects")
		{
			projects.GET("/create", h.createProject)
			projects.POST("/create", h.createProject)

			projects.GET("/list", h.getAllLists)
			projects.GET("/list/:id", h.getProjectById)
			projects.POST("/list/:id/upload", h.uploadFiles)
			projects.POST("list/:id/create-folder", h.createFolder)
			projects.GET("/list/:id/connect", h.connectProject)
		}
	}

	return router
}
