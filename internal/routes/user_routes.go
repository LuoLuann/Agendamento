package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/users")
	{
		// public := userGroup.Group("")

		{
			// public.GET("", listUsers)
			// public.GET("/:id", getUser)
		}

		// protected := userGroup.Group("")
		// // protected.Use(middleware.AuthRequired()) // Descomente quando tiver o middleware
		// {
		// 	protected.POST("", createUser)
		// 	protected.PUT("/:id", updateUser)
		// 	protected.DELETE("/:id", deleteUser)
		// }
	}
	userGroup.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{"status": "ok"})
	})
}
