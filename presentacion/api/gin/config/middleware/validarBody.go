package middleware

import (
	"bytes"
	"errors"
	"fmt"
	"genesis/pos/reportes_pos/dominio/entidades"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func ValidarBody() gin.HandlerFunc {
	return validarBodyVacio
}

func validarBodyVacio(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		mensajeError := entidades.GetMensajeError(http.StatusBadRequest, "JSON mal formado", err.Error())
		log.Println("MENSAJE ERROR", mensajeError)
		ctx.JSON(http.StatusBadRequest, mensajeError)
		ctx.Abort()

		return
	}
	match, err := regexp.Match(`^\s*\{\s*\}\s*$`, bodyBytes)
	if err != nil {
		fmt.Print(err)
		err := errors.New("body no corresponde con el esquema esperado ")
		mensajeError := entidades.GetMensajeError(http.StatusBadRequest, "JSON mal formado", err.Error())
		fmt.Println("MENSAJE ERROR", mensajeError)
		ctx.JSON(http.StatusBadRequest, mensajeError)
		ctx.Abort()
		return
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	if len(bodyBytes) == 0 || match {
		fmt.Println("@@@@@@AQUI NO ES VALIDO")
		err := errors.New("body no corresponde con el esquema esperado ")
		mensajeError := entidades.GetMensajeError(http.StatusBadRequest, "JSON mal formado", err.Error())
		fmt.Println("MENSAJE ERROR", mensajeError)
		ctx.JSON(http.StatusBadRequest, mensajeError)
		ctx.Abort()
		return
	}
	ctx.Next()
}
