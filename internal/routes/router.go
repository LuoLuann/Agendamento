package routes

import "github.com/gin-gonic/gin"

func Initialize() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hehehe funcionou",
		})
	})

	initializeRoutes(router)

	return router
}

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// Exemplo de registro de rotas:
		registerUserRoutes(v1)
		// registerServiceRoutes(v1, cfg)
	}
}
