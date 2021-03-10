package main

import (
	"golang_gorm/api"
	"golang_gorm/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(CORS())

	database.ConnectDatabase()

	r.GET("/", api.GetAllListMethod)
	r.GET("/user", api.GetUserMethod)
	r.POST("/", api.CreateUserMethod)
	r.DELETE("/:id", api.DeleteUserMethod)
	r.POST("/upload", api.UploadFileMethod)
	r.StaticFS("/file", http.Dir("public"))

	r.Run()
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
