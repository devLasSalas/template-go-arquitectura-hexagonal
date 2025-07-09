package entidades_bodegas

type DataResponse struct {
	Id            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Codigo        string `json:"codigo"`
	EstadoId      int    `json:"estado_id"`
	EmpresasId    int    `json:"empresas_id"`
	TipoNegocioId int    `json:"tipo_negocio_id"`
}
