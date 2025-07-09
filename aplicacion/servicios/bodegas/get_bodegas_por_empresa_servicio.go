package aplicacion_servicios_bodegas

import (
	dominio_casosusos_bodegas "genesis/pos/reportes_pos/dominio/casosusos/bodegas"
	entidades_bodegas "genesis/pos/reportes_pos/dominio/entidades/bodegas"
)

type GetBodegasPorEmpresaService struct {
	GetBodegasPorEmpresaUseCase dominio_casosusos_bodegas.GetBodegasPorEmpresaUseCase
}

func (GBPES *GetBodegasPorEmpresaService) Execute(id_empresa int) (*[]entidades_bodegas.DataResponse, error) {
	datos, err := GBPES.GetBodegasPorEmpresaUseCase.Execute(id_empresa)
	if err != nil {
		return nil, err
	}

	return datos, nil
}
