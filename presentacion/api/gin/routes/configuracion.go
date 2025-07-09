package api_routes

import (
	"genesis/pos/reportes_pos/dominio/constantes"
	controllers_categorias_entrada_salida "genesis/pos/reportes_pos/presentacion/api/controllers/categoria_entrada_salida"
	controllers_tipos_mov_entrada_salida "genesis/pos/reportes_pos/presentacion/api/controllers/tipos_movimientos_entrada_salida"
	presentacion_api_middlewares "genesis/pos/reportes_pos/presentacion/api/middlewares"

	"github.com/gin-gonic/gin"
)

func GinConfig() (*gin.Engine, error) {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group(constantes.API_PATH)

	// Definicion de rutas
	api.GET(constantes.API_TIPOS_OPERACIONES, presentacion_api_middlewares.SessionMiddleware(), controllers_categorias_entrada_salida.GetCategoriaEntradaSalidaController)
	api.GET(constantes.API_TIPOS_MOVIMIENTOS+"/:id_categoria", presentacion_api_middlewares.SessionMiddleware(), controllers_tipos_mov_entrada_salida.GetTiposMovEntradaSalidaController)
	return router, nil

}
