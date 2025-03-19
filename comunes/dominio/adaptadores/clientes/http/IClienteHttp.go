package comunes_http_clientes

import comunes_entidades "genesis/pos/reportes_pos/comunes/dominio/entidades"

type IClienteHttp interface {
	Enviar(metodo string, url string, mensaje *comunes_entidades.HttpRequest) (*comunes_entidades.HttpResponse, error)
}
