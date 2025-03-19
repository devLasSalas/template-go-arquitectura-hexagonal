package middleware

import (
	"fmt"
	"genesis/pos/reportes_pos/dominio/entidades"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ValidarDatosInventario() gin.HandlerFunc {
	return validarParamsVacio
}

func validarParamsVacio(ctx *gin.Context) {
	 _,err:= strconv.Atoi(ctx.Param("id_eds"));
 
	if err !=nil {
		mensajeError := entidades.GetMensajeError(http.StatusBadRequest, "JSON mal formado", err.Error())
		fmt.Println("MENSAJE ERROR", mensajeError)
		ctx.JSON(http.StatusBadRequest, mensajeError)
		ctx.Abort()
		return
	}
	ctx.Next()
}
