package routes

import (
	"STDE_proj/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	public := router.Group("/api/public")
	{
		user := public.Group("/user")
		{
			// Маршрут не реализован как регистрация. Реализован, как обычный Create
			user.POST("/register", controllers.PostAuthUserHandler)
		}
	}

}
