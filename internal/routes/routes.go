package routes

import (
	"project-name/internal/config"
	"project-name/internal/controllers"
	"project-name/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	// 全局中间件
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())

	// 公共路由
	public := router.Group("/api")
	{
		public.POST("/login", controllers.Login)
		public.POST("/register", controllers.Register)
	}

	// 需要认证的路由
	protected := router.Group("/api")
	protected.Use(middleware.JWTAuth(cfg.JWT.Secret))
	{
		protected.GET("/users", controllers.GetUsers)
		protected.GET("/users/:id", controllers.GetUser)
		protected.PUT("/users/:id", controllers.UpdateUser)
		protected.DELETE("/users/:id", controllers.DeleteUser)
	}

	return router
}
