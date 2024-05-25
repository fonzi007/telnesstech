package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"telnesstech/api"
	"time"
)

func Run(addr string, handler api.Handler) {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000"},
		AllowMethods:  []string{"POST", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))
	router.POST("/search", handler.Search)
	router.Run(addr)
}
