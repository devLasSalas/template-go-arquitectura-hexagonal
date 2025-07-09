package aplicacion_casosusos_bodegas

import (
	entidades_bodegas "genesis/pos/reportes_pos/dominio/entidades/bodegas"
	dominio_repositorios_bodegas "genesis/pos/reportes_pos/dominio/repositorios/db/bodegas"
)

type GetBodegasPorEmpresaUseCase struct {
	Repositorio dominio_repositorios_bodegas.GetBodegasPorEmpresaInterface
}

func (GCESU *GetBodegasPorEmpresaUseCase) Execute(id_empresa int) (*[]entidades_bodegas.DataResponse, error) {
	return GCESU.Repositorio.Get(id_empresa)
}
