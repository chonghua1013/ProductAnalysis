package routes

import (
	"github.com/chonghua1013/ProductAnalysis/internal/config"
	"github.com/chonghua1013/ProductAnalysis/internal/controllers"
	"github.com/chonghua1013/ProductAnalysis/internal/middleware"
	"github.com/chonghua1013/ProductAnalysis/internal/repositories"
	"github.com/chonghua1013/ProductAnalysis/internal/services"
	"github.com/chonghua1013/ProductAnalysis/pkg/database"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	// 初始化依赖
	db, err := database.NewPostgresDB(cfg.DB)
	if err != nil {
		panic("failed to connect database")
	}

	// 初始化服务层和控制器
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

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
		protected.GET("/users", userController.GetUsers)
		protected.GET("/users/:id", userController.GetUser)
		protected.PUT("/users/:id", userController.UpdateUser)
		protected.DELETE("/users/:id", userController.DeleteUser)
	}

	return router
}
