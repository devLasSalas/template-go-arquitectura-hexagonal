package dominio_casosusos_bodegas

import (
	entidades_bodegas "genesis/pos/reportes_pos/dominio/entidades/bodegas"
)

type GetBodegasPorEmpresaUseCase interface {
	Execute(id_empresa int) (*[]entidades_bodegas.DataResponse, error)
}
