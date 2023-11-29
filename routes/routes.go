package routes

import (
	"tiktok/control"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.POST("/publish", control.Publish)
	r.Run(":9999")
}
