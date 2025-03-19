package entidades

import "time"

type MensajeError struct {
	Codigo       int32     `json:"codigoError"`
	TipoError    string    `json:"tipoError"`
	MensajeError string    `json:"mensajeError"`
	FechaProceso time.Time `json:"fechaProceso"`
}

func GetMensajeError(codigo int32, tipoError string, mensaje string) *MensajeError {
	return &MensajeError{
		Codigo:       codigo,
		TipoError:    tipoError,
		MensajeError: mensaje,
		FechaProceso: time.Now(),
	}
}
