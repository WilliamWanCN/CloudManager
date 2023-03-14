package router

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	server := gin.Default()

	server.Run()
}
