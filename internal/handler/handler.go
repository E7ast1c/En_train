package handler

import (
	"en_train/internal/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repos *repository.Repository
}

func NewHandler(repos *repository.Repository) *Handler {
	return &Handler{repos: repos}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// TODO auth
	// auth := router.Group("/auth")
	// {
	// 	auth.POST("/sign-up")
	// 	auth.POST("/sign-in")
	// }

	api := router.Group("irregular-verbs")
	{
		api.GET("/list", h.GetIrregularVerbsList)
		api.POST("/get-by-id", h.GetIrregularVerbById)
		api.GET("/get-random", h.GetRandomVerb)
	}

	return router
}
