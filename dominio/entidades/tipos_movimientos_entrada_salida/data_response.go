package entidades_tipos_movimientos_entrada_salida

type DataResponse struct {
	Id                       int    `json:"id"`
	Descripcion              string `json:"descripcion"`
	EstadoId                 int    `json:"estado_id"`
	Operacion                string `json:"operacion"`
	CategoriaEntradaSalidaId int    `json:"categoria_entrada_salida_id"`
	OrdenCompra              bool   `json:"orden_compra"`
}
