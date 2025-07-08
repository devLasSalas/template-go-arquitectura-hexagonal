package cliente_infrastruc

import (
	"bytes"
	"genesis/pos/reportes_pos/dominio/entidades"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

type ClienteHttp struct {
	cliente *http.Client
}

func (CHttp *ClienteHttp) Enviar(metodo string, url string, mensaje *entidades.HttpRequest) (*entidades.HttpResponse, error) {

	request, err := http.NewRequest(metodo, "http://"+url, bytes.NewBuffer(mensaje.Body))
	if err != nil {
		log.Printf("ERROR DEL HOS %s : %v", url, err)
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := CHttp.cliente.Do(request)

	if err != nil {
		log.Printf("DESPUES DEL DO %v \n", err)
		errorNet := err.(net.Error)
		return nil, errorNet
	}

	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	responesDominio := &entidades.HttpResponse{
		StatusCode: response.StatusCode,
		Body:       bodyBytes,
		Status:     response.Status,
	}

	log.Printf("RESPUESTA DE PETICION %s : %d \n", url, responesDominio.StatusCode)
	log.Printf("BODY : %v \n", string(bodyBytes))
	return responesDominio, nil
}

func InicializarCliente() (*ClienteHttp, error) {
	httpCliente := &http.Client{
		Timeout: time.Second * 5,
	}
	return &ClienteHttp{
		cliente: httpCliente,
	}, nil
}
