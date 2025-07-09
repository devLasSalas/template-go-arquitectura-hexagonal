package controllers_bodegas

import (
	"errors"
	"genesis/pos/reportes_pos/dominio/entidades"
	dominio_logs "genesis/pos/reportes_pos/dominio/logs"
	"genesis/pos/reportes_pos/presentacion/contenedor"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBodegasPorEmpresaController(c *gin.Context) {
	dominio_logs.Info("Recibida solicitud en GetBodegasPorEmpresaController")

	empresaIdStr := c.Param("id_empresa")
	empresa_id, err := strconv.Atoi(empresaIdStr)
	if err != nil {
		dominio_logs.Error("Error al obtener parametro de la url: %v", err)
		c.JSON(http.StatusBadRequest, entidades.NewErrorResponse("Error obtener parametro", errors.New("no se pudo obtener parametro de la URL")))
		return
	}
	log.Println("🔍 empresa_id:", empresa_id)

	processedData, err := contenedor.GetBodegasPorEmpresaService.Execute(empresa_id)
	if err != nil {
		dominio_logs.Error("Error al procesar datos: %v", err)
		c.JSON(http.StatusInternalServerError, entidades.NewErrorResponse("Error interno del servidor", errors.New("ups, algo salió mal al obtener la información")))
		return
	}

	dominio_logs.Info("datos procesados exitosamente: %s", processedData)
	c.JSON(http.StatusOK, entidades.NewSuccessResponse("Datos procesados exitosamente", processedData))
}
