package main

import (
	"github.com/chonghua1013/ProductAnalysis/internal/config"
	"github.com/chonghua1013/ProductAnalysis/internal/routes"
	"log"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化路由
	router := routes.SetupRouter(cfg)

	// 启动服务
	log.Printf("Starting server on %s...\n", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
