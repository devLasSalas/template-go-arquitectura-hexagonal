package dominio_adaptadores_clientes_http

import "genesis/pos/reportes_pos/dominio/entidades"

type IClienteHttp interface {
	Enviar(metodo string, url string, mensaje *entidades.HttpRequest) (*entidades.HttpResponse, error)
}
