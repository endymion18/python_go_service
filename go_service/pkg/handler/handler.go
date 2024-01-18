package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	greeting := router.Group("/greet")
	{
		greeting.GET("", h.greet)
		greeting.GET("/history", h.greetHistory)
	}

	python := router.Group("/python")
	{
		python.GET("/greet", h.pythonGreet)
		python.GET("/greet/history", h.pythonGreetHistory)
	}

	return router
}
