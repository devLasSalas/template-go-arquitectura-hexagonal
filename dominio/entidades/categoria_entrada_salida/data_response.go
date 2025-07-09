package entidades_categoria_entrada_salida

type DataResponse struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
	Operacion   string `json:"operacion"`
	EstadoId    int    `json:"estado_id"`
}
