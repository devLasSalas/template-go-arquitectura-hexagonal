package dominio_repositorios_bodegas

import entidades_bodegas "genesis/pos/reportes_pos/dominio/entidades/bodegas"

type GetBodegasPorEmpresaInterface interface {
	Get(id_empresa int) (*[]entidades_bodegas.DataResponse, error)
}
