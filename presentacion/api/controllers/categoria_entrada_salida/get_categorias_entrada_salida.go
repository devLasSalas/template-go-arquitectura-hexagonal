package controllers_categorias_entrada_salida

import (
	"errors"
	"genesis/pos/reportes_pos/dominio/entidades"
	dominio_logs "genesis/pos/reportes_pos/dominio/logs"
	"genesis/pos/reportes_pos/presentacion/contenedor"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategoriaEntradaSalidaController(c *gin.Context) {
	dominio_logs.Info("Recibida solicitud en GetCategoriaEntradaSalidaController")

	processedData, err := contenedor.GetCategoriaEntradaSalidaService.Execute()
	if err != nil {
		dominio_logs.Error("Error al procesar datos: %v", err)
		c.JSON(http.StatusInternalServerError, entidades.NewErrorResponse("Error interno del servidor", errors.New("ups, algo salió mal al obtener la información")))
		return
	}

	dominio_logs.Info("datos procesados exitosamente: %s", processedData)
	c.JSON(http.StatusOK, entidades.NewSuccessResponse("Datos procesados exitosamente", processedData))
}
