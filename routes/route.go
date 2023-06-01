package routes

import (
	"assignment-3/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.StaticFS("/public", http.Dir("public"))
	router.GET("/", controller.CreateLog)

	return router
}
