package api_routes

import (
	"github.com/gin-gonic/gin"
)

func GinConfig() (*gin.Engine, error) {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Definicion de rutas
	// router.GET(constantes.API_PATH, middleware(), controladores.Api)
	return router, nil

}
