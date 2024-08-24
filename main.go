package main

import (
	"DailyPunch/config"
	"DailyPunch/internal/controller"
	"DailyPunch/object"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dataSourceName := config.Cfg.Database.DataSourceName
	object.InitDB(dataSourceName)

	router := gin.Default()

	conf := cors.DefaultConfig()
	conf.AllowOrigins = []string{"*"}
	conf.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	conf.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	router.Use(cors.New(conf))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	//auth := router.Group("/api")
	//auth.Use(pkg.JWTAuthMiddleware())
	//{
	//	auth.POST("/login", controller.Login)
	//	auth.POST("/logout", controller.Logout)
	//}
	router.POST("/login", controller.Login)
	router.POST("/logout", controller.Logout)

	router.Run(":8080")
}
